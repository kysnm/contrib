package sessions

import (
	"testing"
)

const memcacheTestServer = "localhost:11211"

var newMemcacheStore = func(_ *testing.T) Store {
	store := NewMemcacheStore(NewMemecacheClient(memcacheTestServer), "", []byte("secret"))
	return store
}

func TestMemcache_SessionGetSet(t *testing.T) {
	sessionGetSet(t, newMemcacheStore)
}

func TestMemcache_SessionDeleteKey(t *testing.T) {
	sessionDeleteKey(t, newMemcacheStore)
}

func TestMemcache_SessionFlashes(t *testing.T) {
	sessionFlashes(t, newMemcacheStore)
}

func TestMemcache_SessionClear(t *testing.T) {
	sessionClear(t, newMemcacheStore)
}

func TestMemcache_SessionOptions(t *testing.T) {
	sessionOptions(t, newMemcacheStore)
}
