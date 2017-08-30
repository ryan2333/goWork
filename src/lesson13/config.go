package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
)

type UserScriptConfig struct {
	Path string
	Step string
}

type SenderConfig struct {
	TransAddr     string `toml:"trans_addr"`
	FlushInterval int    `tomal: "flush_interval"`
	MaxSleepTime  int    `toml: "max_sleep_time"`
}

/*
type config struct {
	TransAddr  string
	UserScript []UserScriptConfig
}
*/

type config struct {
	TransAddr  SenderConfig
	UserScript []UserScriptConfig `toml: "user_script"`
}

var (
	configPath = flag.String("config", "config.toml", "config path")
	gcfg       config
)

func main() {
	flag.Parse()
	_, err := toml.DecodeFile(*configPath, &gcfg)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%#v\n", gcfg)
}
