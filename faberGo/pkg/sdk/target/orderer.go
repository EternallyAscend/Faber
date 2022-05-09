package target

type OrdererConfig struct {
	Key         string       `json:"key"`
	Url         string       `json:"url"`
	GrpcOptions *GrpcOptions `json:"grpcOptions"`
	TlsCACerts  *TlsCACerts  `json:"tlsCaCerts"`
}

func GenerateDefaultOrdererConfig(key string, url string) *OrdererConfig {
	// TODO 需要在生成网络时修改对应证书文件名为tls.pem
	return &OrdererConfig{
		Key:         key,
		Url:         url,
		GrpcOptions: GenerateDefaultGrpcOptions(key),
		TlsCACerts:  &TlsCACerts{Path: "/var/hyperledger/orderer/tls/tlscacerts/tls.pem"},
	}
}