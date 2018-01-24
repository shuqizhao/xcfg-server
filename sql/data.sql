
INSERT INTO `user` VALUES 
('1', 'admin', '123', '0', '1', now(), 0, now(),0, now()),
('2', 'dev', '123', '0', '1', now(), 0, now(),0, now()),
('3', 'testing', '123', '0', '1', now(), 0, now(),0, now()),
('4', 'labs', '123', '0', '1', now(), 0, now(),0, now()),
('5', 'prod', '123', '0', '1', now(), 0, now(),0, now());



INSERT INTO `menu` 
VALUES 
('1', '权限管理', '0', '', '0', '1', '0',now(), 0, now(), 0), 
('2', '用户管理', '1', 'auth/user', '0', '1', '0', now(), 0, now(), 0), 
('3', '角色管理', '1', 'auth/role', '0', '1', '0', now(), 0, now(), 0), 
('4', '资源管理', '1', 'auth/resource', '0', '1', '0', now(), 0, now(), 0), 
('5', '配置文件', '0', '', '0', '1', '0', now(), 0, now(), 0), 
('6', '开发配置文件', '5', 'xcfg/dev?env=dev', '0', '1', '0', now(), 0, now(), 0),
('7', '测试配置文件', '5', 'xcfg/testing?env=testing', '0', '1', '0', now(), 0, now(), 0),
('8', '实验配置文件', '5', 'xcfg/labs?env=labs', '0', '1', '0', now(), 0, now(), 0),
('9', '线上配置文件', '5', 'xcfg/prod?env=prod', '0', '1', '0', now(), 0, now(), 0),
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


INSERT INTO `role` VALUES 
('1', 'dev', '0', '1', now(), '0', now(), '0'), 
('2', 'testing', '0', '1', now(), '0', now(), '0'), 
('3', 'labs', '0', '1', now(), '0', now(), '0'), 
('4', 'prod', '0', '1', now(), '0', now(), '0');


INSERT INTO `rel_user_role` VALUES 
('1', '2', '', '1', 'dev'), 
('2', '3', '', '2', 'testing'),
('3', '4', '', '3', 'labs'),
('4', '5', '', '4', 'prod');



INSERT INTO `rel_role_menu`(id,menu_id,menu_name,role_id,role_name) VALUES 
('45', '5', '配置文件', '1', ''), 
('46', '6', '开发配置文件', '1', ''), 
('47', '25', '配置文件-列表dev', '1', ''), 
('48', '26', '配置文件-添加dev', '1', ''), 
('49', '27', '配置文件-是否存在dev', '1', ''), 
('50', '28', '配置文件-获取dev', '1', ''), 
('51', '29', '配置文件-更新dev', '1', ''), 
('52', '5', '配置文件', '2', ''), 
('53', '7', '测试配置文件', '2', ''), 
('54', '30', '配置文件-列表testing', '2', ''), 
('55', '31', '配置文件-添加testing', '2', ''), 
('56', '32', '配置文件-是否存在testing', '2', ''), 
('57', '33', '配置文件-获取testing', '2', ''), 
('58', '34', '配置文件-更新testing', '2', ''), 
('59', '5', '配置文件', '3', ''), 
('60', '8', '实验配置文件', '3', ''), 
('61', '35', '配置文件-列表labs', '3', ''), 
('62', '36', '配置文件-添加labs', '3', ''), 
('63', '37', '配置文件-是否存在labs', '3', ''), 
('64', '38', '配置文件-获取labs', '3', ''), 
('65', '39', '配置文件-更新labs', '3', ''), 
('66', '5', '配置文件', '4', ''), 
('67', '9', '线上配置文件', '4', ''), 
('68', '40', '配置文件-列表prod', '4', ''), 
('69', '41', '配置文件-添加prod', '4', ''), 
('70', '42', '配置文件-是否存在prod', '4', ''), 
('71', '43', '配置文件-获取prod', '4', ''), 
('72', '44', '配置文件-更新prod', '4', '');


INSERT INTO `cfg` VALUES 
('1', '', 'amqp-settings', '1', '1', '<?xml version=\'1.0\' encoding=\'utf-8\' ?><amqp-settings majorVersion=\"1\" minorVersion=\"1\"><servers><server name=\"default\" uri=\"amqp://abc\"></server><server name=\"abc\" uri=\"amqp://def\"></server><server name=\"def\" uri=\"amqp://ghi\"></server>\n    </servers>\n<exchanges><exchange server=\"default\" name=\"abc\" type=\"direct\" durable=\"true\" autoDelete=\"false\"></exchange><exchange server=\"default\" name=\"def\" type=\"direct\" durable=\"true\" autoDelete=\"false\"></exchange>\n       \n    </exchanges>\n<queues><queue name=\"abc\" server=\"default\" exchange=\"abc\" durable=\"true\" auto-delete=\"false\" need-ack=\"false\" exclusive=\"true\"></queue><queue name=\"def\" server=\"default\" exchange=\"abc\" durable=\"true\" auto-delete=\"false\" need-ack=\"false\" exclusive=\"true\"></queue>\n    </queues>\n\n</amqp-settings>\n', 'dev', now(), '0', now(), '0'), 
('2', 'Cfg4goTest', 'RecruitWechatNoticeConfiguration', '1', '1', '<?xml version=\'1.0\' encoding=\'utf-8\' ?><RecruitWechatNoticeConfiguration majorVersion=\"1\" minorVersion=\"1\"><SenderConfigurations><SenderConfiguration><GroupName>ReumeComplaint</GroupName><CorpSecret>a</CorpSecret><CorpId>wx4aa6233fe9f75f5e</CorpId><AgentId>1</AgentId><ToParty>2|3</ToParty>\n        </SenderConfiguration>\n<SenderConfiguration><GroupName>ReumeRepair</GroupName><CorpSecret>b</CorpSecret><CorpId>wx4aa6233fe9f75f5e</CorpId><AgentId>2</AgentId><ToParty>2</ToParty>\n        </SenderConfiguration>\n<SenderConfiguration><GroupName>ResumeAlert</GroupName><CorpSecret>c</CorpSecret><CorpId>wx4aa6233fe9f75f5e</CorpId><AgentId>3</AgentId><ToParty>2</ToParty>\n        </SenderConfiguration>\n\n    </SenderConfigurations>\n\n</RecruitWechatNoticeConfiguration>\n', 'dev', now(), '0', now(), '0')
