package inmemory

import "testing"

func Test_Register(t *testing.T) {
	inmCache := NewInmemoryCache()
	inmCache.Register("hi", "hi")

	if _, exists := inmCache.Retreive("hi"); !exists {
		t.Fatalf("Expected value when retreive from cache")
	}
	if _, exists := inmCache.Retreive("hello"); exists {
		t.Fatalf("Value not expected when retreive from cache")
	}
}
