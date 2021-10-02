CREATE TABLE `user` (
    `id` varchar(32) NOT NULL,
    `username` varchar(150) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='Discord user';

CREATE TABLE `user_stats` (
    `points` int DEFAULT 0,
    `ban_alerts` int DEFAULT 0,
    `id` varchar(32) NOT NULL,
    FOREIGN KEY (`id`) 
    REFERENCES `user`(`id`)
    ON UPDATE CASCADE
    ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='Statistics for a user';