package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/zhulinw/leetcode-spider/leetcode"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		log.Println("input: ", input)
		titleSlug := formatScanInput(input)
		if !pingLeetCodeQuestionInfo(titleSlug) {
			log.Println("input fail")
			continue
		}
		ltQuestion := FetchLeetCodeQuestionInfo(titleSlug)
		cnQuestionInfo := NewCnQuestion(ltQuestion)

		log.Println("====== fetch success =====")
		log.Printf("题号：%v\n", cnQuestionInfo.QuestionFrontendId)
		log.Printf("题目：%v\n", cnQuestionInfo.TranslatedTitle)
		log.Println("==========================")
		log.Println("开始生成本地文件")

		GenGoCodeFile(cnQuestionInfo)
		log.Println("已生成")
	}
}

func formatScanInput(a string) string {
	split := strings.Split(a, "/")
	for i := len(split) - 1; i >= 0; i-- {
		if len(split[i]) != 0 {
			return split[i]
		}
	}
	return a
}
func pingLeetCodeQuestionInfo(titleSlug string) bool {
	_, err := http.Get(leetcode.ProblemUrlPrefix + titleSlug + "/")
	if err != nil {
		return false
	}
	return true
}

func FetchLeetCodeQuestionInfo(titleSlug string) *leetcode.QuestionInfo {
	postBody := "query{\n" + "  question(titleSlug:\"" + titleSlug + "\") {\n" +
		"    questionFrontendId\n" +
		"    translatedTitle\n" +
		"    titleSlug\n" +
		"    difficulty\n" +
		"    codeDefinition\n" +
		"    translatedContent\n" +
		"  }\n" +
		"}\n"

	var body io.Reader = strings.NewReader(postBody)
	req, err := http.NewRequest(http.MethodPost, leetcode.QueryQuestionUrl, body)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Content-Type", "application/graphql")
	req.Header.Add("Referer", "https://leetcode-cn.com/")
	req.Header.Add("Cookie", "csrftoken=20c1c263bbe7e65c916abd02706b8ffd442599e3c2c1443f15a1a8b89ff20ff0")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	ioData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	q := &leetcode.QuestionData{}
	err = json.Unmarshal(ioData, &q)
	if err != nil {
		panic(err)
	}

	code := make([]*leetcode.CodeDefinition, 0)
	err = json.Unmarshal([]byte(q.Data.Question.CodeDefinitionStr), &code)
	if err != nil {
		panic(err)
	}
	q.Data.Question.CodeDefinition = code

	log.Println(q.Data.Question)

	return q.Data.Question
}

func GenGoCodeFile(q *CnQuestion) {
	title := strings.ReplaceAll(q.TitleSlug, "-", "_")
	dir := fmt.Sprintf("p%v-%v", q.QuestionFrontendId, q.TitleSlug)
	os.Mkdir(dir, os.ModePerm)

	var builder strings.Builder
	mdName := fmt.Sprintf("%v/p%v_%v.md", dir, q.QuestionFrontendId, title)
	mdFile, _ := os.Create(mdName)
	builder.WriteString(fmt.Sprintf("#%v. %v\n\n", q.QuestionFrontendId, q.TranslatedTitle))
	builder.WriteString("##题目描述\n")
	builder.WriteString(q.GetTranslatedContent())
	builder.WriteString("\n##题解\n")
	mdFile.WriteString(builder.String())

	// go
	fileName := fmt.Sprintf("%v/p%v_%v.go", dir, q.QuestionFrontendId, title)
	file, _ := os.Create(fileName)

	builder.Reset()
	builder.WriteString(fmt.Sprintf("package %v\n\n", title))

	codeDefinition := q.GetCodeDefinition("go")
	if isContainStruct(codeDefinition) {
		builder.WriteString(fmt.Sprintf("import . \"github.com/zhulinw/leetcode/ltdata\"\n"))
	}

	builder.WriteString("/*\n")
	builder.WriteString(fmt.Sprintf("%v. %v\n\n", q.QuestionFrontendId, q.TranslatedTitle))
	builder.WriteString(leetcode.CopyRight(q.TitleSlug))
	builder.WriteString("*/\n")
	builder.WriteString(codeDefinition)
	file.WriteString(builder.String())

	// go test
	testFile, _ := os.Create(fmt.Sprintf("%v/p%v_%v_test.go", dir, q.QuestionFrontendId, title))

	builder.Reset()
	builder.WriteString(fmt.Sprintf("package %v\n\n", title))
	builder.WriteString(fmt.Sprintf("%v\n\n", "import \"testing\""))
	builder.WriteString(fmt.Sprintf("func Test%v(t *testing.T) {\n%v\n}\n",
		q.GetFuncName("go"), ""))
	testFile.WriteString(builder.String())
}

func isContainStruct(code string) bool {
	if strings.Contains(code, "TreeNode") || strings.Contains(code, "ListNode") || strings.Contains(code, "Node") {
		return true
	}
	return false
}
