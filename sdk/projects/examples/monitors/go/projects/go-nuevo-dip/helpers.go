package main

import (
	"fmt"

	sdkviper "github.com/devpablocristo/golang/sdk/pkg/configurators/viper"
)

func loadLayerConfig() (*LayersConfig, error) {
	var layersConfig LayersConfig
	err := sdkviper.UnmarshalConfig(&layersConfig)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshaling configuration: %w", err)
	}

	return &layersConfig, nil
}
