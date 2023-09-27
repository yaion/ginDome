/*
 Navicat Premium Data Transfer

 Source Server         : phpstudy
 Source Server Type    : MySQL
 Source Server Version : 80012 (8.0.12)
 Source Host           : localhost:3306
 Source Schema         : gin-demo

 Target Server Type    : MySQL
 Target Server Version : 80012 (8.0.12)
 File Encoding         : 65001

 Date: 28/09/2023 02:42:59
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for gd_menu
-- ----------------------------
DROP TABLE IF EXISTS `gd_menu`;
CREATE TABLE `gd_menu`  (
  `menu_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '菜单ID',
  `menu_name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '菜单名',
  `pid` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '父菜单id',
  `level` tinyint(1) NOT NULL DEFAULT 1 COMMENT '菜单级别',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态：1启用 0停用',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '备注',
  `modifier` bigint(20) NOT NULL DEFAULT 0 COMMENT '修改人id',
  `created_time` datetime NOT NULL COMMENT '创建时间',
  `modified_time` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `creater` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建人id',
  `modifier_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '更新人',
  `creater_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '创建人',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除',
  PRIMARY KEY (`menu_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '菜单' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of gd_menu
-- ----------------------------

-- ----------------------------
-- Table structure for gd_menu_permissions
-- ----------------------------
DROP TABLE IF EXISTS `gd_menu_permissions`;
CREATE TABLE `gd_menu_permissions`  (
  `menu_id` bigint(20) UNSIGNED NOT NULL COMMENT '菜单ID',
  `permissions_id` bigint(20) UNSIGNED NOT NULL COMMENT '权限ID',
  PRIMARY KEY (`permissions_id`, `menu_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '菜单-权限' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of gd_menu_permissions
-- ----------------------------

-- ----------------------------
-- Table structure for gd_menu_role
-- ----------------------------
DROP TABLE IF EXISTS `gd_menu_role`;
CREATE TABLE `gd_menu_role`  (
  `menu_id` bigint(20) UNSIGNED NOT NULL COMMENT '菜单ID',
  `role_id` bigint(20) UNSIGNED NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`role_id`, `menu_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '角色-权限' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of gd_menu_role
-- ----------------------------

-- ----------------------------
-- Table structure for gd_permissions
-- ----------------------------
DROP TABLE IF EXISTS `gd_permissions`;
CREATE TABLE `gd_permissions`  (
  `permissions_id` int(20) NOT NULL AUTO_INCREMENT COMMENT '权限ID',
  `handle` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT 'handle',
  `permissions_name` int(20) NOT NULL COMMENT '权限名',
  `path` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '地址',
  `action` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '请求类型',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态：1启用 0停用',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '备注',
  `modifier` bigint(20) NOT NULL DEFAULT 0 COMMENT '修改人id',
  `created_time` datetime NOT NULL COMMENT '创建时间',
  `modified_time` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `creater` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建人id',
  `modifier_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '更新人',
  `creater_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '创建人',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除',
  PRIMARY KEY (`permissions_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '权限' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of gd_permissions
-- ----------------------------

-- ----------------------------
-- Table structure for gd_permissions_role
-- ----------------------------
DROP TABLE IF EXISTS `gd_permissions_role`;
CREATE TABLE `gd_permissions_role`  (
  `role_id` int(20) NOT NULL COMMENT '角色ID',
  `permissions_id` int(20) NOT NULL COMMENT '权限ID',
  PRIMARY KEY (`permissions_id`, `role_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '权限-角色' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of gd_permissions_role
-- ----------------------------

-- ----------------------------
-- Table structure for gd_role
-- ----------------------------
DROP TABLE IF EXISTS `gd_role`;
CREATE TABLE `gd_role`  (
  `role_id` int(20) NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `role_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '角色名',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态：1启用 0停用',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '备注',
  `modifier` bigint(20) NOT NULL DEFAULT 0 COMMENT '修改人id',
  `created_time` datetime NOT NULL COMMENT '创建时间',
  `modified_time` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `creater` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建人id',
  `modifier_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '更新人',
  `creater_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '创建人',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除',
  PRIMARY KEY (`role_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '角色' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of gd_role
-- ----------------------------

-- ----------------------------
-- Table structure for gd_user
-- ----------------------------
DROP TABLE IF EXISTS `gd_user`;
CREATE TABLE `gd_user`  (
  `user_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `user_name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户名称',
  `password` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '密码',
  `nick_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '昵称',
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '手机号',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '头像',
  `sex` tinyint(1) NULL DEFAULT NULL COMMENT '性别',
  `email` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '邮箱',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态：1启用 0停用',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '备注',
  `modifier` bigint(20) NOT NULL DEFAULT 0 COMMENT '修改人id',
  `created_time` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `modified_time` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `creater` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建人id',
  `modifier_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '更新人',
  `creater_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '创建人',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除',
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of gd_user
-- ----------------------------
INSERT INTO `gd_user` VALUES (1, 'test', 'e10adc3949ba59abbe56e057f20f883e', 'qe', '123123', '1', 1, 'ad', 1, NULL, 0, '2023-09-24 23:12:51', '2023-09-24 23:12:51', 0, 'qe', 'qe', '2023-09-24 23:12:52');
INSERT INTO `gd_user` VALUES (2, 'admin', 'e10adc3949ba59abbe56e057f20f883e', '超级管理员', NULL, NULL, NULL, NULL, 1, NULL, 0, '2023-09-23 14:06:31', '2023-09-23 14:06:39', 0, 'admin', 'admin', NULL);
INSERT INTO `gd_user` VALUES (3, 'john_doe', 'e10adc3949ba59abbe56e057f20f883e', 'john', '17251560913', '/file/img/sajdg.jpg', 2, '12356456123@qq.com', 0, 'tq1a12', 0, '2023-09-24 22:56:23', '2023-09-24 22:56:23', 0, '', '', NULL);

-- ----------------------------
-- Table structure for gd_user_role
-- ----------------------------
DROP TABLE IF EXISTS `gd_user_role`;
CREATE TABLE `gd_user_role`  (
  `role_id` int(20) NOT NULL COMMENT '角色ID',
  `user_id` int(20) NOT NULL COMMENT '用户ID',
  PRIMARY KEY (`user_id`, `role_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户-角色' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of gd_user_role
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
