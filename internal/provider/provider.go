package provider

type Provider interface {
	Translate(source string) (string, error)
}
