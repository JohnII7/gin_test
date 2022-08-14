package main

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type UserInfo struct {
	Id     uint
	Name   string
	Gender string
	Hobby  string
}
type User struct {
	gorm.Model   // 内嵌gorm.Model
	Name         string
	Age          sql.NullInt64 `gorm:"column:user_age"` // 零值类型
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`        // 设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"` // 设置会员号唯一且不为空
	Num          int     `gorm:"AUTO_INCREMENT"`  // 设置自增
	Address      string  `gorm:"index:addr"`      // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"`               // 忽略本字段
}

type Animal struct {
	AnimalID int64 `gorm:"primary_key"`
	Name     string
	Age      int64
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

	db.AutoMigrate(&User{})
	db.AutoMigrate(Animal{})

}
