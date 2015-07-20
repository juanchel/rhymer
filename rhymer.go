package main

import (
    "fmt"
    "io/ioutil"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    fmt.Println("Reading dictionary")
    dat, err := ioutil.ReadFile("cmudict/cmudict")
    check(err)
    fmt.Println(string(dat))
}