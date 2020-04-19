-- 文件表 --
drop table `tbl_file` if exists;
create table `tbl_file`(
    `id` int(11) primary key not null auto_increment,
    `file_sha1` char(40) not null default '' comment '文件hash',
    `file_name` varchar(255) not null default '' comment '文件名',
    `file_size` bigint(20) default '0' comment '文件大小',
  `file_addr` varchar(1024) not null  default '' comment '文件存储位置',
  `create_at` datetime default NOW() comment '创建日期',
  `update_at` datetime default NOW() on update current_timestamp() comment '更新日期',
  `status` int(11) not null default '0' comment '状态(可用/禁用/已删除等状态)',
  `ext1` int(11) default '0' comment '备用字段1',
  `ext2` text comment '备用字段2',
  unique key `idx_file_hash` (`file_sha1`),
  key `idx_status` (`status`)
)engine = InnoDB default charset =utf8;


-- 用户表 --
drop table `tbl_user` if exists;
create table `tbl_user`(
  `id` int(11) not null auto_increment,
  `user_name` varchar(64) not null default '' comment '用户名',
  `user_pwd` varchar(256) not null default '' comment '用户encoded密码',
  `email` varchar(64) default '' comment '邮箱',
  `phone` varchar(128) default '' comment '手机号',
  `email_validated` tinyint(1) default 0 comment '邮箱是否已验证',
  `phone_validated` tinyint(1) default 0 comment '手机号是否已验证',
  `signup_at` datetime default current_timestamp comment '注册日期',
  `last_active` datetime default current_timestamp on update current_timestamp comment '最后活跃时间戳',
  `profile` text comment '用户属性',
  `status` int(11) not null default '0' comment '账户状态(启用/禁用/锁定/标记删除等)',
  primary key (`id`),
  unique key `idx_phone` (`phone`),
  key `idx_status` (`status`)
)engine=InnoDB auto_increment=5 default charset=utf8mb4;
