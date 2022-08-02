package leetcode

import "fmt"

const (
	QueryQuestionUrl = "https://leetcode-cn.com/graphql/"
	ProblemUrlPrefix = "https://leetcode.cn/problems/"
)

func CopyRight(titleSlug string) string {
	return fmt.Sprintf("来源：力扣（LeetCode）\n" +
		"链接：" + ProblemUrlPrefix + titleSlug + "/" + "\n" +
		"著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。\n")
}
