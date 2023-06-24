package initServer

import (
	"github.com/gin-gonic/gin"
	"pay_center/controller/admin"
)

func InitHttpServer() {
	// 初始化一个http服务对象
	router := gin.Default()
	// 添加admin
	adminRouter := new(admin.AdminUserRouter)
	// 注册admin路径请求
	{
		adminRouter.InitManageAdminUserRouter(&router.RouterGroup)
	}

	router.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
