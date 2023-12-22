package jd

import (
	"context"
	"fmt"
	"go.dtapp.net/gorequest"
)

// 请求接口
func (c *Client) request(ctx context.Context, param gorequest.Params) (gorequest.Response, error) {

	// 签名
	c.Sign(param)

	// 创建请求
	client := gorequest.NewHttp()

	// 设置格式
	client.SetContentTypeForm()

	// 设置用户代理
	client.SetUserAgent(gorequest.GetRandomUserAgentSystem())

	// 设置参数
	client.SetParams(param)

	// 发起请求
	request, err := client.Post(ctx)
	if err != nil {
		return gorequest.Response{}, err
	}

	// 记录日志
	if c.gormLog.status {
		go c.gormLog.client.MiddlewareCustom(ctx, fmt.Sprintf("%s", param.Get("method")), request)
	}

	return request, err
}
