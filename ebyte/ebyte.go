package ebyte

import (
	"ginTemp/ebyte/logger"
)

type Config struct {
	Logger logger.Config
}

func New(config Config) error {
	err := logger.New(config.Logger)
	if err != nil {
		return err
	}
	logger.Info("EByte framework logger mod inited")

	return nil

}
