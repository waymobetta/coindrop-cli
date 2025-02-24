// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
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

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/waymobetta/go-coindrop-api/services/ethereum"
)

// deployCmd represents the deploy command
var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy an ERC-721 contract",
	Long: `
deploy an ERC-721 contract to the Ethereum network

usage:
coindrop-cli deploy <token_name> <token_symbol>

example:
coindrop-cli deploy --name=<token_name> --symbol=<token_symbol>
`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		symbol, _ := cmd.Flags().GetString("symbol")
		if name == "" && symbol == "" {
			fmt.Println("[!] usage: coindrop-cli deploy --name=<token_name> --symbol=<token_symbol>")
			os.Exit(1)
		}

		address, err := ethereum.DeployERC721Contract(
			name,
			symbol,
		)
		if err != nil {
			log.Printf("[cmd/deploy] unable to deploy contract: %v\n", err)
		}
		fmt.Println("address:", address.Hex())
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deployCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	deployCmd.Flags().StringP("name", "n", "", "Name of the ERC-721 contract")
	deployCmd.Flags().StringP("symbol", "s", "", "Symbol of the ERC-721")
}
