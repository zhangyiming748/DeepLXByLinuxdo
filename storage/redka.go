package storage

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/nalgeon/redka"
)

var (
	db *redka.DB
)

func GetDatabase() *redka.DB {
	return db
}
func SetDatabase() {
	db, _ = redka.Open("/app/data/cache.db", nil)
}
