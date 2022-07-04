package main

import (
	"os"

	"github.com/bobwanglong/goyaml/command"
	"github.com/spf13/cobra"
)

func init() {
	cobra.EnablePrefixMatching = true
}

var exampleUsage = `  # deployment demo.
  goyaml deploy --name hello --namespace myspace --image="my-app-image:latest" --containerPort=9000 --outpath demo.yaml --nfsFilePath="/data/nfs/demo" --nfsServer=127.0.0.1
`

func main() {
	rootCmd := &cobra.Command{
		Use:        "goyaml",
		Short:      "A command line admin tool for Generate k8s yaml.",
		Example:    exampleUsage,
		SuggestFor: []string{"goyaml"},
		// PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// 	switch command.CommandlineGlobalFlags.OutputFormat {
		// 	case "yaml", "json":
		// 	default:
		// 		command.ExitWithErrorf("unsupported output format: %s",
		// 			command.CommandlineGlobalFlags.OutputFormat)
		// 	}
		// },
	}

	completionCmd := &cobra.Command{
		Use:   "completion bash|zsh",
		Short: "Output shell completion code for the specified shell (bash or zsh)",
		Run: func(cmd *cobra.Command, args []string) {
			switch args[0] {
			case "bash":
				rootCmd.GenBashCompletion(os.Stdout)
			case "zsh":
				rootCmd.GenZshCompletion(os.Stdout)
			default:
				command.ExitWithErrorf("unsupported shell %s, expecting bash or zsh", args[0])
			}
		},
		Args: cobra.ExactArgs(1),
	}

	rootCmd.AddCommand(
		command.DeploymentCmd(),
		// deployCmd(),
		command.VersionCmd,
		completionCmd,
	)

	// rootCmd.PersistentFlags().StringVar(&command.CommandlineGlobalFlags.Server,
	// 	"server", "localhost:2381", "The address of the Easegress endpoint")
	// rootCmd.PersistentFlags().StringVarP(&command.CommandlineGlobalFlags.OutputFormat,
	// 	"output", "o", "yaml", "Output format(json, yaml)")

	err := rootCmd.Execute()
	if err != nil {
		command.ExitWithError(err)
	}
}
