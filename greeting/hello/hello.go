package main

import (
    "fmt"
    "gonzo/greetings"
)

func main(){
    message := greetings.Hello("Muriel")
    fmt.Println(message)
}


