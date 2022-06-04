package controller

import (
	"gin-gorm-token-demo/common"
	"gin-gorm-token-demo/model"
	"gin-gorm-token-demo/response"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func Register(ctx *gin.Context) {
	db := common.GetDB()
	//获取参数
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	//数据验证

	if len(username) == 0 || username == "" {
		response.Response(ctx, http.StatusUnprocessableEntity, 442, nil, "用户名不能为空")
		return
	}
	if len(password) < 8 || password == "" {
		response.Response(ctx, http.StatusUnprocessableEntity, 442, nil, "密码必须大于8位")
		return
	}
	log.Println(username, password)

	//判断手机号是否存在
	if isUsernameExist(db, username) {
		response.Response(ctx, http.StatusUnprocessableEntity, 442, nil, "该用户已注册")
		return
	}

	//创建用户
	hashdPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) //密码hash化
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 500, nil, "加密错误")
		return
	}
	newUser := model.User{
		Username: username,
		Password: string(hashdPassword),
	}
	if err := db.Create(&newUser).Error; err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 442, nil, err.Error())
		return
	}
	response.Success(ctx, nil, "注册成功")
}

func isUsernameExist(db *gorm.DB, username string) bool {
	var user model.User
	db.Where("username=?", username).First(&user)
	return user.ID != 0
}

func Login(ctx *gin.Context) {
	db := common.GetDB()
	//获取参数
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	//数据校验
	if len(username) == 0 || username == "" {
		response.Response(ctx, http.StatusUnprocessableEntity, 442, nil, "用户名不能为空")
		return
	}
	if len(password) < 8 {
		response.Response(ctx, http.StatusUnprocessableEntity, 442, nil, "密码必须大于8位")
		return
	}

	//判断手机号是否存在
	var user model.User
	db.Where("username=?", username).First(&user)
	if user.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 442, nil, "用户不存在，请重新输入")
		return
	}

	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 400, nil, "密码错误")
		return
	}

	//发送token
	token, err := common.GetToken(user)
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 500, nil, "系统token获取异常")
		return
	}

	response.Success(ctx, gin.H{"token": token}, "登录成功")

}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	response.Success(ctx, gin.H{"user": user}, "")
}
