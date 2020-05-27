package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Cfg *ini.File

	RunMode string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize  int
	JwtSecret string
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatal(2, "fail to parse 'conf/app.ini': %v", err)
	}

	LoadBase()
	LoadApp()
	LoadServer()
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("run_mode").MustString("debug")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatal(2, "faile to get sectoin 'server': %v", err)
	}

	HTTPPort = sec.Key("http_port").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("read_timeout").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("write_timeout").MustInt(60)) * time.Second
}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatal(2, "fail to get section 'app': %v", err)
	}

	JwtSecret = sec.Key("jwt_secret").MustString("!!!")
	PageSize = sec.Key("page_size").MustInt(10)
}