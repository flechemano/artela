package cmd

import (
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/version"
	jose "github.com/dvsekhvalnov/jose2go"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/cobra"
)

const flagKeyinfoFile = "file"
const flagKeyinfoPasswd = "passwd"

// Item is a thing stored on the keyring
type Item struct {
	Key         string
	Data        []byte
	Label       string
	Description string

	// Backend specific config
	KeychainNotTrustApplication bool
	KeychainNotSynchronizable   bool
}

func KeyInfoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "keyinfo",
		Short: "keyinfo retrieves private key and address",
		Long: fmt.Sprintf(`Keyinfo print the private key and address to StdOut.

Example:
$ %s keyinfo --file '/root/.artelad/keyring/mykey.info' --passwd 'test'
`, version.AppName),
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			kfile, _ := cmd.Flags().GetString(flagKeyinfoFile)
			kpasswd, _ := cmd.Flags().GetString(flagKeyinfoPasswd)
			privKey, err := readKeyStore(kfile, kpasswd)
			if err != nil {
				return err
			}
			if len(privKey) != 34 {
				return errors.New("read key store failed, priveKey length is not 34")
			}

			privateKey, err := crypto.ToECDSA(privKey[2:])
			if err != nil {
				return err
			}

			privateKeyBytes := privateKey.D.Bytes()
			privateKeyBytesPadded := make([]byte, 32)
			copy(privateKeyBytesPadded[32-len(privateKeyBytes):], privateKeyBytes)
			fmt.Println("private key: 0x", hex.EncodeToString(privateKeyBytesPadded))

			publicKey := privateKey.Public()
			publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
			if !ok {
				return errors.New("error casting public key to ECDSA")
			}
			// publicKeyBytes := crypto.CompressPubkey(publicKeyECDSA)
			// fmt.Println("public key: 0x", hex.EncodeToString(publicKeyBytes))

			fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
			fmt.Println("address: 0x", hex.EncodeToString(fromAddress[:]))
			return nil
		},
	}

	cmd.Flags().String(flagKeyinfoFile, "", "the fullpath of the keyinfo file")
	cmd.Flags().String(flagKeyinfoPasswd, "", "the password of the keyinfo file")

	return cmd
}

func readKeyStore(file string, passwd string) ([]byte, error) {
	fmt.Println("file: ", file)
	if !filepath.IsAbs(file) {
		home := os.Getenv("HOME")
		file = path.Join(home, file)
	}
	fmt.Println("file: ", file)

	bytes, err := os.ReadFile(file)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("file %s not found", file)
	} else if err != nil {
		return nil, err
	}

	payload, _, err := jose.Decode(string(bytes), passwd)
	if err != nil {
		return nil, err
	}

	decoded := &Item{}
	err = json.Unmarshal([]byte(payload), decoded)
	if err != nil {
		return nil, err
	}

	record, err := unmarshalRecord(decoded.Data)
	if err != nil {
		return nil, err
	}
	key := record.GetLocal().PrivKey
	return key.Value, nil
}

func unmarshalRecord(data []byte) (*keyring.Record, error) {
	record := &keyring.Record{}
	if err := record.Unmarshal(data); err != nil {
		return nil, err
	}
	return record, nil
}
