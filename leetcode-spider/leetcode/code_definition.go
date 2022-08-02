package leetcode

import "fmt"

type CodeDefinition struct {
	Value       string `json:"value"`
	Text        string `json:"text"`
	DefaultCode string `json:"defaultCode"`
}

func (c CodeDefinition) String() string {
	return fmt.Sprintf("Value: %v\n\nLanguage: %v\n\nDefautCode: %v\n\n", c.Value, c.Text, c.DefaultCode)
}

func (c *CodeDefinition) GetDefaultCode() string {
	if c == nil {
		return ""
	}
	return c.DefaultCode
}
