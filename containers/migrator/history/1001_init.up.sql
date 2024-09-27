CREATE TABLE task
(
    `id`      CHAR(26)     NOT NULL COMMENT 'ULID 26bytes',
    `title`   VARCHAR(191) NOT NULL COMMENT '타이틀',
    `content` TEXT         NOT NULL COMMENT '내용',
    `status`  ENUM('BACKLOG', 'PROGRESS', 'DONE') NOT NULL COMMENT '상태',
    `created` DATETIME     NOT NULL COMMENT '생성시간',
    `updated` DATETIME     NOT NULL COMMENT '업데이트시간',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;