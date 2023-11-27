package main

import checker "github.com/husamettinarabaci/testproject/checker"

func main() {

    file := File{
        Path:      "asdasd",
        IsFile:    true,
        CheckFunc: CheckFuncHandler(true),
    }

    log := Log{
        Log:       "asdasd",
        IsFile:    false,
        CheckFunc: CheckFuncHandler(false),
    }

    file.CheckFunc(file.Path)
    log.CheckFunc(log.Log)

}

func CheckFuncHandler(isFile bool) func(string) {
    if isFile {
        ch := checker.NewFileChecker()
        return func(value string) {
            ch.Check(value)
        }
    } else {
        ch := checker.NewLogChecker()
        return func(value string) {
            ch.Check(value)
        }
    }
}

type File struct {
    Path      string
    IsFile    bool
    CheckFunc func(string)
}

type Log struct {
    Log       string
    IsFile    bool
    CheckFunc func(string)
}
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