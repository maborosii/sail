package dingtalk

var cityList = map[string]string{
	"jm": "江门",
	"zs": "中山",
	"zq": "肇庆",
	"fs": "佛山",
	"yf": "云浮",
	"hz": "惠州",
	"st": "汕头",
	"yj": "阳江",
	"mm": "茂名",
	"zh": "珠海",
	"dg": "东莞",
	"qy": "清远",
	"cz": "潮州",
	"hy": "河源",
	"sw": "汕尾",
	"gz": "广州",
	"sg": "韶关",
	"mz": "梅州",
	"jy": "揭阳",
	"zj": "湛江",
	"py": "番禺",
}

func formatDomainCity(toCity string, dictCity map[string]string) string {
	for k, v := range dictCity {
		if k+".harbor.com" == toCity {
			return v
		}
	}
	return "UNKNOWN"
}
func formatProjectCity(toCity string, dictCity map[string]string) string {
	for k, v := range dictCity {
		if "chart-"+k == toCity {
			return v
		}
	}
	return "UNKNOWN"
}
