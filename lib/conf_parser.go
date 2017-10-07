package lyndbdump

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	LocalDb LocalDb `json:"localdb"`
	Hosts   []Hosts `json:"hosts"`
}

type LocalDb struct {
	Host    string `json:"host"`
	User    string `json:"user"`
	Pass    string `json:"pass"`
	Options []string
}

type Hosts struct {
	Db struct {
		Host string `json:"host"`
		Name string `json:"name"`
		Pass string `json:"pass"`
		Port int64  `json:"port"`
		User string `json:"user"`
	} `json:"db"`
	LocalName   string `json:"localName"`
	Name        string `json:"name"`
	Protocol    string `json:"protocol"`
	Enabled     bool   `json:"enabled"`
	WriteToFile bool   `json:"writeToFile"`
}

func ConfParse() Config {
	rawConf, err := ioutil.ReadFile("config.json")
	errChk(err)

	conf := Config{}
	json.Unmarshal(rawConf, &conf)
	fmt.Println("Configuration parse complete...")
	return conf
}
