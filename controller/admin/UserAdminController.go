package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"pay_center/config"
	"pay_center/model"
)

type AdminUserRouter struct {
}

func (r *AdminUserRouter) InitManageAdminUserRouter(Router *gin.RouterGroup) {
	AdminUserRouter := Router.Group("/admin")
	{
		AdminUserRouter.GET("/user", handlerExist)
	}
}

func handlerExist(c *gin.Context) {
	//获取name参数, 通过Query获取的参数值是String类型。
	name := c.Query("name")
	id, ok := c.GetQuery("id")
	if !ok {
		c.String(http.StatusBadRequest, fmt.Sprintf("hello %s\n", name))
		return
	}
	var results []model.Activity
	sql := fmt.Sprintf("SELECT * FROM `t` where id = %s", id)
	config.GvaDb.Raw(sql).Scan(&results)
	fmt.Println(results)
	c.String(http.StatusOK, fmt.Sprintf("hello %s %d", name, results[0].Id))
}
