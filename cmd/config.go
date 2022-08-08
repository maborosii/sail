package cmd

import (
	"sail/global"
	"strings"

	"github.com/spf13/cobra"
)

var configDir string

// 定义命令行的主要参数
var sailCmd = &cobra.Command{
	Use:   "config",
	Short: "指定配置路径", // 简短帮助说明
	Long:  sailDesc, // 详细帮助说明
	Run: func(cmd *cobra.Command, args []string) {
		// 主程序，获取自定义配置文件

		global.ConfigPath = configDir
		// fmt.Println(global.MonitorSetting.GetAddress())

	},
}
var sailDesc = strings.Join([]string{
	"该子命令用于指定配置文件路径，流程如下：",
	"1：指定配置文件文件夹即可",
	"2：指定配置名称必须为config.toml",
}, "\n")

// 用于执行main函数前初始化这个源文件里的变量
func init() {
	// 绑定命令行输入，绑定一个参数
	// 参数分别表示，绑定的变量，参数长名(--str)，参数短名(-s)，默认内容，帮助信息
	sailCmd.Flags().StringVarP(&configDir, "config", "c", "configs", "请选择配置文件")
}
