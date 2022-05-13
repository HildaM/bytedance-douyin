# bytedance-douyin

## 项目结构
├─api                   控制层
│  ├─response           仅存放一个response响应工具
│  └─vo                 控制层与客户端互相传递的对象，分为vo与response vo
├─config                项目中的配置结构体，与resource/config.yaml对应。
├─core                  全局配置，很少用。
├─global                全局资源，redis、db连接等
├─initialize            初始化各类工具
├─log                   日志（不提交git）
├─middleware            gin中间件
├─repository            数据持久化层 封装对某个model的查询与更新操作，供service层调用
│  └─model              model与数据库表对应，即dao
├─resource              yaml配置文件
├─router                路由
├─service               服务层
│  └─bo                 service层传给controller层的对象
├─test                  测试文件
└─utils                 工具类

// 持久化层到客户端的数据封装。由于部分业务简单，灵活变通吧。
// 大多数情况下vo、bo、model其实都长得差不多，可以通过json的tag屏蔽掉一些敏感信息。

vo -> bo -> dao
vo <- bo <- dao