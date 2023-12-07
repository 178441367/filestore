package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Upload struct {
		Prefix string
		Dir    string
	}
}
