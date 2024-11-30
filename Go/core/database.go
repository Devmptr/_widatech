package core

import (
	"database/sql"
	"fmt"
	"widatech_interview/golang/config"
)

type Db struct {
	DbConfig   *config.DatabaseConfig
	Connection *sql.DB
}

func (d *Db) MakeConnection() error {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", d.DbConfig.User, d.DbConfig.Password, d.DbConfig.Host, d.DbConfig.Port, d.DbConfig.DbName)
	d.Connection, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	return nil
}

func (d *Db) CloseConnection() {
	d.Connection.Close()
}

func NewDB(conf *config.DatabaseConfig) *Db {
	return &Db{
		DbConfig: conf,
	}
}
