// Copyright 2017
//Created by rubin on 11/2/2017
// subcenter

package process

import (
    "encoding/json"
    // "fmt"
    // "io"
    "io/ioutil"
    // "log"
    ut "SubCenterGo/src/utils"
    // "net/url"
    //"bytes"
    //"compress/zlib"
    // "crypto/tls"
    // "net"
    "net/http"
    // "strings"
    "time"
    // "io/ioutil"
    // "bufio"
    // log "github.com/Sirupsen/logrus"
    // log "github.com/omidnikta/logrus"
    // "gopkg.in/mgo.v2"
    // "gopkg.in/mgo.v2/bson"
    //"strconv"
)

// var redisHost string = "172.16.21.198:6379" //223.202.52.82
// var cacheDB int64

//是否发送结果 到边缘
// var sendFlag = ut.SendFlag
var log = ut.Logger

var Parallel = ut.Parallel

// var uri = ut.URI
// var maxConcurrent_ch = ut.MaxConcurrent_ch

type ReceiveBody struct {
    // Http_status        int    `json:"http_status"`
    // Is_compressed      bool   `json:"is_compressed"`
    Command         string      `json:"command"`
    Callback        string      `json:"callback"`
    Callback_params []string    `json:"callback_params"`
    Edge_list       []edge_body `json: edge_list`
    Is_compressed   bool        `json:"is_compressed"`
    Send_header     map[string]string `json:"send_header"`
}

type edge_body struct {
    Command    string `json:"command"`
    Target_url string `json:"target_url"`
    Callback   string `json:"callback"`
    Callback_params []string `json:"callback_params"` 
}

/*
 ** the body return to refresh center
 */
type ReturnCenter struct {
    Target_url      string   `json:"target_url"`
    Callback_params []string `json:"callback_params"`
    EdgeReturnBody  string   `json:"edgeReturnBody"`
}

//   *  receiver  data from edge
//   *  temporary receive refresh edge data

// func EdgeTaskReturn(writer http.ResponseWriter, r, *http.Request) {

// }

/*TaskRequestPost 处理实时接收汇报的请求，并异步判断请求
post: [{}]
return:{"msg":"ok"}
*/
func TaskRequestPost(writer http.ResponseWriter, r *http.Request) {

    var body ReceiveBody
    var err error
    var codeStatus int = 200
    res, _ := ioutil.ReadAll(r.Body)
    result, z_err := ut.DoZlibUnCompress(res)
    if z_err != nil {
        //no compress
        result = res
    }

    r.Body.Close()
    log.Debugf("TaskRequestPost request_body: %s", result)

    err = json.Unmarshal(result, &body)

    if err != nil {
        log.Error("TaskRequestPost json.Unmarshal[error]", err)
    }

    log.Debugf("result[Command] :%s", body.Command)
    lengthEdgeList := len(body.Edge_list)
    if lengthEdgeList <= 0 {
        // return the length of edge_list is zero
        codeStatus = 201
    }

    send_header := body.Send_header 
    callback_params := body.Callback_params
    
    for i := 0; i < lengthEdgeList; i++ {
        // the type send to edge
        command := body.Command

        if command == "" {
            command = body.Edge_list[i].Command
        }
        //zlib body
        if body.Is_compressed == true {
            command_byte := []byte(command)
            command = string(ut.DoZlibCompress(command_byte))
        }

        // the path of return
        callback := body.Callback
        if callback == "" {
            callback = body.Edge_list[i].Callback
        }

        if len(callback_params) == 0{
            callback_params = body.Edge_list[i].Callback_params
        }

        // add send url judge
        target_url := body.Edge_list[i].Target_url

        log.Debugf("command:%s, target_url:%s, callback:%s , callback_params:%s", command, target_url, callback, callback_params)

        go SendEdge(command, 3, target_url, callback, callback_params, send_header)
    }

    log.Debugf("TaskRequestPost RemoteAddr: %s, URL.Path: %s, body: %+v, callback_params:%s", r.RemoteAddr, r.URL.Path, body, callback_params)

    ResponseToCenter(writer, codeStatus)
}

/*
   *send task to edge, then send the result to refresh center
   command: the body to egde
   timeOut: the time int
   urlSend: the url to the edge
   returnPath: the url to the refresh center
*/

func SendEdge(command string, timeOut int, targetURL string, callback string, callbackParams []string, sendHeader map[string]string) {
    // send data to channel
    Parallel <- 1

    // var resultEdge map[string]string
    resultEdge := make(map[string]string)
    // if send body to edge failed, try three times  at most
    for i := 0; i < 3; i++ {
        resultEdge = ut.SendHttp(command, timeOut, targetURL, "POST", sendHeader)
        log.Debugf("SendEdge send to edge send times:%s, resultEdge content:%s", i, resultEdge)
        if resultEdge["code"] == "200" {
            break
        }

        time.Sleep(2 * time.Second)
    }

    // build return body
    header_m := make(map[string]string)
    header_m["Content-Type"] = "application/json"
    returnCenter := new(ReturnCenter)
    returnCenter.Target_url = targetURL
    returnCenter.Callback_params = callbackParams
    returnCenter.EdgeReturnBody = resultEdge["body"]
    returnCenterJson, err := json.Marshal(returnCenter)
    if err != nil {
        log.Debugf("SendEdgeAndRefreshCenter ")
        <-Parallel
        return
    }
    // log.Debugf("SendEdgeAndRefreshCenter  resultEdge:%s", resultEdge)
    // send result from edge to refresh center
    // if send body to refresh failed, try three times at most
    log.Debugf("returnCenterJson to refresh center:%s", returnCenterJson)
    for i := 0; i < 3; i++ {
        resultFromCenter := ut.SendHttp(returnCenterJson, timeOut, callback, "POST", header_m)
        log.Debugf("SendEdge  send data to center send times:%s, resultCenter content:%s", i, resultFromCenter)
        if resultFromCenter["code"] == "200" {
            break
        }
        time.Sleep(2 * time.Second)
    }
    // Release the contents of the pipeline
    <-Parallel
}
