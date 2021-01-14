package question

import (
	"runtime"

	"github.com/amtoaer/goqut/file"
)

var clearScreen string

func init() {
	switch runtime.GOOS {
	case "windows":
		clearScreen = "cls"
	default:
		clearScreen = "clear"
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
