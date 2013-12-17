package models

import (
    "database/sql"
    _ "mysql"
)

func init() {
	db, err := sql.Open(conf.driver, conf.uname + ":" + conf.pass + "@" + conf.protocol + "(" + conf.host + ":" + conf.port ")/" + conf.database + "?charset=" + conf.charset)
}
