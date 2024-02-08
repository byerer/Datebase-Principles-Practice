package router

import (
	"GradingSystem/global"
	"GradingSystem/internal/dao/mysql"
	"GradingSystem/internal/middleware"
	"GradingSystem/internal/model/api"
	"GradingSystem/internal/model/database"
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func register(c *gin.Context) {
	var registerInfo api.RegisterInfo
	if err := c.BindJSON(&registerInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	global.SugarLogger.Infof("registerInfo: %v", registerInfo)
	// 验证验证码
	if !middleware.ValidateCode(registerInfo.EmailInfo) {
		c.JSON(http.StatusOK, gin.H{"msg": "验证码错误"})
		return
	}
	// 验证邮箱是否存在
	if mysql.FindUserByEmail(registerInfo.EmailInfo.Email) {
		c.JSON(http.StatusOK, gin.H{"msg": "邮箱已存在"})
		return
	}
	// 验证用户名是否存在
	if mysql.FindUserByName(registerInfo.User.Username) {
		c.JSON(http.StatusOK, gin.H{"msg": "用户名已存在"})
		return
	}

	// 使用 bcrypt 对密码进行加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerInfo.User.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法加密密码"})
		return
	}

	node, err := snowflake.NewNode(1)
	if err != nil {
		global.SugarLogger.Errorf("generate snowflake node failed: %v", err)
		return
	}
	id := node.Generate().Int64()
	var duser database.User
	duser.ID = id
	duser.Username = registerInfo.User.Username
	duser.Password = string(hashedPassword)
	duser.Email = registerInfo.EmailInfo.Email
	err = mysql.InsertUser(duser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "注册失败"})
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
	user, err := mysql.GetUserByUsername(loginInfo.Username)
	if err != nil {
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

func forgetPassword(c *gin.Context) {
	var forgetPasswordInfo api.ForgetPasswordInfo
	if err := c.BindJSON(&forgetPasswordInfo); err != nil {
		global.SugarLogger.Infof("forgetPasswordInfo: %v", forgetPasswordInfo)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := mysql.GetUserByEmail(forgetPasswordInfo.EmailInfo.Email)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"msg": "用户不存在"})
		return
	}
	// 验证验证码
	if !middleware.ValidateCode(forgetPasswordInfo.EmailInfo) {
		c.JSON(http.StatusOK, gin.H{"msg": "验证码错误"})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(forgetPasswordInfo.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "无法加密密码"})
		return
	}
	err = mysql.UpdatePassword(forgetPasswordInfo.EmailInfo.Email, string(hashedPassword))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "修改失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "修改成功"})
}

func sendCode(c *gin.Context) {
	email := c.Query("email")
	code := middleware.GenerateCode()
	err := middleware.SendCode(email, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "发送失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "发送成功",
		"code": code,
	})
}
