// Copyright 2025 杜杰 <dycloudlove@163.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/ydcloud-dy/knowledgeBase.git

package apiserver

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/ydcloud-dy/knowledgeBase.git/pkg/log"
	mw "github.com/ydcloud-dy/knowledgeBase.git/pkg/middleware"
	genericoptions "github.com/ydcloud-dy/knowledgeBase.git/pkg/options"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Config 配置结构体，用于存储应用相关的配置.
// 不用 viper.Get，是因为这种方式能更加清晰的知道应用提供了哪些配置项.
type Config struct {
	MySQLOptions *genericoptions.MySQLOptions
	Addr         string
}

// Server 定义一个服务器结构体类型.
type Server struct {
	cfg *Config
	svc *http.Server
}

// NewServer 根据配置创建服务器.
func (cfg *Config) NewServer() (*Server, error) {
	// 创建 Gin 引擎
	engine := gin.New()
	mws := []gin.HandlerFunc{gin.Recovery(), mw.NoCache, mw.Cors, mw.RequestID()}
	engine.Use(mws...)
	// 注册 404 Handler.
	engine.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": "PageNotFound", "message": "Page not found."})
	})

	// 注册 /healthz handler.
	engine.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// 创建 HTTP Server 实例
	httpsrv := &http.Server{Addr: cfg.Addr, Handler: engine}
	return &Server{cfg: cfg, svc: httpsrv}, nil
}

// Run 运行应用.
func (s *Server) Run() error {
	// 运行 HTTP 服务器
	// 打印一条日志，用来提示 HTTP 服务已经起来，方便排障
	log.Infow("Start to listening the incoming requests on http address", "addr", s.cfg.Addr)
	go func() {
		if err := s.svc.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error(err.Error())
			os.Exit(1)
		}
	}()
	// 创建一个 os.Signal 类型的 channel，用于接收系统信号
	quit := make(chan os.Signal, 1)
	// 当执行 kill 命令时（不带参数），默认会发送 syscall.SIGTERM 信号
	// 使用 kill -2 命令会发送 syscall.SIGINT 信号（例如按 CTRL+C 触发）
	// 使用 kill -9 命令会发送 syscall.SIGKILL 信号，但 SIGKILL 信号无法被捕获，因此无需监听和处理
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// 阻塞程序，等待从 quit channel 中接收到信号
	<-quit
	slog.Info("Shutting down server ...")
	// 优雅关闭服务
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// 先关闭依赖的服务，再关闭被依赖的服务
	// 10 秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过 10 秒就超时退出
	if err := s.svc.Shutdown(ctx); err != nil {
		log.Errorw("Insecure Server forced to shutdown", "err", err)
		return err
	}
	log.Infow("Server exited")

	return nil
	// select {} // 调用 select 语句，阻塞防止进程退出
}
