package config

import (
	"flag"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// Config ...
type Config struct {
	Port            string `yaml:"server_port"`
	LogLevel        string `yaml:"log_level"`
	DefaultTemplate string `yaml:"default_template"`
	Templates       []struct {
		Name string `yaml:"name"`
		Path string `yaml:"path"`
	} `yaml:"templates"`
}

var config Config

// Init ...
func Init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			s := strings.Split(f.Function, ".")
			funcname := s[len(s)-1]
			_, filename := path.Split(f.File)
			return funcname, filename
		},
	})
	logrus.SetReportCaller(true)

	logrus.Info("loading configuration from yaml file")
	path := ""
	flag.StringVar(&path, "config", "", "`file` to configuration line notify gateway")
	flag.Parse()

	if path != "" {
		logrus.Info("open configuration file at " + path)
		f, err := os.Open(path)
		if err != nil {
			logrus.Fatal(err)
		}
		defer f.Close()

		decoder := yaml.NewDecoder(f)
		err = decoder.Decode(&config)
		if err != nil {
			logrus.Fatal(err)
		}
	} else {
		logrus.Info("not provide configuration file, server run on default port")
		config.Port = "3000"
	}
	logLevel()
}

func logLevel() {
	switch config.LogLevel {
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "warnning":
		logrus.SetLevel(logrus.WarnLevel)
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}
	logrus.SetOutput(os.Stdout)
}

// Get ..
func Get() Config {
	return config
}
