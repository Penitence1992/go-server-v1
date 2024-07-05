package config

import (
	"bytes"
	"github.com/penitence1992/go-server-v1/pkg/fastconv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strings"
)

type Creator struct {
	fileName      string
	paths         []string
	defaultValues string
	v             *viper.Viper
}

func GetCreator(defaultValues, filename string, paths ...string) (*Creator, error) {
	v := viper.New()

	v.SetConfigName(filename)
	v.SetConfigType("yml")
	v.AddConfigPath(".")
	if paths != nil {
		for _, path := range paths {
			v.AddConfigPath(path)
		}
	}
	v.SetEnvPrefix("ENV")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()
	err := v.ReadConfig(bytes.NewReader(fastconv.String2Byte(defaultValues)))
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			logrus.Errorf("加载配置失败: %v", err)
			return nil, err
		}
	}
	err = v.MergeInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			logrus.Errorf("加载配置失败: %v", err)
			return nil, err
		}
	}
	return &Creator{
		defaultValues: defaultValues,
		paths:         paths,
		fileName:      filename,
		v:             v,
	}, nil
}

func (t *Creator) GetConfig(config interface{}) error {
	return t.v.Unmarshal(config)
}
