package cfg

import (
	"github.com/alecthomas/kong"

	kongtoml "github.com/alecthomas/kong-toml"
	"github.com/rs/zerolog/log"
)

var Cfg struct {
	Target   string            `help:"target file" default:"README.md" type:"path"`
	Provider map[string]string `help:"provider"`
	Cache    map[string]string `help:"cache"`
	Log      struct {
		Level string `enum:"trace,debug,info,warn,error" default:"debug"`
	} `embed:"" prefix:"log-"`
}

func cfgCheck() {

}

func cfgSetDefault() {

}

func Load(fileCfg string) {
	kong.Parse(&Cfg, kong.Configuration(kongtoml.Loader, fileCfg))
	cfgCheck()
	cfgSetDefault()
	log.Info().Interface("cfg", Cfg).Msg("")
}
