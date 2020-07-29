package mysql

import (
	"database/sql"
	"log"
)

func main(){
	db,err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err!= nil{
		log.Print("数据库连接失败")
	}
}
