
INSERT INTO `user` VALUES ('1', 'admin', '123', '1', '0', now(), null, null, null, now(), null)

go

INSERT INTO `menu` 
VALUES 
('1', '权限管理', '0', '', '0', '1', '0',now(), 0, now(), 0), 
('2', '用户管理', '1', 'auth/user', '0', '1', '0', now(), 0, now(), 0), 
('3', '角色管理', '1', 'auth/role', '0', '1', '0', now(), 0, now(), 0), 
('4', '资源管理', '1', 'auth/resource', '0', '1', '0', now(), 0, now(), 0), 
('5', '配置文件', '0', '', '0', '1', '0', now(), 0, now(), 0), 
('6', '开发配置文件', '5', 'xcfg/xcfg?env=dev', '0', '1', '0', now(), 0, now(), 0), 
('7', '测试配置文件', '5', 'xcfg/xcfg?env=testing', '0', '1', '0', now(), 0, now(), 0), 
('8', '实验配置文件', '5', 'xcfg/xcfg?env=labs', '0', '1', '0', now(), 0, now(), 0), 
('9', '线上配置文件', '5', 'xcfg/xcfg?env=prod', '0', '1', '0', now(), 0, now(), 0),
('10', '用户管理-列表', '2', '/user/list', '0', '1', '1', now(), 0, now(), 0),
('11', '用户管理-停用', '2', '/user/disable', '0', '1', '1', now(), 0, now(), 0),
('12', '用户管理-启用', '2', '/user/enable', '0', '1', '1', now(), 0, now(), 0),
('13', '用户管理-添加', '2', '/user/add', '0', '1', '1', now(), 0, now(), 0),
('14', '用户管理-是否存在', '2', '/user/exists', '0', '1', '1', now(), 0, now(), 0),
('15', '用户管理-获取', '2', '/user/get', '0', '1', '1', now(), 0, now(), 0),
('16', '用户管理-角色', '2', '/user/roles', '0', '1', '1', now(), 0, now(), 0),
('17', '用户管理-更新', '2', '/user/update', '0', '1', '1', now(), 0, now(), 0),
('18', '角色管理-列表', '3', '/role/list', '0', '1', '1', now(), 0, now(), 0),
('19', '角色管理-添加', '3', '/role/add', '0', '1', '1', now(), 0, now(), 0),
('20', '角色管理-是否存在', '3', '/role/exists', '0', '1', '1', now(), 0, now(), 0),
('21', '角色管理-获取', '3', '/role/get', '0', '1', '1', now(), 0, now(), 0),
('22', '角色管理-更新', '3', '/role/update', '0', '1', '1', now(), 0, now(), 0),
('23', '角色管理-资源', '3', '/role/resources', '0', '1', '1', now(), 0, now(), 0),
('24', '资源管理-列表', '4', '/menu/list', '0', '1', '1', now(), 0, now(), 0),
('25', '配置文件-列表dev', '5', '/cfg/list?env=dev', '0', '1', '1', now(), 0, now(), 0),
('26', '配置文件-添加dev', '5', '/cfg/add?env=dev', '0', '1', '1', now(), 0, now(), 0),
('27', '配置文件-是否存在dev', '5', '/cfg/exists?env=dev', '0', '1', '1', now(), 0, now(), 0),
('28', '配置文件-获取dev', '5', '/cfg/get?env=dev', '0', '1', '1', now(), 0, now(), 0),
('29', '配置文件-更新dev', '5', '/cfg/update?env=dev', '0', '1', '1', now(), 0, now(), 0),
('30', '配置文件-列表testing', '5', '/cfg/list?env=testing', '0', '1', '1', now(), 0, now(), 0),
('31', '配置文件-添加testing', '5', '/cfg/add?env=testing', '0', '1', '1', now(), 0, now(), 0),
('32', '配置文件-是否存在testing', '5', '/cfg/exists?env=testing', '0', '1', '1', now(), 0, now(), 0),
('33', '配置文件-获取testing', '5', '/cfg/get?env=testing', '0', '1', '1', now(), 0, now(), 0),
('34', '配置文件-更新testing', '5', '/cfg/update?env=testing', '0', '1', '1', now(), 0, now(), 0),
('35', '配置文件-列表labs', '5', '/cfg/list?env=labs', '0', '1', '1', now(), 0, now(), 0),
('36', '配置文件-添加labs', '5', '/cfg/add?env=labs', '0', '1', '1', now(), 0, now(), 0),
('37', '配置文件-是否存在labs', '5', '/cfg/exists?env=labs', '0', '1', '1', now(), 0, now(), 0),
('38', '配置文件-获取labs', '5', '/cfg/get?env=labs', '0', '1', '1', now(), 0, now(), 0),
('39', '配置文件-更新labs', '5', '/cfg/update?env=labs', '0', '1', '1', now(), 0, now(), 0),
('40', '配置文件-列表prod', '5', '/cfg/list?env=prod', '0', '1', '1', now(), 0, now(), 0),
('41', '配置文件-添加prod', '5', '/cfg/add?env=prod', '0', '1', '1', now(), 0, now(), 0),
('42', '配置文件-是否存在prod', '5', '/cfg/exists?env=prod', '0', '1', '1', now(), 0, now(), 0),
('43', '配置文件-获取prod', '5', '/cfg/get?env=prod', '0', '1', '1', now(), 0, now(), 0),
('44', '配置文件-更新prod', '5', '/cfg/update?env=prod', '0', '1', '1', now(), 0, now(), 0),
('45', '配置文件-状态', '5', '/ConfigVersionHandler.ashx', '0', '1', '2', now(), 0, now(), 0),
('46', '配置文件-x获取', '5', '/xcfg/get', '0', '1', '2', now(), 0, now(), 0),
('47', '权限管理-登录', '1', '/auth/login', '0', '1', '2', now(), 0, now(), 0),
('48', '权限管理-检查密码', '1', '/auth/checkpwd', '0', '1', '2', now(), 0, now(), 0),
('49', '权限管理-修改密码', '1', '/auth/mod', '0', '1', '2', now(), 0, now(), 0),
('50', '权限管理-获取菜单', '1', '/auth/menus', '0', '1', '2', now(), 0, now(), 0);
go

INSERT INTO `role` VALUES 
('1', 'dev', '1', '0', now(), '0', now(), '0'), 
('2', 'testing', '1', '0', now(), '0', now(), '0'), 
('3', 'labs', '1', '0', now(), '0', now(), '0'), 
('4', 'prod', '1', '0', now(), '0', now(), '0')
