package utils

import "strings"

func CreateLayer(layer string, subLayers ...string) string {
	layers := []string{GetWorkingDirectory(), layer}
	layers = append(layers, subLayers...)
	return strings.Join(layers, "/")
}
