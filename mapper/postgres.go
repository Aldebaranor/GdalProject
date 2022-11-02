package mapper

import (
	"GdalProject/global"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
)

var SqlSession *gorm.DB

// 连接数据库
func ConnectDB() (err error) {
	DriverName := global.PostgresDbSetting.DriverName
	Host := global.PostgresDbSetting.Host
	Port := global.PostgresDbSetting.Port
	User := global.PostgresDbSetting.User
	Password := global.PostgresDbSetting.Password
	Dbname := global.PostgresDbSetting.Dbname
	Sslmode := global.PostgresDbSetting.Sslmode
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", Host, Port, User, Password, Dbname, Sslmode)
	SqlSession, err = gorm.Open(DriverName, psqlInfo)
	if err != nil {
		log.Println("connectDBError", err)
	}
	return SqlSession.DB().Ping()
}

func Close() {
	SqlSession.Close()
}
