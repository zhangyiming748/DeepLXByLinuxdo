package storage

import "testing"

// go test -v -run TestRedka
func TestRedka(t *testing.T) {
	SetDatabase()
	db := GetDatabase()
	t.Logf("%+v\n", db)
}
