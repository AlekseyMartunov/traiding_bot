package config

type Kucoin struct {
	Key        string `yaml:"key"`
	Secret     string `yaml:"secret"`
	PassPhrase string `yaml:"pass_phrase"`
	Version    string `yaml:"version"`
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
