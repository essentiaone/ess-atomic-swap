package infrastructure

import (
	"context"
	stcrypto "crypto"
	"crypto/ecdsa"
	"math/big"

	"github.com/essentiaone/ess-atomic-swap/sc"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

// HandlerSCDB smart contract handlers
type HandlerSCDB struct {
	Clinet     *ethclient.Client // connect to node
	PrivateKey *ecdsa.PrivateKey
	PublicKey  stcrypto.PublicKey
	Address    common.Address // user address
	Store      *sc.Store      // sc structure
}

// New create new handler for smart contract
func New(nodeAddr, SCDBaddr, userPrivateKey string) (*HandlerSCDB, error) {
	client, err := ethclient.Dial(nodeAddr)
	if err != nil {
		return nil, errors.Wrap(err, "cannot connect to Ethereum node")
	}

	store, err := sc.NewStore(common.HexToAddress(SCDBaddr), client)
	if err != nil {
		return nil, errors.Wrap(err, "cannot connect to SCDB")
	}

	privateKey, err := crypto.HexToECDSA(userPrivateKey)
	if err != nil {
		return nil, errors.Wrap(err, "cannot convert user private key")
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("convert public key to ECDSA")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	return &HandlerSCDB{
		Clinet:     client,
		PrivateKey: privateKey,
		PublicKey:  publicKeyECDSA,
		Address:    address,
		Store:      store,
	}, nil
}

// GenerateTransaction generate data for transaction
func (scdb *HandlerSCDB) GenerateTransaction() (*bind.TransactOpts, error) {
	nonce, err := scdb.Clinet.PendingNonceAt(context.Background(), scdb.Address)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get account nonce")
	}

	gasPrice, err := scdb.Clinet.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "cannot get gas price")
	}

	tx := bind.NewKeyedTransactor(scdb.PrivateKey)
	tx.Nonce = big.NewInt(int64(nonce))
	tx.Value = big.NewInt(0)
	tx.GasLimit = uint64(300000)
	tx.GasPrice = gasPrice // in wei

	return tx, nil
}

// SaveOrder save user to temporary table
func (scdb *HandlerSCDB) SaveOrder(fromAddress, toAddress, password string, amount int64) (string, error) {
	rawTx, err := scdb.GenerateTransaction()
	if err != nil {
		return "", err
	}

	fromAddressHex := common.HexToAddress(fromAddress)
	toAddressHex := common.HexToAddress(toAddress)
	bigAmount := big.NewInt(amount)

	tx, err := scdb.Store.AddInit(rawTx, fromAddressHex, toAddressHex, bigAmount, password)
	if err != nil {
		return "", errors.Wrap(err, "cannot save initial data")
	}

	return tx.Hash().Hex(), nil
}

// InitiateGetByAddress return user by public address from temporary table
func (scdb *HandlerSCDB) InitiateGetByAddress(address string) (map[string]interface{}, error) {
	fromAddressHex, toAddressHex, amount, block, hashHex, err := scdb.Store.GetInit(nil, common.HexToAddress(address))
	if err != nil {
		return nil, errors.Wrap(err, "cannot get initial data")
	}

	return map[string]interface{}{
		"from":           hexutil.Encode(fromAddressHex[:]),
		"to":             hexutil.Encode(toAddressHex[:]),
		"amount":         amount,
		"blockTimestamp": block,
		"hash":           hexutil.Encode(hashHex[:]),
	}, nil
}

// InitiateConfirm move user to constant table
func (scdb *HandlerSCDB) InitiateConfirm(address, password, txHash string, blockTimestamp big.Int) (string, error) {
	var rawTxHash [32]byte
	copy(rawTxHash[:], txHash)

	rawTx, err := scdb.GenerateTransaction()
	if err != nil {
		return "", err
	}

	tx, err := scdb.Store.ConfirmInit(
		rawTx,
		common.HexToAddress(address),
		password,
		rawTxHash,
		&blockTimestamp,
	)
	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil
}
