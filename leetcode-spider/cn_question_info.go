package main

import (
	"encoding/json"
	"fmt"
	"github.com/dlclark/regexp2"
	"github.com/zhulinw/leetcode-spider/leetcode"
	"regexp"
	"strings"
)

type CnQuestion struct {
	QuestionFrontendId string
	TranslatedTitle    string
	TitleSlug          string
	TranslatedContent  string
	ExampleTestCases   string
	CodeDefinition     []*leetcode.CodeDefinition
}

func NewCnQuestion(info *leetcode.QuestionInfo) *CnQuestion {
	code := make([]*leetcode.CodeDefinition, 0)
	err := json.Unmarshal([]byte(info.CodeDefinitionStr), &code)
	if err != nil {
		panic(err)
	}

	return &CnQuestion{
		QuestionFrontendId: info.QuestionFrontendId,
		TranslatedTitle:    info.TranslatedTitle,
		TitleSlug:          info.TitleSlug,
		TranslatedContent:  info.TranslatedContent,
		ExampleTestCases:   info.ExampleTestCases,
		CodeDefinition:     code,
	}
}

func (c *CnQuestion) ParseExampleCaseFromTranslatedContent() ([]string, []string) {
	re := regexp.MustCompile("<(\\S*?)[^>]*>.*?|<.*? />")
	content := re.ReplaceAllString(c.TranslatedContent, "")

	reInput := regexp.MustCompile("输入：.*?\\n")
	reOutput := regexp.MustCompile("输出：.*?\\n")

	findInput := reInput.FindAllString(content, 10)
	findOutput := reOutput.FindAllString(content, 10)
	return findInput, findOutput
}

func (c *CnQuestion) GetTranslatedContent() string {
	return c.TranslatedContent
}

// 有问题
func (c *CnQuestion) FormatTranslatedContent() string {
	re := regexp.MustCompile("<(\\S*?)[^>]*>.*?|<.*? />")
	content := re.ReplaceAllString(c.TranslatedContent, "")

	content = strings.ReplaceAll(content, " ", "")
	content = strings.ReplaceAll(content, "&nbsp;", "")
	content = strings.ReplaceAll(content, "\n", "")

	r := regexp2.MustCompile("(.+(?=示例1))", 0)
	match, err := r.FindStringMatch(content)
	if err != nil {
		return ""
	}

	var builder strings.Builder

	builder.WriteString(match.String())
	builder.WriteString("\n\n")

	input, output := c.ParseExampleCaseFromTranslatedContent()
	for i := range input {
		builder.WriteString(fmt.Sprintf("示例%v：\n", i+1))
		builder.WriteString(input[i])
		builder.WriteString(output[i])
		builder.WriteString("\n")
	}

	return builder.String()
}

func (c *CnQuestion) GetCodeDefinition(language string) string {
	r := regexp2.MustCompile("func.*?\\{", 0)
	for i := range c.CodeDefinition {
		if strings.EqualFold(c.CodeDefinition[i].Text, language) {
			m, err := r.FindStringMatch(c.CodeDefinition[i].DefaultCode)
			if err == nil {
				return m.String() + "\n\n}\n"
			}
		}
	}

	return ""
}

func (c *CnQuestion) GetFuncName(language string) string {
	r := regexp2.MustCompile("(?<=func ).*?(?=\\()", 0)

	for i := range c.CodeDefinition {
		if strings.EqualFold(c.CodeDefinition[i].Text, language) {
			m, err := r.FindStringMatch(c.CodeDefinition[i].DefaultCode)
			if err == nil {
				return 	TitleSlugFirstUpper(m.String())
			}
		}
	}

	return ""
}

func (c *CnQuestion) GenerateTestBodyGolang() string {
	//inputList, outputList := c.ParseExampleCaseFromTranslatedContent()
	//
	//r := regexp2.MustCompile("ss", 0)
	//regexp2FindAllString()
	//match, err := r.FindStringMatch("a")
	//match.Groups()
	//inputTestCases := strings.Split(c.ExampleTestCases, "\n")
	//_, outputTestCases := c.ParseExampleCaseFromTranslatedContent()
	//
	return ""
}

func regexp2FindAllString(re *regexp2.Regexp, s string) []string {
	var matches []string
	m, _ := re.FindStringMatch(s)
	for m != nil {
		matches = append(matches, m.String())
		m, _ = re.FindNextMatch(m)
	}
	return matches
}

func TitleSlugFirstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}