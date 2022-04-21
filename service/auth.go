package service

import "github.com/gin-gonic/gin"

// Register 注册
// @Summary 注册
// @Description 注册
// @Tags 认证
// @Accept  json
// @Produce  json
// @Param body body schema.Register true "注册"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /auth/register [post]
func Register(c *gin.Context) {
	c.JSON(200, gin.H{})
}

// Login 登录
// @Summary 登录
// @Description 登录
// @Tags 认证
// @Accept  json
// @Produce  json
// @Param body json schema.Login true "登录"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /auth/login [post]
func Login(c *gin.Context) {
	c.JSON(200, gin.H{})
}
