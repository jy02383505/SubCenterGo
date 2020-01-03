package main
import(
      "fmt"
      // "bytes"
      "net"
      "time"
      "encoding/json"
      "net/http"
      "net/url"
      "io/ioutil"
      "bytes"

)


func main(){


	// GetProxy("127.0.0.1:9999", "/")

  str_t := "\u003c?xml version=\"1.0\" encoding=\"utf-8\"?\u003e\u003cmethod name=\"url_purge\" sessionid=\"dded8096b7d611e7b4eb000c2946bb3e\"\u003e\u003crecursion\u003e0\u003c/recursion\u003e\u003curl_list\u003e\u003curl id=\"59edba281d41c875c6269924\"\u003ehttps://secure.massmotionmedia.com/common/2.html\u003c/url\u003e\u003curl id=\"59edba281d41c8cd31e787bc\"\u003ehttps://secure.massmotionmedia.com:443/common/2.html\u003c/url\u003e\u003c/url_list\u003e\u003c/method\u003e"
  SendHttp(str_t, 3, "http://175.6.18.76:21108/")
}







func GetProxy(ip_to string, url_t string){


	  tcpAddr, err := net.ResolveTCPAddr("tcp4", ip_to)
	  if err != nil {
	  	fmt.Printf("get Proxy err:%s", err.Error())
	  }

      tcpconn, err2 := net.DialTCP("tcp", nil, tcpAddr);
      if err2 != nil{
      	fmt.Printf("err2:%s", err2.Error())
      	// receiveBody_t.Status = "failed"
        return
      }
        // host, _ := GetHost(receiveBody_t.Url)
	    // if len_str >= 3 {
	    //     req.Header.Set("Host", host)
	    //     }else{
	    //         log.Debugf("host,%s" ,host)
	    //     }
         // host, _ := ""
         str_t := "GET " + url_t + " HTTP/1.1" +  "\r\n\r\n"
         fmt.Printf("str_t:%s\n", str_t)
      	 //向tcpconn中写入数据
        _, err3 := tcpconn.Write([]byte(str_t));
        if err3 != nil{
        	fmt.Printf("err3:%s\n", err3)
	      	// receiveBody_t.Status = "failed"
	       //  go PostRealBody(receiveBody_t)
	       //  <- maxConcurrent_ch
	        return
        }
        buf := make([]byte, 64)
        tcpconn.SetDeadline(time.Now().Add(time.Duration(6) * time.Second))
        for{

        	length, err := tcpconn.Read(buf)
        	if err != nil{
        		tcpconn.Close()
        		fmt.Printf("tcp conn err:%s\n", err)
		        break
        	}
        	Data := buf[:length]
        	// messnager := make(chan byte)
        	if length > 0{  
            // buf[lenght]=0 
               fmt.Printf("length:%s\n", length) 
               fmt.Printf("content:%s\n", Data)
               tcpconn.SetDeadline(time.Now().Add(time.Duration(6) * time.Second))
            }
        //心跳计时  
        // go HeartBeating(tcpconn,messnager, 6)  
        //检测每次Client是否有数据传来  
        // go GravelChannel(Data,messnager)  
        //fmt.Println("Rec[",conn.RemoteAddr().String(),"] Say :" ,string(buf[0:lenght]))  
        // reciveStr :=string(buf[0:lenght]) 
        // if reciveStr != nil {}
        }
        fmt.Printf("post end, url_t:%s\n", url_t)
}



/*
 *  this function is use to send the http request and return the result of the request 
*/
func SendHttp(httpSendBody string, timeOut int, urlDst string) (map[string]string){
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
    httpSendBodyBytes := []byte(httpSendBody)
    postBytesReader := bytes.NewReader(httpSendBodyBytes)
    client := &http.Client{Timeout: timeOut_send}
    req, err := http.NewRequest("POST", urlDst, postBytesReader)
    if (err != nil) {
      // log.Debugf("NewRequest error:%s", err)
      fmt.Println("NewRequest error:%s", err)
      fmt.Println(json.Marshal(map[string]string{"code": "701", "body": "http.NewRequest error"}))
      // return json.Marshal(map[string]string{"code": "701", "body": "http.NewRequest error"})
    }
    // judge success or not, if you don't send the header below
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    // req.Header.Set("Content-Type", "application/json")
    resp, err := client.Do(req)
    if (err != nil) {
      // log.Debugf("SendHttp  client.Do error:%s", err)
      fmt.Println(map[string]string{"code": "702", "body": "client.Do error"})
      // return json.Marshal(map[string]string{"code": "702", "body": "client.Do error"})
    }
    defer resp.Body.Close()

    // read body from edge, after send post to edge
    body, err := ioutil.ReadAll(resp.Body)
    if (err != nil) {
      // log.Debugf("SendHttp ioutil.ReadAll error:%s", err)
      fmt.Println(map[string]string{"code": "703", "body": "ioutil.ReadAll error"})
      // return json.Marshal(map[string]string{"code": "703", "body": "ioutil.ReadAll error"})
    } 
    fmt.Println(map[string]string{"code": "200", "body": string(body)})
    // log.Debugf(string(body))
    // log.Debugf(map[string]string{"code": "200", "body": string(body)})
    // return 0
    
    return map[string]string{"code": "200", "body": string(body)}
}



type ReceiveBody struct {
  // Http_status        int    `json:"http_status"`

  // Is_compressed      bool   `json:"is_compressed"`
  Command string `json:"command"`
  Return_path string `json:"return_path"`

    Edge_list   []edge_body   `json: edge_list`


}

type edge_body struct {
  Host string  `json:"host"`
  Port int  `json:"port"`
  Command string `json:"command"`
  UrlSend string `json:"urlsend"`
  Return_path string `json:"return_path"`

}

//  simulate the central, sub central task 

func  SendBodyToSubCenter(){
    
}


//心跳计时，根据GravelChannel判断Client是否在设定时间内发来信息  
// func HeartBeating(conn net.Conn, readerChannel chan byte,timeout int) {  
//         select {  
//         case fk := <-readerChannel:  
//             fmt.Printf(conn.RemoteAddr().String(), "receive data string:", string(fk))  
//             conn.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Second))  
//             //conn.SetReadDeadline(time.Now().Add(time.Duration(5) * time.Second))  
//             break  
//         case <-time.After(time.Second*2):  
//             fmt.Printf("It's really weird to get Nothing!!!\n")  
//             conn.Close()  
//         }  
  
// } 

// func GravelChannel(n []byte,mess chan byte){  
//     for i , v := range n{ 
//         fmt.Printf("%s byte  v:%s\n", i, v) 
//         mess <- v  
//     }  
//     close(mess)  
// }  