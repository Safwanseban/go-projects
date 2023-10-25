package configs

import (
	"log"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

func NewConfig() *koanf.Koanf {

	k := koanf.New(".")
	err := k.Load(file.Provider("configs/config.yaml"), yaml.Parser())
	if err != nil {
		log.Fatal("error loading configs")
	}
	return k
}
