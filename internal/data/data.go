package data

import (
	"strike_layout_template/internal/conf"
)

type Data struct{}

func NewData(c *conf.Data) (*Data, func(), error) {
	cleanup := func() {
		// clean the data resources
	}
	return &Data{}, cleanup, nil
}
