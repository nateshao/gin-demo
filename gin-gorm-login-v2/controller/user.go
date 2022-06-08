package controller

import (
	"gin-gorm-login-v2/common"
	"gin-gorm-login-v2/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Register(ctx *gin.Context) {

	db := common.GetDB()
	//获取参数
	//此处使用Bind()函数，可以处理不同格式的前端数据
	var requestUser model.User
	ctx.Bind(&requestUser)
	//username := requestUser.UserName
	//password := requestUser.Password

	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	//数据验证
	if len(username) == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "用户名不能为空",
		})
		return
	}
	//if len(telephone) != 11 {
	//	ctx.JSON(http.StatusUnprocessableEntity, gin.H{
	//		"code":    422,
	//		"message": "手机号必须为11位",
	//	})
	//	return
	//}
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "密码不能少于6位",
		})
		return
	}

	//判断手机号是否存在
	var user model.User
	db.Where("username = ?", username).First(&user)
	if user.ID != 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "用户已存在",
		})
		return
	}

	//创建用户
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    500,
			"message": "密码加密错误",
		})
		return
	}
	newUser := model.User{
		UserName: username,
		Password: string(hasedPassword),
	}
	db.Create(&newUser)

	//返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "注册成功",
	})
}

func Login(ctx *gin.Context) {

	db := common.GetDB()

	//获取参数
	//此处使用Bind()函数，可以处理不同格式的前端数据
	var requestUser model.User
	ctx.Bind(&requestUser)
	//username := requestUser.UserName
	//password := requestUser.Password

	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	//数据验证
	if len(username) == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "用户名不能为空",
		})
		return
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "密码不能少于6位",
		})
		return
	}

	//判断手机号是否存在
	var user model.User
	db.Where("username = ?", username).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "用户不存在",
		})
		return
	}

	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "密码错误",
		})
	}

	//返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "登录成功",
	})
}
