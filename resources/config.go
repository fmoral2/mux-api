package resources

import "github.com/spf13/viper"

type CognitoConfig struct {
	Password  string `mapstructure:"PASSWORD"`
	AwsRegion string `mapstructure:"AWS_REGION"`
	PoolID    string `mapstructure:"AWS_USER_POOL_ID"`
	AppID     string `mapstructure:"AWS_APP_CLIENT_ID"`
}

func LoadConfig(path string) (config CognitoConfig, err error) {

	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
