CREATE DATABASE IF NOT EXISTS data_workbench;
USE data_workbench;

CREATE TABLE IF NOT EXISTS `api_config` (
    `api_id` VARCHAR(24) NOT NULL,
    `api_name` VARCHAR(65) NOT NULL,
    `api_path` VARCHAR(128) NOT NULL,
    `api_mode` INT NOT NULL,
    `api_description` VARCHAR(256),
    `space_id` VARCHAR(24) NOT NULL,
    `protocols` VARCHAR(65) ,
    `request_method` INT,
    `response_type` INT ,
    `timeout` INT,
    `visible_range` INT,
    `datasource_Id` VARCHAR(24) NOT NULL,
    `table_name` VARCHAR(65) NOT NULL,
    `script` VARCHAR(20000) ,  -- text
     -- => "enabled", 2 => "disabled", 3 => "deleted"
    `status` TINYINT(1) UNSIGNED DEFAULT 1 NOT NULL,
--     `created` BIGINT(20) UNSIGNED NOT NULL,
--     `updated` BIGINT(20) UNSIGNED NOT NULL,
    PRIMARY KEY (`api_id`)
) ENGINE=InnoDB COMMENT='The dataservice api config';

CREATE TABLE IF NOT EXISTS `api_request_parameters` (
    `param_id` VARCHAR(24) NOT NULL,
    `api_id` VARCHAR(24) NOT NULL,
    `column_name` VARCHAR(65) NOT NULL,
    `default_value` VARCHAR(65),
    `example_value` VARCHAR(65),
    `is_required`  TINYINT(1) UNSIGNED DEFAULT 1,
    `data_type` INT ,
    `param_description` VARCHAR(256),
    `param_name` VARCHAR(65),
    `param_operator` INT,
    `param_position` INT,
--     `created` BIGINT(20) UNSIGNED NOT NULL,
--     `updated` BIGINT(20) UNSIGNED NOT NULL,
    PRIMARY KEY (`param_id`)
    ) ENGINE=InnoDB COMMENT='The dataservice api request parameters';

CREATE TABLE IF NOT EXISTS `api_response_parameters` (
    `param_id` VARCHAR(24) NOT NULL,
    `api_id` VARCHAR(24) NOT NULL,
    `column_name` VARCHAR(65) NOT NULL,
    `default_value` VARCHAR(65),
    `example_value` VARCHAR(65),
    `data_type` INT ,
    `param_description` VARCHAR(256),
    `param_name` VARCHAR(65),
--     `created` BIGINT(20) UNSIGNED NOT NULL,
--     `updated` BIGINT(20) UNSIGNED NOT NULL,
     PRIMARY KEY (`param_id`)
) ENGINE=InnoDB COMMENT='The dataservice api response parameters';