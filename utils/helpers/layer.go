package helpers

import (
	"fmt"
	"strings"
)

type Path struct {
	layers []string
}

func NewPath(layers ...string) *Path {
	return &Path{layers: layers}
}

func (p *Path) Subpath(paths ...string) *Path {
	sublayers := make([]string, len(p.layers)+len(paths))
	copy(sublayers, p.layers)
	copy(sublayers[len(p.layers):], paths)
	return &Path{layers: sublayers}
}

func (p *Path) String() string {
	return fmt.Sprintf("%s/%s", GetWorkingDirectory(), p.JoinLayers("/"))
}

func (p *Path) JoinLayers(separator string) string {
	return JoinStrings(p.layers, separator)
}

func JoinStrings(slice []string, separator string) string {
	return sliceToString(slice, separator)
}

func sliceToString(slice []string, separator string) string {
	if len(slice) == 0 {
		return ""
	}
	if len(slice) == 1 {
		return slice[0]
	}
	var builder strings.Builder
	builder.WriteString(slice[0])
	for _, s := range slice[1:] {
		builder.WriteString(separator)
		builder.WriteString(s)
	}
	return builder.String()
}
