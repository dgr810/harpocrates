package cmd

import (
	"fmt"

	"github.com/BESTSELLER/harpocrates/config"
	"github.com/BESTSELLER/harpocrates/files"
	"github.com/BESTSELLER/harpocrates/util"
	"github.com/spf13/cobra"
	"gopkg.in/gookit/color.v1"
)

var (
	filePath string
	rootCmd  = &cobra.Command{
		Use:   "harpocrates",
		Short: fmt.Sprintf("%sThis application will fetch secrets from Hashicorp Vault", color.Blue.Sprintf("\"Harpocrates was the god of silence, secrets and confidentiality\"\n")),
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			verifyStuff(cmd)

			if filePath != "" {
				config.Config.Input = files.ReadFile(filePath)
			}

			input := util.ReadInput(config.Config.Input)
			allSecrets := util.ExtractSecrets(input)
			fileName := fmt.Sprintf("secrets.%s", input.Format)

			if input.Format == "json" {
				files.WriteFile(input.DirPath, fileName, files.FormatAsJSON(allSecrets))
			}

			if input.Format == "env" {
				files.WriteFile(input.DirPath, fileName, files.FormatAsENV(allSecrets))
			}

			color.Green.Printf("Secrets written to file: %s/%s\n", input.DirPath, fileName)
		},
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	setupFlags()
	syncEnvToFlags()
}

var someSecret util.SecretJSON

func setupFlags() {
	rootCmd.PersistentFlags().StringVar(&config.Config.VaultAddress, "vault_address", "", "Address of you Hashicorp Vault. E.g. https://127:0.0.1")
	rootCmd.PersistentFlags().StringVar(&config.Config.ClusterName, "cluster_name", "", "Used to login to Vault. We should really find a better way than this.")
	rootCmd.PersistentFlags().StringVar(&config.Config.TokenPath, "token_path", "", "Path to where your JWT token is located.")
	rootCmd.PersistentFlags().StringVar(&config.Config.VaultToken, "vault_token", "", "If you already have a Vault token, you can specify it here.")
	rootCmd.PersistentFlags().StringVar(&config.Config.Prefix, "prefix", "", "Prefix of your secret keys. E.g. if you specify 'FOO_' as the prefix and you have a secret called 'BAR' then it will be outputted as 'FOO_BAR'.")
	rootCmd.PersistentFlags().StringVar(&config.Config.Input, "inline", "", "Some JSON or YAML can be inserted here.")
	rootCmd.PersistentFlags().Var(&someSecret, "secrets", "This is how we party")
	rootCmd.PersistentFlags().StringVarP(&filePath, "file", "f", "", "Path to your yaml file containing which secrets you wanna fetch.")
}
