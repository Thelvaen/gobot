package config

import (
	"encoding/base64"
	"errors"
	"math/rand"
	"time"

	"github.com/gorilla/securecookie"
	"github.com/spf13/viper"
)

// Configuration structure holds the configuration
type Configuration struct {
	IsAuth bool
	// Cred holds the Twitch authentication credential
	Cred struct {
		Channel string
		Token   string
	}
	// Aggreg holds the information about the aggregator module
	Aggreg struct {
		StackSize int
		Channels  []string
	}
	// WebConf holds the web server details
	WebConf struct {
		IP       string
		Port     string
		URL      string
		Key      string
		Cert     string
		CSRF     []byte
		HashKey  []byte
		BlockKey []byte
		IsSecure bool
	}
	// MailConf holds the SMTP server details
	MailConf struct {
		Host     string
		Port     string
		From     string
		Username string
		Password string
	}
}

var (
	// ErrorLoading indicates that the config could not be loaded
	ErrorLoading = errors.New("can't load configuration file")
	// ErrorLogin indicates that no login has been provided
	ErrorLogin = errors.New("no login provided in the config file")
	// IsAuth is set to true if oAuth token has been provided in the config file

	// internal vars
	seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
)

// NewConfig gets config from the YAML file and assign it's content to variables
func NewConfig() (*Configuration, error) {
	conf := &Configuration{}

	// Init Viper
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/gobot/")
	viper.AddConfigPath(".")

	// Loading conf in Viper
	err := viper.ReadInConfig()
	if err != nil {
		return nil, ErrorLoading
	}

	// Twitch Authentication information
	if !viper.IsSet("twitch.channel") {
		return nil, ErrorLogin
	}
	conf.Cred.Channel = viper.GetString("twitch.channel")
	if viper.IsSet("twitch.token") {
		conf.IsAuth = true
		conf.Cred.Token = viper.GetString("twitch.token")
	}

	// Config related to Aggregator module
	if viper.IsSet("aggreg.stacksize") {
		conf.Aggreg.StackSize = viper.GetInt("aggreg.stacksize")
	} else {
		conf.Aggreg.StackSize = 40
	}
	if viper.IsSet("aggreg.channels") {
		conf.Aggreg.Channels = viper.GetStringSlice("aggreg.channels")
	}

	// HTTP Config
	if viper.IsSet("http.ip") {
		conf.WebConf.IP = viper.GetString("http.ip")
	} else {
		conf.WebConf.IP = ""
	}
	if viper.IsSet("http.port") {
		conf.WebConf.Port = viper.GetString("http.port")
	} else {
		conf.WebConf.Port = "8090"
	}
	if viper.IsSet("http.url") {
		conf.WebConf.URL = viper.GetString("http.url")
	} else {
		conf.WebConf.URL = ""
	}
	if viper.IsSet("http.cert.key") {
		conf.WebConf.Key = viper.GetString("http.cert.key")
	} else {
		conf.WebConf.Key = ""
	}
	if viper.IsSet("http.cert.cert") {
		conf.WebConf.Cert = viper.GetString("http.cert.cert")
	} else {
		conf.WebConf.Cert = ""
	}

	// Mail Server
	if viper.IsSet("smtp.host") {
		conf.MailConf.Host = viper.GetString("smtp.host")
	}
	if viper.IsSet("smtp.port") {
		conf.MailConf.Port = viper.GetString("smtp.port")
	}
	if viper.IsSet("smtp.from") {
		conf.MailConf.From = viper.GetString("smtp.from")
	}
	if viper.IsSet("smtp.username") {
		conf.MailConf.Username = viper.GetString("smtp.username")
	}
	if viper.IsSet("smtp.password") {
		conf.MailConf.Password = viper.GetString("smtp.password")
	}

	// CSRF Token
	if viper.IsSet("http.csrf") {
		conf.WebConf.CSRF, _ = base64.StdEncoding.DecodeString(viper.GetString("http.csrf"))
	} else {
		// Do some magic to create
		conf.WebConf.CSRF = securecookie.GenerateRandomKey(32)
		viper.Set("http.csrf", string(base64.StdEncoding.EncodeToString([]byte(conf.WebConf.CSRF))))
		viper.WriteConfig()
	}

	// HashKey
	if viper.IsSet("http.hashkey") {
		conf.WebConf.HashKey, _ = base64.StdEncoding.DecodeString(viper.GetString("http.hashkey"))
	} else {
		// Do some magic to create
		conf.WebConf.HashKey = securecookie.GenerateRandomKey(64)
		viper.Set("http.hashkey", string(base64.StdEncoding.EncodeToString([]byte(conf.WebConf.HashKey))))
		viper.WriteConfig()
	}

	// BlockKey
	if viper.IsSet("http.blockkey") {
		conf.WebConf.BlockKey, _ = base64.StdEncoding.DecodeString(viper.GetString("http.blockkey"))
	} else {
		// Do some magic to create
		conf.WebConf.BlockKey = securecookie.GenerateRandomKey(32)
		viper.Set("http.blockkey", string(base64.StdEncoding.EncodeToString([]byte(conf.WebConf.BlockKey))))
		viper.WriteConfig()
	}

	if conf.WebConf.Cert != "" && conf.WebConf.Key != "" {
		conf.WebConf.IsSecure = true
	} else {
		conf.WebConf.IsSecure = false
	}

	return conf, nil
}
