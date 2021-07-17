package tinycache

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

// TestHTTP开启http服务后
// curl http://localhost:9999/_tinycache/scores/Tom
// 结果为：630
// curl http://localhost:9999/_tinycache/scores/tim
// 结果为：tim not exist

func TestHTTP(t *testing.T) {
	_ = NewGroup("scores", 2, GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			if v, ok := db[key]; ok {
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))

	addr := "localhost:9999"
	peers := NewHTTPPool(addr)
	log.Println("tinycache is running at", addr)
	log.Fatal(http.ListenAndServe(addr, peers))
}
