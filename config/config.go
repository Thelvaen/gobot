package config

import (
	"encoding/base64"
	"errors"
	"math/rand"
	"time"

	"github.com/gorilla/securecookie"
	"github.com/spf13/viper"
)

type credential struct {
	Channel string
	Token   string
}

type aggregation struct {
	StackSize int
	Channels  []string
}

type webConfiguration struct {
	Port     int
	URL      string
	CSRF     []byte
	HashKey  []byte
	BlockKey []byte
}

var (
	// ErrorLoading indicates that the config could not be loaded
	ErrorLoading = errors.New("can't load configuration file")
	// ErrorLogin indicates that no login has been provided
	ErrorLogin = errors.New("no login provided in the config file")
	// IsAuth is set to true if oAuth token has been provided in the config file
	IsAuth = false
	// Cred holds the Twitch authentication credential
	Cred credential
	// Aggreg holds the information about the aggregator module
	Aggreg aggregation
	// WebConf holds the web server details
	WebConf webConfiguration

	// internal vars
	seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
)

// LoadAndParse gets config from the YAML file and assign it's content to variables
func LoadAndParse() error {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return ErrorLoading
	}

	// Authentication information
	if !viper.IsSet("Twitch.Channel") {
		return ErrorLogin
	}
	Cred.Channel = viper.GetString("Twitch.Channel")
	if viper.IsSet("Twitch.Token") {
		IsAuth = true
		Cred.Token = viper.GetString("Twitch.Token")
	}

	// Config related to Aggregator module
	if viper.IsSet("Aggreg.StackSize") {
		Aggreg.StackSize = viper.GetInt("Aggreg.StackSize")
	} else {
		Aggreg.StackSize = 40
	}
	if viper.IsSet("Aggreg.Channels") {
		Aggreg.Channels = viper.GetStringSlice("Aggreg.Channels")
	}

	if viper.IsSet("Http.Port") {
		WebConf.Port = viper.GetInt("Http.Port")
	} else {
		WebConf.Port = 8090
	}
	if viper.IsSet("Http.URL") {
		WebConf.URL = viper.GetString("Http.URL")
	} else {
		WebConf.URL = ""
	}

	// CSRF Token
	if viper.IsSet("Http.CSRF") {
		WebConf.CSRF, _ = base64.StdEncoding.DecodeString(viper.GetString("Http.CSRF"))
	} else {
		// Do some magic to create
		WebConf.CSRF = securecookie.GenerateRandomKey(32)
		viper.Set("Http.CSRF", string(base64.StdEncoding.EncodeToString([]byte(WebConf.CSRF))))
		viper.WriteConfig()
	}

	// HashKey
	if viper.IsSet("Http.HashKey") {
		WebConf.HashKey, _ = base64.StdEncoding.DecodeString(viper.GetString("Http.HashKey"))
	} else {
		// Do some magic to create
		WebConf.HashKey = securecookie.GenerateRandomKey(64)
		viper.Set("Http.HashKey", string(base64.StdEncoding.EncodeToString([]byte(WebConf.HashKey))))
		viper.WriteConfig()
	}

	// BlockKey
	if viper.IsSet("Http.BlockKey") {
		WebConf.BlockKey, _ = base64.StdEncoding.DecodeString(viper.GetString("Http.BlockKey"))
	} else {
		// Do some magic to create
		WebConf.BlockKey = securecookie.GenerateRandomKey(32)
		viper.Set("Http.BlockKey", string(base64.StdEncoding.EncodeToString([]byte(WebConf.BlockKey))))
		viper.WriteConfig()
	}

	return nil
}
