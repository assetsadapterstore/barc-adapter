package barc

import (
	"fmt"

	"github.com/denkhaus/bitshares/config"
	"github.com/denkhaus/bitshares/types"
)

//Decrypt calculates a shared secret by the receivers private key
//and the senders public key, then decrypts the given memo message.
func Decrypt(msg, fromPub, toPub, wif string) (string, error) {
	var buf types.Buffer
	ret := config.FindByID(ChainIDBAR)
	if ret == nil {
		config.Add(config.ChainConfig{
			Name:      "BAR",
			CoreAsset: "BAR",
			Prefix:    "BAR",
			ID:        ChainIDBAR,
		})
	}
	config.SetCurrent(ChainIDBAR)

	from, err := types.NewPublicKeyFromString(fromPub)
	if err != nil {
		return "", fmt.Errorf("NewPublicKeyFromString: %v", err)
	}
	to, err := types.NewPublicKeyFromString(toPub)
	if err != nil {
		return "", fmt.Errorf("NewPublicKeyFromString: %v", err)
	}

	buf.FromString(msg)

	memo := types.Memo{
		From:    *from,
		To:      *to,
		Message: buf,
	}

	priv, err := types.NewPrivateKeyFromWif(wif)
	if err != nil {
		return "", fmt.Errorf("NewPrivateKeyFromWif: %v", err)
	}

	m, err := memo.Decrypt(priv)
	if err != nil {
		return "", fmt.Errorf("Decrypt: %v", err)
	}

	return m, nil
}
