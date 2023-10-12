/*
 *  * Copyright (c) 2023 guojia99 All rights reserved.
 *  * Created: 2023/7/17 下午5:06.
 *  * Author: guojia(https://github.com/guojia99)
 */

package api

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/guojia99/my-cubing-api/src/api/common"
)

type Admin struct {
	gorm.Model
	UserName string    `gorm:"column:user_name"`
	Password string    `gorm:"column:password"`
	Token    string    `gorm:"column:token"`
	Timeout  time.Time `gorm:"column:timeout"`
}

func (c *Client) initAuth() {
	_ = c.svc.DB.AutoMigrate(&Admin{})
}

type (
	GetTokenRequest struct {
		UserName string `json:"user_name"`
		PassWord string `json:"password"`
	}

	GetTokenResponse struct {
		Ts    int64  `json:"Ts"`
		Token string `json:"Token"`
	}
)

// ValidToken 获取合法token
func (c *Client) ValidToken(ctx *gin.Context) {
	var req GetTokenRequest
	if err := ctx.Bind(&req); err != nil {
		common.Error(ctx, http.StatusNetworkAuthenticationRequired, 20, "密码格式错误")
		return
	}
	if req.UserName == "" || req.PassWord == "" {
		common.Error(ctx, http.StatusNetworkAuthenticationRequired, 20, "密码格式错误")
		return
	}

	var admin Admin
	if err := c.svc.DB.Where("user_name = ?", req.UserName).First(&admin).Error; err != nil {
		common.Error(ctx, http.StatusNetworkAuthenticationRequired, 20, "查询不到角色")
		return
	}
	if admin.Password != req.PassWord {
		common.Error(ctx, http.StatusNetworkAuthenticationRequired, 20, "密码错误")
		return
	}
	if time.Now().Sub(admin.Timeout) > 0 || admin.Token == "" {
		admin.Token = generateUniqueToken(req.UserName, time.Now().Unix())
		admin.Timeout = time.Now().Add(time.Hour * 48)
		c.svc.DB.Save(&admin)
	}

	resp := GetTokenResponse{
		Ts:    admin.Timeout.Unix(),
		Token: admin.Token,
	}
	ctx.JSON(http.StatusOK, resp)
}

func generateUniqueToken(username string, ts int64) string {
	uniqueStr := username + strconv.FormatInt(ts, 10)
	uniqueBytes := []byte(uniqueStr)
	randomBytes := make([]byte, 255-len(uniqueBytes))
	_, _ = rand.Read(randomBytes)
	tokenBytes := append(uniqueBytes, randomBytes...)
	return base64.URLEncoding.EncodeToString(tokenBytes)
}

// AuthMiddleware 授权中间件
func (c *Client) AuthMiddleware(ctx *gin.Context) {
	log.Printf("auth in %s", ctx.ClientIP())
	token := ctx.GetHeader("Authorization")
	if token == "" {
		common.Error(ctx, http.StatusUnauthorized, 21, "权限不足")
		return
	}

	// token 可用，直接通过
	if _, ok := c.svc.Cache.Get(token); ok {
		ctx.Next()
		return
	}

	// 查token 是否合法
	var admin Admin
	if err := c.svc.DB.Where("token = ?", token).First(&admin).Error; err != nil {
		common.Error(ctx, http.StatusUnauthorized, 21, "权限不足")
		return
	}
	if time.Now().Sub(admin.Timeout) > 0 {
		common.Error(ctx, http.StatusUnauthorized, 21, "权限过期")
		return
	}

	_ = c.svc.Cache.Add(admin.Token, admin.UserName, time.Minute)
	ctx.Next()
}
