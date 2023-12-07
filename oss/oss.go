package main

import (
	"filestorage/oss/internal/config"
	"filestorage/oss/internal/handler"
	"filestorage/oss/internal/svc"
	"filestorage/oss/libs"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/oss.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()
	handler.RegisterHandlers(server, ctx)
	libs.StaticFileHandler(server, c.Upload.Prefix, c.Upload.Dir)
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
