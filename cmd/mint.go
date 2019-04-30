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
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"github.com/waymobetta/go-coindrop-api/services/ethereum"
)

// mintCmd represents the mint command
var mintCmd = &cobra.Command{
	Use:   "mint",
	Short: "Mint an ERC-721 for a specific address",
	Long: `
mint an ERC-721 for an address

usage:
coindrop-cli mint <token_id> <contract_address> <recipient_address> <token_uri>

example:
coindrop-cli mint 0 0x600ec79f2B258d7cc625AE80267Eb23689be417b 0xfedc485ab2c87529fb13414c57e391a98fd113ef bob-archaeologist
`,
	Run: func(cmd *cobra.Command, args []string) {
		tokenId, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}
		contractAddress := common.HexToAddress(args[1])
		recipientAddress := common.HexToAddress(args[2])
		tokenURI := args[3]
		tx, err := ethereum.MintERC721Token(
			tokenId,
			contractAddress,
			recipientAddress,
			tokenURI,
		)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("transaction:", tx)
	},
}

func init() {
	rootCmd.AddCommand(mintCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mintCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mintCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
