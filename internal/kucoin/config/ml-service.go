package config

type MlService struct {
	addr string `yaml:"addr"`
}

func (ml *MlService) MlAddr() string {
	return ml.addr
}
