package leetcode

import (
	"fmt"
	"strings"
)

type QuestionData struct {
	Data *Question `json:"data"`
}

type Question struct {
	Question *QuestionInfo `json:"question"`
}

type QuestionInfo struct {
	QuestionFrontendId string `json:"questionFrontendId"`
	Title              string `json:"title"`
	TranslatedTitle    string `json:"translatedTitle"`
	TitleSlug          string `json:"titleSlug"`
	Content            string `json:"content"`
	TranslatedContent  string `json:"translatedContent"`
	CodeDefinitionStr  string `json:"codeDefinition"`
	Difficulty         string `json:"difficulty"`
	Likes              int    `json:"likes"`
	Dislikes           int    `json:"dislikes"`
	ExampleTestCases   string `json:"exampleTestcases"`

	CodeDefinition []*CodeDefinition
}

func (q QuestionInfo) String() string {
	var builder strings.Builder

	builder.WriteString(fmt.Sprintf("QuestionFrontendId: %v\n", q.QuestionFrontendId))
	builder.WriteString(fmt.Sprintf("Title: %v\n", q.Title))
	builder.WriteString(fmt.Sprintf("TranslatedTitle: %v\n", q.TranslatedTitle))
	builder.WriteString(fmt.Sprintf("TitleSlug: %v\n", q.TitleSlug))
	builder.WriteString(fmt.Sprintf("Content: %v\n", q.Content))
	//builder.WriteString(fmt.Sprintf("TranslatedContent: %v\n", q.TranslatedContent))
	//builder.WriteString(fmt.Sprintf("CodeDefinitionStr: %v\n", q.CodeDefinitionStr))
	builder.WriteString(fmt.Sprintf("Difficulty: %v\n", q.Difficulty))
	builder.WriteString(fmt.Sprintf("Likes: %v\n", q.Likes))
	builder.WriteString(fmt.Sprintf("Dislikes: %v\n", q.Dislikes))
	return builder.String()
}
