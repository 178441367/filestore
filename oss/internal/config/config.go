package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Upload struct {
		Driver string
		Prefix string
		Dir    string
	}
	Authorization string
}
