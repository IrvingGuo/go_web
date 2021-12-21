package config

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	var err error
	db, _ := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:root@tcp(10.70.9.111:3306)/test?charset=utf8mb4&parseTime=True&loc=UTC",
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		SkipDefaultTransaction:                   false,
		DisableForeignKeyConstraintWhenMigrating: true, //逻辑外键(代码里自动建立外键关系)
	})

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)           //设置空闲连接池中连接的最大数量
	sqlDB.SetMaxOpenConns(100)          //sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour) //SetConnMaxLifetime 设置了连接可复用的最大时间
	Db = db
	if err != nil {
		log.Fatal(err.Error())
	}
}
