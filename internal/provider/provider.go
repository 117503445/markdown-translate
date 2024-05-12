package provider

type Provider interface {
	Translate(source string) string
}

type MockProvider struct {
}

func NewMockProvider() *MockProvider {
	return &MockProvider{}
}

func (p *MockProvider) Translate(source string) string {
	return source
}