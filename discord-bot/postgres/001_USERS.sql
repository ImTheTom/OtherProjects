\c discord;

CREATE TABLE users (
    user_id VARCHAR(128) NOT NULL,
	guild_id VARCHAR(128) NOT NULL,
    username VARCHAR(128) NOT NULL,
    nickname VARCHAR(128),
    points INT NOT NULL DEFAULT 0,
    PRIMARY KEY (user_id, guild_id)
);
