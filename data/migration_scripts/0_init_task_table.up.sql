CREATE TABLE `tasks` (
     `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
     `name` varchar(64) NOT NULL DEFAULT '',
     `status` tinyint(4) NOT NULL DEFAULT '0',
     PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
