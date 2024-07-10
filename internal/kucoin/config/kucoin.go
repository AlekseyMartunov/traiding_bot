package config

type Kucoin struct {
	Key          string `yaml:"key"`
	Secret       string `yaml:"secret"`
	PassPhrase   string `yaml:"pass_phrase"`
	Version      string `yaml:"version"`
	BaseEndpoint string `yaml:"base_endpoint"`
}

func (k *Kucoin) GetKey() string {
	return k.Key
}

func (k *Kucoin) GetSecret() string {
	return k.Secret
}

func (k *Kucoin) GetPassPhrase() string {
	return k.PassPhrase
}

func (k *Kucoin) GetVersion() string {
	return k.Version
}

func (k *Kucoin) GetBaseEndpoint() string {
	return k.BaseEndpoint
}
