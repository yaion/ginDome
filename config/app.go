package config

import "time"

// APP 应用配置信息
type APP struct{
	AppName string `yaml:"app_name"`
	Version  string `yaml:"version"`
	Addr string  `yaml:"addr"`
	Env   string `yaml:"env"`
	Cron  bool `yaml:"cron"`
	ConfigFile string  `yaml:"config_file"`
}

// MySQL信息
type Mysql struct {
	Host                      string        `yaml:"host"`
	Port                      string        `yaml:"port"`
	User                      string        `yaml:"user"`
	Password                  string        `yaml:"password"`
	Database                  string        `yaml:"database"`
	Charset                   string        `yaml:"charset"`                   //要支持完整的UTF-8编码,需设置成: utf8mb4
	AutoMigrate               bool          `yaml:"autoMigrate"`               // 初始化时调用数据迁移
	ParseTime                 bool          `yaml:"parseTime"`                 //解析time.Time类型
	TimeZone                  string        `yaml:"timeZone"`                  // 时区,若设置 Asia/Shanghai,需写成: Asia%2fShanghai
	DefaultStringSize         uint          `yaml:"defaultStringSize"`         // string 类型字段的默认长度
	DisableDatetimePrecision  bool          `yaml:"disableDatetimePrecision"`  // 禁用 datetime 精度
	SkipInitializeWithVersion bool          `yaml:"skipInitializeWithVersion"` // 根据当前 MySQL 版本自动配置
	Gorm                      gorm          `yaml:"gorm"`
	SlowSql                   time.Duration `yaml:"slowSql"`                   //慢SQL
	LogLevel                  string        `yaml:"logLevel"`                  // 日志记录级别
	IgnoreRecordNotFoundError bool          `yaml:"ignoreRecordNotFoundError"` // 是否忽略ErrRecordNotFound(未查到记录错误)
}
// gorm 配置信息
type gorm struct {
	SkipDefaultTx   bool   `yaml:"skipDefaultTx"`                            //是否跳过默认事务
	CoverLogger     bool   `yaml:"coverLogger"`                              //是否覆盖默认logger
	PreparedStmt    bool   `yaml:"prepareStmt"`                              // 设置SQL缓存
	CloseForeignKey bool   `yaml:"disableForeignKeyConstraintWhenMigrating"` // 禁用外键约束
	TablePrefix     string `yaml:"tablePrefix"`                              // 表前缀
	SingularTable   bool   `yaml:"singularTable"`                            //是否使用单数表名(默认复数)，启用后，User结构体表将是user
}

type Redis struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Password string `yaml:"password"`
	DefaultDB int `yaml:"defaultDB"`
	DialTimeOut time.Duration `yaml:"dialTimeOut"`
}

// 日志信息
type Log struct {
	Path       string     `yaml:"path"`
	Level      string     `yaml:"level"`
	FilePrefix string     `yaml:"filePrefix"`
	FileFormat time.Time     `yaml:"fileFormat"`
	OutFormat  string     `yaml:"outFormat"`
	LumberJack lumberJack `yaml:"lumberJack"`
}
// 日志切割
type lumberJack struct {
	MaxSize    int  `yaml:"maxSize"`    //单文件最大容量(单位MB)
	MaxBackups int  `yaml:"maxBackups"` // 保留旧文件的最大数量
	MaxAge     int  `yaml:"maxAge"`     // 旧文件最多保存几天
	Compress   bool `yaml:"compress"`   // 是否压缩/归档旧文件
}

// Jwt jwt
type Jwt struct{
	Secret string `yaml:"secret"`
	Issuer string `yaml:"issuer"`
	Expire time.Duration `yaml:"expire"`
}

type ServerConfig struct {
	App APP `yaml:"app"`
	Mysql Mysql `yaml:"mysql"`
	Redis Redis `yaml:"redis"`
	Log Log `yaml:"log"`
	Jwt Jwt `yaml:"jwt"`
}

