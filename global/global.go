package global

import "gorm.io/gorm"

var Db *gorm.DB
var MysqlCahn = make(chan string)
