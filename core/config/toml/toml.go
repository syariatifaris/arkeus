package toml

import (
	"github.com/BurntSushi/toml"
)

type Toml struct {
	tomlFile string
}

func NewTomlConfiguration(tomlFile string) *Toml {
	return &Toml{
		tomlFile: tomlFile,
	}
}

func (t *Toml) DecodeConfig(c interface{}) error {
	if _, err := toml.DecodeFile(t.tomlFile, c); err != nil {
		return err
	}

	return nil
}
