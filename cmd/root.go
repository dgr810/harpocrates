package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/BESTSELLER/harpocrates/config"
	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "harpocrates",
		Short: "A generator for Cobra based Applications",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Echo: " + strings.Join(args, " "))
			fmt.Printf("%+v\n", config.Config)
		},
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

var VaultAddress string
var ClusterName string
var TokenPath string
var Prefix string
var VaultToken string
var Secrets string

func init() {
	config.LoadConfig()

	rootCmd.MarkPersistentFlagRequired("format")
	rootCmd.MarkPersistentFlagRequired("dirPath")
	rootCmd.MarkPersistentFlagRequired("prefix")
	rootCmd.MarkPersistentFlagRequired("secret")
	// rootCmd.MarkPersistentFlagRequired("")
	// rootCmd.MarkPersistentFlagRequired("")
	// rootCmd.MarkPersistentFlagRequired("")

	rootCmd.PersistentFlags().StringVarP(&VaultAddress, "vault_address", "a", "", "name of license for the project")
	rootCmd.PersistentFlags().StringVarP(&ClusterName, "cluster_name", "b", "", "name of license for the project")
	rootCmd.PersistentFlags().StringVarP(&TokenPath, "token_path", "c", "", "name of license for the project")
	rootCmd.PersistentFlags().StringVarP(&Prefix, "prefix", "d", "", "name of license for the project")
	rootCmd.PersistentFlags().StringVarP(&VaultToken, "vault_token", "e", "", "name of license for the project")
	rootCmd.PersistentFlags().StringVarP(&Secrets, "secrets", "f", "", "name of license for the project")

	setConfigFromFlags()
	checkRequiredFlags()
	// * harpocrates --format env --dirPath /tmp/secrets.env --prefix K8S_CLUSTER_ 'ES/data/someSecret'
	// * harpocrates --format env --dirPath /tmp/secrets.env 'ES/data/someSecret:DOCKER_,ES/data/something:K8S_CLUSTER_'
	// * harpocrates '{}'
	// * harpocrates --file /path/to/[yaml or json]
}

func setConfigFromFlags() {
	if VaultAddress != "" {
		config.Config.VaultAddress = VaultAddress
	}
	if ClusterName != "" {
		config.Config.ClusterName = ClusterName
	}
	if TokenPath != "" {
		config.Config.TokenPath = TokenPath
	}
	if VaultToken != "" {
		config.Config.VaultToken = VaultToken
	}
}

func checkRequiredFlags() {
	fmt.Println("checking...")
	fmt.Printf("%+v\n", &config.Config)

	var missing []string

	if config.Config.VaultAddress == "" {
		missing = append(missing, "vault_address")
	}
	if config.Config.ClusterName == "" {
		missing = append(missing, "cluster_name")
	}
	if config.Config.TokenPath == "" {
		missing = append(missing, "token_path")
	}
	if config.Config.VaultToken == "" {
		missing = append(missing, "vault_token")
	}

	fmt.Printf("The following config(s) are missing:\n%s\n", strings.Join(missing, ","))
	os.Exit(1)
}
