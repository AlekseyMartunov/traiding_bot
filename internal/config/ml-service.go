package config

type closeMLService struct {
	Addr string `yaml:"addr"`
}

func (cs *closeMLService) toExternalStruct() MlService {
	return MlService{
		addr: cs.Addr,
	}
}

type MlService struct {
	addr string `yaml:"addr"`
}

func (ml *MlService) Addr() string {
	return ml.addr
}
