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
	"log"

	"github.com/spf13/cobra"
	"github.com/waymobetta/go-coindrop-api/services/ethereum"
)

// deployCmd represents the deploy command
var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "deploy an ERC-721 contract",
	Long: `
deploy an ERC-721 contract to the Ethereum network

usage:
coindrop-cli deploy <token_name> <token_symbol>

example:
coindrop-cli deploy coindrop-1 cd1
`,
	Run: func(cmd *cobra.Command, args []string) {
		tokenName := args[0]
		tokenSymbol := args[1]

		address, err := ethereum.DeployERC721Contract(
			tokenName,
			tokenSymbol,
		)
		if err != nil {
			log.Fatal(err)
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
	// deployCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
