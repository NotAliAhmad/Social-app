-- 000003_add_fk_to_users.up.sql
ALTER TABLE posts ADD CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id);