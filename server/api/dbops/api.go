package dbops

import(
	"database/sql"
	"github.com/go-sql-driver/mysql"
)


func AddUserCredential (loginName string, pwd string) error{}

func GetUserCredential (loginName string) (string, error){}

