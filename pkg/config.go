package ikuai

type RetryConfig struct {
	Enable        bool `yaml:"enable"`
	MaxRetryTimes int  `yaml:"max_retry_times"`
}

type Config struct {
	Url      string      `yaml:"url"`
	Username string      `yaml:"username"`
	Password string      `yaml:"password"`
	Log      bool        `yaml:"log"`
	Retry    RetryConfig `yaml:"retry"`
}
