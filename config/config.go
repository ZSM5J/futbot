package config

import (
	"encoding/json"
	"fmt"
	"os"
)

//Configuration setup our database
type Configuration struct {
	Database struct {
		Host     string `json:"Host"`
		User     string `json:"User"`
		Dbname   string `json:"Dbname"`
		Password string `json:"Password"`
	} `json:"database"`
	Domain struct {
		URL string `json:"url"`
	} `json:"domain"`
	App struct {
		Client_id     string `json:"client_id"`
		Client_secret string `json:"client_secret"`
		Redirect_uri  string `json:"redirect_uri"`
	} `json:"app"`
	Group struct {
		Token     string `json:"token"`
		Id string `json:"id"`
    } `json:"group"`
}

//Config is a global variable
var Config Configuration

//LoadConfiguration setup config
func LoadConfiguration(file string) Configuration {

	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&Config)
	return Config
}
