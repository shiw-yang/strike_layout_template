package data

import (
	"strike_layout_template/internal/conf"

	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData)

type Data struct{}

func NewData(c *conf.Data) (*Data, func(), error) {
	cleanup := func() {
		// clean the data resources
	}
	return &Data{}, cleanup, nil
}
