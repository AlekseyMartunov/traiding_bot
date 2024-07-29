package config

type closeLogger struct {
	Env     string `yaml:"env"`
	LogAddr string `yaml:"addr"`
	Level   string `yaml:"level"`
}

func (cl *closeLogger) toExternalStruct() Logger {
	return Logger{
		env:     cl.Env,
		logAddr: cl.LogAddr,
		level:   cl.Level,
	}
}

type Logger struct {
	env     string
	logAddr string
	level   string
}

func (l *Logger) Env() string {
	return l.env
}

func (l *Logger) LogAddr() string {
	return l.logAddr
}

func (l *Logger) Level() string {
	return l.level
}
