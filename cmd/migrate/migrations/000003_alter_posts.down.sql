-- 000003_add_fk_to_users.down.sql
ALTER TABLE posts DROP CONSTRAINT IF EXISTS fk_user;