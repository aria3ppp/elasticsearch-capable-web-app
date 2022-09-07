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
	
	p_column_contributed_by_name => 'user_id',
	p_column_contributed_by_type => 'INT NOT NULL',
	p_column_contributed_by_fk_constraint_name => concat_ws('_', 'posts', 'fk', 'users'),
	
	p_column_contributed_at_name => 'contributed_at',
	p_column_contributed_at_type => 'TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP',
	
	p_column_deleted_name => 'deleted',
	p_column_deleted_type => 'BOOLEAN NOT NULL DEFAULT FALSE'
);

call create_audit_table(
	p_table => 'posts',
	p_audit_table_name => concat_ws('_', 'posts', 'audit'),
	p_audit_table_pk_columns_order_sep_by_comma => 'id, user_id, contributed_at'
);

call create_audit_update_trigger_on_table(
	p_table => 'posts',
	p_table_contributed_at_column => 'contributed_at',
	p_audit_table_name => concat_ws('_', 'posts', 'audit'),
	p_trigger_name => concat_ws('_', 'trigger', 'posts', 'update', 'audit'),
	p_trigger_function_name => concat_ws('_', 'update', 'posts', 'trigger', 'audit')
);

COMMIT;