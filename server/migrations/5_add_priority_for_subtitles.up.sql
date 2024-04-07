ALTER TABLE `subtitles`
    ADD COLUMN `priority` INT(11) NOT NULL DEFAULT 0 AFTER `link`;