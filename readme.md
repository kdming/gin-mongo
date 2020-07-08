#### go gin web框架
##### 数据库为mongo,各层含义如下
1.api 加载路由提供外部接口  
2.dao 数据库驱动以及crud工具层  
3.middleware gin中间件  
4.models 模型，数据持久化层  
5.pkg 工具层  
6.service 数据处理/逻辑层  
7.static 静态文件目录  
8.config.yaml 全局配置文件  
9.main.go 项目入口  

config.yaml填写完配置  
执行: `go run main.go`
