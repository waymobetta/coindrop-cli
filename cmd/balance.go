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

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"github.com/waymobetta/go-coindrop-api/services/ethereum"
)

// balanceCmd represents the balance command
var balanceCmd = &cobra.Command{
	Use:   "balance",
	Short: "Get the ERC-721 balance of a specific address",
	Long: ` 
get the ERC-721 balance of a specific address

usage:
coindrop-cli balance <owner_wallet_address> <contract_address>

example:
coindrop-cli balance 0xfedc485ab2c87529fb13414c57e391a98fd113ef 0x600ec79f2B258d7cc625AE80267Eb23689be417b
`,
	Run: func(cmd *cobra.Command, args []string) {

		ownerAddress := common.HexToAddress(args[0])
		contractAddress := common.HexToAddress(args[1])

		ownerBalance, err := ethereum.GetBalanceOf(
			ownerAddress,
			contractAddress,
		)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("balance:", ownerBalance)
	},
}

func init() {
	rootCmd.AddCommand(balanceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// balanceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// balanceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
