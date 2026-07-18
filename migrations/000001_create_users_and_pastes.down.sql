DROP INDEX IF EXISTS pastes_user_updated_at_index;
DROP INDEX IF EXISTS pastes_user_favorite_index;
DROP INDEX IF EXISTS pastes_user_id_index;
DROP INDEX IF EXISTS pastes_public_id_unique;
DROP TABLE IF EXISTS pastes;

DROP INDEX IF EXISTS users_email_lower_unique;
DROP TABLE IF EXISTS users;