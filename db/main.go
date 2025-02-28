package db

import (
	"github.com/verlinof/fiber-project-structure/configs/db_config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error

	dsnMysql := db_config.Config.DbUser + ":" + db_config.Config.DbPassword + "@tcp(" + db_config.Config.Host + ":" + db_config.Config.Port + ")/" + db_config.Config.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsnMysql), &gorm.Config{})

	//========Config Postgres========
	// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	// dsn := "host=" + db_config.DB_HOST + " user=" + db_config.DB_USER + " password=" + db_config.DB_PASSWORD + " dbname=" + db_config.DB_NAME + " port=" + db_config.DB_PORT + " sslmode=disable TimeZone=Asia/Shanghai"
	// DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}
