package configs

import "github.com/spf13/viper"

type Configs struct {
    Port string
    Mode string
}


func InitConfigs() *Configs {
	port, mode := readEnv()

	configs := &Configs{Port: port, Mode: mode}
	return configs
}


func readEnv() (string, string) {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("configs")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	port := viper.GetString("dev.port")
	mode := viper.GetString("dev.mode")

	return port, mode
}
