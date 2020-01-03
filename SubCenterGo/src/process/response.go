//subcenter 
// create time 2017/20/25
// author  rubin
package process
import(
    // "fmt"
    // ut "SubCenterGo/src/utils"
    "net/http"
    "encoding/json"
)

// var log = ut.Logger


// return something to refresh center
func ResponseToCenter(writer http.ResponseWriter, codeStatus int){
	var err error
	writer.WriteHeader(200)
	var param = make(map[string]interface{})
	if (codeStatus == 200) {
          param["msg"] = "ok"
	} else if (codeStatus == 201){
		param["msg"] = "the length of edge_list is zero"
	}
	log.Debugf("return body:%s", param)
	msg, _ := json.Marshal(param)
	_, err = writer.Write([]byte(msg))
	if err != nil {
		http.Error(writer, "Interal ERROR: ", 500)
	}
}