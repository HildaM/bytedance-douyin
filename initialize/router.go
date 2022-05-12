package initialize

import (
	"bytedance-douyin/global"
	"bytedance-douyin/middleware"
	"bytedance-douyin/router"
	"github.com/gin-gonic/gin"
)

/**
 * @Author: 1999single
 * @Description: 初始化gin路由
 * @File: router
 * @Version: 1.0.0
 * @Date: 2022/5/6 15:25
 */

func Routers() *gin.Engine {
	Router := gin.Default()

	router := router.GroupApp
	// 如果想要不使用nginx代理前端网页，可以修改 web/.env.production 下的
	// VUE_APP_BASE_API = /
	// VUE_APP_BASE_PATH = http://localhost
	// 然后执行打包命令 npm run build。在打开下面4行注释
	// Router.LoadHTMLGlob("./dist/*.html") // npm打包成dist的路径
	// Router.Static("/favicon.ico", "./dist/favicon.ico")
	// Router.Static("/static", "./dist/assets")   // dist里面的静态资源
	// Router.StaticFile("/", "./dist/index.html") // 前端网页入口页面

	// Router.StaticFS(global.GVA_CONFIG.Local.Path, http.Dir(global.GVA_CONFIG.Local.Path)) // 为用户头像和文件提供静态地址
	global.GVA_LOG.Info("use middleware logger")
	// 跨域，如需跨域可以打开下面的注释
	// Router.Use(middleware.Cors()) // 直接放行全部跨域请求
	//Router.Use(middleware.CorsByRules()) // 按照配置的规则放行跨域请求
	global.GVA_LOG.Info("use middleware cors")
	// 方便统一添加路由组前缀 多服务器上线使用
	PublicGroup := Router.Group("douyin")
	{
		// 不做鉴权
		router.InitBaseUserRouter(PublicGroup)  // 注册、登录
		router.InitVideoFeedRouter(PublicGroup) // 视频流
		router.InitFollowRouter(PublicGroup)
	}
	PrivateGroup := Router.Group("douyin")
	PrivateGroup.Use(middleware.JWTAuth()) //.Use(middleware.CasbinHandler())
	{
		// 鉴权
		router.InitUserInfoRouter(PrivateGroup) // 查看用户信息
		router.InitVideoRouter(PrivateGroup)
		router.InitCommentRouter(PrivateGroup)
		//router.InitFollowRouter(PrivateGroup)
		router.InitLikeRouter(PrivateGroup)
	}

	// InstallPlugin(PublicGroup, PrivateGroup) // 安装插件

	global.GVA_LOG.Info("router register success")
	return Router
}
