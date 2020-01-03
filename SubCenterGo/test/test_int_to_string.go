package  main 

import(
   "fmt"
   "strconv"
)

func test1(){
       a := strconv.Itoa(21108)
       fmt.Println("a:%s", string(a))
}

func main(){
	test1()
}
