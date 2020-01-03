package main

import(
    "fmt"
    "reflect"
    "encoding/json"
)

type person struct {
    name string
    age string
}

var P person


func main() {
    P.name="a"
    P.age= "1234"
    test, _:= json.Marshal(P)
    strint_t := "1234"
    fmt.Println("The person's name is %s", reflect.TypeOf(test).String())
    fmt.Println("The person's name is %v", P)
    fmt.Println("The person's name is %s", reflect.TypeOf(strint_t).String() == "string")
    if reflect.TypeOf(strint_t).String() == "string"{
    	fmt.Println("string")
    }
}
