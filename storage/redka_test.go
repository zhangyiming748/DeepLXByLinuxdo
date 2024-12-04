package storage

import "testing"

func TestMain(m *testing.M) {

}
func init() {
	SetDatabase()
}

// docker run -dit --name trans -v C:\Users\zen\Github\DeepLXByLinuxdo:/data golang:1.23.3-alpine3.20 ash
// docker exec -it trans ash
// go test -v -run TestRedka
func TestRedka(t *testing.T) {
	SetDatabase()
	db := GetDatabase()
	t.Logf("%+v\n", db)
}

// go test -v -run TestSet
func TestSet(t *testing.T) {
	SetDatabase()
	err := GetDatabase().Str().Set("hello", "你好")
	if err != nil {
		t.Logf("插入错误%+v\n", err)

		return
	}
}

// go test -v -run TestGet
func TestGet(t *testing.T) {

	get, err := GetDatabase().Str().Get("hello")
	if err != nil {
		return
	}
	t.Logf("get: %+v\n", get)
}
