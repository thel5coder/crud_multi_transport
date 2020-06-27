package main

import (
	"crud_multi_transport/cli/command"
	"flag"
	"github.com/spf13/viper"
)

func main(){
	var config string
	flag.StringVar(&config, "config","config.yml","Defines the path, name and extension of the config file")
	flag.Parse()

	if config != ""{
		viper.SetConfigFile(config)
		viper.ReadInConfig()
	}

	command.Execute()
}
