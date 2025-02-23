package cache

type DisableCache struct {
}

func NewDisableCache() *DisableCache {
	return &DisableCache{}
}

func (b *DisableCache) Get(source string) string {
	return ""
}

func (b *DisableCache) Set(source string, result string) {
}
