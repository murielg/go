package greetings

import (
    "fmt"
    "errors"
)

func Hello(name string) (string, error) {
    if name == "" {
        return "", errors.New("empty name")
    }
    return fmt.Sprintf("Hi, %v. Welcome!", name), nil
}
