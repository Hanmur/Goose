/*
Navicat MySQL Data Transfer

Source Server         : mysqltest
Source Server Version : 80028
Source Host           : localhost:3306
Source Database       : goose

Target Server Type    : MYSQL
Target Server Version : 80028
File Encoding         : 65001

Date: 2022-03-30 21:10:35
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for goose_article
-- ----------------------------
DROP TABLE IF EXISTS `goose_article`;
CREATE TABLE `goose_article` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '文章标题',
  `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '文章简述',
  `cover_image_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '封面图片地址',
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '文章内容',
  `created_on` int unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `created_by` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '创建人',
  `modified_on` int unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '修改人',
  `deleted_on` int unsigned DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除（0为未删除，1为已删除）',
  `state` tinyint(3) unsigned zerofill DEFAULT '001' COMMENT '状态为0为禁用，1为已用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='文章管理';

-- ----------------------------
-- Records of goose_article
-- ----------------------------
INSERT INTO `goose_article` VALUES ('1', 'Hanmur', 'Hanmur', 'www.baidu.com', 'HHHHHHHHHHHHH，Hanmur被爆杀拉！！！', '1646734595', 'Hanmur', '0', '', '0', '0', '001');
INSERT INTO `goose_article` VALUES ('2', '第二次了 我还是不开心', 'Not Happy', 'dsa', '不想再写内容了', '1646857245', 'Hanmur', '1648400704', 'Hanmur', '0', '0', '001');
INSERT INTO `goose_article` VALUES ('3', 'Not Happyaa ', 'Not Happya a ', '\'\'', '我不开心....', '1646857258', 'Hanmur', '1646857258', '', '1646858555', '1', '001');
INSERT INTO `goose_article` VALUES ('4', 'Test', 'Test', '', 'Test', '1646858610', 'Test', '1646858610', '', '1646858625', '1', '001');

-- ----------------------------
-- Table structure for goose_auth
-- ----------------------------
DROP TABLE IF EXISTS `goose_auth`;
CREATE TABLE `goose_auth` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `auth_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT 'Key',
  `auth_code` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT 'Secret',
  `email` varchar(50) NOT NULL DEFAULT '' COMMENT 'Email',
  `created_on` int unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `created_by` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '创建人',
  `modified_on` int unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '修改人',
  `deleted_on` int unsigned DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除（0为未删除，1为已删除）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='认证管理';

-- ----------------------------
-- Records of goose_auth
-- ----------------------------
INSERT INTO `goose_auth` VALUES ('1', 'Hanmur', 'Hanmur_Goose', '1466046208@qq.com', '0', 'Hanmur', '0', '', '0', '0');
INSERT INTO `goose_auth` VALUES ('22', 'Hanmuryeah', 'Goose_007', 'hanmur@foxmail.com', '1648328257', 'Hanmuryeah', '1648443342', 'Hanmuryeah', '0', '0');

-- ----------------------------
-- Table structure for goose_tags
-- ----------------------------
DROP TABLE IF EXISTS `goose_tags`;
CREATE TABLE `goose_tags` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT '' COMMENT '标签名称',
  `created_on` int unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `deleted_on` int unsigned DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除（0为未删除，1为已删除）',
  `state` tinyint unsigned DEFAULT '1' COMMENT '状态为0为禁用，1为已用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='标签管理';

-- ----------------------------
-- Records of goose_tags
-- ----------------------------
INSERT INTO `goose_tags` VALUES ('1', 'HHH', '1646733331', 'HanmurFirstTest', '1646734267', 'HanmurFirstChangeTest', '1646734325', '1', '1');
INSERT INTO `goose_tags` VALUES ('2', 'Hmm', '1646734505', 'HanmurFirstTestTest', '1646734595', 'Hanmur', '0', '0', '0');
INSERT INTO `goose_tags` VALUES ('3', 'KKNoBiss', '1646810898', 'WZH', '1646811420', 'Hanmur', '1646811483', '1', '1');
INSERT INTO `goose_tags` VALUES ('4', 'Fish40', '1646819588', '吴彦青', '1646819588', '', '0', '0', '1');
INSERT INTO `goose_tags` VALUES ('5', 'Fish41', '1646819594', '吴彦青', '1646819594', '', '0', '0', '1');
INSERT INTO `goose_tags` VALUES ('6', 'Fish42', '1646819597', '吴彦青', '1646819597', '', '0', '0', '1');
INSERT INTO `goose_tags` VALUES ('7', 'Fish43', '1646819599', '吴彦青', '1646819599', '', '0', '0', '1');
INSERT INTO `goose_tags` VALUES ('8', 'Fish44', '1646819602', '吴彦青', '1646819602', '', '0', '0', '1');
INSERT INTO `goose_tags` VALUES ('9', 'Fish45', '1646819604', '吴彦青', '1646819604', '', '0', '0', '1');
INSERT INTO `goose_tags` VALUES ('10', 'Fish46', '1646819611', '吴彦青', '1646819611', '', '0', '0', '1');
INSERT INTO `goose_tags` VALUES ('11', 'Fish47', '1646819613', '吴彦青', '1646819613', '', '0', '0', '1');
INSERT INTO `goose_tags` VALUES ('12', 'Fish48', '1646819616', '吴彦青', '1646819616', '', '0', '0', '1');
INSERT INTO `goose_tags` VALUES ('13', 'fish300', '1646825354', 'Hanmur', '1646825354', '', '0', '0', '1');
INSERT INTO `goose_tags` VALUES ('14', 'fish302', '1646825372', 'Hanmur', '1646825851', 'Hanmur', '1646825862', '1', '1');
INSERT INTO `goose_tags` VALUES ('15', 'hanmurfish', '1646828073', '33333', '1646828273', 'hanm', '1646828394', '1', '1');
INSERT INTO `goose_tags` VALUES ('16', 'WuzihouBiss', '1646828627', 'WuziHou', '1646828627', '', '0', '0', '1');
INSERT INTO `goose_tags` VALUES ('17', 'A', '1648324918', 'AF', '1648324918', '', '0', '0', '1');

-- ----------------------------
-- Table structure for goose_tag_article
-- ----------------------------
DROP TABLE IF EXISTS `goose_tag_article`;
CREATE TABLE `goose_tag_article` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `article_id` int NOT NULL COMMENT '文章ID',
  `tag_id` int NOT NULL DEFAULT '0' COMMENT '标签ID',
  `created_on` int unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `deleted_on` int unsigned DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除（0为未删除，1为已删除）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='文章标签关联';

-- ----------------------------
-- Records of goose_tag_article
-- ----------------------------
