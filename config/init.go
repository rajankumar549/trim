package Config

var config = map[string]string{
	"APPNAME":    "Trim App",
	"PORT":       "8012",
	"APIPREFIX":  "",
	"DEBUGLEVEL": "1",
}

func GetPORT() string {
	return ":" + config["PORT"]
}