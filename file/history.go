package file

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"

	"github.com/amtoaer/goqut/text"
	"github.com/amtoaer/goqut/utils"
	"github.com/mitchellh/go-homedir"
)

// History 存储顺序刷题的进度和错误题目的题号
type History struct {
	Progress      int
	ErrorProblems []int
}

var defaultHistory = &History{
	Progress:      0,
	ErrorProblems: []int{},
}

func writeInitArchive(path string) {
	content, _ := json.Marshal(defaultHistory)
	ioutil.WriteFile(path, content, os.FileMode(0777))
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func getArchiveDir() string {
	path, err := homedir.Expand("~/.config/goqut")
	utils.CheckError(err, text.HomeError)
	return path
}

func getArchiveFileDir(sub Subject) string {
	var fileName string
	switch sub {
	case Mayuan:
		fileName = "mayuan.archive"
	case Maogai:
		fileName = "maogai.archive"
	default:
		panic(text.SubjectError)
	}
	path := path.Join(getArchiveDir(), fileName)
	return path
}

// WriteArchive 将历史写入对应文件
func WriteArchive(sub Subject, history History) {
	path := getArchiveFileDir(sub)
	content, _ := json.Marshal(history)
	ioutil.WriteFile(path, content, os.FileMode(0644))
}

// ReadArchive 读取指定科目的历史文件
func ReadArchive(sub Subject) History {
	path := getArchiveFileDir(sub)
	if !isExist(path) {
		os.MkdirAll(getArchiveDir(), os.FileMode(0755))
		writeInitArchive(path)
	}
	content, err := ioutil.ReadFile(path)
	utils.CheckError(err, text.ReadError)
	history := History{}
	err = json.Unmarshal(content, &history)
	utils.CheckError(err, text.ParseArchiveError)
	return history
}
