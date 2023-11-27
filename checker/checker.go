package checkher

import (
    "fmt"
    "strings"
)

type FileChecker struct {
}

func NewFileChecker() FileChecker {
    return &FileChecker{}
}

func (cFileChecker) Check(path string) {
    if strings.Contains(path, " ") {
        fmt.Println(false)
    } else {
        fmt.Println(true)
    }
}

type LogChecker struct {
}

func NewLogChecker() LogChecker {
    return &LogChecker{}
}

func (cLogChecker) Check(log string) {
    if strings.Contains(log, "MESSAGE") {
        fmt.Println(true)
    } else {
        fmt.Println(false)
    }
}