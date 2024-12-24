package etc

import (
	"ginTemp/ebyte"
	"ginTemp/ebyte/apiserver"
)

type Config struct {
	EByteConfig   ebyte.Config
	DEMOAPIServer apiserver.Config
	FileServer    fileServer
	WsServer      wsConfig
}

type fileServer struct {
	PrefixPath      string
	RelativePath    string
	AudioSamplePath string //声音采样资源
}

type wsConfig struct {
	Host string
	Port string
}
