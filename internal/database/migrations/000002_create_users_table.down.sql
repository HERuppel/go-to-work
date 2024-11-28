ALTER TABLE users DROP CONSTRAINT users_address_id_fk;
DROP TRIGGER IF EXISTS trigger_user_updated_at ON users;
DROP FUNCTION updated_at();
DROP TABLE IF EXISTS users;
DROP TYPE role_type;