package setting

import "github.com/spf13/viper"

type Setting struct {
	vp *viper.Viper
}

// 读取配置文件
func NewSetting(filepath string) (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath(filepath)
	vp.SetConfigType("toml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Setting{vp}, nil
}
