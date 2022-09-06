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
	p_column_contributed_by_fk_constraint_name => concat_ws('_', 'posts', 'fk', 'users'),
	
	p_column_contributed_at_name => 'contributed_at',
	p_column_contributed_at_type => 'TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP',
	
	p_column_deleted_name => 'deleted',
	p_column_deleted_type => 'BOOLEAN NOT NULL DEFAULT FALSE'
);

call create_audit_table(
	p_table => 'posts',
	p_audit_table_name => concat_ws('_', 'posts', 'audit'),
	p_audit_table_pk_columns_order_sep_by_comma => 'id, contributed_by, contributed_at'
);

call create_audit_update_trigger_on_table(
	p_table => 'posts',
	p_table_contributed_at_column => 'contributed_at',
	p_audit_table_name => concat_ws('_', 'posts', 'audit'),
	p_trigger_name => concat_ws('_', 'trigger', 'posts', 'update', 'audit'),
	p_trigger_function_name => concat_ws('_', 'update', 'posts', 'trigger', 'audit')
);

--------------------------------------------------------------------------------------------------------------------

-- CREATE TABLE IF NOT EXISTS vidoes(
-- 	id SERIAL PRIMARY KEY,
-- 	post_id INT NOT NULL,
-- 	post_modified_at TIMESTAMPTZ NOT NULL,
-- 	name TEXT NOT NULL,
-- 	deleted BOOLEAN NOT NULL DEFAULT FALSE,

-- 	FOREIGN KEY (post_id, post_modified_at) REFERENCES posts_store (id,modified_at)
--  );

--  CREATE TABLE IF NOT EXISTS tags (
-- 	id SERIAL PRIMARY KEY,
-- 	name TEXT NOT NULL,
-- 	deleted BOOLEAN NOT NULL DEFAULT FALSE
--  );

--  CREATE TABLE IF NOT EXISTS tags_videos(
-- 	video_id INT NOT NULL,
-- 	tag_id INT NOT NULL,

-- 	PRIMARY KEY (video_id, tag_id),
-- 	FOREIGN KEY (video_id) REFERENCES vidoes (id),
-- 	FOREIGN KEY (tag_id) REFERENCES tags (id)
--  );

COMMIT;