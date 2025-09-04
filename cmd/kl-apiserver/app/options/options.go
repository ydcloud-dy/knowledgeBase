// Copyright 2025 杜杰 <dycloudlove@163.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/ydcloud-dy/knowledgeBase.git

package options

import (
	"github.com/ydcloud-dy/knowledgeBase.git/internal/apiserver"
	"github.com/ydcloud-dy/knowledgeBase.git/pkg/errorsx"
	"github.com/ydcloud-dy/knowledgeBase.git/pkg/log"
	genericoptions "github.com/ydcloud-dy/knowledgeBase.git/pkg/options"
	"net"
	"strconv"
)

type ServerOptions struct {
	MySQLOptions *genericoptions.MySQLOptions `json:"mysql" mapstructure:"mysql"`
	Addr         string                       `json:"addr" mapstructure:"addr"`
}

// NewServerOptions 创建带有默认值的 ServerOptions 实例.
func NewServerOptions() *ServerOptions {
	return &ServerOptions{
		MySQLOptions: genericoptions.NewMySQLOptions(),
		Addr:         "0.0.0.0:8888",
	}
}

// Validate 校验 ServerOptions 中的选项是否合法.
// 提示：Validate 方法中的具体校验逻辑可以由 Claude、DeepSeek、GPT 等 LLM 自动生成。
func (o *ServerOptions) Validate() error {
	// 验证服务器地址
	if o.Addr == "" {
		log.Errorw("Server address cannot be empty", "Addr", o.Addr)
		return errorsx.ErrInvalidArgument
	}

	// 检查地址格式是否为host:port
	_, portStr, err := net.SplitHostPort(o.Addr)
	if err != nil {
		log.Errorw("invalid server address format", "Addr", err)
		return errorsx.ErrInvalidArgument
	}

	// 验证端口是否为数字且在有效范围内
	port, err := strconv.Atoi(portStr)
	if err != nil || port < 1 || port > 65535 {
		log.Errorw("invalid server address format", "Port", err)
		return errorsx.ErrInvalidArgument
	}
	if err := o.MySQLOptions.Validate(); err != nil {
		return err
	}

	return nil
}

// Config 基于 ServerOptions 构建 apiserver.Config.
func (o *ServerOptions) Config() (*apiserver.Config, error) {
	return &apiserver.Config{
		MySQLOptions: o.MySQLOptions,
		Addr:         o.Addr,
	}, nil
}

// Config 基于 ServerOptions 构建 apiserver.Config.
