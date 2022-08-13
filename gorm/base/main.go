package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserInfo struct {
	Id     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	// 连接数据库
	db, err := gorm.Open("mysql", "root:root@(47.106.181.104)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer func(db *gorm.DB) {
		_ = db.Close()
	}(db)

	// 创建表 自动迁移(把结构体和数据库表进行对应)
	db.AutoMigrate(&UserInfo{})

	// 创建数据行
	//u1 := UserInfo{1, "John", "男", "code"}
	//db.Create(u1)

	// 查询
	var u UserInfo
	db.First(&u)
	fmt.Printf("u:%#v\n", u)

	// 更新
	db.Model(&u).Update("hobby", "双色球")

	// 删除
	db.Delete(&u)
}
