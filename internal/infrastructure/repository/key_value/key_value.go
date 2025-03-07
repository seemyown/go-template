package key_value

// Интерфейс ключ значение хранилища

import (
	"go-fiber-template/pkg/logging"
)

var log = logging.New(logging.Config{
	FileName: "key_val",
	Name:     "key_val",
})

type IKeyValueRepository interface {
	Get(key string) (string, error)
	Set(key string, val interface{}, ttl int64) error
	Del(key string) error
}
