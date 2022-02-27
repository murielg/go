package main

import (
    "fmt"
    "log"

    "gonzo/greetings"
)

func main() {
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    message, err := greetings.Hello("")
    // message, err := greetings.Hello("Muriel")

    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println(message)
}


