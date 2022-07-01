package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"

	"github.com/bobwanglong/goyaml/command"
	"github.com/spf13/cobra"
)

//go:embed deploy.yaml
var deploy string

var generateDeployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "deploy 子命令.",
	Long:  "这是一个deploy 子命令",
	Run:   generateDeploy,
}
var (
	availableData = map[string]interface{}{
		"Namespace":     "my-namespace",
		"Name":          "my-app-name",
		"Image":         "my-app-image",
		"ContainerPort": 8080,
		"NfsFilePath":   "/data/nfs",
		"NfsServer":     "127.0.0.1",
	}

	availableFunc = template.FuncMap{
		"GeneratePassword": GeneratePasswordFunc,
	}
)
var (
	nfsFilePath string
	nfsServer   string
	outpath     string

	name      string
	image     string
	namespace string

	containerPort int
)

func init() {
	cobra.EnablePrefixMatching = true
}

var exampleUsage = `  # List APIs.
  goyaml deploy -name my-app -namespace my-namespace -image my-app-image -containerPort 8080 
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
		// command.ObjectCmd(),
		deployCmd(),
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

func ReadEmbed() ([]byte, error) {
	tmpl, err := template.New("deploy.yaml").Funcs(availableFunc).Parse(deploy)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, availableData); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
func Read(filePath string) ([]byte, error) {
	name := filepath.Base(filePath)
	tmpl, err := template.New(name).Funcs(availableFunc).ParseFiles(filePath)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, availableData); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func GeneratePasswordFunc() (string, error) {
	return "bobwang", nil
}

func GeneralWrite(fileName string, buf []byte) {
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println("open file error :", err)
		return
	}
	defer f.Close()

	_, err = f.Write(buf)
	if err != nil {
		log.Println(err)
		return
	}
}

func generateDeploy(cmd *cobra.Command, args []string) {
	fmt.Println("name:", name)
	fmt.Println("name:", namespace)
	fmt.Println("name:", image)
	fmt.Println("name:", containerPort)
	fmt.Println("filePath:", outpath)
	// 赋值
	availableData["Namespace"] = namespace
	availableData["Name"] = name

	availableData["Image"] = image
	availableData["ContainerPort"] = containerPort
	availableData["NfsFilePath"] = nfsFilePath
	availableData["NfsServer"] = nfsServer

	buf, _ := ReadEmbed()
	// GeneralWrite("./cm.yaml", buf)
	GeneralWrite(outpath, buf)

}

func deployCmd() *cobra.Command {
	generateDeployCmd.Flags().IntVar(&containerPort, "containerPort", 8080, "Set containerPort Int")
	generateDeployCmd.Flags().StringVar(&name, "name", "my-app", "Set app name")
	generateDeployCmd.Flags().StringVar(&namespace, "namespace", "defalt", "Set app namespace")
	generateDeployCmd.Flags().StringVar(&image, "image", "app-image", "Set app image")
	generateDeployCmd.Flags().StringVar(&outpath, "outpath", "./app.yaml", "Set yaml file path")

	generateDeployCmd.Flags().StringVar(&nfsFilePath, "nfsFilePath", "/data/nfs", "Set nfs file path")
	generateDeployCmd.Flags().StringVar(&nfsServer, "nfsServer", "127.0.0.1", "Set nfs server address")

	return generateDeployCmd
}
