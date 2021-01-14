package file

import (
	"encoding/json"
	"io/ioutil"

	// 题库文件
	_ "github.com/amtoaer/goqut/statik"
	"github.com/amtoaer/goqut/text"
	"github.com/amtoaer/goqut/utils"
	"github.com/rakyll/statik/fs"
)

// Subject 代表科目
type Subject int

const (
	// Mayuan 马原
	Mayuan = iota
	// Maogai 毛概
	Maogai
)

// Problem 题目结构体
type Problem struct {
	Description string
	Answer      string
	Choice      []string
}

// LoadData 通过科目获取题目切片
func LoadData(sub Subject) []Problem {
	db, _ := fs.New()
	var subject string
	switch sub {
	case Mayuan:
		subject = "/mayuan.json"
	case Maogai:
		subject = "/maogai.json"
	default:
		panic(text.SubjectError)
	}
	file, _ := db.Open(subject)
	defer file.Close()
	questions, _ := ioutil.ReadAll(file)
	var result []Problem
	err := json.Unmarshal(questions, &result)
	utils.CheckError(err, text.ParseProblemError)
	return result
}
