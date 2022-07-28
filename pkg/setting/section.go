package setting

type Config struct {
	Senders   map[string]interface{} `toml:"sender" mapstructure:"sender"`
	Receiver  Receiver               `toml:"receiver"`
	Template  map[string]interface{} `toml:"template" mapstructure:"template"`
	LogConfig *LogConf               `toml:"logconfig"`
	JobQueue  JobQueueConfig         `toml:"job_queue"`
}
type DingTalkConfig struct {
	Delay       int    `mapstructure:"delay"`
	AccessToken string `mapstructure:"access_token"`
	Secret      string `mapstructure:"secret"`
	Domain      string `mapstructure:"domain"`
}
type DingTalkMessageTemplate struct {
	MsgType string `mapstructure:"type"`
	Title   string `mapstructure:"title"`
	Content string `mapstructure:"content"`
}

func (dtt *DingTalkMessageTemplate) GetSentence() []string {
	return []string{dtt.Title, dtt.Content}
}

type Receiver struct {
	Port int `toml:"port"`
}
type JobQueueConfig struct {
	Size int `toml:"size"`
}

type LogConf struct {
	Level      string `toml:"level"`
	LogFile    string `toml:"logfile"`
	MaxSize    int    `toml:"maxsize"`
	MaxAge     int    `toml:"maxage"`
	MaxBackups int    `toml:"maxbackups"`
}

// 将配置文件数据映射到结构体中
func (s *Setting) ReadConfig(value interface{}) error {
	err := s.vp.Unmarshal(value)
	if err != nil {
		return err
	}
	return nil
}
