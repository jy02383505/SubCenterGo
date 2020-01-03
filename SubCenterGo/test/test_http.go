// package utils
package main
import (
     "encoding/json"
     "net/http"
     // ut "SubCenterGo/src/utils"
     "fmt"
     "net/url"
     "time"
     "io/ioutil"
     "bytes"
     // "reflect"
)

// var log = Logger
/*
 *  this function is use to send the http request and return the result of the request 
*/
func SendHttp(httpSendBody interface{}, timeOut int, urlDst string) (map[string]string){
	fmt.Printf("httpSendBody:%s\n", httpSendBody)
    fmt.Printf("timeOut:%s\n", timeOut)
    fmt.Printf("urlDst:%s\n", urlDst)
    // client :=&http.Client{};
    // fmt.Printf(":%s\n", client)
    uri, err := url.Parse(urlDst)
    fmt.Printf("error,%s, uri:%s\n", err, uri)
    // log.Debugf("error, %s", err)
    // fmt.Printf("timeOut %s", timeOut)
    timeOut_send := time.Duration(timeOut) * time.Second
    var httpSendBodyBytes []byte
    // if reflect.TypeOf(httpSendBody).String() == "string"{
    //     httpSendBodyBytes := []byte(httpSendBody)
    // }else{
    //     httpSendBodyBytes = httpSendBody
    // }
    httpSendBody_t, ok := httpSendBody.(string)
    if ok {
        httpSendBodyBytes = []byte(httpSendBody_t)
    }
    httpSendBody_t1, ok := httpSendBody.([]byte)
    if ok {
        httpSendBodyBytes = httpSendBody_t1
    }
    fmt.Println("httpSendBodyBytes:%s", httpSendBodyBytes)
      
    

    
    postBytesReader := bytes.NewReader(httpSendBodyBytes)
    client := &http.Client{Timeout: timeOut_send}
    req, err := http.NewRequest("POST", urlDst, postBytesReader)
    if (err != nil) {
    	// log.Debugf("NewRequest error:%s", err)
    	fmt.Println("NewRequest error:%s", err)
    	// format the err
    	errorTemp := fmt.Sprintf("%s", err)
    	fmt.Println(map[string]string{"code": "701", "body": "http.NewRequest error" + errorTemp})
    	// log.Debugf("%s", map[string]string{"code": "701", "body": "http.NewRequest error" + errorTemp})
    	return map[string]string{"code": "701", "body": "http.NewRequest error" + errorTemp}
    }
    // judge success or not, if you don't send the header below
    // req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    req.Header.Set("Content-Type", "application/json")
    resp, err := client.Do(req)
    if (err != nil) {
    	// log.Debugf("SendHttp  client.Do error:%s", err)
    	errorTemp := fmt.Sprintf("%s", err)
    	fmt.Println(map[string]string{"code": "702", "body": "client.Do error" + errorTemp})
    	// log.Debugf("%s", map[string]string{"code": "702", "body": "client.Do error" + errorTemp})
    	return map[string]string{"code": "702", "body": "client.Do error" + errorTemp}
    }
    defer resp.Body.Close()

    // read body from edge, after send post to edge
    body, err := ioutil.ReadAll(resp.Body)
    if (err != nil) {
    	// log.Debugf("SendHttp ioutil.ReadAll error:%s", err)
    	errorTemp := fmt.Sprintf("%s", err)
    	fmt.Println(map[string]string{"code": "703", "body": "ioutil.ReadAll error"})
    	// log.Debugf("%s", map[string]string{"code": "703", "body": "ioutil.ReadAll error" + errorTemp})
    	return map[string]string{"code": "703", "body": "ioutil.ReadAll error" + errorTemp}
    } 
    fmt.Println(map[string]string{"code": "200", "body": string(body)})
    // log.Debugf(map[string]string{"code": "200", "body": string(body)})
    // return 0
    
    return map[string]string{"code": "200", "body": string(body)}
}

// func SendHttp(urlDst string){
//     fmt.Printf("urlDst:%s", urlDst)


// }

func test_map(){
	map_t := map[string]string{"nihao": "234"}
	fmt.Println(json.Marshal(map_t))
}

type test_struct struct {
    Name string `json:name`
}

func main() {
    // build return body 
    returnRefreshCenter := new(test_struct)
    returnRefreshCenter.Name = "nihao"
   
    returnRefreshCenterJson, err:= json.Marshal(returnRefreshCenter)
    if err != nil {
        // log.Debugf("SendEdgeAndRefreshCenter ")
        
        return
    }
    fmt.Println("returnRefreshCenterJson:%s", returnRefreshCenterJson)

	result := SendHttp(returnRefreshCenterJson, 3, "http://127.0.0.1/receiveSubCenterResult_new");
	// result := SendHttp("nihao", 3, "http://127.0.0.1/receiveSubCenterResult_new");
	fmt.Println("result:%s", result)
	// SendHttp("1234")
	// test_map()
	// SendHttp("nihao", 3, "http://www.baidu.com/")
	fmt.Printf("nihao\n")
}