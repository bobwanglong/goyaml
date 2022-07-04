package command

import (
	"bytes"
	_ "embed"
	"log"
	"os"
	"text/template"

	"fmt"

	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

//go:embed deploy.yaml
var deploy string

var (
	nfsFilePath string
	nfsServer   string
	outpath     string

	name      string
	image     string
	namespace string

	containerPort int

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

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "version 子命令.",
	Long:  "这是一个version 子命令",
	Run:   runVersion,
}
var generateDeployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "deploy 子命令.",
	Long:  "这是一个deploy 子命令",
	Run:   generateDeploy,
}

func DeploymentCmd() *cobra.Command {
	generateDeployCmd.Flags().IntVar(&containerPort, "containerPort", 8080, "Set containerPort Int")
	generateDeployCmd.Flags().StringVar(&name, "name", "my-app", "Set app name")
	generateDeployCmd.Flags().StringVar(&namespace, "namespace", "defalt", "Set app namespace")
	generateDeployCmd.Flags().StringVar(&image, "image", "app-image", "Set app image")
	generateDeployCmd.Flags().StringVar(&outpath, "outpath", "./app.yaml", "Set yaml file path")
	generateDeployCmd.Flags().StringVar(&nfsFilePath, "nfsFilePath", "/data/nfs", "Set nfs file path")
	generateDeployCmd.Flags().StringVar(&nfsServer, "nfsServer", "127.0.0.1", "Set nfs server address")

	return generateDeployCmd
}
func runVersion(cmd *cobra.Command, args []string) {
	// TODO 这里处理version子命令

	fmt.Println("version is 1.0.0")
}

func GeneratePasswordFunc() (string, error) {
	return "yourPasswd", nil
}

func GeneralWrite(fileName string, buf []byte) {
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println("open file error :", err)
		return
	}
	// 关闭文件
	defer f.Close()
	// 字节方式写入
	_, err = f.Write(buf)
	if err != nil {
		log.Println(err)
		return
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

func generateDeploy(cmd *cobra.Command, args []string) {
	// fmt.Println("name:", name)
	// fmt.Println("name:", namespace)
	// fmt.Println("name:", image)
	// fmt.Println("name:", containerPort)
	// fmt.Println("filePath:", outpath)
	// 赋值
	availableData["Namespace"] = namespace
	availableData["Name"] = name

	availableData["Image"] = image
	availableData["ContainerPort"] = containerPort
	availableData["NfsFilePath"] = nfsFilePath
	availableData["NfsServer"] = nfsServer

	buf, _ := ReadEmbed()
	GeneralWrite(outpath, buf)
	// 打印彩色
	fmt.Println(aurora.Green("Done."))

}
