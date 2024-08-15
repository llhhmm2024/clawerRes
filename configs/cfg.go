package configs

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
)

// Mysql 数据库连接信息
type Mysql struct {
	Host         string        `json:"host" toml:"host"`
	User         string        `json:"user" toml:"user"`
	Password     string        `json:"password" toml:"password"`
	Dbname       string        `json:"dbname" toml:"dbname"`
	Port         string        `json:"port" toml:"port"`
	Debug        bool          `json:"debug" toml:"debug"`
	DryRun       bool          `json:"dry_run" toml:"dry_run"`
	Charset      string        `json:"charset" toml:"charset"`
	MaxIdleConns time.Duration `json:"max_idle_conns" toml:"max_idle_conns"`
	MaxOpenConns int           `json:"max_open_conns" toml:"max_open_conns"`
}

func (m *Mysql) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", m.User, m.Password, m.Host, m.Port, m.Dbname, m.Charset)
}

type Jobs struct {
	Name      string `json:"name"`
	State     string `json:"state"`
	Model     string `json:"model"`
	Page      int    `json:"page"`
	OverWrite int    `json:"overwrite"`
}

type config struct {
	YzzyList      []string `mapstructure:"yzzy" yaml:"yzzy"`
	FfzyList      []string `mapstructure:"ffzy" yaml:"ffzy"`
	Database      Mysql    `mapstructure:"db" yaml:"db"`
	Jobs          []Jobs   `mapstructure:"jobs" yaml:"jobs"`
	Downloader    string   `mapstructure:"downloader" yaml:"downloader"`
	MaxConcurrent int      `mapstructure:"max_concurrent" yaml:"max_concurrent"`
	Proxy         string   `mapstructure:"proxy" yaml:"proxy"`
	SaveDir       string   `mapstructure:"save_dir" yaml:"save_dir"`
}

var Cfg *config

func init() {
	// 初始化 Viper
	viper.SetConfigName("cfg") // 配置文件的名称（无扩展名）
	viper.AddConfigPath(".")   // 例如：在当前目录中查找配置文件

	// 设置配置文件类型
	viper.SetConfigType("yaml")

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s\n", err)
	}

	// 将读取的配置信息映射到 Config 结构体
	// var config Config
	if err := viper.Unmarshal(&Cfg); err != nil {
		fmt.Printf("Unable to decode into struct, %v\n", err)
		os.Exit(1)
	}
}

func GetMaxConcurrent() int {
	return Cfg.MaxConcurrent
}

func GetProxy() string {
	return Cfg.Proxy
}

func GetSaveDir() string {
	return Cfg.SaveDir
}
