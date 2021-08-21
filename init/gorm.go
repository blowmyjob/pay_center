package init

import (
	"github.com/douyu/jupiter/pkg/store/gorm"
)

func GormMysql(name string) *gorm.DB {
	return gorm.StdConfig(name).Build()
}
