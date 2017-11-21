package mapper

import (
	"github.com/rancher/norman/types"
	"github.com/rancher/norman/types/mapper"
)

type NodeAddressMapper struct {
}

func (n NodeAddressMapper) FromInternal(data map[string]interface{}) {
	addresses, _ := mapper.GetSlice(data, "addresses")
	for _, address := range addresses {
		t := address["type"]
		a := address["address"]
		if t == "InternalIP" {
			data["IpAddress"] = a
		} else if t == "Hostname" {
			data["hostname"] = a
		}
	}
}

func (n NodeAddressMapper) ToInternal(data map[string]interface{}) {
}

func (n NodeAddressMapper) ModifySchema(schema *types.Schema, schemas *types.Schemas) error {
	return nil
}
