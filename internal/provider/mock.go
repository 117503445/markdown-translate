package provider

import "fmt"

type MockProvider struct {
}

func NewMockProvider() *MockProvider {
	return &MockProvider{}
}

func (p *MockProvider) Translate(source string) (string, error) {
	return fmt.Sprintf("[翻译开始]%s[翻译结束]", source), nil
}
