package inmemory

type inmemoryCache struct {
	data map[string]*interface{}
}

func NewInmemoryCache() *inmemoryCache {
	imm := &inmemoryCache{}
	imm.data = make(map[string]*interface{})
	return imm
}

func (in *inmemoryCache) Register(key string, tok interface{}) {
	in.data[key] = &tok
}
func (in *inmemoryCache) Retreive(key string) (interface{}, bool) {
	val, exist := in.data[key]
	return val, exist
}
