package conf

// Config 配置文件结构体
type Config struct {
	ListenAddr string `json:"listen_addr"`
	FLogger    struct {
		Level            int    `json:"level"`
		File             string `json:"file"`
		SlowlogThreshold int    `json:"slowlog_threshold"`
	} `json:"f_logger"`
}

var _cfg *Config

// SetConfig 设置配置
func SetConfig(cfg *Config) {
	_cfg = cfg
}

// GetConfig 获取配置文件并返回
func GetConfig() *Config {
	return _cfg
}
