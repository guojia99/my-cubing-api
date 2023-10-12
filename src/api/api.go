/*
 *  * Copyright (c) 2023 guojia99 All rights reserved.
 *  * Created: 2023/7/10 下午5:59.
 *  * Author: guojia(https://github.com/guojia99)
 */

package api

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/guojia99/my-cubing-api/src/api/middleware"
	"github.com/guojia99/my-cubing-api/src/svc"
)

func NewClient(s *svc.Context) *Client {
	return &Client{svc: s}
}

type Client struct {
	e   *gin.Engine
	svc *svc.Context
}

func (c *Client) Run() error {
	c.initAuth()

	gin.SetMode(c.svc.Cfg.GinMode)
	c.e = gin.New()
	c.e.Use(
		gin.Logger(),
		gin.Recovery(),
		middleware.CorsMiddleware(),
		middleware.NewRateMiddleware(10),
		//middleware.NewStatusCodeGreaterThan(500),
	)
	c.initRoute()

	return c.e.Run(fmt.Sprintf("0.0.0.0:%d", c.svc.Cfg.Port))
}
