package model

import "fmt"

//执行数据迁移

func migration() {
	// 自动迁移模式
	fmt.Println("执行数据库迁移")
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Video{})
}
