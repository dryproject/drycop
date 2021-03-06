/* This is free and unencumbered software released into the public domain. */

package util

import "github.com/spf13/viper"

type File struct {
	File    string `mapstructure:"file"`
	Markup  string `mapstructure:"markup"`
	Builder string `mapstructure:"builder"`
}

type CheckConfig struct {
	Dirs  []string `mapstructure:"dirs"`
	Files []File   `mapstructure:"files"`
}

type TemplatesConfig struct {
	Templates map[string]string
}

func LoadCheckConfig() (*CheckConfig, error) {
	var config CheckConfig
	err := viper.UnmarshalKey("check", &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func LoadTemplatesConfig() (*TemplatesConfig, error) {
	var config TemplatesConfig
	err := viper.UnmarshalKey("templates", &config.Templates)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
