/*
 *  * Copyright (c) 2023 guojia99 All rights reserved.
 *  * Created: 2023/7/10 下午6:36.
 *  * Author: guojia(https://github.com/guojia99)
 */

package api

import (
	"github.com/guojia99/my-cubing-api/src/api/auth"
	"github.com/guojia99/my-cubing-api/src/api/report"
	"github.com/guojia99/my-cubing-api/src/api/result"
	"github.com/guojia99/my-cubing-api/src/api/stat"
	"github.com/guojia99/my-cubing-api/src/api/xlog"
)

func (c *Client) initRoute() {
	api := c.e.Group("/v2/api")

	api.POST("/auth/token", auth.ValidToken(c.svc)) // 获取授权
	api.POST("/auth/token/:player_id")              // 获取玩家授权

	// todo player middleware
	result.AddResultRoute(api, nil, auth.ValidMiddleware(c.svc), c.svc)
	xlog.AddXLogRoute(api, auth.ValidMiddleware(c.svc), c.svc)
	report.AddReportRoute(api, c.svc)
	stat.AddStatRoute(api, c.svc)
}
