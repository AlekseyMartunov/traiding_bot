package config

type MlService struct {
	Addr string `yaml:"addr"`
}

func (ml *MlService) GetMlAddr() string {
	return ml.Addr
}
