package cache

import (
	"testing"
	"time"
)

const (
	err string = "Expected result %v != \"%v\" real result"
)

var (
	toAdd time.Duration
)

func EqualSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestCache(t *testing.T) {

	t.Run("Get with Put", func(t *testing.T) {

		var cache = NewCache()
		cache.items = make(map[string]Item)

		toAdd = 1 * time.Second

		cache.Put("put", "put")

		val, found := cache.Get("put")

		if !found {
			t.Errorf(err, "", val)
		}

	})

	t.Run("Get with PutTill negative,", func(t *testing.T) {

		var cache = NewCache()
		cache.items = make(map[string]Item)

		toAdd = 1 * time.Second

		cache.PutTill("put_till", "put_till", time.Now().Add(toAdd))

		time.Sleep(1 * time.Second)

		val, found := cache.Get("put_till")

		if found {
			t.Errorf(err, "", val)
		}

	})

	t.Run("Get with PutTill positive", func(t *testing.T) {

		var cache = NewCache()
		cache.items = make(map[string]Item)

		toAdd = 1 * time.Second

		cache.PutTill("put_till", "put_till", time.Now().Add(toAdd))

		val, found := cache.Get("put_till")

		if !found {
			t.Errorf(err, "", val)
		}

	})

	t.Run("Keys negative", func(t *testing.T) {

		var cache = NewCache()
		var realResult = []string{"put"}

		cache.items = make(map[string]Item)

		toAdd = 1 * time.Second

		cache.PutTill("put_till", "put_till", time.Now().Add(toAdd))
		cache.Put("put", "put")

		time.Sleep(1 * time.Second)

		keys := cache.Keys()

		if !EqualSlices(keys, realResult) {
			t.Errorf(err, realResult, keys)
		}

	})

	t.Run("Keys positive", func(t *testing.T) {

		var cache = NewCache()
		var realResult = []string{"put_till", "put"}

		cache.items = make(map[string]Item)

		cache.PutTill("put_till", "put_till", time.Now().Add(toAdd))
		cache.Put("put", "put")

		keys := cache.Keys()

		if !EqualSlices(keys, realResult) {
			t.Errorf(err, realResult, keys)
		}

	})

	t.Run("Delete", func(t *testing.T) {

		var cache = NewCache()
		var realResult = "put_till"

		cache.items = make(map[string]Item)

		cache.PutTill("put_till", "put_till", time.Now().Add(toAdd))
		cache.Put("put", "put")
		cache.Delete("put")

		val, found := cache.Get("put")

		if found {
			t.Errorf(err, realResult, val)
		}

	})

}
