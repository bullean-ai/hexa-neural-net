package ports

// IRedisRepository Whale Hunter Domain redis interface
type IRedisRepository interface {
	SetCache(string, any) error
	DeleteCache(string) error
}
