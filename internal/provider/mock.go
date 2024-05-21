package provider

import "fmt"

type MockProvider struct {
}

func NewMockProvider() *MockProvider {
	return &MockProvider{}
}

func (p *MockProvider) Translate(source string) (string, error) {
	return fmt.Sprintf("[翻译]%s[结束]", source), nil
}
