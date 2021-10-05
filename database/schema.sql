CREATE TABLE "discord_user" (
    id varchar(32) NOT NULL,
    username varchar(150) NOT NULL,
    points int DEFAULT 0,
    ban_alerts int DEFAULT 0,
    PRIMARY KEY (id)
);

CREATE TABLE "role" (
    id varchar(32) NOT NULL,
    name varchar(32) NOT NULL,
    min_points int,
    PRIMARY KEY (id)
);

CREATE TABLE "rank" (
    user_id varchar(32) NOT NULL,
    role_id varchar(32) NOT NULL,
    PRIMARY KEY (user_id, role_id),
    FOREIGN KEY (user_id) REFERENCES "discord_user" (id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (role_id) REFERENCES "role" (id) ON UPDATE CASCADE ON DELETE CASCADE
);