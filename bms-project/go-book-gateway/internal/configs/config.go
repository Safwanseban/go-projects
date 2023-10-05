package configs

import (
	"log"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

func NewConfig() *koanf.Koanf {

	k := koanf.New(".")
	if err := k.Load(file.Provider("./internal/configs/config.yaml"), yaml.Parser()); err != nil {
		log.Fatal("error loading configurations")
	}
	return k
}
