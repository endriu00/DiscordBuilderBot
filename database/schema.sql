CREATE TABLE `user` (
    `id` varchar(32) NOT NULL,
    `username` varchar(150) NOT NULL,
    `points` int DEFAULT 0,
    `ban_alerts` int DEFAULT 0,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB COMMENT='Discord user';

CREATE TABLE `role` (
    `id` varchar(32) NOT NULL,
    `name` varchar(32) NOT NULL,
    `min_points` int,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB COMMENT='Role for the server';

CREATE TABLE `rank` (
    `user_id` varchar(32) NOT NULL,
    `role_id` varchar(32) NOT NULL,
    FOREIGN KEY (`user_id`)
    REFERENCES `user`(`id`)
    ON UPDATE CASCADE 
    ON DELETE CASCADE,
    FOREIGN KEY (`role_id`)
    REFERENCES `role`(`id`)
    ON UPDATE CASCADE
    ON DELETE CASCADE
) ENGINE=InnoDB COMMENT='Rank of the user in the server';