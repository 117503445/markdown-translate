package model

type Provider interface {
	Translate(source string) (string, error)
}

type Cache interface {
	Get(source string) string
	Set(source string, result string) 
}