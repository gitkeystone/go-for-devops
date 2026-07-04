package main

import (
	"database/sql"
	"fmt"
)

var db *sql.DB
var err error

func getPrimaryModePort() (string, error) {
	driverName := "dm"
	dataSourceName := "dm://dwc?dwc=(18.22.45.6:31821,18.22.45.6:32366)&TIME_ZONE=(+480)&LOGIN_MODE=1&characterEncoding=utf8"

	// 创建数据库连接
	if db, err = connect(driverName, dataSourceName); err != nil {
		return "", err
	}

	// 查询 PRIMARY MODE 端口号
	var port string
	if port, err = queryPrimaryModePort(); err != nil {
		return "", err
	}

	// 断开数据库连接
	if err = disconnect(); err != nil {
		return "", err
	}

	// 返回 PRIMARY MODE 端口号
	return port, nil
}

// 查询 PRIMARY MODE 端口号
func queryPrimaryModePort() (string, error) {
	var port string
	err = db.QueryRow(`SELECT  VALUE 
			FROM V$PARAMETER
			WHERE NAME = 'PORT_NUM'`).
		Scan(&port)
	if err != nil {
		return "", err
	}
	return port, nil
}

// 创建数据库连接
func connect(driverName, dataSourceName string) (*sql.DB, error) {
	var db *sql.DB
	var err error
	if db, err = sql.Open(driverName, dataSourceName); err != nil {
		return nil, fmt.Errorf("open database failed, err:%v\n", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("ping database failed, err:%v\n", err)
	}

	// 测试数据库连接成功
	fmt.Printf("connect successed, database is %v\n", db)
	return db, nil
}

// 关闭数据库连接
func disconnect() error {
	if err := db.Close(); err != nil {
		return fmt.Errorf("close database failed, err:%v\n", err)
	}
	fmt.Println("close database successed")
	return nil
}
