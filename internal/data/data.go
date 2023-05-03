package data

import (
	"strike_layout_template/internal/conf"

	redis "github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

// Data .
type Data struct {
	mysql *gorm.DB
	redis *redis.Client
}

// NewData .
func NewData(c *conf.Data) (*Data, func(), error) {
	cleanup := func() {
		// clean the data resources
	}
	// TODO: init data source.
	// mysql & redis
	return &Data{}, cleanup, nil
}
