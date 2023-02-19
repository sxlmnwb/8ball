package keeper

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"

	"eightball/x/eightball/types"

	cosmosbase58 "github.com/cosmos/btcutil/base58"
	"github.com/cosmos/btcutil/bech32"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/mr-tron/base58"
	"github.com/tendermint/tendermint/libs/log"
	"golang.org/x/crypto/ripemd160"
)

type (
	Keeper struct {
		cdc           codec.BinaryCodec
		storeKey      storetypes.StoreKey
		memKey        storetypes.StoreKey
		paramstore    paramtypes.Subspace
		stakingKeeper types.StakingKeeper
		accountKeeper types.AccountKeeper
		bankKeeper    types.BankKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	stakingKeeper types.StakingKeeper,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
) *Keeper {

	// ensure module account is set
	if addr := accountKeeper.GetModuleAddress(types.ModuleName); addr == nil {

		panic("the module account has not been set")
	}

	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		accountKeeper: accountKeeper,
		stakingKeeper: stakingKeeper,
		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		paramstore:    ps,
		bankKeeper:    bankKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func getZclassicAddress(xPoint string, yPoint string) (string, bool) {
	pkBytes := pubKeyBytes(xPoint, yPoint)

	pk, _ := secp256k1.ParsePubKey(pkBytes)

	compressedPk := pk.SerializeCompressed()
	pubKeyHex := hexutil.Encode(compressedPk)
	bitcoinAddress, err := generateZclassicAddress(pubKeyHex)
	if err != nil {
		return "", false
	}
	return bitcoinAddress, true
}

func generateZclassicAddress(publicKeyHex string) (string, error) {
	verPublicKeyHash, err := hashZclassicPublicKey(publicKeyHex)
	if err != nil {
		return "", err
	}

	checkSum := checksumKeyHash(verPublicKeyHash)
	zclassicAddressHex := encodeKeyHash(verPublicKeyHash, checkSum)

	return zclassicAddressHex, nil
}

func getTronAddress(xPoint string, yPoint string) (string, bool) {
	pkBytes := pubKeyBytes(xPoint, yPoint)
	pk, err := secp256k1.ParsePubKey(pkBytes)
	if err != nil {
		return "", false
	}
	eKey := pk.ToECDSA()
	addressBytes := ethcrypto.PubkeyToAddress(*eKey).Bytes()

	encoded := cosmosbase58.CheckEncode(addressBytes, byte(0x41))
	return encoded, true
}

// Convert any TRON address to an 8ball Address
func tronTo8ballAddress(tronAddr string) string {
	result, _, _ := cosmosbase58.CheckDecode(tronAddr)

	//result = result[1:]
	conv, err := bech32.ConvertBits(result, 8, 5, true)
	if err != nil {
		fmt.Println("Error:", err)
	}
	encoded, err := bech32.Encode("8ball", conv)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return encoded
}

func getEightballAddress(xPoint string, yPoint string) (string, bool) {
	tronAddr, ok := getTronAddress(xPoint, yPoint)
	if !ok {
		return "", false
	}
	eballAddr := tronTo8ballAddress(tronAddr)
	return eballAddr, true
}

func getEthereumAddress(xPoint string, yPoint string) (string, bool) {
	pkBytes := pubKeyBytes(xPoint, yPoint)
	pk, err := secp256k1.ParsePubKey(pkBytes)
	if err != nil {
		return "", false
	}
	eKey := pk.ToECDSA()
	return ethcrypto.PubkeyToAddress(*eKey).String(), true
}

func hashZclassicPublicKey(publicKeyHex string) ([]byte, error) {

	// 1 - SHA-256 HASH
	publicKeyByte, _ := hex.DecodeString(publicKeyHex)
	publicSHA256 := sha256.Sum256(publicKeyByte)

	// 2 - RIPEMD-160 HASH
	RIPEMD160Hasher := ripemd160.New()
	_, err := RIPEMD160Hasher.Write(publicSHA256[:])
	if err != nil {
		return nil, err
	}
	publicKeyHash := RIPEMD160Hasher.Sum(nil)

	// VERSION (ZCLASSIC VERSION BYTES FOR MAIN NET)
	version := make([]byte, 2)
	version[0] = 0x1C
	version[1] = 0xB8

	// 3 - CONCAT
	verPublicKeyHash := append(version, publicKeyHash...)

	return verPublicKeyHash, err
}

func checkValidPublicKey(xPoint string, yPoint string) bool {
	n := new(big.Int)
	n, ok := n.SetString(xPoint, 10)
	if !ok {
		return false
	}

	m := new(big.Int)
	m, ok = m.SetString(yPoint, 10)
	if !ok {
		return false
	}

	xHexStr := n.Text(16)
	yHexStr := m.Text(16)

	xBytes, err := hex.DecodeString(xHexStr)
	if err != nil {
		return false
	}

	yBytes, err := hex.DecodeString(yHexStr)
	if err != nil {
		return false
	}

	pubKeyBytes := []byte{0x04}
	pubKeyBytes = append(pubKeyBytes, xBytes...)
	pubKeyBytes = append(pubKeyBytes, yBytes...)
	_, err = secp256k1.ParsePubKey(pubKeyBytes)

	return err == nil
}

func getBitcoinAddress(xPoint string, yPoint string) (string, bool) {
	pkBytes := pubKeyBytes(xPoint, yPoint)

	pk, _ := secp256k1.ParsePubKey(pkBytes)
	key := pk.ToECDSA()

	pubKeyHex := hex.EncodeToString(ethcrypto.FromECDSAPub(key))

	bitcoinAddress, err := generateBitcoinAddress(pubKeyHex)
	if err != nil {
		return "", false
	}
	return bitcoinAddress, true
}

func pubKeyBytes(xPoint string, yPoint string) []byte {
	n := new(big.Int)
	n, _ = n.SetString(xPoint, 10)

	m := new(big.Int)
	m, _ = m.SetString(yPoint, 10)

	xHexStr := n.Text(16)
	yHexStr := m.Text(16)

	xBytes, _ := hex.DecodeString(xHexStr)

	yBytes, _ := hex.DecodeString(yHexStr)

	pubKeyBytes := []byte{0x04}
	pubKeyBytes = append(pubKeyBytes, xBytes...)
	pubKeyBytes = append(pubKeyBytes, yBytes...)

	return pubKeyBytes
}

func generateBitcoinAddress(publicKeyHex string) (string, error) {
	verPublicKeyHash, err := hashPublicKey(publicKeyHex)
	if err != nil {
		return "", err
	}

	checkSum := checksumKeyHash(verPublicKeyHash)
	bitcoinAddressHex := encodeKeyHash(verPublicKeyHash, checkSum)

	return bitcoinAddressHex, nil
}

func hashPublicKey(publicKeyHex string) ([]byte, error) {

	publicKeyByte, _ := hex.DecodeString(publicKeyHex)
	publicSHA256 := sha256.Sum256(publicKeyByte)

	RIPEMD160Hasher := ripemd160.New()
	_, err := RIPEMD160Hasher.Write(publicSHA256[:])
	if err != nil {
		return nil, err
	}
	publicKeyHash := RIPEMD160Hasher.Sum(nil)

	version := make([]byte, 1)
	version[0] = 0x00

	verPublicKeyHash := append(version, publicKeyHash...)

	return verPublicKeyHash, err
}

func checksumKeyHash(verPublicKeyHash []byte) []byte {
	first := sha256.Sum256(verPublicKeyHash)
	second := sha256.Sum256(first[:])
	checkSum := second[:4]
	return checkSum
}

func encodeKeyHash(verPublicKeyHash []byte, checkSum []byte) string {
	addressHex := append(verPublicKeyHash, checkSum...)
	bitcoinAddressHex := base58.Encode(addressHex)

	return bitcoinAddressHex
}
