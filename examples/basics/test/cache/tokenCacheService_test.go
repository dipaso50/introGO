package cache

import (
	"testing"
	"time"
)

func Test_cacheservice_should_add_token(t *testing.T) {

	mockedRecorder := &RecorderMock{
		RegisterFunc: func(key string, tok interface{}) {
			//do nothing
		},
		RetreiveFunc: func(key string) (interface{}, bool) {
			return Token{Value: "1234"}, true
		},
	}

	service := NewTokenCacheService(mockedRecorder)

	if err := service.AddToken(Token{Value: "1234"}); err != nil {
		t.Fatalf("Error setting token %v", err)
	}

	if _, err := service.GetToken(""); err != nil {
		t.Fatalf("Error getting token %v", err)
	}
}

func Test_expect_error_if_token_doesn_exist(t *testing.T) {
	mockedRecorder := &RecorderMock{
		RegisterFunc: func(key string, tok interface{}) {
			//do nothing
		},
		RetreiveFunc: func(key string) (interface{}, bool) {
			return Token{}, false
		},
	}

	service := NewTokenCacheService(mockedRecorder)

	if _, err := service.GetToken(""); err == nil {
		t.Fatalf("Error expected if token doesn exist")
	}
}

func Test_expect_error_if_token_invalid(t *testing.T) {
	mockedRecorder := &RecorderMock{
		RegisterFunc: func(key string, tok interface{}) {
			//do nothing
		},
		RetreiveFunc: func(key string) (interface{}, bool) {
			return Token{}, false
		},
	}

	service := NewTokenCacheService(mockedRecorder)

	if err := service.AddToken(Token{}); err == nil {
		t.Fatalf("Error expected token must have a value")
	}

	if err := service.AddToken(Token{Value: ""}); err == nil {
		t.Fatalf("Error expected token must have a value")
	}
}

func Test_token_expired(t *testing.T) {
	mockedRecorder := &RecorderMock{
		RegisterFunc: func(key string, tok interface{}) {
			//do nothing
		},
		RetreiveFunc: func(key string) (interface{}, bool) {
			t := Token{}
			//fecha de creación de hace 5 min
			t.creation = time.Now().Add(-5 * time.Minute)
			return t, true
		},
	}

	service := NewTokenCacheService(mockedRecorder)

	var expired bool
	var err error

	expired, err = service.IsExpired("1234")

	if err != nil {
		t.Fatalf("Error getting token %v", err)
	}

	if !expired {
		t.Fatalf("Expected expired = true, got %v", expired)
	}
}

func Test_token_not_expired(t *testing.T) {
	mockedRecorder := &RecorderMock{
		RegisterFunc: func(key string, tok interface{}) {
			//do nothing
		},
		RetreiveFunc: func(key string) (interface{}, bool) {
			t := Token{}
			t.creation = time.Now()
			return t, true
		},
	}

	service := NewTokenCacheService(mockedRecorder)

	var expired bool
	var err error

	expired, err = service.IsExpired("1234")

	if err != nil {
		t.Fatalf("Error getting token %v", err)
	}

	if expired {
		t.Fatalf("Expected expired = false, got %v", expired)
	}
}
