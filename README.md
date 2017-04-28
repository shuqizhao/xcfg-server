# xcfg-server
golang 远程配置文件服务器

# 安装说明
    1,xcfg-server使用了mysql和redis,请在启动前请在app.config中修改对应的连接接字符串
    2,xcfg-server是基于beego开发的,应用xcfg-server请使用beego命令启动
    3,启用xcfg-server后请在mysql中执行data.sql脚本初始化数据
    4,初始化用户名和密码分别为admin和123
    5,客户端案列请参考我的资源xcfg-demo
