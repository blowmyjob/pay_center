package initServer

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"pay_center/config"
)

func ReadProperties(env string) {
	configPath := fmt.Sprintf("./config-%s.toml", env)
	configs := make(map[string]map[string]interface{})
	_, err := toml.DecodeFile(configPath, &configs)
	if err != nil {
		fmt.Println(err.Error())
	}
	sqlConfig := configs["mysql_config"]
	config.DB_USER_NAME = sqlConfig["username"].(string)
	config.DB_USER_PASWD = sqlConfig["password"].(string)
	config.DB_USER_HOST = sqlConfig["host"].(string)
	config.DB_USER_PORT = sqlConfig["port"].(int64)
	config.DB_NAME = sqlConfig["Dbname"].(string)
	config.DB_TIMEOUT = sqlConfig["timeout"].(string)

	redisConfig := configs["redis_config"]
	config.REDIS_ARR = redisConfig["addr"].(string)
	config.REDIS_PASS = redisConfig["password"].(string)

	mongoConfig := configs["mongo-config"]
	config.MONGO_ARR = mongoConfig["addr"].(string)
	config.MONGO_POOL = mongoConfig["poolNum"].(int64)

}
