package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"strings"
)

var Viper *viper.Viper

func LoadConfig(path string) {
	log.Println("[VIPER] Loading environment variables")

	viperSetup := viper.GetViper()
	viper.Set("Verbose", true)
	viperSetup.SetConfigType("yml")
	viperSetup.SetConfigName("application")
	viperSetup.AddConfigPath(path)
	viperSetup.AllowEmptyEnv(true)
	viperSetup.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viperSetup.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	Viper = viperSetup
}
