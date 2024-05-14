package provider

type MockProvider struct {
}

func NewMockProvider() *MockProvider {
	return &MockProvider{}
}

func (p *MockProvider) Translate(source string) string {
	return source
}
