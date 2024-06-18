package cache

type Cache struct {
	Type string
}

type Option func(*Cache)

func WithType(t string) Option {
	return func(h *Cache) {
		h.Type = t
	}
}
