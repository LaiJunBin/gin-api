package setting

import "github.com/spf13/viper"

type Setting struct {
	viper *viper.Viper
}

func NewSetting() (*Setting, error) {
	v := viper.New()
	v.AddConfigPath("./")
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()

	if err != nil {
		return nil, err
	}

	return &Setting{v}, nil
}
