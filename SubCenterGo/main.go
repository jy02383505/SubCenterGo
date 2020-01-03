// Copyright 201７
//Created by longjun.zhao on 5/19/2016
// 预加载

package main

import (
	rec "SubCenterGo/src/process"
	ut "SubCenterGo/src/utils"
	"fmt"
	"net/http"
	_ "net/http/pprof"

	// "gopkg.in/ini.v1"
	// log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"time"
	// "strconv"
)

var serverPort = ut.ServerPort
var log = ut.Logger
var interval int = ut.Interval



func main() {
	fmt.Println("interval:", interval)
	var interval1 int = 20
	ticker :=  time.NewTicker(time.Duration(interval1) * time.Second)
	log.Debugf("start subcenter, register subcenter first time")
	rec.RegisterSubcenter() 
	go func(){
		     for _ = range ticker.C {
		     	rec.RegisterSubcenter()
		     }
	}()
	
	fmt.Println("ticked at %v", time.Now())
	router := mux.NewRouter()
	router.HandleFunc("/", rec.TaskRequestPost).Methods("POST")
	// router.HandleFunc("/forward", rec.EdgeTaskReturn).Methods("POST")
	log.Printf("Starting on port: %d", serverPort)
	// log.Printf("maxConcurrent_ch:%s", maxConcurrent_ch)
	fmt.Println("listen on port %s", serverPort)
	http.ListenAndServe(fmt.Sprintf(":%d", serverPort), router)
}
