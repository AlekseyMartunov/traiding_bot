package config

type Logger struct {
	Env     string `yaml:"env"`
	LogAddr string `yaml:"addr"`
	Level   string `yaml:"level"`
}

func (l *Logger) GetEnv() string {
	return l.Env
}

func (l *Logger) GetLogAddr() string {
	return l.LogAddr
}

func (l *Logger) GetLevel() string {
	return l.Level
}
