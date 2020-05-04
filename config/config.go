package config

// GlobalConfig defines the structure of the global configuration parameters
type GlobalConfig struct {
	VaultAddress string
	ClusterName  string
	TokenPath    string
	Prefix       string
	VaultToken   string
	Input        string
}

// Config stores the Global Configuration.
var Config GlobalConfig
