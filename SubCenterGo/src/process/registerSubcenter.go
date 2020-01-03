package process

import(

    "fmt"
    ut "SubCenterGo/src/utils"
    "os"
    "net"
    // "fmt"
    "strings"
)

var urlToCenter = ut.UrlToCenter
var ttl = ut.Ttl
// var log = ut.Logger

func RegisterSubcenter(){
     method := "PUT"
     header_m := map[string]string{"Content-Type": "application/x-www-form-urlencoded"}
     host :=  getHostName()
     ipOfHost := getIp()
     urlToCenterT := urlToCenter + host
     body_send := "ttl=" + ttl + "&value=" + ipOfHost + ":21109"
     //log.Debugf("host:%s, ip:%s, urlToCenter:%s, body_send:%s", host, ipOfHost, urlToCenter, body_send)
     ut.SendHttp(body_send, 3, urlToCenterT, method, header_m)
}


// get the name of the host
func getHostName() (string){
     host, err := os.Hostname()
     if err != nil {
        fmt.Println("%s", err)
        log.Error("get os.HostName  error ", err)
    } else {
        //fmt.Println("%s", host)
        log.Debugf("get os.HostName success host:", host)
    }
    return host
}

// get the ip the host
func getIp() (string){
	conn, err := net.Dial("udp", "google.com:80")
    if err != nil {
	    fmt.Println(err.Error())
	    log.Debugf("getIp udp error ", err.Error())
	    return ""
    }
    defer conn.Close()
    // fmt.Println("conn.LocalAddr().String(),", conn.LocalAddr().String())
    //log.Debugf("conn.LocalAddr().String(),", conn.LocalAddr().String())
    // fmt.Println(strings.Split(conn.LocalAddr().String(), ":")[0])
    return strings.Split(conn.LocalAddr().String(), ":")[0]
}