package common

import (
	std "gitlab.wallstcn.com/wscnbackend/ivankastd"
)

var (
	GlobalConf *Config
)

type Config struct {
	Micro     std.ConfigService `yaml:"micro"`
	Log       std.ConfigLog     `yaml:"log"`
	Mysql     std.ConfigMysql   `yaml:"mysql"`
	EtcdAddrs EtcdAddrsDetail   `yaml:"etcd_addrs"`
	Qcloud    QcloudDetail      `yaml:"qcloud"`
	Gitlab    GitlabDetail      `yaml:"gitlab"`
	Redis     Redis             `yaml:"redis"`
}

type GitlabDetail struct {
	PrivateToken string `yaml:"private_token"`
}

type QcloudDetail struct {
	SecretId  string `yaml:"secret_id"`
	SecretKey string `yaml:"secret_key"`
	ApiDomain string `yaml:"api_domain"`
	ApiPath   string `yaml:"api_path"`
}

type EtcdAddrsDetail struct {
	SH []string `yaml:"sh"`
	BJ []string `yaml:"bj"`
	GZ []string `yaml:"gz"`
}

func LoadConfig(filePath string) {
	println("loading config")
	GlobalConf = &Config{}
	std.LoadConf(GlobalConf, filePath)
}

func Initalise() {
	std.InitLog(GlobalConf.Log)
	//models.InitModel(GlobalConf.Mysql)
}
