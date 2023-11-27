package main

import checker "github.com/Sddilora/DecoratorPattern/checker"

func main() {

	file := File{
		Path:      "asdasd",
		IsFile:    true,
		CheckFunc: CheckFuncHandler(true),
	}

	log := Log{
		Log:       "asda sd",
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
