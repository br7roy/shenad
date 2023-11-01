package opts

import (
	"fmt"
	"github.com/BurntSushi/toml"
	//"gopkg.in/yaml.v2"
	//"io/ioutil"
	"path/filepath"
	"sync"
)

type TomlConfig struct {
	App struct {
		Test bool `toml:"test"`

		Db struct {
			DbFilePath string `toml:"dbFilePath"`
		} `toml:"db"`
	} `toml:"app"`
}

var (
	Cfg  *TomlConfig
	once sync.Once
)

func LoadConfig(cPath string) *TomlConfig {
	once.Do(func() {
		Cfg = loadConfByPath(cPath)
	})
	return Cfg

}

func loadConfByPath(cPath string) *TomlConfig {
	filePath, err := filepath.Abs(cPath)
	if err != nil {
		panic(err)
	}

	conf := new(TomlConfig)

	if _, err := toml.DecodeFile(filePath, &conf); err != nil {
		panic(err)
	}

	return conf
}

// reload config
func TriggerReload(confPath string) bool {
	fmt.Println("reload config start")
	Cfg = loadConfByPath(confPath)
	fmt.Println("reload config done")
	return true
}
