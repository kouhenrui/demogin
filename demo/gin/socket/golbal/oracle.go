package golbal

import (
	"database/sql"
	//_ "github.com/go-goracle/goracle"
)

/**
 * @ClassName oracle
 * @Description TODO
 * @Author khr
 * @Date 2023/6/13 10:28
 * @Version 1.0
 */

var (
	OracleClient *sql.DB
	err          error
)

func initOracle() {
	//OracleClient, err = sql.Open("goracle", "user/password@//hostname:port/service_name")
	//if err != nil {
	//	panic(err)
	//}
}
