package repo

import (
	"fmt"

	"github.com/myste1tainn/msnet"
	"github.com/spf13/viper"
)

func NewGameProfileRepoConfig() *msnet.Config {
	var key = "integration.gameprofile"
	var cfg msnet.Config
	if err := viper.UnmarshalKey(key, &cfg); err != nil {
		panic(err)
	}
	fmt.Printf("[info] config with key = %s, loaded cfg = %v\n", key, cfg)
	return &cfg
}
