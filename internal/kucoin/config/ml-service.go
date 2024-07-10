package config

type MlService struct {
	Addr string `yaml:"addr"`
}

func (ml *MlService) MlAddr() string {
	return ml.Addr
}
