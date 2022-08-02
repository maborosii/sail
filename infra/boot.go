package infra

import "sail/pkg/setting"

// 全局Starter配置及启动器，全局单例
type bootApplication struct {
	conf *setting.Config
}

var conf *setting.Config

// 传入全局配置文件conf，并获取bootApplication
func NewBootApplication(vipConf *setting.Config) *bootApplication {
	if vipConf == nil {
		panic("conf is nil")
	}
	conf = vipConf
	application := &bootApplication{conf: vipConf}
	return application
}

// 获取全局配置文件
func Conf() *setting.Config {
	if conf == nil {
		panic("conf is nil")
	}
	return conf
}

// 启动bootApplication，初始化Starter并启动
func (b *bootApplication) Run() {
	// 初始化
	initStarters()
	// 配置
	setupStarters()
	// 启动
	startStarters()
}

func startStarters() {
	for _, s := range AllStarters() {
		s.Start(conf)
	}
}

func setupStarters() {
	for _, s := range AllStarters() {
		s.Setup(conf)
	}
}

func initStarters() {
	for _, s := range AllStarters() {
		s.Init(conf)
	}
}
