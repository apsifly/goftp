// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

var ftpUser, ftpPassword, ftpServer, ftpType string
var ftpPort int
var ftpPassive bool

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "client",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initMode)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.client.yaml)")
	RootCmd.PersistentFlags().StringVarP(&ftpUser, "user", "u", "nobody", "ftp user")
	RootCmd.PersistentFlags().StringVarP(&ftpPassword, "password", "P", "nopass", "ftp password")
	RootCmd.PersistentFlags().StringVarP(&ftpServer, "server", "s", "localhost", "ftp server host")
	RootCmd.PersistentFlags().IntVarP(&ftpPort, "port", "p", 21, "ftp server port")
	RootCmd.PersistentFlags().StringVarP(&ftpType, "type", "t", "server", "ftp server mode (ascii/binary, default provided by server)")
	ftpPassive = *RootCmd.PersistentFlags().Bool("passive", false, "passive mode")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initMode() {
	switch ftpType { //convert type to protocol standard values
	case "binary":
		ftpType = "L"
	case "ascii":
		ftpType = "A"
	default:
		ftpType = "server" //will be converted after SYST command
	}
}
