package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	ID   int64
	Name string
	Age  int64
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

	// 2. 映射
	db.AutoMigrate(&User{})

	// 3. 创建
	u := User{
		Name: "John",
		Age:  18,
	}
	fmt.Println(db.NewRecord(&u))
	db.Create(u)
	fmt.Println(db.NewRecord(&u))

}
