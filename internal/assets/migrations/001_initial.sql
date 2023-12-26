-- +migrate Up

-- +migrate StatementBegin
DO $$
    BEGIN
        IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'oauth2_provider_enum') THEN
            CREATE TYPE oauth2_provider_enum AS ENUM
                ('google', 'facebook', 'twitter', 'github', 'aws', 'azure');
        END IF;
    END $$;
-- +migrate StatementEnd

CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY,
    author_id INTEGER NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255),
    name VARCHAR(255) NOT NULL,
    email_verified BOOLEAN NOT NULL DEFAULT FALSE,
    oauth2_user BOOLEAN NOT NULL DEFAULT FALSE,
    oauth2_provider oauth2_provider_enum,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX IF NOT EXISTS users_author_id_uindex ON users (author_id);
CREATE UNIQUE INDEX IF NOT EXISTS users_email_uindex ON users (email);

-- +migrate StatementBegin
CREATE OR REPLACE FUNCTION TRIGGER_SET_TIMESTAMP() RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$  LANGUAGE plpgsql;
-- +migrate StatementEnd

-- +migrate StatementBegin
DO $$
    BEGIN
        IF NOT EXISTS (SELECT 1 FROM pg_trigger WHERE tgname = 'users_update_trigger'
                                                  AND tgrelid = 'users'::regclass) THEN
            CREATE TRIGGER users_update_trigger
                AFTER UPDATE
                ON users
                FOR EACH ROW
            EXECUTE PROCEDURE TRIGGER_SET_TIMESTAMP();
        END IF;
    END $$;
-- +migrate StatementEnd

-- +migrate StatementBegin
DO $$
    BEGIN
        IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'post_status_enum') THEN
            CREATE TYPE post_status_enum AS ENUM
                ('new', 'pending', 'confirmed', 'failed');
        END IF;
    END $$;
-- +migrate StatementEnd

CREATE TABLE IF NOT EXISTS posts(
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    title VARCHAR(512) NOT NULL,
    body VARCHAR(2048) NOT NULL,
    status post_status_enum NOT NULL,
    tx_hash TEXT,
    tx_timestamp TIMESTAMP
);

CREATE TABLE IF NOT EXISTS starred_posts(
    id SERIAL NOT NULL,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    post_id INTEGER NOT NULL REFERENCES posts(id) ON DELETE CASCADE
);

CREATE UNIQUE INDEX starred_posts_user_post_uindex ON starred_posts(user_id, post_id);

CREATE TABLE IF NOT EXISTS refresh_tokens(
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    token VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    valid_till TIMESTAMP NOT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS refresh_tokens_token_uindex ON refresh_tokens (token);

CREATE TABLE IF NOT EXISTS post_transactions(
    id SERIAL PRIMARY KEY,
    post_id INTEGER NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
    tx bytea
);

CREATE TABLE IF NOT EXISTS oauth2_states(
    id SERIAL PRIMARY KEY,
    state VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    valid_till TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS email_verification_tokens(
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    token VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    token_expires_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS password_recovery_tokens(
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    token VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    token_expires_at TIMESTAMP NOT NULL
);


-- +migrate Down
DROP TABLE IF EXISTS password_recovery_tokens;

DROP TABLE IF EXISTS email_verification_tokens;

DROP TABLE IF EXISTS oauth2_states;

DROP TABLE IF EXISTS post_transactions;

DROP INDEX IF EXISTS refresh_tokens_user_id_index;
DROP INDEX IF EXISTS refresh_tokens_token_uindex;
DROP TABLE IF EXISTS refresh_tokens;

DROP INDEX IF EXISTS starred_posts_user_post_uindex;
DROP TABLE IF EXISTS starred_posts;

DROP TABLE IF EXISTS posts;
DROP TYPE IF EXISTS post_status_enum;

DROP TRIGGER IF EXISTS users_update_trigger ON users;
DROP INDEX IF EXISTS users_email_uindex;
DROP INDEX IF EXISTS users_author_id_uindex;
DROP TABLE IF EXISTS users;