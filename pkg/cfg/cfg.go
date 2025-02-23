package cfg

import (
	"github.com/alecthomas/kong"

	kongtoml "github.com/alecthomas/kong-toml"
	"github.com/rs/zerolog/log"
)

var Cfg struct {
	Target   string `help:"target file" default:"README.md" type:"path"`
	Provider string `help:"provider" default:"mock"`
}

func cfgCheck() {

}

func cfgSetDefault() {

}

func Load() {
	kong.Parse(&Cfg, kong.Configuration(kongtoml.Loader, "config.toml"))
	cfgCheck()
	cfgSetDefault()
	log.Info().Interface("cfg", Cfg).Msg("")
}
