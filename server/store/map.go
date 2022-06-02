package store

func NewStore() StoreHandler {
	return &storeHandler{
		Map: make(map[string]string),
	}
}

type storeHandler struct {
	Map map[string]string
}

type StoreHandler interface {
	Set(key, value string)
	Get(Key string) string
	Delete(key string)
}

func (s *storeHandler) Set(key, value string) {
	s.Map[key] = value
}

func (s *storeHandler) Get(key string) string {
	return s.Map[key]
}

func (s *storeHandler) Delete(key string) {
	delete(s.Map, key)
}
