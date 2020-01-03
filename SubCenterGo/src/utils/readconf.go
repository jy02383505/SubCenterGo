package utils

import (
	// "fmt"
	// "time"

	r "gopkg.in/ini.v1"
	// "gopkg.in/mgo.v2"
	// "gopkg.in/redis.v3"
)

// var cacheDB int64
var SendFlag = true
// var URI = ""
// var DatabaseName = ""
var ServerPort = 9999
// the number of sent http, max parallel  
var Parallel_num = 1000
var Parallel = make(chan int, 1000)
var UrlToCenter = "http://rep.chinacache.com:80/v2/keys/subcenter/"
var Ttl = "60"
var Interval int =    20



func init() {
	cfg, err := r.Load("config.ini")
	if err != nil {
		panic(err)
	}
	loglevel := cfg.Section("log").Key("level").String()
	logfile := cfg.Section("log").Key("filename").String()
	// logfile = fmt.Sprintf("%s/%s", "logs", logfile)
	SetLog(loglevel, logfile)

	// SetRedis(cfg)
	// SetDB(cfg)
	ServerPort, err = cfg.Section("server").Key("port").Int()
	SendFlag, err = cfg.Section("server").Key("send").Bool()
	Parallel_num, err = cfg.Section("server").Key("parallel").Int()
	if err != nil{
		panic(err)
	}else{
		Parallel = make(chan int, Parallel_num)
	}

	UrlToCenter = cfg.Section("register").Key("urlToCenter").String()
	// the expire time of key in etcd
	Ttl = cfg.Section("register").Key("ttl").String()
	// interval time of register to etcd
	Interval, err = cfg.Section("register").Key("interval").Int()
	if err != nil{
		log.Debugf("config not have interval time")
	}


}
