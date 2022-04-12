package db

import (
	"database/sql"
	"fmt"
	// mysql驱动 必须引入
	_ "github.com/go-sql-driver/mysql"
	"sendMsg/log"
)

func initDb(dataSourceName string) error {
	// 创建sql.db连接池
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Err("database|initDb sql.Open is err:%v", err)
		return err
	}
	// 连接池最多同时打开的连接数(包括使用中+空闲连接)
	db.SetMaxOpenConns(maxConn)
	// 设置最大空闲连接数
	db.SetMaxIdleConns(maxIdle)
	// 连接池里最大连接存活时长 不能超过mysql设置的wait_timeout 否则导致golang保留了mysql已经关闭的连接
	// mysql 默认是8h 可以使用 show VARIABLES like "wait_timeout" 来查看mysql的设置时间
	db.SetConnMaxLifetime(maxLifeTime)
	// 该方法在go 1.15版本之后引入 设置连接池里面的连接最大空闲时长
	db.SetConnMaxIdleTime(maxIdleTime)
	if err = db.Ping(); err != nil {
		log.Err("database|db.Ping is err:%v", err)
		return err
	}
	return nil
}

func getDataSourceName(dbName, userName, passWord, host string, port int) string {
	str := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", userName, passWord, host, port, dbName)
	str += "?charset=utf8&timeout=5s&parseTime=true&loc=Asia%2FShanghai"
	return str
}
