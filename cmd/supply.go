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
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"github.com/waymobetta/go-coindrop-api/services/ethereum"
)

// supplyCmd represents the supply command
var supplyCmd = &cobra.Command{
	Use:   "supply",
	Short: "Get the total supply of ERC-721s minted",
	Long: `
get the total supply of ERC-721s minted from a contract

usage:
coindrop-cli supply --contract=<contract_address>

example:
coindrop-cli supply --contract=0x600ec79f2B258d7cc625AE80267Eb23689be417b
`,
	Run: func(cmd *cobra.Command, args []string) {
		contract, _ := cmd.Flags().GetString("contract")
		if contract == "" {
			fmt.Println("[!] usage: coindrop-cli supply --contract=<contract_address>")
			os.Exit(1)
		}

		contractAddress := common.HexToAddress(contract)
		totalSupply, err := ethereum.GetTotalSupply(
			contractAddress,
		)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("supply:", totalSupply)
	},
}

func init() {
	rootCmd.AddCommand(supplyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// supplyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	supplyCmd.Flags().StringP("contract", "c", "", "Address of the ERC-721 contract")
}
