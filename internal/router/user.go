package router

import (
	"GradingSystem/global"
	"GradingSystem/internal/middleware"
	"GradingSystem/internal/model/api"
	"GradingSystem/internal/model/database"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func register(c *gin.Context) {
	var user api.UserCreate
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 使用 bcrypt 对密码进行加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法加密密码"})
		return
	}
	user.Password = string(hashedPassword)

	node, err := snowflake.NewNode(1)
	if err != nil {
		log.Println(err)
		return
	}
	id := node.Generate().Int64()
	var duser database.User
	duser.ID = id
	duser.Username = user.Username
	duser.Password = user.Password
	duser.Email = user.Email
	result := global.DB.Create(&duser)
	if result.Error != nil {
		log.Println(result.Error)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": 0,
	})
}

func login(c *gin.Context) {
	var loginInfo api.User
	if err := c.BindQuery(&loginInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user database.User
	fmt.Println(loginInfo.Username, loginInfo.Password)
	result := global.DB.Where("username = ?", loginInfo.Username).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{"msg": "用户名或密码错误"})
		return
	}

	// 使用 bcrypt.CompareHashAndPassword 比较提交的密码和数据库中的哈希
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginInfo.Password)); err != nil {
		// 密码不匹配
		c.JSON(http.StatusOK, gin.H{"msg": "用户名或密码错误"})
		return
	}
	token, err := middleware.GenerateJWT(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "无法生成token"})
		return
	}
	// 密码匹配
	c.JSON(http.StatusOK, gin.H{
		"msg":   "登录成功",
		"token": token,
	})
}

func sendCode(c *gin.Context) {
	email := c.Query("email")
	code := middleware.GenerateCode()
	err := middleware.SendCode(email, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "发送失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "发送成功"})
}
