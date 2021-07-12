package cache

type Cache interface {
	AddIntegerKey(key int) error
	IsIntegerKeyExist(key int) bool
}
