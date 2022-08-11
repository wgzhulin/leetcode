package main

import (
	"log"
	"testing"
)

func TestGenGoCode(t *testing.T) {
	ltQuestion := FetchLeetCodeQuestionInfo("maximum-number-of-groups-entering-a-competition")
	cnQuestionInfo := NewCnQuestion(ltQuestion)

	log.Println("====== fetch success =====")
	log.Printf("题号：%v\n", cnQuestionInfo.QuestionFrontendId)
	log.Printf("题目：%v\n", cnQuestionInfo.TranslatedTitle)
	log.Println("==========================")
	log.Println("开始生成本地文件")

	GenGoCodeFile(cnQuestionInfo)
	log.Println("已生成")
}
