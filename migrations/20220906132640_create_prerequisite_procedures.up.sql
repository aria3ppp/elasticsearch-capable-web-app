-- add complement contribution and deleted columns to table
create or replace procedure add_contribution_and_delete_columns(
	p_table text,
	p_contributer_table text,
	p_contributer_table_pk text,
	
	p_column_contributed_by_name text,
	p_column_contributed_by_type text,
	p_column_contributed_by_fk_constraint_name text,
	
	p_column_contributed_at_name text,
	p_column_contributed_at_type text,
	
	p_column_deleted_name text,
	p_column_deleted_type text
)
language plpgsql
as $$
begin
	-- check table exists
	perform from information_schema.tables
	where table_name = p_table and table_type = 'BASE TABLE';
	
	if not found then
		raise exception 'table name "%" not found', p_table;
	end if;
	
	-- check contributer referenced table exists
	perform from information_schema.tables
	where table_name = p_contributer_table and table_type = 'BASE TABLE';
	
	if not found then
		raise exception 'contributer table name "%" not found', p_contributer_table;
	end if;
	
	-- add columns
	execute 'ALTER TABLE ' || quote_ident(p_table) || ' '
		|| 'ADD COLUMN ' || quote_ident(p_column_contributed_by_name) || ' ' || p_column_contributed_by_type || ', '
		|| 'ADD COLUMN ' || quote_ident(p_column_contributed_at_name) || ' ' || p_column_contributed_at_type || ', '
		|| 'ADD COLUMN ' || quote_ident(p_column_deleted_name) || ' ' || p_column_deleted_type;
				   
	-- add contributed_by foreign key
	execute 'ALTER TABLE ' || quote_ident(p_table) || ' '
				|| 'ADD CONSTRAINT ' || quote_ident(p_column_contributed_by_fk_constraint_name) || ' '
				|| 'FOREIGN KEY (' || quote_ident(p_column_contributed_by_name) || ') '
				|| 'REFERENCES ' || quote_ident(p_contributer_table) || ' (' || quote_ident(p_contributer_table_pk) || ')';
				
end;
$$;

-- create an audit table with index on table primary key column
create or replace procedure create_audit_table(
	p_table text,
	p_audit_table_name text,
	p_audit_table_index_name text,
	p_audit_table_index_column text
)
language plpgsql
as $$
declare
	v_row RECORD;
    v_CREATE_AUDIT_TABLE_BODY TEXT;
	v_CREATE_AUDIT_TABLE_CMD TEXT;
begin
	perform from information_schema.tables
	where table_name = p_table and table_type = 'BASE TABLE';
	
	if not found then
		raise exception 'table name "%" not found', p_table;
	end if;

    v_CREATE_AUDIT_TABLE_BODY = '';
	
	for v_row in
		select column_name, data_type, is_nullable, ordinal_position
		from information_schema.columns
		where table_name = p_table
		order by ordinal_position
	loop
	
		v_CREATE_AUDIT_TABLE_BODY = v_CREATE_AUDIT_TABLE_BODY || quote_ident(v_row.column_name) || ' ' || v_row.data_type;
		
		if v_row.is_nullable = 'NO' then
			v_CREATE_AUDIT_TABLE_BODY = v_CREATE_AUDIT_TABLE_BODY || ' NOT NULL';
		end if;
		
		-- if this column is not the last one append a comma (,) at the end
		if not (
			select MAX(ordinal_position) = v_row.ordinal_position
			from information_schema.columns
			where table_name = p_table
		) then
			v_CREATE_AUDIT_TABLE_BODY = v_CREATE_AUDIT_TABLE_BODY || ', ';
		end if;
		
	end loop;
    
	v_CREATE_AUDIT_TABLE_CMD = 'CREATE TABLE IF NOT EXISTS ' || quote_ident(p_audit_table_name) || ' (' || v_CREATE_AUDIT_TABLE_BODY || ')';
		
	-- create the audit table
	execute v_CREATE_AUDIT_TABLE_CMD;
	
	-- create index
	execute 'CREATE INDEX ' || quote_ident(p_audit_table_index_name)
			|| ' ON ' || quote_ident(p_audit_table_name) || '(' || quote_ident(p_audit_table_index_column) || ')';
	
end;
$$;

-- create a trigger on update to save old record into audit table
create or replace procedure create_audit_update_trigger_on_table(
	p_table text,
	p_table_contributed_at_column text,
	p_audit_table_name text,
	p_trigger_name text,
	p_trigger_function_name text
)
language plpgsql
as $body$
declare
	v_trigger_func_body text;
	v_trigger_func_cmd text;
	v_create_trigger_on_table_cmd text;
begin
	-- build trigger function
	v_trigger_func_body = 'BEGIN '
			|| 'INSERT INTO ' || quote_ident(p_audit_table_name) || ' SELECT OLD.*; '
			|| 'NEW.' || p_table_contributed_at_column || ' = CURRENT_TIMESTAMP; '
			|| 'RETURN NEW; '
			|| 'END;';
	
	v_trigger_func_cmd = 'CREATE OR REPLACE FUNCTION ' || p_trigger_function_name || '() RETURNS TRIGGER ' 
						|| 'LANGUAGE plpgsql AS $$ ' || v_trigger_func_body || ' $$';
	
	raise notice 'trigger func cmd: %', v_trigger_func_cmd;
	
	-- create trigger function
	execute v_trigger_func_cmd;
	
	-- build trigger on table
	v_create_trigger_on_table_cmd = 'CREATE TRIGGER ' || p_trigger_name || ' '
									|| 'BEFORE UPDATE ON ' || p_table || ' '
									|| 'FOR EACH ROW EXECUTE FUNCTION ' || p_trigger_function_name || '()';
									
	raise notice 'create trigger cmd: %', v_create_trigger_on_table_cmd;
	
	-- create trigger on table
	execute v_create_trigger_on_table_cmd;
	
end;
$body$;