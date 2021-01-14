package main

import (
	"fmt"

	"github.com/amtoaer/goqut/file"
	"github.com/amtoaer/goqut/question"
	"github.com/amtoaer/goqut/text"
)

func main() {
	var choice int
	var subject file.Subject
	handler := question.NewHandler()
	fmt.Println(`请输入科目：
	1. 马原
	2. 毛概`)
	fmt.Scanf("%d", &choice)
	switch choice {
	case 1:
		subject = file.Mayuan
	case 2:
		subject = file.Maogai
	default:
		panic(text.InputError)
	}
	handler.Init(subject)
	fmt.Println(`请输入功能：
	1. 顺序刷题
	2. 随机刷题
	3. 模拟考试
	4. 错题本`)
	fmt.Scanf("%d", &choice)
	switch choice {
	case 1:
		handler.ProblemInOrder()
	case 2:
		handler.ProblemInRandomOrder()
	case 3:
		handler.ProblemForExam()
	case 4:
		handler.ProblemWrongBefore()
	default:
		panic(text.InputError)
	}
}
