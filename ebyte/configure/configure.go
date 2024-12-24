package configure

import "github.com/spf13/viper"

var Getter *viper.Viper

func Load(configFile string, configStruct any) error {
	viper.SetConfigFile(configFile)
	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	// 将配置文件解码到结构体中
	if err := viper.Unmarshal(configStruct); err != nil {
		return err
	}
	Getter = viper.GetViper()
	return nil
}
