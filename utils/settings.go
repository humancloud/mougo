package utils

import (
	"log"

	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbName     string
	DbPassWord string
)

func InitConfig() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		log.Fatalln("配置文件读取错误，检查文件路径" + err.Error())
	}

	LoadServer(file)
	LoadDb(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug") // the default value is "debug"
	HttpPort = file.Section("server").Key("HttpPort").MustString(":8888")
}

func LoadDb(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbName = file.Section("database").Key("DbName").MustString("mougo")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("mougo")
}
