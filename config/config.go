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
	IP       string
	Port     string
	URL      string
	Key      string
	Cert     string
	CSRF     []byte
	HashKey  []byte
	BlockKey []byte
}

// SMTP struct gives the package the SMTP details to send token to user to initialize password or to change them when lost
type SMTP struct {
	Host     string
	Port     string
	From     string
	Username string
	Password string
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
	// MailConf holds the SMTP server details
	MailConf SMTP

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

	// Twitch Authentication information
	if !viper.IsSet("twitch.channel") {
		return ErrorLogin
	}
	Cred.Channel = viper.GetString("twitch.channel")
	if viper.IsSet("twitch.token") {
		IsAuth = true
		Cred.Token = viper.GetString("twitch.token")
	}

	// Config related to Aggregator module
	if viper.IsSet("aggreg.stacksize") {
		Aggreg.StackSize = viper.GetInt("aggreg.stacksize")
	} else {
		Aggreg.StackSize = 40
	}
	if viper.IsSet("aggreg.channels") {
		Aggreg.Channels = viper.GetStringSlice("aggreg.channels")
	}

	// HTTP Config
	if viper.IsSet("http.ip") {
		WebConf.IP = viper.GetString("http.ip")
	} else {
		WebConf.IP = ""
	}
	if viper.IsSet("http.port") {
		WebConf.Port = viper.GetString("http.port")
	} else {
		WebConf.Port = "8090"
	}
	if viper.IsSet("http.url") {
		WebConf.URL = viper.GetString("http.url")
	} else {
		WebConf.URL = ""
	}
	if viper.IsSet("http.cert.key") {
		WebConf.Key = viper.GetString("http.cert.key")
	} else {
		WebConf.Key = ""
	}
	if viper.IsSet("http.cert.cert") {
		WebConf.Cert = viper.GetString("http.cert.cert")
	} else {
		WebConf.Cert = ""
	}

	// Mail Server
	if viper.IsSet("smtp.host") {
		MailConf.Host = viper.GetString("smtp.host")
	}
	if viper.IsSet("smtp.port") {
		MailConf.Port = viper.GetString("smtp.port")
	}
	if viper.IsSet("smtp.from") {
		MailConf.From = viper.GetString("smtp.from")
	}
	if viper.IsSet("smtp.username") {
		MailConf.Username = viper.GetString("smtp.username")
	}
	if viper.IsSet("smtp.password") {
		MailConf.Password = viper.GetString("smtp.password")
	}

	// CSRF Token
	if viper.IsSet("http.csrf") {
		WebConf.CSRF, _ = base64.StdEncoding.DecodeString(viper.GetString("http.csrf"))
	} else {
		// Do some magic to create
		WebConf.CSRF = securecookie.GenerateRandomKey(32)
		viper.Set("http.csrf", string(base64.StdEncoding.EncodeToString([]byte(WebConf.CSRF))))
		viper.WriteConfig()
	}

	// HashKey
	if viper.IsSet("http.hashkey") {
		WebConf.HashKey, _ = base64.StdEncoding.DecodeString(viper.GetString("http.hashkey"))
	} else {
		// Do some magic to create
		WebConf.HashKey = securecookie.GenerateRandomKey(64)
		viper.Set("http.hashkey", string(base64.StdEncoding.EncodeToString([]byte(WebConf.HashKey))))
		viper.WriteConfig()
	}

	// BlockKey
	if viper.IsSet("http.blockkey") {
		WebConf.BlockKey, _ = base64.StdEncoding.DecodeString(viper.GetString("http.blockkey"))
	} else {
		// Do some magic to create
		WebConf.BlockKey = securecookie.GenerateRandomKey(32)
		viper.Set("http.blockkey", string(base64.StdEncoding.EncodeToString([]byte(WebConf.BlockKey))))
		viper.WriteConfig()
	}

	return nil
}
