package sdk

type ChannelConfig struct {
	Orderers []OrdererConfig `json:"orderers"`
	Peers    []PeerConfig    `json:"peers"`
}
