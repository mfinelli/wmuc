package main

import "github.com/spf13/viper"

import "github.com/mfinelli/wmuc/cmd"

func main() {
	viper.SetDefault("debug", false)
	viper.SetDefault("verbose", false)
	cmd.Execute()
}
