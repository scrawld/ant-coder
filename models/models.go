/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2018-08-29 08:17:31
# File Name: models.go
# Description:
####################################################################### */

package models

import (
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/ant-libs-go/util"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	Orm *xorm.Engine
)

func Init(dsn string) error {
	user, pawd, host, port, name, err := parseDsn(dsn)
	if err != nil {
		return fmt.Errorf("parse dsn err: %s", err)
	}
	Orm, err = xorm.NewEngine("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&autocommit=true",
		user, pawd, host, port, name))
	if err != nil {
		return fmt.Errorf("db connect err: %s", err)
	}
	Orm.SetConnMaxLifetime(3000)

	if os.Getenv("VERBOSE") == "true" {
		Orm.ShowSQL(true)
	} else {
		Orm.ShowSQL(false)
	}
	Orm.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
	return nil
}

func parseDsn(dsn string) (user, pawd, host, port, name string, err error) {
	re, _ := regexp.Compile(`(?P<user>[\w]+):(?P<pawd>[\w]+)@(?P<host>[\w\.]+):(?P<port>[\w]+)/(?P<name>[\w]+)`)
	r, err := util.FindStringSubmatch(re, dsn)
	for k, v := range r {
		switch k {
		case "user":
			user = v
		case "pawd":
			pawd = v
		case "host":
			host = v
		case "port":
			port = v
		case "name":
			name = v
		}
	}
	return
}
