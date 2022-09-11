BEGIN;

CREATE TABLE IF NOT EXISTS posts (
    id SERIAL PRIMARY KEY,
    title VARCHAR(150) NOT NULL,
    body TEXT NOT NULL
);

call add_contribution_and_delete_columns(
	p_table => 'posts',
	p_contributer_table => 'users',
	p_contributer_table_pk => 'id',
	
	p_column_contributed_by_name => 'contributed_by',
	p_column_contributed_by_type => 'INT NOT NULL',
	p_column_contributed_by_fk_constraint_name => 'posts_contributed_by_fk_users',
	
	p_column_contributed_at_name => 'contributed_at',
	p_column_contributed_at_type => 'TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP',
	
	p_column_deleted_name => 'deleted',
	p_column_deleted_type => 'BOOLEAN NOT NULL DEFAULT FALSE'
);

call create_audit_table(
	p_table => 'posts',
	p_audit_table_name => 'posts_audit',
	p_audit_table_pk_columns_order_sep_by_comma => 'id, contributed_by, contributed_at'
);

call build_trigger_audit_on_update(
	p_table => 'posts',
	p_table_contributed_at_column => 'contributed_at',
	p_audit_table_name => 'posts_audit',
	p_trigger_name => 'posts_trigger_audit_on_update',
	p_trigger_function_name => 'posts_function_triggers_on_update'
);

COMMIT;