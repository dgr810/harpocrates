package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/BESTSELLER/harpocrates/config"
	"github.com/spf13/cobra"
	"gopkg.in/gookit/color.v1"
)

const (
	NOT_REQUIRED = false
	REQUIRED     = true
)

var (
	// Used for flags.
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "harpocrates",
		Short: "A generator for Cobra based Applications",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			// fmt.Println("filePath", filePath)
			verifyStuff(cmd)
		},
	}
)

func verifyStuff(cmd *cobra.Command) {
	verifySecretInput(cmd)
	verifyToken(cmd)
}

func verifyToken(cmd *cobra.Command) {
	if config.Config.TokenPath == "" && config.Config.VaultToken == "" {
		color.Red.Println("You need to specify either [--vault_token] or [--token_path]")
		cmd.Usage()
		os.Exit(1)
	}

	if config.Config.TokenPath != "" && config.Config.VaultToken != "" {
		color.Red.Println("You can't use [--vault_token] and [--token_path] at the same time")
		cmd.Usage()
		os.Exit(1)
	}
}

func verifySecretInput(cmd *cobra.Command) {
	if filePath == "" && config.Config.SecretJSON == "" {
		color.Red.Println("You need to specify either [--file] or [--json]")
		cmd.Usage()
		os.Exit(1)
	}

	if filePath != "" && config.Config.SecretJSON != "" {
		color.Red.Println("You can't use [--file] and [--json] at the same time")
		cmd.Usage()
		os.Exit(1)
	}
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

var filePath string

func init() {
	rootCmd.PersistentFlags().StringVar(&config.Config.VaultAddress, "vault_address", "", "Address of you Hashicorp Vault. E.g. https://127:0.0.1")
	rootCmd.PersistentFlags().StringVar(&config.Config.ClusterName, "cluster_name", "", "Used to login to Vault. We should really find a better way than this.")
	rootCmd.PersistentFlags().StringVar(&config.Config.TokenPath, "token_path", "", "Path to where your JWT token is located.")
	rootCmd.PersistentFlags().StringVar(&config.Config.VaultToken, "vault_token", "", "If you already have a Vault token, you can specify it here.")
	rootCmd.PersistentFlags().StringVar(&config.Config.Prefix, "prefix", "", "Prefix of your secret keys. E.g. if you specify 'FOO_' as the prefix and you have a secret called 'BAR' then it will be outputted as 'FOO_BAR'.")
	rootCmd.PersistentFlags().StringVar(&config.Config.SecretJSON, "json", "", "Some json/yaml can be inserted here.")
	rootCmd.PersistentFlags().StringVarP(&filePath, "file", "f", "", "Path to your yaml file containing which secrets you wanna fetch.")

	syncEnvToFlags()
}

// We should be able to do this better!
func syncEnvToFlags() {
	if config.Config.VaultAddress == "" {
		tryEnv("vault_address", &config.Config.VaultAddress, REQUIRED)
	}
	if config.Config.ClusterName == "" {
		tryEnv("cluster_name", &config.Config.ClusterName, REQUIRED)
	}
	if config.Config.TokenPath == "" {
		tryEnv("token_path", &config.Config.TokenPath, REQUIRED)
	}
	if config.Config.Prefix == "" {
		tryEnv("prefix", &config.Config.Prefix, NOT_REQUIRED)
	}
	if config.Config.VaultToken == "" {
		tryEnv("vault_token", &config.Config.VaultToken, REQUIRED)
	}
}

func tryEnv(env string, some *string, required bool) {
	envPrefix := "harpocrates"

	envVar, ok := os.LookupEnv(strings.ToUpper(fmt.Sprintf("%s_%s", envPrefix, env)))
	if ok == true && envVar != "" {
		*some = envVar
	} else {
		if required {
			rootCmd.MarkPersistentFlagRequired(env)
		}
	}
}
