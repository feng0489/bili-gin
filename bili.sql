/*
 Navicat Premium Data Transfer

 Source Server         : 192.168.1.177
 Source Server Type    : MySQL
 Source Server Version : 50560
 Source Host           : localhost:3306
 Source Schema         : bili

 Target Server Type    : MySQL
 Target Server Version : 50560
 File Encoding         : 65001

 Date: 13/08/2019 15:51:45
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for bili_head_nav
-- ----------------------------
DROP TABLE IF EXISTS `bili_head_nav`;
CREATE TABLE `bili_head_nav`  (
  `id` int(10) UNSIGNED NOT NULL COMMENT '头部菜单',
  `name` varchar(40) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '菜单名称',
  `action` varchar(40) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '菜单路径',
  `icon` varchar(225) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '菜单小图标',
  `nav_id` int(10) NULL DEFAULT 0 COMMENT '上级菜单',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `nav_id`(`nav_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of bili_head_nav
-- ----------------------------
INSERT INTO `bili_head_nav` VALUES (1, '首页', '/', 'index.icon', 0);
INSERT INTO `bili_head_nav` VALUES (2, '动漫', 'dongm', '', 0);
INSERT INTO `bili_head_nav` VALUES (3, '番剧', 'fanj', '', 0);
INSERT INTO `bili_head_nav` VALUES (4, '国创', 'guoc', '', 0);
INSERT INTO `bili_head_nav` VALUES (5, '音乐', 'yiny', '', 0);
INSERT INTO `bili_head_nav` VALUES (6, '舞蹈', 'wud', '', 0);
INSERT INTO `bili_head_nav` VALUES (7, '游戏', 'youx', '', 0);
INSERT INTO `bili_head_nav` VALUES (8, '科技', 'kej', '', 0);
INSERT INTO `bili_head_nav` VALUES (9, '数码', 'shum', '', 0);
INSERT INTO `bili_head_nav` VALUES (10, '生活', 'shengh', '', 0);
INSERT INTO `bili_head_nav` VALUES (11, '鬼畜', 'guic', '', 0);
INSERT INTO `bili_head_nav` VALUES (12, '时尚', 'shis', '', 0);
INSERT INTO `bili_head_nav` VALUES (13, '广告', 'guangg', '', 0);
INSERT INTO `bili_head_nav` VALUES (14, '娱乐', 'yul', '', 0);
INSERT INTO `bili_head_nav` VALUES (15, '影视', 'yings', '', 0);
INSERT INTO `bili_head_nav` VALUES (16, '放映室', 'fangys', '', 0);
INSERT INTO `bili_head_nav` VALUES (17, '专栏', 'zhuanl', 'zhuanlan.icon', 0);
INSERT INTO `bili_head_nav` VALUES (18, '广场', 'guangc', 'guangchang.icon', 0);
INSERT INTO `bili_head_nav` VALUES (19, '直播', 'zhib', 'zhibo.icon', 0);
INSERT INTO `bili_head_nav` VALUES (20, '小黑屋', 'xiaohw', 'xiaohewu.icon', 0);
INSERT INTO `bili_head_nav` VALUES (301, '连载动漫', 'lianzdm', '', 3);
INSERT INTO `bili_head_nav` VALUES (302, '完结动漫', 'wanjdm', '', 3);
INSERT INTO `bili_head_nav` VALUES (303, '咨询', 'zix', '', 3);

-- ----------------------------
-- Table structure for bili_top_nav
-- ----------------------------
DROP TABLE IF EXISTS `bili_top_nav`;
CREATE TABLE `bili_top_nav`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '顶部菜单',
  `name` varchar(40) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '菜单名称',
  `action` varchar(40) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '菜单路径',
  `icon` varchar(225) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '菜单小图标',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `name`(`name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of bili_top_nav
-- ----------------------------
INSERT INTO `bili_top_nav` VALUES (1, '主站', '/', 'home.icon');
INSERT INTO `bili_top_nav` VALUES (2, '音频', 'video', '');
INSERT INTO `bili_top_nav` VALUES (3, '游戏中心', 'game', '');
INSERT INTO `bili_top_nav` VALUES (4, '直播', 'direct', '');
INSERT INTO `bili_top_nav` VALUES (5, '漫画', 'cartoon', '');
INSERT INTO `bili_top_nav` VALUES (6, 'app下载', 'app/download', '');

SET FOREIGN_KEY_CHECKS = 1;
