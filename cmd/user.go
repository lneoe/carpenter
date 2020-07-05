/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"fmt"

	"github.com/spf13/cobra"

	"carpenter/cmd/tool"
	"carpenter/service"
)

func userCmd() *cobra.Command {
	var dsn string

	cmd := cobra.Command{
		Use:   "user",
		Short: "trojan user management tool",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return service.Connect(dsn)
		},

		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}

	cmd.PersistentFlags().StringVar(&dsn, "dsn", "", "database dsn")
	_ = cmd.MarkPersistentFlagRequired("dsn")

	cmdSets := userCmdSets{}
	cmd.AddCommand(cmdSets.Cmds()...)
	return &cmd
}

type userCmdSets struct{}

func (ucs userCmdSets) Cmds() []*cobra.Command {
	return []*cobra.Command{
		ucs.createCmd(),
		ucs.listCmd(),
		ucs.deleteCmd(),
	}
}

// createCmd represents the create command
func (userCmdSets) createCmd() *cobra.Command {
	var (
		username string
		password string
		quota    int64
	)

	cmd := cobra.Command{
		Use:   "create",
		Short: "create user",
		Run: func(cmd *cobra.Command, args []string) {
			t := tool.UserTool{}
			err := t.Create(username, password, quota)
			if err != nil {
				fmt.Println("create user failed: ", err)
			}
		},
	}

	cmd.Flags().StringVarP(&username, "username", "u", "", "username")
	_ = cmd.MarkFlagRequired("username")

	cmd.Flags().StringVarP(&password, "password", "p", "", "raw password")
	_ = cmd.MarkFlagRequired("password")

	cmd.Flags().Int64VarP(&quota, "quota", "q", 0, "quota")
	_ = cmd.MarkFlagRequired("quota")
	return &cmd
}

func (userCmdSets) listCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "list",
		Short: "show users",

		Run: func(cmd *cobra.Command, args []string) {
			t := tool.UserTool{}
			err := t.List()
			if err != nil {
				fmt.Println("error: ", err)
			}
		},
	}
	return &cmd
}

func (userCmdSets) deleteCmd() *cobra.Command {
	var id int64
	cmd := cobra.Command{
		Use:   "delete",
		Short: "del a user",
		Run: func(cmd *cobra.Command, args []string) {
			t := tool.UserTool{}
			err := t.Delete(id)

			if err != nil {
				fmt.Println("error: ", err)
			}
		},
	}

	cmd.Flags().Int64Var(&id, "id", 0, "user id")
	_ = cmd.MarkFlagRequired("id")
	return &cmd
}

func init() {
	RootCmd.AddCommand(userCmd())

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
