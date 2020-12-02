package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func resolveDsn(configs *DbConfigs) string {
	host := dbConfigs.DbHostIP + ":" + dbConfigs.DbHostPort
	cred := dbConfigs.LoginUsername + ":" + dbConfigs.LoginPassword
	return cred + "@tcp(" + host + ")/"
}

func connectDb(dbName string) *sql.DB {
	fmt.Println("... 正在连接数据库")
	db, err := sql.Open("mysql", resolveDsn(&dbConfigs) + dbName)
	if err != nil {
		log.Printf("Open db error: %v", err)
	}
	fmt.Println("*** 数据库已连接 ***")
	return db
}

/*func main() {
        db := connectDb(dbConfigs.DbName)
        defer db.Close()

        fmt.Println(db)
}*/