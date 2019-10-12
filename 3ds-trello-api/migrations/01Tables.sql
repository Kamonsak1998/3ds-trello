CREATE DATABASE IF NOT EXISTS `3dsManagementTrello` CHARACTER SET utf8 COLLATE utf8_unicode_ci;

USE `3dsManagementTrello`;

CREATE TABLE `3dsManagementTrello`.`boards`
(
    `board_id`   VARCHAR(30) NOT NULL,
    `board_name` VARCHAR(50) NOT NULL,
    PRIMARY KEY (`board_id`)
) ENGINE = InnoDB
  DEFAULT CHARACTER SET = utf8
  COLLATE = utf8_unicode_ci;


CREATE TABLE `3dsManagementTrello`.`members`
(
    `member_id`   VARCHAR(30)  NOT NULL,
    `member_name` VARCHAR(50)  NOT NULL,
    `avatar_hash` VARCHAR(200) NULL,
    PRIMARY KEY (`member_id`)
) ENGINE = InnoDB
  DEFAULT CHARACTER SET = utf8
  COLLATE = utf8_unicode_ci;


CREATE TABLE `3dsManagementTrello`.`cards`
(
    `card_id`     varchar(30)  NOT NULL,
    `board_id`    VARCHAR(30)  NOT NULL,
    `member_id`   VARCHAR(30)  NOT NULL,
    `card_name`   VARCHAR(200) NOT NULL,
    `card_time`   DATE         NOT NULL,
    `card_point`  FLOAT,
    `sprint_name` VARCHAR(9)   NULL,
    PRIMARY KEY (`card_id`)
) ENGINE = InnoDB
  DEFAULT CHARACTER SET = utf8
  COLLATE = utf8_unicode_ci;

ALTER TABLE `cards`
    ADD KEY `board_id` (`board_id`),
    ADD KEY `member_id` (`member_id`),
    ADD CONSTRAINT `card_tb_ibfk_1` FOREIGN KEY (`board_id`) REFERENCES `boards` (`board_id`),
    ADD CONSTRAINT `card_tb_imfk_1` FOREIGN KEY (`member_id`) REFERENCES `members` (`member_id`);


CREATE TABLE `3dsManagementTrello`.`sprints`
(
    `sprint_id`   INT(10)     NOT NULL,
    `board_id`    VARCHAR(30) NOT NULL,
    `sprint_name` VARCHAR(9)  NOT NULL,
    `start_date`  DATE        NOT NULL,
    `end_date`    DATE        NOT NULL,
    `sprint_day`  INT(10),
    PRIMARY KEY (`sprint_id`)
) ENGINE = InnoDB
  DEFAULT CHARACTER SET = utf8
  COLLATE = utf8_unicode_ci;

ALTER TABLE `sprints`
    MODIFY `sprint_id` int(10) NOT NULL AUTO_INCREMENT,
    ADD KEY `board_id` (`board_id`),
    ADD CONSTRAINT `sprint_tb_ibfk_1` FOREIGN KEY (`board_id`) REFERENCES `boards` (`board_id`);

CREATE TABLE `3dsManagementTrello`.`score_sizes`
(
    `size_id`    varchar(5) NOT NULL,
    `size_point` FLOAT      NOT NULL,
    PRIMARY KEY (`size_id`)
) ENGINE = InnoDB
  DEFAULT CHARACTER SET = utf8
  COLLATE = utf8_unicode_ci;

CREATE TABLE `3dsManagementTrello`.`lists`
(
    `list_id`       varchar(30)  NOT NULL,
    `list_name`     varchar(100) NOT NULL,
    `board_id`      VARCHAR(30)  NOT NULL,
    `list_priority` int(1)       NOT NULL,
    PRIMARY KEY (`list_id`)
) ENGINE = InnoDB
  DEFAULT CHARACTER SET = utf8
  COLLATE = utf8_unicode_ci;

ALTER TABLE `lists`
    ADD KEY `board_id` (`board_id`),
    ADD CONSTRAINT `lists_tb_ibfk_1` FOREIGN KEY (`board_id`) REFERENCES `boards` (`board_id`);


