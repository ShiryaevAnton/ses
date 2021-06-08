package config

import (
	"os"

	"github.com/BurntSushi/toml"
)

type SheetConfig struct {
	Name     string `toml:"name"`
	StartRow int    `toml:"start_row"`
	NameCol  string `toml:"name_col"`
	CodeCol  string `toml:"code_col"`
	PhoneCol string `toml:"phone_col"`
}

type KeySequence struct {
	Seq string `toml:"seq"`
}

type Config struct {
	SheetConfig SheetConfig
	Seqs        []KeySequence
}

func GetConfig(path string) (*Config, error) {

	var c Config

	f, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = toml.Unmarshal(f, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil

}
