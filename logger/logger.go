package logger

import (
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var zlog *zap.Logger

func init() {
	//defer zlog.Sync() //defer 종료될 때에 로그에 쌓인 버퍼를 지워줍니다.
	

	var err error	

	config := zap.NewProductionConfig()
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.StacktraceKey = ""
	config.EncoderConfig = encoderConfig
	//config.Level = NewAto

	zlog, err = config.Build(zap.AddCallerSkip(1))

	if err != nil {
		panic(err)
	}
}

func Info(msg string, fields ...zap.Field) {
	zlog.Info(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	zlog.Debug(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	zlog.Error(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	zlog.Warn(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	zlog.Fatal(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	zlog.Panic(msg, fields...)
}











func Test(){
	// The bundled Config struct only supports the most common configuration
		// options. More complex needs, like splitting logs between multiple files
		// or writing to non-file outputs, require use of the zapcore package.
		//
		// In this example, imagine we're both sending our logs to Kafka and writing
		// them to the console. We'd like to encode the console output and the Kafka
		// topics differently, and we'd also like special treatment for
		// high-priority logs.
	
		// First, define our level-handling logic.
		highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zapcore.ErrorLevel
		})
		lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl < zapcore.ErrorLevel
		})
	
		// Assume that we have clients for two Kafka topics. The clients implement
		// zapcore.WriteSyncer and are safe for concurrent use. (If they only
		// implement io.Writer, we can use zapcore.AddSync to add a no-op Sync
		// method. If they're not safe for concurrent use, we can add a protecting
		// mutex with zapcore.Lock.)
		topicDebugging := zapcore.AddSync(io.Discard)
		topicErrors := zapcore.AddSync(io.Discard)
	
		// High-priority output should also go to standard error, and low-priority
		// output should also go to standard out.
		consoleDebugging := zapcore.Lock(os.Stdout)
		consoleErrors := zapcore.Lock(os.Stderr)
	
		// Optimize the Kafka output for machine consumption and the console output
		// for human operators.
		kafkaEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	
		// Join the outputs, encoders, and level-handling functions into
		// zapcore.Cores, then tee the four cores together.
		core := zapcore.NewTee(
			zapcore.NewCore(kafkaEncoder, topicErrors, highPriority),
			zapcore.NewCore(consoleEncoder, consoleErrors, highPriority),
			zapcore.NewCore(kafkaEncoder, topicDebugging, lowPriority),
			zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority),
		)
	
		// From a zapcore.Core, it's easy to construct a Logger.
		logger := zap.New(core)
		defer logger.Sync()
		logger.Info("constructed a logger")
	}
	