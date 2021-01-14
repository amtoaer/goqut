package question

import (
	"fmt"

	"os"
	"os/exec"

	"strings"

	"github.com/amtoaer/goqut/file"
)

func showProblem(problem file.Problem, begin, end int) status {
	var stats status = correct
	var answer, getchar string
	for {
		clear()
		fmt.Printf("%d/%d\n", begin, end)
		if len(problem.Answer) > 1 {
			fmt.Println("注意：此题为多选题")
		}
		fmt.Println(problem.Description)
		for _, item := range problem.Choice {
			fmt.Println(item)
		}
		if stats != correct {
			fmt.Println("答案错误")
		}
		fmt.Printf("请输入答案：")
		fmt.Scanf("%s", &answer)
		if strings.ToUpper(answer) == problem.Answer {
			break
		} else if answer == "quit" {
			return quit
		} else if answer == "ans" {
			fmt.Printf("答案为%s，回车继续...\n", problem.Answer)
			fmt.Scanln(&getchar)
			// 直接看答案当错误处理
			return wrong
		}
		stats = wrong
	}
	return stats
}

func isIn(num int, slice []float64) bool {
	for _, item := range slice {
		if num == int(item) {
			return true
		}
	}
	return false
}

func remove(num float64, slice []float64) []float64 {
	for index, value := range slice {
		if value == num {
			return append(slice[:index], slice[index+1:]...)
		}
	}
	return slice
}

func clear() {
	cmd := exec.Command(clearScreen)
	cmd.Stdout = os.Stdout
	cmd.Run()
}
