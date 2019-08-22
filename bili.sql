/*
 Navicat Premium Data Transfer

 Source Server         : 192.168.1.176
 Source Server Type    : MySQL
 Source Server Version : 50560
 Source Host           : localhost:3306
 Source Schema         : bili

 Target Server Type    : MySQL
 Target Server Version : 50560
 File Encoding         : 65001

 Date: 22/08/2019 17:58:42
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for bili_carousel
-- ----------------------------
DROP TABLE IF EXISTS `bili_carousel`;
CREATE TABLE `bili_carousel`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '轮播',
  `title` varchar(225) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '轮播标题',
  `photo` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '轮播图片',
  `jump_url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '轮播图片',
  `tab` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '类型标识',
  `tab_name` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '类型标识名称',
  `sort` int(8) NULL DEFAULT 0 COMMENT '排序',
  `opne` tinyint(2) NULL DEFAULT 0 COMMENT '是否开启:0否1是',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `sort`(`sort`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of bili_carousel
-- ----------------------------
INSERT INTO `bili_carousel` VALUES (1, '两位少女缔造的音乐奇迹', 'https://i0.hdslb.com/bfs/archive/dd0eb89cc79030ba2fbab2ba722562a416611913.jpg@880w_440h.jpg', 'https://www.bilibili.com/blackboard/topic/activity-ktmkwSLyt.html', 'tabTop', '头部推荐', 1, 1);
INSERT INTO `bili_carousel` VALUES (2, '数款泳装福利来袭', 'https://i0.hdslb.com/bfs/sycp/creative_img/201908/ee742a73fb27930886319bfc4399683d.jpg@880w_440h.jpg', 'https://game.bilibili.com/finalgear/summermemories/', 'tabTop', '头部推荐', 2, 1);
INSERT INTO `bili_carousel` VALUES (3, '音乐区【联合投稿功能】全量使用！', 'https://i0.hdslb.com/bfs/archive/a55b7cc86ae6d054547b1d43f49afac7d52e0d81.jpg@880w_440h.webp', 'https://www.bilibili.com/blackboard/activity-Rly74JX7B.html', 'tabTop', '头部推荐', 3, 1);
INSERT INTO `bili_carousel` VALUES (4, '百万粉丝的UP是如何制作一期美食视频的', 'https://i0.hdslb.com/bfs/archive/567a57f60d442889f0906e681d4f36cc7f63159c.jpg@880w_440h.webp', 'https://www.bilibili.com/video/av63980680', 'tabTop', '头部推荐', 4, 1);
INSERT INTO `bili_carousel` VALUES (5, '全新时代开启？LPL季后赛前瞻', 'https://i0.hdslb.com/bfs/archive/954b775b345cffd8800f2a76fffc8190b7fdaa87.png@880w_440h.webp', 'https://www.bilibili.com/read/cv3401362', 'tabTop', '头部推荐', 5, 1);
INSERT INTO `bili_carousel` VALUES (6, '寻找单机区玩家之星', 'https://i0.hdslb.com/bfs/live/54916278aaf5cb60f0985fbec4f0d280447863c3.png@260w_248h.webp', 'https://www.bilibili.com/blackboard/live/activity-gamestar01.html', 'recomTop', '为你推荐', 1, 1);
INSERT INTO `bili_carousel` VALUES (7, '寻梦异世镜，带你奇幻大冒险', 'https://i0.hdslb.com/bfs/live/2565e012cac7f17b5ad25d494891a7174e13892f.jpg@260w_248h.webp', 'https://www.bilibili.com/blackboard/live/activity-xmysj1qh.html', 'recomTop', '为你推荐', 2, 1);
INSERT INTO `bili_carousel` VALUES (8, 'DOTA2-TI9国际邀请赛淘汰赛正在火热进行', 'https://i0.hdslb.com/bfs/live/ff5dccf60502fec14ec323ace2b64d536e78ed8d.jpg@260w_248h.webp', 'https://live.bilibili.com/13', 'recomTop', '为你推荐', 3, 1);
INSERT INTO `bili_carousel` VALUES (9, '化为熊熊烈焰吧', 'https://i0.hdslb.com/bfs/bangumi/96bf2a60c88eda1fd082477d456660a340c5445b.png@520w_536h.webp', 'https://www.bilibili.com/bangumi/play/ss27959', 'specialTop', '特别推荐', 1, 1);
INSERT INTO `bili_carousel` VALUES (10, '甜甜的狗粮', 'https://i0.hdslb.com/bfs/bangumi/f8f6b7741cfe6c7d819495d325698bc2b9f64bcc.jpg@520w_536h.webp', 'https://www.bilibili.com/bangumi/play/ss28006', 'specialTop', '特别推荐', 2, 1);
INSERT INTO `bili_carousel` VALUES (11, '站在异端上的君主', 'https://i0.hdslb.com/bfs/bangumi/5edc7086e4635dfff63bfce1ccb0ca8faa6c870f.jpg@520w_536h.webp', 'https://www.bilibili.com/bangumi/play/ss26363', 'specialTop', '特别推荐', 3, 1);
INSERT INTO `bili_carousel` VALUES (12, '骨头社二十周年纪念作品', 'https://i0.hdslb.com/bfs/bangumi/ad5d016d809ce163645b74b078c7c16e2639bb46.jpg@520w_536h.webp', 'https://www.bilibili.com/bangumi/play/ss26875', 'specialTop', '特别推荐', 4, 1);

-- ----------------------------
-- Table structure for bili_carousel_nav
-- ----------------------------
DROP TABLE IF EXISTS `bili_carousel_nav`;
CREATE TABLE `bili_carousel_nav`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '轮播类型',
  `name` varchar(40) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '名称',
  `tab` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '标识',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of bili_carousel_nav
-- ----------------------------
INSERT INTO `bili_carousel_nav` VALUES (1, '头部推荐', 'tabTop');
INSERT INTO `bili_carousel_nav` VALUES (2, '直播排行', 'onlineTop');
INSERT INTO `bili_carousel_nav` VALUES (3, '关注主播', 'followTop');
INSERT INTO `bili_carousel_nav` VALUES (4, '为你推荐', 'recomTop');
INSERT INTO `bili_carousel_nav` VALUES (5, '特别推荐', 'specialTop');
INSERT INTO `bili_carousel_nav` VALUES (6, '国创漫画', 'nationalTo');

-- ----------------------------
-- Table structure for bili_head_nav
-- ----------------------------
DROP TABLE IF EXISTS `bili_head_nav`;
CREATE TABLE `bili_head_nav`  (
  `id` int(10) UNSIGNED NOT NULL COMMENT '头部菜单',
  `name` varchar(40) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '菜单名称',
  `action` varchar(40) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '菜单路径',
  `new_count` int(10) NULL DEFAULT 0 COMMENT '更新数量',
  `icon` varchar(225) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '菜单小图标',
  `nav_id` int(10) NULL DEFAULT 0 COMMENT '上级菜单',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `nav_id`(`nav_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of bili_head_nav
-- ----------------------------
INSERT INTO `bili_head_nav` VALUES (1, '首页', '/', 0, 'index.icon', 0);
INSERT INTO `bili_head_nav` VALUES (2, '动漫', 'dongm', 0, '', 0);
INSERT INTO `bili_head_nav` VALUES (3, '番剧', 'fanj', 0, '', 0);
INSERT INTO `bili_head_nav` VALUES (4, '国创', 'guoc', 0, '', 0);
INSERT INTO `bili_head_nav` VALUES (5, '音乐', 'yiny', 0, '', 0);
INSERT INTO `bili_head_nav` VALUES (6, '舞蹈', 'wud', 0, '', 0);
INSERT INTO `bili_head_nav` VALUES (7, '游戏', 'youx', 0, '', 0);
INSERT INTO `bili_head_nav` VALUES (8, '科技', 'kej', 0, '', 0);
INSERT INTO `bili_head_nav` VALUES (9, '数码', 'shum', 0, '', 0);
INSERT INTO `bili_head_nav` VALUES (10, '生活', 'shengh', 0, '', 0);
INSERT INTO `bili_head_nav` VALUES (11, '鬼畜', 'guic', 0, '', 0);
INSERT INTO `bili_head_nav` VALUES (12, '时尚', 'shis', 0, '', 0);
INSERT INTO `bili_head_nav` VALUES (13, '广告', 'guangg', 0, '', 0);
INSERT INTO `bili_head_nav` VALUES (14, '娱乐', 'yul', 0, '', 0);
INSERT INTO `bili_head_nav` VALUES (15, '影视', 'yings', 0, '', 0);
INSERT INTO `bili_head_nav` VALUES (16, '放映室', 'fangys', 0, '', 0);
INSERT INTO `bili_head_nav` VALUES (17, '专栏', 'zhuanl', 0, 'zhuanlan.icon', 0);
INSERT INTO `bili_head_nav` VALUES (18, '广场', 'guangc', 0, 'guangchang.icon', 0);
INSERT INTO `bili_head_nav` VALUES (19, '直播', 'zhib', 0, 'zhibo.icon', 0);
INSERT INTO `bili_head_nav` VALUES (20, '小黑屋', 'xiaohw', 0, 'xiaohewu.icon', 0);
INSERT INTO `bili_head_nav` VALUES (301, '连载动漫', 'lianzdm', 0, '', 3);
INSERT INTO `bili_head_nav` VALUES (302, '完结动漫', 'wanjdm', 0, '', 3);
INSERT INTO `bili_head_nav` VALUES (303, '咨询', 'zix', 0, '', 3);

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
