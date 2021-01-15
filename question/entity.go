package question

import (
	"os"
	"os/exec"
	"runtime"

	"github.com/amtoaer/goqut/file"
)

var clear func()

func init() {
	switch runtime.GOOS {
	case "windows":
		clear = func() {
			cmd := exec.Command("cmd", "/c", "cls")
			cmd.Stdout = os.Stdout
			cmd.Run()
		}
	default:
		clear = func() {
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()
		}
	}
}

// Variable 题库实体内部的变量
type entity struct {
	subject     file.Subject
	problemList []file.Problem
	history     file.History
}

// Handler 完成程序功能的函数
type Handler interface {
	Init(file.Subject)
	SaveHistory()
	ProblemInOrder()
	ProblemInRandomOrder()
	ProblemForExam()
	ProblemWrongBefore()
}

// NewHandler 获取新的Handler
func NewHandler() Handler {
	return &entity{}
}
