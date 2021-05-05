package main

import (
	"fmt"
	//"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 定义模型
type User struct {
	gorm.Model
	Name string `gorm:"default:'千羽'"` // 设置默认值
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
	// 创建
	//u1 := User{Name: "qianyu", Age: 21}
	//db.Create(&u1)
	//u2 := User{Name: "qianyue", Age: 21}
	//
	//db.Create(&u2)

	// 查询
	//var user User

	user := new(User)

	//db.First(&user)
	fmt.Println(db.First(user))
	fmt.Printf("user: %#v\n",user)

	var users []User
	db.Find(&users)
	fmt.Printf("users: %#v\n",users)

	// 未找到
	db.Where(User{Name: "non_existing"}).Assign(User{Age: 20}).FirstOrInit(&user)
	//// user -> User{Name: "non_existing", Age: 20}
	fmt.Printf("users: %#v\n",users)


}





















