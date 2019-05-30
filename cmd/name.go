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

	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/waymobetta/go-coindrop-api/services/ethereum"
)

// nameCmd represents the name command
var nameCmd = &cobra.Command{
	Use:   "name",
	Short: "Get the token name of ERC-721s minted",
	Long: `
get the token name of ERC-721s minted from a given contract

usage:
coindrop-cli name --contract=<contract_address>

example:
coindrop-cli name --contract=0x600ec79f2B258d7cc625AE80267Eb23689be417b
`,
	Run: func(cmd *cobra.Command, args []string) {
		contract, _ := cmd.Flags().GetString("contract")
		if contract == "" {
			fmt.Println("[!] usage: coindrop-cli name --contract=<contract_address>")
			os.Exit(1)
		}

		contractAddress := common.HexToAddress(contract)
		tokenName, err := ethereum.GetTokenName(
			contractAddress,
		)
		if err != nil {
			log.Printf("[cmd/name] unable to retrieve token name: %v\n", err)
		}
		fmt.Println("name:", tokenName)
	},
}

func init() {
	rootCmd.AddCommand(nameCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nameCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nameCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	nameCmd.Flags().StringP("contract", "c", "", "Address of the ERC-721 contract")
}
