/*
 *  * Copyright (c) 2023 guojia99 All rights reserved.
 *  * Created: 2023/7/17 下午5:06.
 *  * Author: guojia(https://github.com/guojia99)
 */

package auth

import (
	"crypto/rand"
	"encoding/base64"
	"strconv"
	"time"

	"gorm.io/gorm"

	"github.com/guojia99/my-cubing-api/src/svc"
)

func InitAuth(svc *svc.Context) {
	_ = svc.DB.AutoMigrate(&Admin{})
}

type Admin struct {
	gorm.Model
	UserName string    `gorm:"column:user_name"`
	Password string    `gorm:"column:password"`
	Token    string    `gorm:"column:token"`
	Timeout  time.Time `gorm:"column:timeout"`
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

func generateUniqueToken(username string, ts int64) string {
	uniqueStr := username + strconv.FormatInt(ts, 10)
	uniqueBytes := []byte(uniqueStr)
	randomBytes := make([]byte, 255-len(uniqueBytes))
	_, _ = rand.Read(randomBytes)
	tokenBytes := append(uniqueBytes, randomBytes...)
	return base64.URLEncoding.EncodeToString(tokenBytes)
}
