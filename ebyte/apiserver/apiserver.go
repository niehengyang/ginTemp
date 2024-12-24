package apiserver

import (
	"github.com/gin-gonic/gin"
)

type ApiRouteBinder func(r *gin.Engine)

type Config struct {
	Host string
	Port string
	Mode string
}

type ApiServer struct {
	Config Config
	Route  *gin.Engine
}

func NewApiServer(config Config) *ApiServer {
	return &ApiServer{
		Config: config,
		Route:  gin.Default(),
	}
}

func (gs *ApiServer) BindRouter(routes ApiRouteBinder) {
	routes(gs.Route)
}

func (gs *ApiServer) Start() error {
	return gs.Route.Run(gs.Config.Host + ":" + gs.Config.Port)
}

func (gs *ApiServer) Stop() {}
