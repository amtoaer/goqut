package question

import (
	"fmt"
	"math/rand"
	"os"

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
}

func (e *entity) SaveHistory() {
	file.WriteArchive(e.subject, e.history)
}

func (e *entity) ProblemInOrder() {
	currentNum := int(e.history.Progress)
	for index, problem := range e.problemList[currentNum:] {
		// 判断是否一次答对
		switch showProblem(problem, index+1, len(e.problemList)+1) {
		case quit:
			e.SaveHistory()
			fmt.Println("保存刷题记录成功，正在退出...")
			os.Exit(0)
		case wrong:
			if !isIn(index, e.history.ErrorProblems) {
				e.history.ErrorProblems = append(e.history.ErrorProblems, float64(index))
			}
		}
		e.history.Progress = float64(index)
	}
	fmt.Println("题目已刷完，是否清空以便二刷？(y/n)")
	var answer string
	fmt.Scanf("%s", &answer)
	if answer == "y" {
		e.history.Progress = 0
		fmt.Println("已清空，正在退出...")
	}
	e.SaveHistory()
	os.Exit(0)
}

func (e *entity) ProblemInRandomOrder() {
	var index int
	for {
		index = rand.Intn(len(e.problemList))
		switch showProblem(e.problemList[index], 0, 0) {
		case quit:
			e.SaveHistory()
			fmt.Println("保存刷题记录成功，正在退出...")
			os.Exit(0)
		case wrong:
			if !isIn(index, e.history.ErrorProblems) {
				e.history.ErrorProblems = append(e.history.ErrorProblems, float64(index))
			}
		}
	}
}

func (e *entity) ProblemForExam() {
	var index int
	for i := 0; i < 50; i++ {
		for {
			index = rand.Intn(len(e.problemList))
			if (i < 40 && len(e.problemList[index].Answer) == 1) || (i >= 40 && len(e.problemList[index].Answer) > 1) {
				break
			}
		}
		switch showProblem(e.problemList[index], i+1, 50) {
		case quit:
			e.SaveHistory()
			fmt.Println("保存刷题记录成功，正在退出...")
			os.Exit(0)
		case wrong:
			if !isIn(index, e.history.ErrorProblems) {
				e.history.ErrorProblems = append(e.history.ErrorProblems, float64(index))
			}
		}
	}
}

func (e *entity) ProblemWrongBefore() {
	length := len(e.history.ErrorProblems)
	tmp := make([]float64, length)
	copy(tmp, e.history.ErrorProblems)
	for num, index := range tmp {
		switch showProblem(e.problemList[int(index)], num+1, length) {
		case quit:
			e.SaveHistory()
			fmt.Println("保存刷题记录成功，正在退出...")
			os.Exit(0)
		case correct:
			e.history.ErrorProblems = remove(index, e.history.ErrorProblems)
		}
	}
	fmt.Println("错题本已刷完，正在退出...")
	e.SaveHistory()
	os.Exit(0)
}
