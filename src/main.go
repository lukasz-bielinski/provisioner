package main                        //[1]

import (
    "fmt"
    "hello"
    "git-sync"
)



func main() {                       //[3]
    fmt.Println("Witaj świecie!")
    hello.Hi()
    gitsync.Sync()
}
