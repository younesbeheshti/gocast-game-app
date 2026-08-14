package config

import (
	"fmt"
	"strings"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/env/v2"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

func Load(configPath string) (*Config, error) {
	var k = koanf.New(".")

	k.Load(confmap.Provider(defaultConfig, "."), nil)

	//// Load JSON config on top of the default values.
	//if err := k.Load(file.Provider("mock/mock.json"), json.Parser()); err != nil {
	//	log.Fatalf("error loading config: %v", err)
	//}

	// Load YAML config and merge into the previously loaded config (because we can).
	k.Load(file.Provider(configPath), yaml.Parser())

	k.Load(env.Provider(".", env.Opt{
		Prefix: "GAMEAPP_",
		TransformFunc: func(k, v string) (string, any) {
			// Transform the key.
			k = strings.ReplaceAll(strings.ToLower(strings.TrimPrefix(k, "GAMEAPP_")), "_", ".")
			k = strings.Replace(k, "..", "_", -1)

			// Transform the value into slices, if they contain spaces.
			// Eg: MYVAR_TAGS="foo bar baz" -> tags: ["foo", "bar", "baz"]
			// This is to demonstrate that string values can be transformed to any type
			// where necessary.
			if strings.Contains(v, " ") {
				return k, strings.Split(v, " ")
			}

			return k, v
		},
	}), nil)

	var cfg Config
	if err := k.Unmarshal("", &cfg); err != nil {
		panic(err)
	}

	fmt.Println(cfg)

	return nil, nil
}
