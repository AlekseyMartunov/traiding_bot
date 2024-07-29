package config

type closeKucoin struct {
	Key          string `yaml:"key"`
	Secret       string `yaml:"secret"`
	PassPhrase   string `yaml:"pass_phrase"`
	Version      string `yaml:"version"`
	BaseEndpoint string `yaml:"base_endpoint"`
}

func (ck *closeKucoin) toExternalStruct() Kucoin {
	return Kucoin{
		key:          ck.Key,
		secret:       ck.Secret,
		passPhrase:   ck.PassPhrase,
		version:      ck.Version,
		baseEndpoint: ck.BaseEndpoint,
	}
}

type Kucoin struct {
	key          string
	secret       string
	passPhrase   string
	version      string
	baseEndpoint string
}

func (k *Kucoin) Key() string {
	return k.key
}

func (k *Kucoin) Secret() string {
	return k.secret
}

func (k *Kucoin) PassPhrase() string {
	return k.passPhrase
}

func (k *Kucoin) Version() string {
	return k.version
}

func (k *Kucoin) BaseEndpoint() string {
	return k.baseEndpoint
}
