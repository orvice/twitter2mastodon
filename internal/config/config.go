package config

import (
	"github.com/BurntSushi/toml"
	"github.com/weeon/mod"
)

const (
	DefaultCachePath = "/var/lib/tw2mst-cache.db"
)

type Config struct {
	InstanceID         string
	CacheDriver        string
	CachePath          string
	Redis              mod.Redis
	SkipReply          bool
	TwitterKey         string
	TwitterSecretKey   string
	TwitterToken       string
	TwitterTokenSecret string
	TwitterScreenName  string
	Mastodons          []Mastodon
}

type Mastodon struct {
	Enable       bool
	Server       string
	ClientID     string
	ClientSecret string
	Token        string
}

var (
	Conf = new(Config)
)

func Init(path string) error {
	if _, err := toml.DecodeFile(path, Conf); err != nil {
		return err
	}
	return nil
}
