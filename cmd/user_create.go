/*
Copyright © 2021 Joeyscat <zhouyu.fun@qq.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	icmd "github.com/joeyscat/mongoman/internal/cmd"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var password string
var roles string

// userCreateCmd represents the userCreate command
var userCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a user",
	Long: `Create a user

Example:
mongoman user create --uri mongodb://test_01:123456@127.0.0.1:27017/test --roles readWrite:test, --username test_03 --password 123456`,
	Run: func(cmd *cobra.Command, args []string) {
		if uri == "" {
			uri = viper.GetString("uri")
		}
		icmd.CreateUser(username, password, roles, uri, dbname)
	},
}

func init() {
	userCmd.AddCommand(userCreateCmd)

	userCreateCmd.Flags().StringVar(&password, "password", "", "password")
	userCreateCmd.Flags().StringVar(&roles, "roles", "", "roles")
}
