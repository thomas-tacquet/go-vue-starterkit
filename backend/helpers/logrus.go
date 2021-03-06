package helpers

import (
	"errors"

	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// default log path, if not set in config file
const defaultLogPath = "."

// Logger
type Logger struct {
	// Logs is main pointer to logrus
	Logs *logrus.Entry
	// name is used to determine log file name
	name string
	// isInit is used to avoid double initialization
	isInit bool
}

var (
	// ErrInitAlreadyDone is returned when multiplit init function is called on same object
	ErrInitAlreadyDone = errors.New("this instance is already initialized")

	// ErrFieldCantBeEmpty is returned when a necessary field is empty
	ErrFieldCantBeEmpty = errors.New("field can't be empty")
)

// Init create and setup logrus
// available log levels : panic, fatal, error, warn, info, debug, trace (ordered)
func (l *Logger) Init(name, logLevel, logPath string) error {
	if l.isInit {
		return ErrInitAlreadyDone
	}
	l.isInit = true

	if name == "" {
		return ErrFieldCantBeEmpty
	}
	l.name = name

	log := logrus.New()

	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return err
	}

	log.SetLevel(level)

	if logPath == "" {
		logPath = defaultLogPath
	}
	if err := SetupLogrus(logPath+"/"+l.name+".log", log); err != nil {
		return err
	}

	l.Logs = log.WithFields(logrus.Fields{
		"service": l.name,
	})

	return nil
}

// InitWithViper allows to init log by just passing viper instance instead of all parameters
// You can use InitWithViper or Init
func (l *Logger) InitWithViper(vpr *viper.Viper) error {
	return l.Init(
		vpr.GetString(EnvLogName),
		vpr.GetString(EnvLogLevel),
		vpr.GetString(EnvLogPath))
}

// SetupLogrus
func SetupLogrus(path string, logger *logrus.Logger) error {
	logger.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})

	// Log as JSON instead of the default ASCII formatter.
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Out = &lumberjack.Logger{
		Filename:   path,
		MaxSize:    10, // megabytes
		MaxBackups: 10,
		MaxAge:     7, //days
		Compress:   true,
	}

	return nil
}
