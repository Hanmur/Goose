package setting

import "github.com/spf13/viper"

//Setting 全局设置
type Setting struct {
	vp *viper.Viper
}

//NewSetting 初始化全局设置
func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Setting{vp}, nil
}
