package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "version 子命令.",
	Long:  "这是一个version 子命令",
	Run:   runVersion,
}

// func ObjectCmd() *cobra.Command {
// 	// cmd := &cobra.Command{
// 	// 	Use:   "deploy",
// 	// 	Short: "View k8s deploy yaml",
// 	// }

// 	generateDeployCmd.Flags().IntVar(&containerPort, "containerPort", 8080, "Set containerPort Int")
// 	generateDeployCmd.Flags().StringVar(&name, "name", "my-app", "Set app name")
// 	generateDeployCmd.Flags().StringVar(&namespace, "namespace", "defalt", "Set app namespace")
// 	generateDeployCmd.Flags().StringVar(&image, "image", "app-image", "Set app image")
// 	generateDeployCmd.Flags().StringVar(&outpath, "outpath", "./app.yaml", "Set yaml file path")

// 	return generateDeployCmd
// }
func runVersion(cmd *cobra.Command, args []string) {
	// TODO 这里处理version子命令

	fmt.Println("version is 1.0.0")
}

// func generateDeploy(cmd *cobra.Command, args []string) {
// 	fmt.Println("name:", name)
// 	fmt.Println("name:", namespace)
// 	fmt.Println("name:", image)
// 	fmt.Println("name:", containerPort)
// 	fmt.Println("filePath:", outpath)
// 	// 赋值
// 	availableData["Namespace"] = namespace
// 	availableData["Name"] = name

// 	availableData["Image"] = image
// 	availableData["ContainerPort"] = containerPort

// 	buf, _ := ReadEmbed()
// 	// GeneralWrite("./cm.yaml", buf)
// 	GeneralWrite(outpath, buf)

// }

// func Read(filePath string) ([]byte, error) {
// 	name := filepath.Base(filePath)
// 	tmpl, err := template.New(name).Funcs(availableFunc).ParseFiles(filePath)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var buf bytes.Buffer
// 	if err := tmpl.Execute(&buf, availableData); err != nil {
// 		return nil, err
// 	}

// 	return buf.Bytes(), nil
// }

// func GeneratePasswordFunc() (string, error) {
// 	return "bobwang", nil
// }

// func GeneralWrite(fileName string, buf []byte) {
// 	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
// 	if err != nil {
// 		log.Println("open file error :", err)
// 		return
// 	}
// 	// 关闭文件
// 	defer f.Close()
// 	// 字节方式写入
// 	_, err = f.Write(buf)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// }

// func ReadEmbed() ([]byte, error) {
// 	fmt.Println("deploy content:", deploy)
// 	tmpl, err := template.New("deploy.yaml").Funcs(availableFunc).Parse(deploy)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var buf bytes.Buffer
// 	if err := tmpl.Execute(&buf, availableData); err != nil {
// 		return nil, err
// 	}

// 	return buf.Bytes(), nil
// }
