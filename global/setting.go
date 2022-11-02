package global

import (
	"GdalProject/global/viper"
	"log"
)

// 和yml对应
type ServerSettings struct {
	RunMode  string
	HttpPort string
}

type PostgresDbSettings struct {
	DriverName string
	Host       string
	Port       string
	User       string
	Password   string
	Dbname     string
	Sslmode    string
}

type FileSettings struct {
	Location string
	Num      int
}

// 定义全局变量
var (
	ServerSetting     *ServerSettings
	PostgresDbSetting *PostgresDbSettings
	FileSetting       *FileSettings
)

// 读取配置到全局便量
func SetupSetting() error {
	s, err := viper.NewSetting()
	if err != nil {
		return err
	}
	err = s.ReadSection("Server", &ServerSetting)
	if err != nil {
		return err
	}

	pgdb, err := viper.NewSetting()
	if err != nil {
		return err
	}
	err = pgdb.ReadSection("Postgres", &PostgresDbSetting)
	if err != nil {
		return err
	}

	fl, err := viper.NewSetting()
	if err != nil {
		return err
	}
	err = fl.ReadSection("File", &FileSetting)
	if err != nil {
		return err
	}

	log.Printf("PgsqlSetting:")
	log.Printf("%+v", PostgresDbSetting)

	return nil
}
