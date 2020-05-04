package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/BESTSELLER/harpocrates/config"
	"github.com/BESTSELLER/harpocrates/util"
	"github.com/spf13/cobra"
	"gopkg.in/gookit/color.v1"
)

const (
	NOT_REQUIRED = false
	REQUIRED     = true
)

// We should be able to do this better!
func syncEnvToFlags() {
	if config.Config.VaultAddress == "" {
		tryEnv("vault_address", &config.Config.VaultAddress, REQUIRED)
	}
	if config.Config.ClusterName == "" {
		tryEnv("cluster_name", &config.Config.ClusterName, REQUIRED)
	}
	if config.Config.TokenPath == "" {
		tryEnv("token_path", &config.Config.TokenPath, NOT_REQUIRED)
	}
	if config.Config.Prefix == "" {
		tryEnv("prefix", &config.Config.Prefix, NOT_REQUIRED)
	}
	if config.Config.VaultToken == "" {
		tryEnv("vault_token", &config.Config.VaultToken, NOT_REQUIRED)
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

func verifyStuff(cmd *cobra.Command) {
	verifySecretInput(cmd)
	verifyToken(cmd)
}

func verifySecretInput(cmd *cobra.Command) {
	if len(someSecret.Secrets) > 0 {
		fmt.Println("LOTS of secrets !")
		fmt.Println(someSecret.Secrets)
		return
	} else {
		fmt.Println("no secrets")
	}

	if filePath == "" && config.Config.Input == "" {
		color.Red.Println("You need to specify either [--file] or [--inline]")
		cmd.Usage()
		os.Exit(1)
	}

	if filePath != "" && config.Config.Input != "" {
		color.Red.Println("You can't use [--file] and [--inline] at the same time")
		cmd.Usage()
		os.Exit(1)
	}
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

	util.GetVaultToken()
}
