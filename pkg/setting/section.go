package setting

import "time"

//ServerSettingS 服务器设置
type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

//AppSettingS 应用设置
type AppSettingS struct {
	DefaultPageSize       int
	MaxPageSize           int
	DefaultContextTimeout time.Duration
	LogSavePath           string
	LogFileName           string
	LogFileExt            string
	UploadSavePath        string
	UploadServerUrl       string
	UploadImageMaxSize    int
	UploadImageAllowExts  []string
}

//DatabaseSettingS 数据库设置
type DatabaseSettingS struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

//JWTSettingS JWT设置
type JWTSettingS struct {
	Secret string
	Issuer string
	Expire time.Duration
}

//EmailSettingS 邮箱设置
type EmailSettingS struct {
	Host     string
	Port     int
	UserName string
	Password string
	IsSSL    bool
	From     string
	To       []string
}

//RedisPoolSettingS Redis池设置
type RedisPoolSettingS struct {
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
	Host        string
	Protocol    string
}

//ReadSection 设置段读取
func (setting *Setting) ReadSection(key string, value interface{}) error {
	err := setting.vp.UnmarshalKey(key, value)
	if err != nil {
		return err
	}

	return nil
}
