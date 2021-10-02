CREATE TABLE `user` (
    `id` varchar(32) NOT NULL,
    `username` varchar(150) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='Discord user';

CREATE TABLE `user_stat` (
    `points` int NOT NULL,
    `ban_alerts` int,
    `id` varchar(32) NOT NULL,
    FOREIGN KEY (`id`) 
    REFERENCES `user`(`id`)
    ON UPDATE CASCADE
    ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='Statistics for a user';