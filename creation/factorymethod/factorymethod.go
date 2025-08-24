package factorymethod

import "fmt"

type Logger interface {
	Log(message string)
}

// FileLogger 文件记录
type FileLogger struct {
}

func (f *FileLogger) Log(message string) {
	fmt.Println("Log to file: ", message)
}

// ConsoleLogger 控制台日志记录
type ConsoleLogger struct {
}

func (c *ConsoleLogger) Log(message string) {
	fmt.Println("Log to console: " + message)
}

type LoggerFactory interface {
	CreateLogger() Logger
}

// FileLoggerFactory 是文件日志记录器工厂实现
type FileLoggerFactory struct{}

func (f *FileLoggerFactory) CreateLogger() Logger {
	return &FileLogger{}
}

// ConsoleLoggerFactory 是控制台日志记录器工厂实现
type ConsoleLoggerFactory struct{}

func (c *ConsoleLoggerFactory) CreateLogger() Logger {
	return &ConsoleLogger{}
}
