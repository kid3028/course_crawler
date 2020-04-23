/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50726
Source Host           : localhost:3306
Source Database       : course

Target Server Type    : MYSQL
Target Server Version : 50726
File Encoding         : 65001

Date: 2020-04-23 23:00:27
*/

SET FOREIGN_KEY_CHECKS=0;

create database `course` default charset=utf8;

use `course`;

-- ----------------------------
-- Table structure for course_pkg
-- ----------------------------
DROP TABLE IF EXISTS `course_pkg`;
CREATE TABLE `course_pkg` (
  `package_id` varchar(128) NOT NULL COMMENT '主键id',
  `title` varchar(128) NOT NULL DEFAULT '' COMMENT '名称',
  `grade` int(11) NOT NULL COMMENT '年级',
  `subject` int(11) NOT NULL COMMENT '科目',
  `course_bg_time` datetime NOT NULL COMMENT '开课时间',
  `course_end_time` datetime NOT NULL COMMENT '结课时间',
  `course_sign_end_time` datetime NOT NULL COMMENT '创建时间',
  `sold_count` int(11) NOT NULL DEFAULT '0' COMMENT '销量',
  `season` int(11) NOT NULL DEFAULT '0' COMMENT '期数',
  `course_min_price` int(11) NOT NULL DEFAULT '0' COMMENT '最低价',
  `course_max_price` int(11) NOT NULL DEFAULT '0' COMMENT '最高价',
  `discount_price` int(11) NOT NULL DEFAULT '0' COMMENT '折扣价',
  `create_time` datetime NOT NULL COMMENT '新增时间',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`package_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='系统课程包';

-- ----------------------------
-- Table structure for course_pkg_relate
-- ----------------------------
DROP TABLE IF EXISTS `course_pkg_relate`;
CREATE TABLE `course_pkg_relate` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `cid` int(11) NOT NULL COMMENT '课程id',
  `pkg_id` varchar(128) NOT NULL COMMENT '课程包id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_cid` (`cid`),
  KEY `idx_pkg_id` (`pkg_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='课程包课程';

-- ----------------------------
-- Table structure for grade_subject
-- ----------------------------
DROP TABLE IF EXISTS `grade_subject`;
CREATE TABLE `grade_subject` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `grade` int(11) NOT NULL COMMENT '年级',
  `subject` int(11) NOT NULL COMMENT '科目',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='年级科目表';

-- ----------------------------
-- Table structure for spe_course
-- ----------------------------
DROP TABLE IF EXISTS `spe_course`;
CREATE TABLE `spe_course` (
  `cid` int(11) NOT NULL COMMENT '主键id',
  `name` varchar(256) NOT NULL DEFAULT '' COMMENT '课程名称',
  `cover_url` varchar(256) NOT NULL DEFAULT '' COMMENT '封面',
  `grade` int(11) NOT NULL COMMENT '年级',
  `subject` int(11) NOT NULL COMMENT '科目',
  `has_discount` tinyint(2) NOT NULL DEFAULT '0' COMMENT '是否折扣 0：否 1：是',
  `pre_amount` int(11) NOT NULL COMMENT '原价',
  `af_amount` int(11) NOT NULL DEFAULT '0' COMMENT '折后价',
  `record_time` date NOT NULL COMMENT '记录时间',
  `first_sub_bg_time` datetime NOT NULL COMMENT '开课时间',
  `first_sub_end_time` datetime NOT NULL COMMENT '结课时间',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`cid`),
  KEY `idx_subject` (`subject`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='专题课';

-- ----------------------------
-- Table structure for teacher
-- ----------------------------
DROP TABLE IF EXISTS `teacher`;
CREATE TABLE `teacher` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `cid` int(11) NOT NULL COMMENT '教授课程',
  `name` varchar(128) NOT NULL DEFAULT '' COMMENT '名称',
  `cover_url` varchar(256) NOT NULL DEFAULT '' COMMENT '头像',
  `introduce` varchar(256) NOT NULL DEFAULT '' COMMENT '介绍',
  `tutor` tinyint(2) NOT NULL COMMENT '是否是辅导 0：教师 1：辅导',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_cid` (`cid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='课程教师';
