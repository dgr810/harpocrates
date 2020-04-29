package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// GlobalConfig defines the structure of the global configuration parameters
type GlobalConfig struct {
	VaultAddress string `required:"false" envconfig:"vault_address"`
	ClusterName  string `required:"false" envconfig:"cluster_name"`
	TokenPath    string `required:"false" envconfig:"token_path"`
	Prefix       string `required:"false" envconfig:"prefix"`
	VaultToken   string `required:"false" envconfig:"vault_token"`
	Secrets      string `required:"false" envconfig:"secrets"`
}

// Config stores the Global Configuration.
var Config GlobalConfig

//LoadConfig Loads config from env
func LoadConfig() {

	configErr := envconfig.Process("HARPOCRATES", &Config)
	if configErr != nil {
		log.Fatal(configErr)
	}
}
