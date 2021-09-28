package config

import (
	"github.com/nordix/meridio/pkg/configuration/reader"
)

// Input is a slice of Vips.
// Return a map with key as vip names.
func MakeMapFromVipList(lst []*reader.Vip) map[string]reader.Vip {
	ret := make(map[string]reader.Vip)
	for _, item := range lst {
		ret[item.Name] = *item
	}
	return ret
}

// Input is a slice of Gateways.
// Return a map with key as gateway names.
func MakeMapFromGWList(lst []*reader.Gateway) map[string]reader.Gateway {
	ret := make(map[string]reader.Gateway)
	for _, item := range lst {
		ret[item.Name] = *item
	}
	return ret
}

// Input is a slice of Attractors.
// Return a map with key as attractor names.
func MakeMapFromAttractorList(lst []*reader.Attractor) map[string]reader.Attractor {
	ret := make(map[string]reader.Attractor)
	for _, item := range lst {
		ret[item.Name] = *item
	}
	return ret
}
