/*
 *  * Copyright (c) 2023 guojia99 All rights reserved.
 *  * Created: 2023/7/10 下午6:36.
 *  * Author: guojia(https://github.com/guojia99)
 */

package api

import (
	"github.com/guojia99/my-cubing-api/src/api/report"
	"github.com/guojia99/my-cubing-api/src/api/result"
	"github.com/guojia99/my-cubing-api/src/api/xlog"
)

func (c *Client) initRoute() {
	api := c.e.Group("/v2/api")

	api.POST("/auth/token", c.ValidToken) // 获取授权

	result.AddResultRoute(api, c.AuthMiddleware, c.svc)
	xlog.AddXLogRoute(api, c.AuthMiddleware, c.svc)
	report.AddReportRoute(api, c.svc)
}
