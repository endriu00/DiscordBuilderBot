CREATE TABLE "discord_user" (
    id TEXT NOT NULL,
    username TEXT NOT NULL,
    points int DEFAULT 0,
    ban_alerts int DEFAULT 0,
    PRIMARY KEY (id)
);

CREATE TABLE "role" (
    id TEXT NOT NULL,
    name TEXT NOT NULL,
    min_points int,
    PRIMARY KEY (id)
);

CREATE TABLE "rank" (
    user_id TEXT NOT NULL,
    role_id TEXT NOT NULL,
    PRIMARY KEY (user_id, role_id),
    FOREIGN KEY (user_id) REFERENCES "discord_user" (id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (role_id) REFERENCES "role" (id) ON UPDATE CASCADE ON DELETE CASCADE
);