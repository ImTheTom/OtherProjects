\c discord;

CREATE TABLE gambles (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(128) NOT NULL,
	guild_id VARCHAR(128) NOT NULL,
    amount INT NOT NULL DEFAULT 0,
    winner boolean NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

ALTER TABLE gambles
    ADD CONSTRAINT FK_USER_GUILD
    FOREIGN KEY(user_id, guild_id)
    REFERENCES users(user_id, guild_id);

CREATE INDEX gambles_user_guild_idx ON gambles(user_id, guild_id);