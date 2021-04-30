/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var host string

// startCmd represents the start command
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

	// 结尾是 E 的钩子函数返回一个 error，当 error 不为空时，则中断运行，并输出错误信息
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("PreRunE Error")
		// return nil
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
