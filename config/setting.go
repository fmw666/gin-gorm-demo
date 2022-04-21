package config

import (
	"fmt"
	"time"

	"github.com/go-ini/ini"
)

type App struct {
	DefaultPage     string
	DefaultPageSize string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     string
	PrefixUrl    string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type     string
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	Url      string
}

var DatabaseSetting = &Database{}

var cfg *ini.File

func init() {
	var err error
	cfg, err = ini.Load("config/app.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("db", DatabaseSetting)

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second

	DatabaseSetting.Url = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DatabaseSetting.User,
		DatabaseSetting.Password,
		DatabaseSetting.Host,
		DatabaseSetting.Port,
		DatabaseSetting.Name,
	)
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		panic(err)
	}
}
