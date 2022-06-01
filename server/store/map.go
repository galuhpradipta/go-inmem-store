package store

func NewDB() DBHandler {
	return &dbHandler{
		Map: make(map[string]string),
	}
}

type dbHandler struct {
	Map map[string]string
}

type DBHandler interface {
	Set(key, value string)
	Get(Key string) string
	Delete(key string)
}

func (db *dbHandler) Set(key, value string) {
	db.Map[key] = value
}

func (db *dbHandler) Get(key string) string {
	return db.Map[key]
}

func (db *dbHandler) Delete(key string) {
	delete(db.Map, key)
}
