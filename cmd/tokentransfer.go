// Copyright © 2017 Weald Technology Trading
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

	"github.com/orinocopay/go-etherutils/cli"
	"github.com/orinocopay/go-etherutils/ens"
	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/util"
)

var tokenTransferAmount string
var tokenTransferToAddress string
var tokenTransferData string

// tokenTransferCmd represents the token transfer command
var tokenTransferCmd = &cobra.Command{
	Use:   "transfer",
	Short: "Transfer tokens to a given address",
	Long: `Transfer token from one address to another.  For example:

    ethereal token transfer --token=x --to=x --amount=y --passphrase=secret 0x5FfC014343cd971B7eb70732021E26C35B744cc4

In quiet mode this will return 0 if the transfer transaction is successfully sent, otherwise 1.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(len(args) == 1, quiet, "Requires a single address from which to transfer funds")
		cli.Assert(args[0] != "", quiet, "Sender address is required")

		fromAddress, err := ens.Resolve(client, args[0])
		cli.ErrCheck(err, quiet, "Failed to obtain sender for transfer")

		token, err := tokenContract(tokenStr)
		cli.ErrCheck(err, quiet, "Failed to obtain token address")

		decimals, err := token.Decimals(nil)
		cli.ErrCheck(err, quiet, "Failed to obtain token decimals")

		toAddress, err := ens.Resolve(client, tokenTransferToAddress)
		cli.ErrCheck(err, quiet, "Failed to obtain recipient for transfer")

		cli.Assert(tokenTransferAmount != "", quiet, "Require an amount to transfer with --to")
		amount, err := util.StringToTokenValue(tokenTransferAmount, decimals)
		cli.ErrCheck(err, quiet, "Invalid amount")

		// Obtain the balance of the address
		balance, err := token.BalanceOf(nil, fromAddress)
		cli.ErrCheck(err, quiet, "Failed to obtain balance of address from which to send funds")
		cli.Assert(balance.Cmp(amount) > 0, quiet, fmt.Sprintf("Balance of %s insufficient for transfer", util.TokenValueToString(balance, decimals, false)))

		fmt.Printf("Sending %s from %s to %s (balance %s)\n", util.TokenValueToString(amount, decimals, false), fromAddress.Hex(), toAddress.Hex(), util.TokenValueToString(balance, decimals, false))
		// Create and sign the transaction
		//		signedTx, err := createSignedTransaction(fromAddress, &toAddress, amount, data)
		//		cli.ErrCheck(err, quiet, "Failed to create transaction")

		//		err = client.SendTransaction(context.Background(), signedTx)
		//		cli.ErrCheck(err, quiet, "Failed to send transaction")

		if quiet {
			os.Exit(0)
		}
		//		fmt.Println(signedTx.Hash().Hex())
	},
}

func init() {
	tokenCmd.AddCommand(tokenTransferCmd)
	tokenFlags(tokenTransferCmd)
	tokenTransferCmd.Flags().StringVar(&tokenTransferAmount, "amount", "", "Amount of Ether to transfer")
	tokenTransferCmd.Flags().StringVar(&tokenTransferToAddress, "to", "", "Address to which to transfer Ether")
	addTransactionFlags(tokenTransferCmd, "Passphrase for the address that holds the tokens")
}
