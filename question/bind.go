package question

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/amtoaer/goqut/file"
)

type status int

const (
	// 第一次答对
	correct = iota
	// 第一次答错
	wrong
	// 退出
	quit
)

func (e *entity) Init(sub file.Subject) {
	e.subject = sub
	e.problemList = file.LoadData(sub)
	e.history = file.ReadArchive(sub)
	rand.Seed(time.Now().UnixNano())
}

func (e *entity) SaveHistory() {
	file.WriteArchive(e.subject, e.history)
}

func (e *entity) ProblemInOrder() {
	currentNum := int(e.history.Progress)
	for index, problem := range e.problemList[currentNum:] {
		// 判断是否一次答对
		switch showProblem(problem, currentNum+index+1, len(e.problemList)) {
		case quit:
			fmt.Println("正在退出...")
			os.Exit(0)
		case wrong:
			if !isIn(index, e.history.ErrorProblems) {
				e.history.ErrorProblems = append(e.history.ErrorProblems, index)
			}
		}
		e.history.Progress = currentNum + index + 1
		e.SaveHistory()
	}
	fmt.Println("题目已刷完，是否清空刷题记录以便二刷？(y/n)")
	var answer string
	fmt.Scan(&answer)
	if answer == "y" {
		e.history.Progress = 0
		e.SaveHistory()
		fmt.Println("已清空，正在退出...")
	}
}

func (e *entity) ProblemInRandomOrder() {
	var index int
	for {
		index = rand.Intn(len(e.problemList))
		switch showProblem(e.problemList[index], 0, 0) {
		case quit:
			fmt.Println("正在退出...")
			os.Exit(0)
		case wrong:
			if !isIn(index, e.history.ErrorProblems) {
				e.history.ErrorProblems = append(e.history.ErrorProblems, index)
			}
		}
		e.SaveHistory()
	}
}

func (e *entity) ProblemForExam() {
	var index int

	// 模拟考试分数记录
	var score int
	// 模拟考试错题对应题号记录
	var wrongProblems []int

	for i := 0; i < 50; i++ {
		for {
			index = rand.Intn(len(e.problemList))
			if (i < 40 && len(e.problemList[index].Answer) == 1) || (i >= 40 && len(e.problemList[index].Answer) > 1) {
				break
			}
		}
		switch showProblem(e.problemList[index], i+1, 50) {
		// 第一次答对才加分
		case correct:
			score++
		case quit:
			fmt.Println("保存刷题记录成功，正在退出...")
			os.Exit(0)
		case wrong:
			wrongProblems = append(wrongProblems, index)
			if !isIn(index, e.history.ErrorProblems) {
				e.history.ErrorProblems = append(e.history.ErrorProblems, index)
			}
		}
		e.SaveHistory()
	}

	clear()
	fmt.Println("模拟考试结束，本次的得分是：", score)
	// 打印本次考试错题的题号
	fmt.Println("错题在题库中的编号为：")
	for _, item := range wrongProblems {
		fmt.Printf("     No.%d\n", item+1)
	}

	// 选择是否查看本次模拟的错题
	fmt.Println("是否查看错题？(y 查看 / n 退出程序)")
	var answer, getchar string

	// 循环读取输入，直到输入y或n
	for {
		_, _ = fmt.Scan(&answer)
		switch answer {
		case "y":
			for num, index := range wrongProblems {
				showProblem(e.problemList[index], num+1, len(wrongProblems))
			}
			fmt.Println("本次测试的错题已练习完毕，按回车退出...")
			if isWindows {
				_, _ = fmt.Scanln(&getchar)
			}
			os.Exit(0)
		case "n":
			fmt.Println("模拟考试结束，正在退出...")
			time.Sleep(time.Second * 2)
			os.Exit(0)
		}
		fmt.Println("输入错误，请重新输入(y 查看 / n 退出程序)：")
	}
}

func (e *entity) ProblemWrongBefore() {
	length := len(e.history.ErrorProblems)
	tmp := make([]int, length)
	copy(tmp, e.history.ErrorProblems)
	for num, index := range tmp {
		switch showProblem(e.problemList[int(index)], num+1, length) {
		case quit:
			fmt.Println("保存刷题记录成功，正在退出...")
			os.Exit(0)
		case correct:
			e.history.ErrorProblems = remove(index, e.history.ErrorProblems)
		}
		e.SaveHistory()
	}
	fmt.Println("错题本已刷完，正在退出...")
}
