package addrdec

import (
	"fmt"
	"strings"

	"github.com/blocktree/go-owcdrivers/addressEncoder"
)

var (
	BARPublicKeyPrefix       = "PUB_"
	BARPublicKeyK1Prefix     = "PUB_K1_"
	BARPublicKeyR1Prefix     = "PUB_R1_"
	BARPublicKeyPrefixCompat = "BAR"

	//BAR stuff
	BAR_mainnetPublic = addressEncoder.AddressType{"bar", addressEncoder.BTCAlphabet, "ripemd160", "", 33, []byte(BARPublicKeyPrefixCompat), nil}
	// BAR_mainnetPrivateWIF           = AddressType{"base58", BTCAlphabet, "doubleSHA256", "", 32, []byte{0x80}, nil}
	// BAR_mainnetPrivateWIFCompressed = AddressType{"base58", BTCAlphabet, "doubleSHA256", "", 32, []byte{0x80}, []byte{0x01}}

	Default = AddressDecoderV2{}
)

//AddressDecoderV2
type AddressDecoderV2 struct {
	IsTestNet bool
}

//NewAddressDecoder 地址解析器
func NewAddressDecoderV2() *AddressDecoderV2 {
	decoder := AddressDecoderV2{}
	return &decoder
}

// AddressDecode decode address
func (dec *AddressDecoderV2) AddressDecode(pubKey string, opts ...interface{}) ([]byte, error) {

	var pubKeyMaterial string
	if strings.HasPrefix(pubKey, BARPublicKeyR1Prefix) {
		pubKeyMaterial = pubKey[len(BARPublicKeyR1Prefix):] // strip "PUB_R1_"
	} else if strings.HasPrefix(pubKey, BARPublicKeyK1Prefix) {
		pubKeyMaterial = pubKey[len(BARPublicKeyK1Prefix):] // strip "PUB_K1_"
	} else if strings.HasPrefix(pubKey, BARPublicKeyPrefixCompat) { // "BAR"
		pubKeyMaterial = pubKey[len(BARPublicKeyPrefixCompat):] // strip "BAR"
	} else {
		return nil, fmt.Errorf("public key should start with [%q | %q] (or the old %q)", BARPublicKeyK1Prefix, BARPublicKeyR1Prefix, BARPublicKeyPrefixCompat)
	}

	ret, err := addressEncoder.Base58Decode(pubKeyMaterial, addressEncoder.NewBase58Alphabet(BAR_mainnetPublic.Alphabet))
	if err != nil {
		return nil, addressEncoder.ErrorInvalidAddress
	}
	if addressEncoder.VerifyChecksum(ret, BAR_mainnetPublic.ChecksumType) == false {
		return nil, addressEncoder.ErrorInvalidAddress
	}

	return ret[:len(ret)-4], nil
}

// AddressEncode encode address
func (dec *AddressDecoderV2) AddressEncode(hash []byte, opts ...interface{}) (string, error) {
	data := addressEncoder.CatData(hash, addressEncoder.CalcChecksum(hash, BAR_mainnetPublic.ChecksumType))
	return string(BAR_mainnetPublic.Prefix) + addressEncoder.EncodeData(data, "base58", BAR_mainnetPublic.Alphabet), nil
}
