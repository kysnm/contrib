package sessions

import (
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/gorilla/sessions"
	gsm "github.com/bradleypeabody/gorilla-sessions-memcache"
)

type MemcacheStore interface {
	Store
}

func NewMemecacheClient(server ...string) (*memcache.Client) {
	memcacheClent := memcache.New(server...)
	return memcacheClent
}

func NewMemcacheStore(client *memcache.Client, keyPrefix string, keyPairs ...[]byte) (MemcacheStore) {
	store := gsm.NewMemcacheStore(client, keyPrefix, keyPairs...)
	return &memcacheStore{store}
}

type memcacheStore struct {
	*gsm.MemcacheStore
}

func (c *memcacheStore) Options(options Options) {
	c.MemcacheStore.Options = &sessions.Options{
		Path:     options.Path,
		Domain:   options.Domain,
		MaxAge:   options.MaxAge,
		Secure:   options.Secure,
		HttpOnly: options.HttpOnly,
	}
}
