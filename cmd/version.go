package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/q191201771/naza/pkg/bininfo"
	"github.com/spf13/cobra"
)

// 定义子命令行的主要参数
var versionCmd = &cobra.Command{
	Use:   "version",   // 子命令的标识
	Short: "获取当前版本",    // 简短帮助说明
	Long:  versionDesc, // 详细帮助说明
	Run: func(cmd *cobra.Command, args []string) {
		_, _ = fmt.Fprint(os.Stderr, bininfo.StringifyMultiLine())
		os.Exit(0)
	},
}
var versionDesc = strings.Join([]string{
	"该子命令支持获取当前版本信息:",
	"GitTag:        tag号",
	"GitCommitLog:  当前commit日志",
	"GitStatus:     stash状态",
	"BuildTime:     构建时间",
	"GoVersion:     golang版本",
	"runtime:       运行环环境",
}, "\n")

func init() {
}
