package cache

//go:generate moq -out recorder_mock.go . Recorder
type Recorder interface {
	Register(key string, tok interface{})
	Retreive(key string) (interface{}, bool)
}
