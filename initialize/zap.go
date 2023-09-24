package initialize

import (
	"ginDome/global"
	"ginDome/utils"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"path"
	"strings"
	"time"
)

func InitLogger()  {
	logConfig := global.GvaConfig.Log
	//判断目录是否存在，
	if result:=utils.DirExist(logConfig.Path);!result{
		_=utils.CreateDir(logConfig.Path)
	}
	//设置输出格式
	var encoder zapcore.Encoder
	if logConfig.OutFormat == logConfig.OutFormat{
		encoder = zapcore.NewJSONEncoder(getEncoderConfig())
	}else {
		encoder= zapcore.NewConsoleEncoder(getEncoderConfig())
	}
	//设置日志文件切割
	writeSyncer:=zapcore.AddSync(getLumberJackWriteSyncer())
	//创建newCore
	zapCore:=zapcore.NewCore(encoder,writeSyncer,getLevel())
	//创建logger
	logger := zap.New(zapCore)
	defer logger.Sync()
	//赋值全局变量
	global.GvaLogger=logger
}

//自定义日志输出字段
func getEncoderConfig()zapcore.EncoderConfig{
	config:=zapcore.EncoderConfig{
		TimeKey: "time",
		LevelKey: "level",
		NameKey: "logger",
		CallerKey: "caller",
		FunctionKey: zapcore.OmitKey,
		MessageKey: "msg",
		StacktraceKey: "S",
		LineEnding: zapcore.DefaultLineEnding,
		EncodeLevel: zapcore.CapitalLevelEncoder,
		EncodeTime: getEncodeTime,//自定输出时间格式
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
	}
	return config
}

func getEncodeTime(time time.Time,encoder zapcore.PrimitiveArrayEncoder){
	encoder.AppendString(time.Format("2006/01/02 - 15:04:05.000"))
}

//获取文件切割和归档配置信息
func getLumberJackWriteSyncer()zapcore.WriteSyncer{
	lumberjackConfig:=global.GvaConfig.Log.LumberJack
	lumberjackLogger:=&lumberjack.Logger{
		Filename: getLogFile(),
		MaxSize: lumberjackConfig.MaxSize,
		MaxBackups: lumberjackConfig.MaxBackups,
		MaxAge: lumberjackConfig.MaxAge,
		Compress: lumberjackConfig.Compress,
	}
	return zapcore.AddSync(lumberjackLogger)
}

//获取日志文件名
func getLogFile() string{
	fileFormat:=time.Now().Format("2006-01-02")
	fileName:=strings.Join([]string{
		global.GvaConfig.Log.FilePrefix,
		fileFormat,"log"},".")
	return path.Join(global.GvaConfig.Log.Path,fileName)
}

//获取最低记录日志级别
func getLevel()zapcore.Level{
	levelMap:=map[string]zapcore.Level{
		"debug":zapcore.DebugLevel,
		"info":zapcore.InfoLevel,
		"warn":zapcore.WarnLevel,
		"error":zapcore.ErrorLevel,
		"dPanic":zapcore.DPanicLevel,
		"fatal":zapcore.FatalLevel,
	}
	if level,ok:=levelMap[global.GvaConfig.Log.Level];ok{
		return level
	}
	return zapcore.InfoLevel
}
