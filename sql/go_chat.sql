/*
 Navicat MySQL Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50726
 Source Host           : localhost:3306
 Source Schema         : go_chat

 Target Server Type    : MySQL
 Target Server Version : 50726
 File Encoding         : 65001

 Date: 12/12/2022 17:35:47
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for contact
-- ----------------------------
DROP TABLE IF EXISTS `contact`;
CREATE TABLE `contact` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `owner_id` bigint(20) DEFAULT NULL COMMENT '谁的关系id',
  `target_id` bigint(20) DEFAULT NULL COMMENT '对应id',
  `type` tinyint(1) DEFAULT NULL COMMENT '聊天类型(0群聊，1私聊，2广播)',
  `description` text COMMENT '描述',
  PRIMARY KEY (`id`),
  KEY `idx_contact_deleted_at` (`deleted_at`),
  KEY `fid` (`owner_id`),
  KEY `tid` (`target_id`)
) ENGINE=InnoDB AUTO_INCREMENT=26 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for group_basic
-- ----------------------------
DROP TABLE IF EXISTS `group_basic`;
CREATE TABLE `group_basic` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `owner_id` bigint(20) DEFAULT NULL COMMENT '拥有者id',
  `name` varchar(20) DEFAULT NULL COMMENT '名称',
  `type` tinyint(1) DEFAULT NULL COMMENT '类型',
  `icon` varchar(128) DEFAULT NULL COMMENT '图片',
  `description` text COMMENT '描述',
  PRIMARY KEY (`id`),
  KEY `idx_group_basic_deleted_at` (`deleted_at`),
  KEY `fid` (`owner_id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for message
-- ----------------------------
DROP TABLE IF EXISTS `message`;
CREATE TABLE `message` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` bigint(20) DEFAULT NULL COMMENT '发送者id',
  `target_id` bigint(20) DEFAULT NULL COMMENT '接收者id',
  `type` tinyint(1) DEFAULT NULL COMMENT '聊天类型(0群聊，1私聊，2广播)',
  `media` tinyint(1) DEFAULT NULL COMMENT '消息类型(文字，图片，音频)',
  `content` text COMMENT '内容',
  `pic` text COMMENT '图片',
  `url` text COMMENT 'url',
  `description` text COMMENT '描述',
  `amount` int(10) DEFAULT NULL COMMENT '其他统计字段',
  PRIMARY KEY (`id`),
  KEY `idx_message_deleted_at` (`deleted_at`),
  KEY `fid` (`user_id`),
  KEY `tid` (`target_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for user_basic
-- ----------------------------
DROP TABLE IF EXISTS `user_basic`;
CREATE TABLE `user_basic` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(20) DEFAULT NULL,
  `password` varchar(128) DEFAULT NULL,
  `phone` varchar(20) DEFAULT NULL,
  `email` varchar(128) DEFAULT NULL,
  `client_ip` varchar(32) DEFAULT NULL,
  `identity` varchar(32) DEFAULT '',
  `client_port` varchar(32) DEFAULT NULL,
  `login_time` datetime(3) DEFAULT NULL,
  `heartbeat_time` datetime(3) DEFAULT NULL,
  `login_out_time` datetime(3) DEFAULT NULL,
  `is_login_out` tinyint(1) DEFAULT NULL,
  `device_info` varchar(128) DEFAULT NULL,
  `icon` varchar(128) DEFAULT NULL,
  `nickname` varchar(128) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `idx_name` (`name`),
  KEY `idx_user_basic_deleted_at` (`deleted_at`),
  KEY `idx_phone` (`phone`)
) ENGINE=InnoDB AUTO_INCREMENT=1601854192220114945 DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;
