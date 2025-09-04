// Copyright 2025 杜杰 <dycloudlove@163.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/ydcloud-dy/knowledgeBase.git

package options

import (
	"fmt"
	"github.com/ydcloud-dy/knowledgeBase.git/pkg/errorsx"
	"github.com/ydcloud-dy/knowledgeBase.git/pkg/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net"
	"strconv"
	"time"
)

type MySQLOptions struct {
	Addr                  string        `json:"addr,omitempty" mapstructure:"addr"`
	Username              string        `json:"username,omitempty" mapstructure:"username"`
	Password              string        `json:"-" mapstructure:"password"`
	Database              string        `json:"database" mapstructure:"database"`
	MaxIdleConnections    int           `json:"max-idle-connections,omitempty" mapstructure:"max-idle-connections,omitempty"`
	MaxOpenConnections    int           `json:"max-open-connections,omitempty" mapstructure:"max-open-connections"`
	MaxConnectionLifeTime time.Duration `json:"max-connection-life-time,omitempty" mapstructure:"max-connection-life-time"`
}

// NewMySQLOptions create a `zero` value instance.
func NewMySQLOptions() *MySQLOptions {
	return &MySQLOptions{
		Addr:                  "127.0.0.1:3306",
		Username:              "onex",
		Password:              "onex(#)666",
		Database:              "onex",
		MaxIdleConnections:    100,
		MaxOpenConnections:    100,
		MaxConnectionLifeTime: time.Duration(10) * time.Second,
	}
}

// Validate verifies flags passed to MySQLOptions.
func (o *MySQLOptions) Validate() error {
	// 验证MySQL地址格式
	if o.Addr == "" {
		log.Errorw("MySQL server address cannot be empty")
		return errorsx.ErrInvalidArgument
	}
	// 检查地址格式是否为host:port
	host, portStr, err := net.SplitHostPort(o.Addr)
	if err != nil {
		log.Errorw("Invalid MySQL address format", "addr", err)
		return errorsx.ErrInvalidArgument
	}
	// 验证端口是否为数字
	port, err := strconv.Atoi(portStr)
	if err != nil || port < 1 || port > 65535 {
		log.Errorw("Invalid MySQL port:", "mysqlPort", err)
		return errorsx.ErrInvalidArgument
	}
	// 验证主机名是否为空
	if host == "" {
		log.Errorw("MySQL hostname cannot be empty", "mysqlHostname", err)
		return errorsx.ErrInvalidArgument
	}

	// 验证凭据和数据库名
	if o.Username == "" {
		log.Errorw("MySQL username cannot be empty")
		return errorsx.ErrInvalidArgument
	}

	if o.Password == "" {
		log.Errorw("MySQL password cannot be empty")
		return errorsx.ErrInvalidArgument
	}

	if o.Database == "" {
		log.Errorw("MySQL database name cannot be empty")
		return errorsx.ErrInvalidArgument
	}

	// 验证连接池参数
	if o.MaxIdleConnections <= 0 {
		log.Errorw("MySQL max idle connections must be greater than 0", "maxIdleConnections", o.MaxIdleConnections)
		return errorsx.ErrInvalidArgument
	}

	if o.MaxOpenConnections <= 0 {
		log.Errorw("MySQL max open connections must be greater than 0", "maxOpenConnections", o.MaxOpenConnections)
		return errorsx.ErrInvalidArgument
	}

	if o.MaxIdleConnections > o.MaxOpenConnections {
		log.Errorw("MySQL max idle connections cannot be greater than max open connections",
			"maxIdleConnections", o.MaxIdleConnections, "maxOpenConnections", o.MaxOpenConnections)
		return errorsx.ErrInvalidArgument
	}

	if o.MaxConnectionLifeTime <= 0 {
		log.Errorw("MySQL max connection lifetime must be greater than 0", "maxConnectionLifetime", o.MaxConnectionLifeTime)
		return errorsx.ErrInvalidArgument
	}

	return nil
}

// DSN return DSN from MySQLOptions.
func (o *MySQLOptions) DSN() string {
	return fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s`,
		o.Username,
		o.Password,
		o.Addr,
		o.Database,
		true,
		"Local")
}

// NewDB create mysql store with the given config.
func (o *MySQLOptions) NewDB() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(o.DSN()), &gorm.Config{
		// PrepareStmt executes the given query in cached statement.
		// This can improve performance.
		PrepareStmt: true,
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(o.MaxOpenConnections)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(o.MaxConnectionLifeTime)

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(o.MaxIdleConnections)

	return db, nil
}
