package nonecode

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// var DB *gorm.DB

type GormConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

func NewGorm(conf *GormConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Database)
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名，启用该选项后，`User` 表将是`user`
		},
	})

	return db
}
