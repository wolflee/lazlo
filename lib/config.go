package lib

import (
	"github.com/ccding/go-logging/logging"
	"github.com/danryan/env"
	"os"
	"time"
)

// Config struct
type Config struct {
	Name     string `env:"key=CHATBOT_NAME default=lazlo"`
	Token    string `env:"key=CHATBOT_TOKEN"`
	URL      string `env:"key=CHATBOT_URL default=http://localhost"`
	LogLevel string `env:"key=CHATBOT_LOG_LEVEL default=info"`
	RedisURL string `env:"key=CHATBOT_REDIS_URL"`
	RedisPW  string `env:"key=CHATBOT_REDIS_PW"`
	Port     string `env:"key=CHATBOT_PORT"`
}

func newConfig() *Config {
	c := &Config{}
	env.MustProcess(c)
	return c
}

func newLogger() *logging.Logger {
	format := "%25s [%s] %8s: %s\n time,name,levelname,message"
	timeFormat := time.RFC3339
	level := logging.GetLevelValue(`INFO`)
	logger, _ := logging.WriterLogger("chatbot", level, format, timeFormat, os.Stdout, true)
	return logger
}
