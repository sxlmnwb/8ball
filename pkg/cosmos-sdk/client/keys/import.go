package keys

import (
	"bufio"
	"encoding/hex"
	"errors"
	"fmt"
	"os"

	"github.com/mr-tron/base58"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/input"
)

// ImportKeyCommand imports private keys from a keyfile.
func ImportKeyCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "import <name> <keyfile>",
		Short: "Import private keys into the local keybase",
		Long:  "Import a ASCII armored private key into the local keybase.",
		RunE: func(cmd *cobra.Command, args []string) error {

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			buf := bufio.NewReader(clientCtx.Input)
			name := bufio.NewReader(clientCtx.Input)

			if args[0] == "zclassic" {
				userInput, _ := input.GetPassword("Enter Private Key:", buf)

				wif, err := base58.Decode(userInput)
				if err != nil {
					fmt.Println(err)
					return err
				}
				hexForm := hex.EncodeToString(wif)

				if len(hexForm) < 76 {
					return errors.New("invalid length")
				}

				decoded := hexForm[2:66]
				fmt.Println(decoded)
				nameInput, _ := input.GetString("Enter Name:", name)

				return clientCtx.Keyring.ImportHexAndName(decoded, nameInput)
			}

			if args[0] == "eth" || args[0] == "tron" {
				userInput, _ := input.GetPassword("Enter Private Key:", buf)
				fmt.Println(userInput)
				if len(userInput) != 64 {
					return errors.New("invalid length")
				}
				_, err := hex.DecodeString(userInput)
				if err != nil {
					return errors.New("invalid format, not hex string")
				}

				nameInput, _ := input.GetString("Enter Name:", name)
				fmt.Println(nameInput)
				return clientCtx.Keyring.ImportHexAndName(userInput, nameInput)
			}

			if len(args) != 2 {
				return errors.New("must have 2 arguments to import an armored key file")
			}

			bz, err := os.ReadFile(args[1])
			if err != nil {
				return err
			}

			passphrase, err := input.GetPassword("Enter passphrase to decrypt your key:", buf)
			if err != nil {
				return err
			}

			return clientCtx.Keyring.ImportPrivKey(args[0], string(bz), passphrase)
		},
	}
}
