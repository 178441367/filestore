package libs

import (
	"github.com/zeromicro/go-zero/rest"
	"net/http"
	"strings"
)

func dirhandler(patern, filedir string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		handler := http.StripPrefix(patern, http.FileServer(http.Dir(filedir)))
		handler.ServeHTTP(w, req)

	}
}

func StaticFileHandler(engine *rest.Server, prifix string, dirpath string) {
	//这里注册
	dirlevel := []string{":1", ":2", ":3", ":4", ":5", ":6", ":7", ":8"}
	for i := 1; i < len(dirlevel); i++ {
		path := prifix + "/" + strings.Join(dirlevel[:i], "/")
		//最后生成 /asset
		engine.AddRoute(
			rest.Route{
				Method:  http.MethodGet,
				Path:    path,
				Handler: dirhandler(prifix, dirpath),
			})
	}

}
