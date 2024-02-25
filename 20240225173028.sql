/*
MySQL Backup
Database: one
Backup Time: 2024-02-25 17:30:28
*/

SET FOREIGN_KEY_CHECKS=0;
DROP TABLE IF EXISTS `one`.`casbin_rule`;
DROP TABLE IF EXISTS `one`.`file`;
DROP TABLE IF EXISTS `one`.`group`;
DROP TABLE IF EXISTS `one`.`member`;
DROP TABLE IF EXISTS `one`.`mfile`;
DROP TABLE IF EXISTS `one`.`sys_auth_rule`;
DROP TABLE IF EXISTS `one`.`sys_config`;
DROP TABLE IF EXISTS `one`.`sys_dept`;
DROP TABLE IF EXISTS `one`.`sys_menu`;
DROP TABLE IF EXISTS `one`.`sys_post`;
DROP TABLE IF EXISTS `one`.`sys_role`;
DROP TABLE IF EXISTS `one`.`sys_user`;
CREATE TABLE `casbin_rule` (
  `ptype` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v0` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v1` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v2` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v3` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v4` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v5` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=COMPACT;
CREATE TABLE `file` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `name` varchar(45) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '文件名称',
  `src` varchar(500) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL COMMENT '本地文件存储路径',
  `url` varchar(500) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT 'URL地址，可能为空',
  `user_id` bigint unsigned NOT NULL COMMENT '操作用户',
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin ROW_FORMAT=DYNAMIC COMMENT='文件列表';
CREATE TABLE `group` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `type` int DEFAULT NULL COMMENT '1开发商2银行',
  `keyword` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `mark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `is_del` tinyint DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
CREATE TABLE `member` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT '',
  `realname` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT '',
  `nickname` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `idcard` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `group` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT '0',
  `bigclass` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT '0',
  `smallclass` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT '0',
  `mobile` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `web_auth` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT '0',
  `pwd` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT '',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `sex` tinyint DEFAULT '1' COMMENT '0woman1man',
  `verify_type` int DEFAULT '0',
  `verify_photo` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `verify_time` datetime DEFAULT NULL,
  `status` tinyint DEFAULT '1' COMMENT '2dongjie',
  `openid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
CREATE TABLE `mfile` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `name` varchar(45) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '文件名称',
  `src` varchar(500) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL COMMENT '本地文件存储路径',
  `url` varchar(500) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT 'URL地址，可能为空',
  `member_id` bigint unsigned NOT NULL COMMENT '操作用户',
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `user_id` (`member_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin ROW_FORMAT=DYNAMIC COMMENT='文件列表';
CREATE TABLE `sys_auth_rule` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `pid` int unsigned NOT NULL DEFAULT '0' COMMENT '父ID',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '规则名称',
  `jump` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '规则名称',
  `icon` varchar(300) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '图标',
  `small_auth` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '接口细化权限说明，用来写入casbin第二个权限',
  `condition` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '条件：nocheck就不用检测权限',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `type` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '类型1菜单 2接口',
  `sort` int NOT NULL DEFAULT '0' COMMENT '权重',
  `is_hide` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '显示状态',
  `created_at` datetime DEFAULT NULL COMMENT '创建日期',
  `updated_at` datetime DEFAULT NULL COMMENT '修改日期',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `pid` (`pid`) USING BTREE,
  KEY `weigh` (`sort`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=60 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='菜单节点表';
CREATE TABLE `sys_config` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `val` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `remark` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
CREATE TABLE `sys_dept` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '部门id',
  `pid` bigint DEFAULT '0' COMMENT '父部门id',
  `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '部门名称',
  `sort` int DEFAULT '0' COMMENT '显示顺序',
  `tel` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '联系电话',
  `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '邮箱',
  `status` tinyint unsigned DEFAULT '0' COMMENT '部门状态（0正常 1停用）',
  `created_by` bigint unsigned DEFAULT '0' COMMENT '创建人',
  `updated_by` bigint DEFAULT NULL COMMENT '修改人',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '修改时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=206 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=COMPACT COMMENT='部门表';
CREATE TABLE `sys_menu` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(20) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '菜单名称',
  `name` varchar(20) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `jump` varchar(100) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `url` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `icon` varchar(20) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `remark` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `pid` int DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
CREATE TABLE `sys_post` (
  `post_id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '岗位ID',
  `post_code` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '岗位编码',
  `post_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '岗位名称',
  `post_sort` int NOT NULL COMMENT '显示顺序',
  `status` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '状态（0正常 1停用）',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '备注',
  `created_by` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建人',
  `updated_by` bigint unsigned NOT NULL DEFAULT '0' COMMENT '修改人',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '修改时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`post_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=COMPACT COMMENT='岗位信息表';
CREATE TABLE `sys_role` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `status` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '状态;0:禁用;1:正常',
  `sort` int unsigned NOT NULL DEFAULT '0' COMMENT '排序',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '角色名称',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `status` (`status`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=COMPACT COMMENT='角色表';
CREATE TABLE `sys_user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `user_name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `mobile` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '中国手机不带国家代码，国际手机号格式为：国家代码-手机号',
  `user_nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户昵称',
  `birthday` int NOT NULL DEFAULT '0' COMMENT '生日',
  `user_password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '登录密码;cmf_password加密',
  `user_salt` char(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '加密盐',
  `user_status` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '用户状态;0:禁用,1:正常,2:未验证',
  `user_email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户登录邮箱',
  `sex` tinyint NOT NULL DEFAULT '0' COMMENT '性别;0:保密,1:男,2:女',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户头像',
  `dept_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '部门id',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `post_id` bigint unsigned NOT NULL DEFAULT '1',
  `open_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '',
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '联系地址',
  `describe` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT ' 描述信息',
  `last_login_ip` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '最后登录ip',
  `last_login_time` datetime DEFAULT NULL COMMENT '最后登录时间',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `user_login` (`user_name`,`deleted_at`) USING BTREE,
  UNIQUE KEY `mobile` (`mobile`,`deleted_at`) USING BTREE,
  KEY `user_nickname` (`user_nickname`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=46 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=COMPACT COMMENT='用户表';
BEGIN;
LOCK TABLES `one`.`casbin_rule` WRITE;
DELETE FROM `one`.`casbin_rule`;
INSERT INTO `one`.`casbin_rule` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`) VALUES ('p', '2', '2', 'All', '', '', ''),('p', '3', '2', 'All', '', '', ''),('p', '3', '3', 'All', '', '', ''),('p', '2', '3', 'All', '', '', ''),('p', '5', '4', 'All', '', '', ''),('p', '2', '10', 'All', '', '', ''),('g', '42', '1', '', '', '', ''),('g', '42', '2', '', '', '', ''),('g', '1', '1', '', '', '', ''),('g', '1', '2', '', '', '', ''),('g', '2', '3', '', '', '', ''),('g', '2', '2', '', '', '', ''),('g', '4', '2', '', '', '', ''),('g', '5', '2', '', '', '', ''),('g', '7', '2', '', '', '', ''),('g', '8', '2', '', '', '', ''),('g', '10', '2', '', '', '', ''),('g', '14', '2', '', '', '', ''),('g', '15', '2', '', '', '', ''),('g', '16', '2', '', '', '', ''),('p', '2', '1', 'All', '', '', ''),('g', '6', '2', '', '', '', ''),('g', '3', '2', '', '', '', ''),('p', '1', '27', 'All', '', '', ''),('p', '1', '1', 'All', '', '', ''),('p', '1', '10', 'All', '', '', ''),('p', '1', '12', 'All', '', '', ''),('p', '1', '13', 'All', '', '', ''),('p', '1', '14', 'All', '', '', ''),('p', '8', '1', 'All', '', '', ''),('p', '8', '2', 'All', '', '', ''),('p', '8', '4', 'All', '', '', ''),('p', '2', '22', 'All', '', '', ''),('p', '4', '1', 'All', '', '', ''),('p', '4', '2', 'All', '', '', ''),('g', 'u_5', '2', 'd', '', '', ''),('p', 'r_5', '1', 'd', '', '', ''),('p', 'r_5', '2', 'd', '', '', ''),('p', 'r_5', '3', 'd', '', '', ''),('p', 'r_5', '4', 'd', '', '', ''),('p', 'r_5', '11', 'd', '', '', ''),('p', 'r_5', '34', 'd', '', '', ''),('p', 'r_5', '10', 'd', '', '', ''),('p', 'r_5', '12', 'd', '', '', ''),('p', 'r_5', '13', 'd', '', '', ''),('p', 'r_5', '14', 'd', '', '', ''),('p', 'r_5', '15', 'd', '', '', ''),('p', 'r_5', '19', 'd', '', '', ''),('p', 'r_5', '20', 'd', '', '', ''),('p', 'r_5', '21', 'd', '', '', ''),('p', 'r_5', '22', 'd', '', '', ''),('p', 'r_5', '23', 'd', '', '', ''),('p', 'r_5', '24', 'd', '', '', ''),('p', 'r_5', '25', 'd', '', '', ''),('p', 'r_5', '26', 'd', '', '', ''),('p', 'r_5', '27', 'd', '', '', ''),('p', 'r_5', '28', 'd', '', '', ''),('p', 'r_5', '35', 'd', '', '', ''),('p', 'r_5', '36', 'd', '', '', ''),('p', 'r_2', '1', 'd', '', '', ''),('p', 'r_2', '2', 'd', '', '', ''),('p', 'r_2', '3', 'd', '', '', ''),('p', 'r_2', '4', 'admin', '', '', ''),('p', 'r_2', '34', 'd', '', '', ''),('g', 'u_44', 'r_3', '', '', '', ''),('g', 'u_44', 'r_4', '', '', '', ''),('g', 'u_45', 'r_4', '', '', '', ''),('g', 'u_31', 'r_1', '', '', '', ''),('g', 'u_34', 'r_4', '', '', '', ''),('g', 'u_42', 'r_2', '', '', '', ''),('p', 'r_1', '1', 'd', '', '', ''),('p', 'r_1', '2', 'd', '', '', ''),('p', 'r_1', '3', 'd', '', '', ''),('p', 'r_1', '4', 'd', '', '', ''),('p', 'r_1', '11', 'd', '', '', ''),('p', 'r_1', '34', 'd', '', '', ''),('p', 'r_1', '36', 'd', '', '', ''),('p', 'r_1', '10', 'd', '', '', ''),('p', 'r_1', '12', 'd', '', '', ''),('p', 'r_1', '13', 'd', '', '', ''),('p', 'r_1', '14', 'd', '', '', ''),('p', 'r_1', '35', 'd', '', '', ''),('p', 'r_1', '37', 'd', '', '', ''),('p', 'r_1', '38', 'd', '', '', ''),('p', 'r_1', '39', 'd', '', '', ''),('p', 'r_1', '40', 'd', '', '', ''),('p', 'r_1', '41', 'd', '', '', '');
UNLOCK TABLES;
COMMIT;
BEGIN;
LOCK TABLES `one`.`file` WRITE;
DELETE FROM `one`.`file`;
UNLOCK TABLES;
COMMIT;
BEGIN;
LOCK TABLES `one`.`group` WRITE;
DELETE FROM `one`.`group`;
UNLOCK TABLES;
COMMIT;
BEGIN;
LOCK TABLES `one`.`member` WRITE;
DELETE FROM `one`.`member`;
UNLOCK TABLES;
COMMIT;
BEGIN;
LOCK TABLES `one`.`mfile` WRITE;
DELETE FROM `one`.`mfile`;
UNLOCK TABLES;
COMMIT;
BEGIN;
LOCK TABLES `one`.`sys_auth_rule` WRITE;
DELETE FROM `one`.`sys_auth_rule`;
INSERT INTO `one`.`sys_auth_rule` (`id`,`pid`,`name`,`jump`,`title`,`icon`,`small_auth`,`condition`,`remark`,`type`,`sort`,`is_hide`,`created_at`,`updated_at`) VALUES (1, 0, 'sys', '', '权限管理', 'ele-Stamp', 'd', '', '', 1, 30, 0, '2022-03-24 15:03:37', '2022-12-03 12:51:50'),(2, 1, 'menu', '', '菜单管理', 'ele-Calendar', 'd', '', '', 1, 0, 0, '2022-03-24 17:24:13', '2022-12-03 12:52:20'),(3, 2, 'sys/menu/add', '', '添加菜单', '', 'd,edit', '', '', 2, 0, 0, '2022-03-29 16:48:43', '2023-06-03 20:11:33'),(4, 2, 'sys/menu/update', '', '修改菜单', '', 'd,admin,edit,read,thExt', '', '', 2, 0, 0, '2022-03-29 17:04:25', '2023-02-19 10:12:06'),(10, 1, 'role', '', '角色管理', 'iconfont icon-juxingkaobei', 'd', '', '', 1, 0, 0, '2022-03-29 18:15:03', '2022-12-23 19:23:44'),(11, 2, 'sys/menu/delete', '', '删除菜单', '', 'd', '', '', 2, 0, 0, '2022-04-06 14:49:10', '2022-12-03 15:05:10'),(12, 10, 'sys/role/add', '', '添加角色', '', 'd', '', '', 2, 0, 0, '2022-04-06 14:49:46', '2022-12-23 22:46:59'),(13, 10, 'sys/role/update', '', '修改角色', '', 'd', '', '', 2, 0, 0, '2022-04-06 14:50:08', '2022-12-23 23:06:29'),(14, 10, 'sys/role/delete', '', '删除角色', '', 'd', '', '', 2, 0, 0, '2022-04-06 14:50:22', '2022-12-23 22:46:39'),(15, 1, 'dept', '', '部门管理', 'iconfont icon-siweidaotu', 'd', '', '', 1, 0, 0, '2022-04-06 14:52:23', '2022-12-23 19:25:54'),(19, 15, 'sys/dept/add', '', '添加部门', '', 'd', '', '', 2, 0, 0, '2022-04-07 22:56:39', '2022-12-24 10:56:42'),(20, 15, 'sys/dept/update', '', '修改部门', '', 'd', '', '', 2, 0, 0, '2022-04-07 22:57:00', '2022-12-24 10:57:20'),(21, 15, 'sys/dept/delete', '', '删除部门', '', 'd', '', '', 2, 0, 0, '2022-04-07 22:57:30', '2022-12-24 10:57:43'),(22, 1, 'post', '', '岗位管理', 'iconfont icon-neiqianshujuchucun', 'd', '', '', 1, 0, 0, '2022-04-07 22:58:46', '2022-12-23 19:26:09'),(23, 22, 'sys/post/add', '', '添加岗位', '', 'd', '', '', 2, 0, 0, '2022-04-09 14:14:49', '2022-12-24 10:59:20'),(24, 22, 'sys/post/edit', '', '修改岗位', '', 'd', '', '', 2, 0, 0, '2022-04-09 14:15:25', '2022-12-24 10:59:59'),(25, 22, 'sys/post/delete', '', '删除岗位', '', 'd', '', '', 2, 0, 0, '2022-04-09 14:15:47', '2022-12-24 11:00:19'),(26, 1, 'user', '', '用户管理', 'ele-User', 'd', '', '', 1, 0, 0, '2022-04-09 14:19:10', '2022-12-23 19:26:42'),(27, 0, 'home', '', '主页', 'iconfont icon-shuxingtu', 'd', '', '', 1, 40, 0, '2022-04-14 16:28:51', '2022-12-23 20:36:13'),(28, 27, '', '/', '控制台', 'iconfont icon-crew_feature', 'd', '', '', 1, 0, 0, '2022-04-14 16:32:10', '2022-12-23 20:37:21'),(34, 2, 'sys/menu/list', '', '全节点列表', '', 'd', '', '', 2, 0, 0, '2022-12-03 16:07:12', '2022-12-23 20:49:00'),(35, 10, 'sys/role/del_one_rule', '', '删除角色单个权限', '', 'd', '', '', 2, 0, 0, '2022-12-06 21:28:27', '2022-12-23 20:20:50'),(36, 2, 'sys/menu/auth_list', '', '获取我的权限菜单', '', 'd', 'nocheck', '', 2, 0, 0, '2022-12-06 21:29:03', '2022-12-23 20:43:12'),(37, 10, 'sys/role/list', '', '角色列表', '', 'd', '', '', 2, 0, 0, '2022-12-23 23:09:09', '2022-12-23 23:09:09'),(38, 10, 'sys/role/nodes', '', '角色对应节点', '', 'd', '', '', 2, 0, 0, '2022-12-23 23:10:30', '2022-12-23 23:10:30'),(39, 10, 'sys/role/save_nodes', '', '更新角色节点', '', 'd', '', '', 2, 0, 0, '2022-12-23 23:11:11', '2022-12-23 23:11:11'),(40, 10, 'sys/role/update_old_rule', '', '批量更新角色权限', '', 'd', '', '', 2, 0, 0, '2022-12-23 23:12:09', '2022-12-23 23:12:09'),(41, 10, 'sys/role/update_single_rule', '', '更新单条角色权限', '', 'd', '', '', 2, 0, 0, '2022-12-23 23:12:48', '2022-12-23 23:12:48'),(42, 15, 'sys/dept/list', '', '部门列表', '', 'd', '', '', 2, 0, 0, '2022-12-24 10:58:43', '2022-12-24 10:58:43'),(43, 22, 'sys/post/list', '', '职位列表', '', 'd', '', '', 2, 0, 0, '2022-12-24 11:01:02', '2022-12-24 11:01:02'),(44, 26, 'sys/user/add', '', '添加用户', '', 'd', '', '', 2, 0, 0, '2022-12-24 11:01:41', '2022-12-24 11:01:41'),(45, 26, 'sys/user/update', '', '更新用户', '', 'd', '', '', 2, 0, 0, '2022-12-24 11:02:11', '2022-12-24 11:02:11'),(46, 26, 'sys/user/delete', '', '删除用户', '', 'd', '', '', 2, 0, 0, '2022-12-24 11:02:36', '2022-12-24 11:02:36'),(47, 26, 'sys/user/list', '', '用户列表', '', 'd', '', '', 2, 0, 0, '2022-12-24 11:03:15', '2022-12-24 11:03:15'),(48, 26, 'sys/user/get_dept_post_role', '', '获取部门职位角色', '', 'd', '', '', 2, 0, 0, '2022-12-24 11:12:40', '2022-12-24 11:12:40'),(49, 26, 'sys/user/get_role_ids', '', '获取用户角色ids', '', 'd', '', '', 2, 0, 0, '2022-12-24 11:13:15', '2022-12-24 11:13:15'),(50, 0, '/test/', '', '测试', '', 'd', '', '', 1, 0, 0, '2022-12-26 20:23:07', '2022-12-26 20:23:07'),(51, 0, 'set', '', '系统设置', '', 'd', '', '', 1, 0, 0, '2023-01-08 14:49:14', '2023-01-08 14:49:14'),(52, 51, 'conf', '', '参数设置', '', 'd', '', '', 1, 0, 0, '2023-01-08 14:49:53', '2023-01-08 14:49:53'),(53, 52, 'sys/conf/get_val', '', '获取参数值', '', 'd', 'nocheck', '', 2, 0, 0, '2023-01-24 19:15:50', '2023-01-24 19:15:50'),(54, 52, 'sys/conf/add', '', '添加参数', '', 'd', '', '', 2, 0, 0, '2023-01-24 19:58:26', '2023-01-24 19:58:26'),(55, 52, 'sys/conf/delete', '', '删除参数', '', 'd', '', '', 2, 0, 0, '2023-01-24 19:59:27', '2023-01-24 19:59:27'),(56, 52, 'sys/conf/update', '', '更新参数', '', 'd', '', '', 2, 0, 0, '2023-01-24 20:00:09', '2023-01-24 20:00:09'),(57, 52, 'sys/conf/list', '', '参数列表', '', 'd', '', '', 2, 0, 0, '2023-01-24 20:00:52', '2023-01-24 20:00:52');
UNLOCK TABLES;
COMMIT;
BEGIN;
LOCK TABLES `one`.`sys_config` WRITE;
DELETE FROM `one`.`sys_config`;
INSERT INTO `one`.`sys_config` (`id`,`name`,`val`,`remark`,`created_at`,`updated_at`,`deleted_at`) VALUES (1, 'appversion', '1.045', '111', '2023-01-08 20:31:31', '2023-02-12 18:43:23', NULL),(2, 'wgtpath', 'http://124.238.116.123:8900/single/wgt/one1.wgt', '', '2023-01-08 20:50:32', '2023-01-24 17:55:13', NULL),(3, 'test', 'aa', 'dasd', '2023-01-08 21:24:31', '2023-01-08 21:24:31', '2023-01-08 21:33:09'),(4, 'sprint', 'ok', 'test!!', '2023-01-09 23:28:39', '2023-01-09 23:29:05', '2023-01-09 23:30:06');
UNLOCK TABLES;
COMMIT;
BEGIN;
LOCK TABLES `one`.`sys_dept` WRITE;
DELETE FROM `one`.`sys_dept`;
INSERT INTO `one`.`sys_dept` (`id`,`pid`,`name`,`sort`,`tel`,`email`,`status`,`created_by`,`updated_by`,`created_at`,`updated_at`,`deleted_at`) VALUES (100, 0, '小T的公司', 0, '15888888888', 'ry@qq.com', 1, 0, 1, '2021-07-13 15:56:52', '2024-02-25 16:29:49', NULL),(101, 100, '棉被总公司', 1, '15888888888', 'ry@qq.com', 1, 0, 1, '2021-07-13 15:56:52', '2024-02-25 16:30:18', NULL),(102, 100, '北京分公司', 2, '15888888888', 'ry@qq.com', 1, 0, 1, '2021-07-13 15:56:52', '2024-02-25 16:30:50', NULL),(103, 101, '研发部门', 1, '15888888888', 'rey@qq.com', 1, 0, 31, '2021-07-13 15:56:52', '2023-01-08 20:27:44', NULL),(104, 101, '市场部门', 2, '15888888881', 'ry@qq.com', 1, 0, 31, '2021-07-13 15:56:52', '2022-12-11 17:27:15', NULL),(105, 101, '测试部门', 3, '15888888888', 'ry@qq.com', 1, 0, NULL, '2021-07-13 15:56:52', '2021-07-13 15:56:52', NULL),(106, 101, '财务部门', 4, '15888888888', 'ry@qq.com', 1, 0, NULL, '2021-07-13 15:56:52', '2021-07-13 15:56:52', NULL),(107, 101, '运维部门', 5, '15888888881', 'ry@qq.com', 1, 0, 31, '2021-07-13 15:56:52', '2022-12-26 20:21:31', NULL),(108, 102, '市场部门', 1, '15888888888', 'ry@qq.com', 1, 0, NULL, '2021-07-13 15:56:52', '2021-07-13 15:56:52', NULL),(109, 102, '财务部门', 2, '15888888888', 'ry@qq.com', 1, 0, NULL, '2021-07-13 15:56:52', '2021-07-13 15:56:52', NULL);
UNLOCK TABLES;
COMMIT;
BEGIN;
LOCK TABLES `one`.`sys_menu` WRITE;
DELETE FROM `one`.`sys_menu`;
INSERT INTO `one`.`sys_menu` (`id`,`title`,`name`,`jump`,`url`,`icon`,`remark`,`pid`,`created_at`,`updated_at`,`deleted_at`) VALUES (1, '权限管理', 'system', NULL, NULL, NULL, '1', 0, '2022-11-20 14:08:20', NULL, NULL),(2, '角色管理', 'role', '', '', '', '', 1, '2022-11-20 15:15:13', '2022-11-20 15:15:13', NULL),(3, '菜单管理', 'menu', '', '', '', 'dddd', 1, '2022-11-20 15:38:15', '2022-11-20 15:38:15', NULL),(4, '接口管理', 'api', '', '', '', '', 1, '2022-11-20 15:41:23', '2022-11-20 15:41:23', NULL),(5, '用户管理', 'uer', '', '', '', 'ok', 1, '2022-11-20 15:42:12', '2022-11-20 15:42:12', NULL),(6, '学习管理', 'study', '', '', '', '重点要做的1', 0, '2022-11-20 15:43:57', '2022-11-20 17:32:41', NULL);
UNLOCK TABLES;
COMMIT;
BEGIN;
LOCK TABLES `one`.`sys_post` WRITE;
DELETE FROM `one`.`sys_post`;
INSERT INTO `one`.`sys_post` (`post_id`,`post_code`,`post_name`,`post_sort`,`status`,`remark`,`created_by`,`updated_by`,`created_at`,`updated_at`,`deleted_at`) VALUES (1, 'ceo', '董事长', 1, 1, '', 0, 0, '2021-07-11 11:32:58', NULL, NULL),(2, 'seo', '项目经理', 2, 1, '', 0, 31, '2021-07-12 11:01:26', '2022-12-15 12:37:49', NULL),(3, 'hr', '人力资源', 3, 1, '', 0, 0, '2021-07-12 11:01:30', NULL, NULL),(4, 'user', '普通员工', 4, 0, '普通员工', 0, 31, '2021-07-12 11:01:33', '2022-04-08 15:32:23', NULL),(5, 'it', 'IT部', 5, 1, '信息部', 31, 31, '2021-07-12 11:09:42', '2022-04-09 12:59:12', NULL);
UNLOCK TABLES;
COMMIT;
BEGIN;
LOCK TABLES `one`.`sys_role` WRITE;
DELETE FROM `one`.`sys_role`;
INSERT INTO `one`.`sys_role` (`id`,`status`,`sort`,`name`,`remark`,`created_at`,`updated_at`) VALUES (1, 1, 0, '超级管理员', '好的ok', '2022-04-01 11:38:39', '2022-12-04 14:11:01'),(2, 1, 0, '普通管理员', '备注', '2022-04-01 11:38:39', '2022-04-01 11:38:39'),(3, 1, 0, '站点管理员', '站点管理人员', '2022-04-01 11:38:39', '2022-04-01 11:38:39'),(4, 1, 0, '初级管理员', '初级管理员', '2022-04-01 11:38:39', '2022-10-24 22:03:27'),(5, 1, 0, '高级管理员', '高级管理员', '2022-04-01 11:38:39', '2022-04-01 11:38:39'),(8, 0, 0, '区级管理员', '', '2022-04-01 11:38:39', '2022-12-04 14:08:32');
UNLOCK TABLES;
COMMIT;
BEGIN;
LOCK TABLES `one`.`sys_user` WRITE;
DELETE FROM `one`.`sys_user`;
INSERT INTO `one`.`sys_user` (`id`,`user_name`,`mobile`,`user_nickname`,`birthday`,`user_password`,`user_salt`,`user_status`,`user_email`,`sex`,`avatar`,`dept_id`,`remark`,`post_id`,`open_id`,`address`,`describe`,`last_login_ip`,`last_login_time`,`created_at`,`updated_at`,`deleted_at`) VALUES (1, 'admin', '13578342363', '超级管理员', 0, 'c567ae329f9929b518759d3bea13f492', 'f9aZTAa8yz', 1, 'yxh669@qq.com', 1, 'https://yxh-1301841944.cos.ap-chongqing.myqcloud.com/gfast/2021-07-19/ccwpeuqz1i2s769hua.jpeg', 101, '', 1, '', 'asdasfdsaf大发放打发士大夫发按时', '描述信息', '::1', '2022-04-19 16:38:37', '2021-06-22 17:58:00', '2022-04-19 16:38:37', NULL),(2, 'fanshurui', '13699885599', '范淑瑞', 0, '997f4f3ad786c26fb3dc5986ffb2c01d', 'dlqVVBTADg', 1, 'yxh@qq.com', 1, 'pub_upload/2020-11-02/c6sntzg7r96c7p9gqf.jpeg', 102, '备注', 1, '', '', '', '[::1]', '2022-02-14 18:10:40', '2021-06-22 17:58:00', '2023-02-12 11:00:38', NULL),(3, 'terry', '16399669855', 'terry', 0, 'b0b4392240f4d6e5f356198602e8c29e', 'dlqVVBTADg', 1, 'zs@qq.com', 0, 'https://yxh-1301841944.cos.ap-chongqing.myqcloud.com/gfast/2021-08-02/cd8nif79egjg9kbkgk.jpeg', 101, '', 1, '', '', '', '127.0.0.1', '2022-03-18 15:22:13', '2021-06-22 17:58:00', '2023-02-12 11:00:31', NULL),(4, 'huahua', '13758596696', '花花', 0, '28ff662072c627792c394a86e530bea1', 'dlqVVBTADg', 1, 'qlgl@qq.com', 0, '', 102, '', 1, '', '', '', '127.0.0.1', NULL, '2021-06-22 17:58:00', '2023-02-12 11:00:21', NULL),(5, 'test', '13845696696', '测试2', 0, 'ce669ab10db08c0a1bd202169a9e0ab2', 'dlqVVBTADg', 1, '123@qq.com', 0, '', 101, '', 0, '', '', '', '::1', '2022-03-30 10:50:39', '2021-06-22 17:58:00', '2023-02-12 11:00:15', NULL);
UNLOCK TABLES;
COMMIT;
