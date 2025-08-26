// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package db

import (
	"fmt"
	"log"
	"os"
	"time"

	mysql_driver "gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlog "gorm.io/gorm/logger"
)

type MySQLCfg struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Schema   string `json:"schema"`
}

type DBConfig struct {
	MaxIdleConn     int `mapstructure:"maxIdleConn"`
	MaxOpenConn     int `mapstructure:"maxOpenConn"`
	ConnMaxLifetime int `mapstructure:"connMaxLifetime"`
}

func InitMySQLDB(cfg *MySQLCfg, dbCfg *DBConfig) (*gorm.DB, error) {
	DSN := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Schema)
	slow := gormlog.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		gormlog.Config{
			SlowThreshold:             200 * time.Millisecond,
			LogLevel:                  gormlog.Error,
			IgnoreRecordNotFoundError: true, // 忽略 Record Not Found 错误
		},
	)

	db, err := gorm.Open(
		mysql_driver.New(mysql_driver.Config{
			DSN:                       DSN,
			DefaultStringSize:         64,
			DisableDatetimePrecision:  true,
			DontSupportRenameIndex:    true,
			DontSupportRenameColumn:   true,
			SkipInitializeWithVersion: false,
		}),
		&gorm.Config{
			QueryFields: true,
			Logger:      slow})
	if err != nil {
		return nil, err
	}
	c, err := db.DB()
	if err != nil {
		return nil, err
	}

	c.SetMaxIdleConns(dbCfg.MaxIdleConn)
	c.SetMaxOpenConns(dbCfg.MaxOpenConn)
	c.SetConnMaxLifetime(time.Second * time.Duration(dbCfg.ConnMaxLifetime))
	return db, nil
}
