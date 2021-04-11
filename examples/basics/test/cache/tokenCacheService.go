package cache

import (
	"fmt"
	"time"
)

type Token struct {
	Value    string
	creation time.Time
}

type TokenCacheService struct {
	record Recorder
}

func NewTokenCacheService(recorder Recorder) TokenCacheService {
	return TokenCacheService{recorder}
}

func (cs *TokenCacheService) AddToken(tok Token) error {
	if len(tok.Value) == 0 {
		return fmt.Errorf("Token must have valid value")
	}
	tok.creation = time.Now()

	cs.record.Register(tok.Value, tok)
	return nil
}

func (cs *TokenCacheService) GetToken(key string) (*Token, error) {
	t, exists := cs.record.Retreive(key)

	if !exists {
		return nil, fmt.Errorf("Token not found with key %s", key)
	}

	tok := (t).(Token)

	return &tok, nil
}

func (cs *TokenCacheService) IsExpired(key string) (bool, error) {

	tok, err := cs.GetToken(key)

	if err != nil {
		return false, err
	}

	now := time.Now()

	//token lives only one minute
	expired := now.Sub(tok.creation).Minutes() > 1

	return expired, nil
}
