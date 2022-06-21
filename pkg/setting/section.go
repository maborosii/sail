package setting

type MonitorConfig struct {
	Address   string         `toml:"address"`
	TimeOut   int            `toml:"timeout"`
	Items     MonitorItems   `toml:"items"`
	Output    *MonitorOutput `toml:"output"`
	LogConfig *LogConf       `toml:"logconfig"`
}
type LogConf struct {
	Level      string `toml:"level"`
	LogFile    string `toml:"logfile"`
	MaxSize    int    `toml:"maxsize"`
	MaxAge     int    `toml:"maxage"`
	MaxBackups int    `toml:"maxbackups"`
}

func (conf *MonitorConfig) GetTimeOut() int {
	return conf.TimeOut
}
func (conf *MonitorConfig) GetAddress() string {
	return conf.Address
}
func (conf *MonitorConfig) GetMonitorItems() map[string]string {
	return conf.Items.ConvertToMap()
}
func (conf *MonitorConfig) GetOutputFileAndSheetName() (string, string) {
	return conf.Output.Project + "_" + conf.Output.FileName, conf.Output.SheetName
}
func (conf *MonitorConfig) GetOutputTitle() []string {
	return conf.Output.Title
}

func (conf *MonitorConfig) GetLogConfig() *LogConf {
	return conf.LogConfig
}

func NewMonitorConfig() *MonitorConfig {
	return &MonitorConfig{}
}

type MonitorItems []*MonitorItem

type MonitorItem struct {
	Metrics string `toml:"metrics"`
	Promql  string `toml:"promql"`
}

func (i MonitorItems) ConvertToMap() map[string]string {
	promsqls := make(map[string]string)
	for _, item := range i {
		promsqls[item.Metrics] = item.Promql
	}
	return promsqls
}

type MonitorOutput struct {
	Project   string   `toml:"project"`
	FileName  string   `toml:"filename"`
	SheetName string   `toml:"sheetname"`
	Title     []string `toml:"title"`
}

// 将配置文件数据映射到结构体中
func (s *Setting) ReadConfig(value interface{}) error {
	err := s.vp.Unmarshal(value)
	if err != nil {
		return err
	}
	return nil
}
