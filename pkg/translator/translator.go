package translator

type Translator struct {
}

func NewTranslator() *Translator {
	return &Translator{}
}

func (t *Translator) Translate(source string) string {
	return source
}
