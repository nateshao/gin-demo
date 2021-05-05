package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 定义模型
type User struct {
	ID   int64
	Name string
	Age  int64
}

func main() {
	// 连接数据库
	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 把模型与数据库的表对应起来
	db.AutoMigrate(&User{})

}
