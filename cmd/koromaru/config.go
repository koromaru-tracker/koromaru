package main

import (
	"errors"
	"io"
	"os"

	"github.com/koromaru-tracker/koromaru/index/types"

	yaml "gopkg.in/yaml.v2"
)

func ParseConfig(path string) (*types.Config, error) {
	if path == "" {
		return nil, errors.New("no config path specified")
	}

	f, err := os.Open(os.ExpandEnv(path))
	if err != nil {
		return nil, err
	}
	defer f.Close()

	contents, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	var cfgFile types.Config
	err = yaml.Unmarshal(contents, &cfgFile)
	if err != nil {
		return nil, err
	}

	return &cfgFile, nil
}
