create database gin-demo;

use gin-demo;

/*
*  用戶
*/
CREATE TABLE `gd_user` (
    `user_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `user_name` varchar(32) NOT NULL COMMENT '用户名称',
    `password` varchar(128) DEFAULT NULL COMMENT '密码',
    `nick_name` varchar(128) DEFAULT NULL COMMENT '昵称',
    `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '手机号',
    `avatar` varchar(255) DEFAULT NULL COMMENT '头像',
    `sex` tinyint(1) DEFAULT NULL COMMENT '性别: 1 男 2 女',
    `email` varchar(128) DEFAULT NULL COMMENT '邮箱',
    `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态：1启用 2停用',
    `remark` varchar(255) DEFAULT NULL COMMENT '备注',
    `modifier` bigint(20) NOT NULL DEFAULT '0' COMMENT '修改人id',
    `created_time` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
    `modified_time` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    `creater` bigint(20) NOT NULL DEFAULT '0' COMMENT '创建人id',
    `modifier_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '更新人',
    `creater_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '创建人',
    `deleted_at` datetime DEFAULT NULL COMMENT '删除',
    PRIMARY KEY (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户';

/*
*  角色 Role
*/
create table `gd_role`(
  `role_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT'角色ID',
  `role_name` varchar(20) NOT NULL COMMENT'角色名',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT'状态：1启用 0停用',
  `remark` varchar(255) DEFAULT NULL COMMENT'备注',
  `modifier` bigint(20) NOT NULL DEFAULT '0' COMMENT '修改人id',
  `created_time` datetime NOT NULL COMMENT '创建时间',
  `modified_time` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `creater` bigint(20) NOT NULL DEFAULT '0' COMMENT '创建人id',
  `modifier_name` varchar(20) NOT NULL COMMENT '更新人',
  `creater_name` varchar(20) NOT NULL COMMENT '创建人',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除',
  PRIMARY KEY (`role_id`) USING BTREE
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色';

/*
*  用戶-角色
*/
create table `gd_user_role`(
  `role_id` bigint(20) unsigned NOT NULL COMMENT'角色ID',
  `user_id` bigint(20) unsigned NOT NULL COMMENT'用户ID',
  PRIMARY KEY (`user_id`,`role_id`) USING BTREE
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户-角色';

/*
*  权限 permissions
*/
create table `gd_permissions`(
  `permissions_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT'权限ID',
  `handle` varchar(128) DEFAULT NULL COMMENT 'handle',
  `permissions_name` int(20) NOT NULL COMMENT'权限名',
  `path` varchar(64) NOT NULL COMMENT'地址',
  `action` varchar(16) DEFAULT NULL COMMENT '请求类型',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT'状态：1启用 0停用',
  `remark` varchar(255) DEFAULT NULL COMMENT'备注',
  `modifier` bigint(20) NOT NULL DEFAULT '0' COMMENT '修改人id',
  `created_time` datetime NOT NULL COMMENT '创建时间',
  `modified_time` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `creater` bigint(20) NOT NULL DEFAULT '0' COMMENT '创建人id',
  `modifier_name` varchar(20) NOT NULL COMMENT '更新人',
  `creater_name` varchar(20) NOT NULL COMMENT '创建人',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除',
  PRIMARY KEY (`permissions_id`) USING BTREE
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='权限';

/*
*  权限-角色
*/
create table `gd_permissions_role`(
  `role_id` bigint(20) unsigned NOT NULL COMMENT'角色ID',
  `permissions_id` bigint(20) unsigned NOT NULL COMMENT'权限ID',
  PRIMARY KEY (`permissions_id`,`role_id`) USING BTREE
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='权限-角色';

/*
*  菜单
*/
create table `gd_menu`(
  `menu_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT'菜单ID',
  `menu_name` varchar (32)  NOT NULL COMMENT'菜单名',
  `pid` bigint(20) unsigned NOT NULL default 0 COMMENT'父菜单id',
  `level` tinyint(1) NOT NULL DEFAULT 1 COMMENT'菜单级别',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT'状态：1启用 0停用',
  `remark` varchar(255) DEFAULT NULL COMMENT'备注',
  `modifier` bigint(20) NOT NULL DEFAULT '0' COMMENT '修改人id',
  `created_time` datetime NOT NULL COMMENT '创建时间',
  `modified_time` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `creater` bigint(20) NOT NULL DEFAULT '0' COMMENT '创建人id',
  `modifier_name` varchar(20) NOT NULL COMMENT '更新人',
  `creater_name` varchar(20) NOT NULL COMMENT '创建人',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除',
  PRIMARY KEY (`menu_id`) USING BTREE
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='菜单';

/*
*  菜单-权限
*/
create table `gd_menu_permissions`(
  `menu_id` bigint(20) unsigned NOT NULL COMMENT'菜单ID',
  `permissions_id` bigint(20) unsigned NOT NULL COMMENT'权限ID',
  PRIMARY KEY (`permissions_id`,`menu_id`) USING BTREE
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='菜单-权限';

/*
*  菜单-权限
*/
create table `gd_menu_role`(
   `menu_id` bigint(20) unsigned NOT NULL COMMENT'菜单ID',
   `role_id` bigint(20) unsigned NOT NULL COMMENT'角色ID',
    PRIMARY KEY (`role_id`,`menu_id`) USING BTREE
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色-权限';

/*
 * casbin 表
 */
 CREATE TABLE `gd_casbin_rule` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) DEFAULT NULL,
  `v0` varchar(100) DEFAULT NULL,
  `v1` varchar(100) DEFAULT NULL,
  `v2` varchar(100) DEFAULT NULL,
  `v3` varchar(100) DEFAULT NULL,
  `v4` varchar(100) DEFAULT NULL,
  `v5` varchar(100) DEFAULT NULL,
  `v6` varchar(25) DEFAULT NULL,
  `v7` varchar(25) DEFAULT NULL,
  PRIMARY KEY (`id`)USING BTREE,
  UNIQUE KEY `idx_sys_casbin_rule` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`,`v6`,`v7`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/**
  * 插入超级管理员
 */

INSERT INTO `gin-demo`.`gd_user` (`user_name`, `password`, `nick_name`, `created_time`, `modified_time`, `modifier_name`, `creater_name`) VALUES ('admin', 'e10adc3949ba59abbe56e057f20f883e', '超级管理员', '2023-09-23 14:06:31', '2023-09-23 14:06:39', 'admin', 'admin')