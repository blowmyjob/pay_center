package initServer

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"pay_center/config"
)

func GormMysql() *gorm.DB {
	username := config.DB_USER_NAME  //账号
	password := config.DB_USER_PASWD //密码
	host := config.DB_USER_HOST      //数据库地址，可以是Ip或者域名
	port := config.DB_USER_PORT      //数据库端口
	Dbname := config.DB_NAME         //数据库名
	timeout := config.DB_TIMEOUT     //连接超时，10秒

	//拼接下dsn参数, dsn格式可以参考上面的语法，这里使用Sprintf动态拼接dsn参数，因为一般数据库连接参数，我们都是保存在配置文件里面，需要从配置文件加载参数，然后拼接dsn。
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}
	return db
}
