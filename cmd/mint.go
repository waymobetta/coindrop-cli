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
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
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
coindrop-cli mint --id=<token_id> --contract=<contract_address> --recipient=<recipient_address> --uri=<token_uri>

example:
coindrop-cli mint --id=0 --contract=0x600ec79f2B258d7cc625AE80267Eb23689be417b --recipient=0xfedc485ab2c87529fb13414c57e391a98fd113ef --uri=bob-archaeologist
`,
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString("id")
		contract, _ := cmd.Flags().GetString("contract")
		recipient, _ := cmd.Flags().GetString("recipient")
		uri, _ := cmd.Flags().GetString("uri")
		if id == "" && contract == "" && recipient == "" && uri == "" {
			fmt.Println("[!] usage: coindrop-cli mint --id=0 --contract=0x600ec79f2B258d7cc625AE80267Eb23689be417b --recipient=0xfedc485ab2c87529fb13414c57e391a98fd113ef --uri=bob-archaeologist")
			os.Exit(1)
		}

		tokenId, err := strconv.Atoi(id)
		if err != nil {
			log.Fatal(err)
		}
		contractAddress := common.HexToAddress(contract)
		recipientAddress := common.HexToAddress(recipient)
		tokenURI := uri
		tx, err := ethereum.MintERC721Token(
			tokenId,
			contractAddress,
			recipientAddress,
			tokenURI,
		)
		if err != nil {
			log.Printf("[cmd/mint] unable to mint token: %v\n", err)
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
	mintCmd.Flags().StringP("id", "i", "", "Unique ERC-721 Token ID")
	mintCmd.Flags().StringP("contract", "c", "", "Address of the ERC-721 contract")
	mintCmd.Flags().StringP("recipient", "r", "", "Address of the ERC-721 recipient")
	mintCmd.Flags().StringP("uri", "u", "", "ERC-721 Token URI")
}
