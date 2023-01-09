package wechat

import "strings"

func colorMessage(jobStatus string) string {
	switch strings.ToLower(jobStatus) {
	case "success", "succeeded", "succeed":
		// green
		return "#3CB371"
	default:
		// red
		return "#FF0000"
	}
}
