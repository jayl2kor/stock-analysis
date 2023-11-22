package routing

import (
	"fmt"
	"stock-anaysis/pkg/config"
)

func Serve() {
	r := GetRouter()

	configs := config.Get()

	err := r.Run(fmt.Sprintf("%s:%s", configs.Server.Host, configs.Server.Port))
	if err != nil {
		panic(err)
	}
}
