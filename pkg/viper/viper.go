package viper

import(
	"github.com/spf13/viper"
)

func defaultViperVal() {
	viper.SetDefault("serverPort", ":50000")
	viper.SetDefault("firestoreAccountKey", "configs/serviceAccountKey.json")
	viper.SetDefault("host", "localhost")
	viper.SetDefault("dbPort", 5432)
	viper.SetDefault("user", "postgres")
	viper.SetDefault("dbname", "simple-http-server")
}

func ReadConfig(filename string, configPath string) error {
	defaultViperVal()
	//viper.AutomaticEnv()
	viper.AddConfigPath(configPath)
	viper.SetConfigName(filename)
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		return err
	}
	return nil
}