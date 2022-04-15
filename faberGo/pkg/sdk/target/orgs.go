package target

import "strings"

type OrgConfig struct {
	Key   string    `json:"key"`
	MspId string    `json:"mspId"`
	Peers *[]string `json:"peers"`
}

func GenerateDefaultOrgConfig(name string, mspId string) *OrgConfig {
	return &OrgConfig{
		Key:   name,
		MspId: mspId,
		Peers: &[]string{},
	}
}

func MspId(key string) string {
	return strings.Split(key, ".")[0] + "MSP"
}

func (that *OrgConfig) AddPeer(key string) {
	for _, element := range *that.Peers {
		if key == element {
			return
		}
	}
	*that.Peers = append(*that.Peers, key)
}
