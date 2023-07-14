package common

type Config struct {
	Email EmailConfig
}

type EmailConfig struct {
	PublicKey  string `json:"public_key"`
	PrivateKey string `json:"private_key"`
}
