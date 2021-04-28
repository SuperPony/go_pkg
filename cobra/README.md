# Index

- 介绍
- 核心概念
- 安装
- 使用

  - 创建应用
  - 添加子命令

- 常用选项说明

# 介绍

一款应用程序生成框架，用于创建自己的应用程序或命令行（Command）程序，从而开发以 Cobra 为基础的应用。目前 Docker、Kubernetes、Hugo 等著名项目都使用了 Cobra。

# 核心概念

- commands：命令行，代表行为。可细分为 rootCmd 和 subCmd。
- arguments：命令行参数。
- flags：命令行选型，代表对行为的改变。通常以 - 或 – 标识。

# 安装

`go get github.com/spf13/cobra`

# 使用

## 创建应用

`cobra init --pkg-name name`

## 添加子命令

`cobra add subCmdName`

# 常用选项说明

```
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var host string

var startCmd = &cobra.Command{
	Use:          "start",
	Short:        "启动服务 Short",
	Long:         `启动服务 Long`,
	SilenceUsage: true, // 当发生错误时，不输出 help
	// 自定义验证
	Args: func(cmd *cobra.Command, args []string) error {

		// 验证传入的 Arguments 是否为有效的 arg
		if err := cobra.OnlyValidArgs(cmd, args); err != nil {
			return err
		}

		return nil
	},

	// 设置允许的 args
	ValidArgs: []string{"host", "port"},

	// 在 Run 之前执行
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("PersistentPreRun")
	},

	// args 保存着 cmd 的 Arguments
	// flag 的 Argument 保存在对应的变量或通过 Flages().GetType(name) 获取
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetString("port")
		fmt.Println(port, host, args)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Type 系列用于注册当前 cmd 的 flags, --name
	// TypeP --name -shorthand
	startCmd.Flags().StringP("port", "p", "9090", "启动端口")
	// 获取指定 flag 输入的值
	// startCmd.Flags().GetString("key")

	// TypeVar 系列用于将 flag 绑定到指定变量,但是只有 --name 的格式
	// TypeVarP --name  -shorthand
	startCmd.Flags().StringVarP(&host, "host", "a", "127.0.0.1", "地址")

	// MarkFlagRequired 设置指定的 flag 为必选项
	// startCmd.MarkFlagRequired("port")

	// PersistentFlags 用于注册全局的 flags，通常在 root 中注册
	// rootCmd.PersistentFlags().String("name", "pony", "姓名")

	// 设置新的帮助文档
	// startCmd.SetHelpFunc(func(c *cobra.Command, s []string) {
	// 	fmt.Println("新的帮助文档")
	// })

}
```
