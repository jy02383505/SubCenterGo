/*
  获取当前PC的主机名
*/
package main

import (
    "fmt"
    "os"
)

func main() {
    host, err := os.Hostname()
    if err != nil {
        fmt.Println("%s", err)
    } else {
        fmt.Println("%s", host)
    }
}