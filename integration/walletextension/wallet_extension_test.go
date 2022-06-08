//nolint:unused // TODO - Remove once tests are unskipped.
package walletextension

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/obscuronet/obscuro-playground/tools/walletextension"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/obscuronet/obscuro-playground/integration"
	"github.com/obscuronet/obscuro-playground/integration/ethereummock"
	"github.com/obscuronet/obscuro-playground/integration/simulation/network"
	"github.com/obscuronet/obscuro-playground/integration/simulation/params"
	"github.com/obscuronet/obscuro-playground/integration/simulation/stats"
)

const (
	chainID    = 1337
	chainIDHex = "0x539"                // Chain ID in hex.
	allocHex   = "0x3635c9adc5dea00000" // Default account allocation in hex.
	// Generated using Hardhat from https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/token/ERC20/ERC20.sol.
	erc20Bytecode    = "0x60806040523480156200001157600080fd5b5060405162001620380380620016208339818101604052810190620000379190620002be565b81600390805190602001906200004f92919062000071565b5080600490805190602001906200006892919062000071565b505050620003a8565b8280546200007f9062000372565b90600052602060002090601f016020900481019282620000a35760008555620000ef565b82601f10620000be57805160ff1916838001178555620000ef565b82800160010185558215620000ef579182015b82811115620000ee578251825591602001919060010190620000d1565b5b509050620000fe919062000102565b5090565b5b808211156200011d57600081600090555060010162000103565b5090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6200018a826200013f565b810181811067ffffffffffffffff82111715620001ac57620001ab62000150565b5b80604052505050565b6000620001c162000121565b9050620001cf82826200017f565b919050565b600067ffffffffffffffff821115620001f257620001f162000150565b5b620001fd826200013f565b9050602081019050919050565b60005b838110156200022a5780820151818401526020810190506200020d565b838111156200023a576000848401525b50505050565b6000620002576200025184620001d4565b620001b5565b9050828152602081018484840111156200027657620002756200013a565b5b620002838482856200020a565b509392505050565b600082601f830112620002a357620002a262000135565b5b8151620002b584826020860162000240565b91505092915050565b60008060408385031215620002d857620002d76200012b565b5b600083015167ffffffffffffffff811115620002f957620002f862000130565b5b62000307858286016200028b565b925050602083015167ffffffffffffffff8111156200032b576200032a62000130565b5b62000339858286016200028b565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200038b57607f821691505b60208210811415620003a257620003a162000343565b5b50919050565b61126880620003b86000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80633950935111610071578063395093511461016857806370a082311461019857806395d89b41146101c8578063a457c2d7146101e6578063a9059cbb14610216578063dd62ed3e14610246576100a9565b806306fdde03146100ae578063095ea7b3146100cc57806318160ddd146100fc57806323b872dd1461011a578063313ce5671461014a575b600080fd5b6100b6610276565b6040516100c39190610b22565b60405180910390f35b6100e660048036038101906100e19190610bdd565b610308565b6040516100f39190610c38565b60405180910390f35b61010461032b565b6040516101119190610c62565b60405180910390f35b610134600480360381019061012f9190610c7d565b610335565b6040516101419190610c38565b60405180910390f35b610152610364565b60405161015f9190610cec565b60405180910390f35b610182600480360381019061017d9190610bdd565b61036d565b60405161018f9190610c38565b60405180910390f35b6101b260048036038101906101ad9190610d07565b6103a4565b6040516101bf9190610c62565b60405180910390f35b6101d06103ec565b6040516101dd9190610b22565b60405180910390f35b61020060048036038101906101fb9190610bdd565b61047e565b60405161020d9190610c38565b60405180910390f35b610230600480360381019061022b9190610bdd565b6104f5565b60405161023d9190610c38565b60405180910390f35b610260600480360381019061025b9190610d34565b610518565b60405161026d9190610c62565b60405180910390f35b60606003805461028590610da3565b80601f01602080910402602001604051908101604052809291908181526020018280546102b190610da3565b80156102fe5780601f106102d3576101008083540402835291602001916102fe565b820191906000526020600020905b8154815290600101906020018083116102e157829003601f168201915b5050505050905090565b60008061031361059f565b90506103208185856105a7565b600191505092915050565b6000600254905090565b60008061034061059f565b905061034d858285610772565b6103588585856107fe565b60019150509392505050565b60006012905090565b60008061037861059f565b905061039981858561038a8589610518565b6103949190610e04565b6105a7565b600191505092915050565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6060600480546103fb90610da3565b80601f016020809104026020016040519081016040528092919081815260200182805461042790610da3565b80156104745780601f1061044957610100808354040283529160200191610474565b820191906000526020600020905b81548152906001019060200180831161045757829003601f168201915b5050505050905090565b60008061048961059f565b905060006104978286610518565b9050838110156104dc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104d390610ecc565b60405180910390fd5b6104e982868684036105a7565b60019250505092915050565b60008061050061059f565b905061050d8185856107fe565b600191505092915050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610617576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161060e90610f5e565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610687576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161067e90610ff0565b60405180910390fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925836040516107659190610c62565b60405180910390a3505050565b600061077e8484610518565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146107f857818110156107ea576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107e19061105c565b60405180910390fd5b6107f784848484036105a7565b5b50505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141561086e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610865906110ee565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156108de576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108d590611180565b60405180910390fd5b6108e9838383610a7f565b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490508181101561096f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161096690611212565b60405180910390fd5b8181036000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610a029190610e04565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610a669190610c62565b60405180910390a3610a79848484610a84565b50505050565b505050565b505050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610ac3578082015181840152602081019050610aa8565b83811115610ad2576000848401525b50505050565b6000601f19601f8301169050919050565b6000610af482610a89565b610afe8185610a94565b9350610b0e818560208601610aa5565b610b1781610ad8565b840191505092915050565b60006020820190508181036000830152610b3c8184610ae9565b905092915050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610b7482610b49565b9050919050565b610b8481610b69565b8114610b8f57600080fd5b50565b600081359050610ba181610b7b565b92915050565b6000819050919050565b610bba81610ba7565b8114610bc557600080fd5b50565b600081359050610bd781610bb1565b92915050565b60008060408385031215610bf457610bf3610b44565b5b6000610c0285828601610b92565b9250506020610c1385828601610bc8565b9150509250929050565b60008115159050919050565b610c3281610c1d565b82525050565b6000602082019050610c4d6000830184610c29565b92915050565b610c5c81610ba7565b82525050565b6000602082019050610c776000830184610c53565b92915050565b600080600060608486031215610c9657610c95610b44565b5b6000610ca486828701610b92565b9350506020610cb586828701610b92565b9250506040610cc686828701610bc8565b9150509250925092565b600060ff82169050919050565b610ce681610cd0565b82525050565b6000602082019050610d016000830184610cdd565b92915050565b600060208284031215610d1d57610d1c610b44565b5b6000610d2b84828501610b92565b91505092915050565b60008060408385031215610d4b57610d4a610b44565b5b6000610d5985828601610b92565b9250506020610d6a85828601610b92565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680610dbb57607f821691505b60208210811415610dcf57610dce610d74565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610e0f82610ba7565b9150610e1a83610ba7565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115610e4f57610e4e610dd5565b5b828201905092915050565b7f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760008201527f207a65726f000000000000000000000000000000000000000000000000000000602082015250565b6000610eb6602583610a94565b9150610ec182610e5a565b604082019050919050565b60006020820190508181036000830152610ee581610ea9565b9050919050565b7f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460008201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b6000610f48602483610a94565b9150610f5382610eec565b604082019050919050565b60006020820190508181036000830152610f7781610f3b565b9050919050565b7f45524332303a20617070726f766520746f20746865207a65726f20616464726560008201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b6000610fda602283610a94565b9150610fe582610f7e565b604082019050919050565b6000602082019050818103600083015261100981610fcd565b9050919050565b7f45524332303a20696e73756666696369656e7420616c6c6f77616e6365000000600082015250565b6000611046601d83610a94565b915061105182611010565b602082019050919050565b6000602082019050818103600083015261107581611039565b9050919050565b7f45524332303a207472616e736665722066726f6d20746865207a65726f20616460008201527f6472657373000000000000000000000000000000000000000000000000000000602082015250565b60006110d8602583610a94565b91506110e38261107c565b604082019050919050565b60006020820190508181036000830152611107816110cb565b9050919050565b7f45524332303a207472616e7366657220746f20746865207a65726f206164647260008201527f6573730000000000000000000000000000000000000000000000000000000000602082015250565b600061116a602383610a94565b91506111758261110e565b604082019050919050565b600060208201905081810360008301526111998161115d565b9050919050565b7f45524332303a207472616e7366657220616d6f756e742065786365656473206260008201527f616c616e63650000000000000000000000000000000000000000000000000000602082015250565b60006111fc602683610a94565b9150611207826111a0565b604082019050919050565b6000602082019050818103600083015261122b816111ef565b905091905056fea26469706673582212201954fc1d02a14b0a4f73a70fb6faaec7035b39b47ae9b9ae074064af1c889c1364736f6c63430008090033"
	dummyAccountAddr = "0x8D97689C9818892B700e27F316cc3E41e17fBeb9"

	startPort              = 3000
	httpProtocol           = "http://"
	pathGenerateViewingKey = "/generateviewingkey/"
	pathSubmitViewingKey   = "/submitviewingkey/"

	reqJSONMethodChainID    = "eth_chainId"
	reqJSONMethodGetBalance = "eth_getBalance"
	reqJSONMethodCall       = "eth_call"
	respJSONKeyResult       = "result"
	reqJSONKeyTo            = "to"
	reqJSONKeyFrom          = "from"
	emptyResult             = "0x"
	errInsecure             = "enclave could not respond securely to %s request because there is no viewing key for the account"
	signedMsgPrefix         = "vk"
)

func TestCanMakeNonSensitiveRequestWithoutSubmittingViewingKey(t *testing.T) {
	t.Skip() // Skipping while support for viewing keys is being implemented.

	stopHandle, err := createObscuroNetwork(int(integration.StartPortWalletExtensionTest + 1))
	defer stopHandle()
	if err != nil {
		t.Fatalf("failed to create test Obscuro network. Cause: %s", err)
	}

	nodeRPCPort := integration.StartPortWalletExtensionTest + 1 + network.DefaultHostRPCOffset
	walletExtensionConfig := walletextension.Config{
		WalletExtensionPort: startPort,
		NodeRPCAddress:      fmt.Sprintf("%s:%d", network.Localhost, nodeRPCPort),
	}
	walletExtensionAddr := fmt.Sprintf("%s:%d", network.Localhost, integration.StartPortWalletExtensionTest)

	walletExtension := walletextension.NewWalletExtension(walletExtensionConfig)
	defer walletExtension.Shutdown()
	go walletExtension.Serve(walletExtensionAddr)
	waitForWalletExtension(t, walletExtensionAddr)

	respJSON := makeEthJSONReqAsJSON(t, walletExtensionAddr, reqJSONMethodChainID, []string{})

	if respJSON[respJSONKeyResult] != chainIDHex {
		t.Fatalf("Expected chainId of %s, got %s", "1337", respJSON[respJSONKeyResult])
	}
}

func TestCannotGetBalanceWithoutSubmittingViewingKey(t *testing.T) {
	t.Skip() // Skipping while support for viewing keys is being implemented.

	stopHandle, err := createObscuroNetwork(int(integration.StartPortWalletExtensionTest + 1))
	defer stopHandle()
	if err != nil {
		t.Fatalf("failed to create test Obscuro network. Cause: %s", err)
	}

	nodeRPCPort := integration.StartPortWalletExtensionTest + 1 + network.DefaultHostRPCOffset
	walletExtensionConfig := walletextension.Config{
		WalletExtensionPort: startPort,
		NodeRPCAddress:      fmt.Sprintf("%s:%d", network.Localhost, nodeRPCPort),
	}
	walletExtensionAddr := fmt.Sprintf("%s:%d", network.Localhost, integration.StartPortWalletExtensionTest)

	walletExtension := walletextension.NewWalletExtension(walletExtensionConfig)
	defer walletExtension.Shutdown()
	go walletExtension.Serve(walletExtensionAddr)
	waitForWalletExtension(t, walletExtensionAddr)

	respBody := makeEthJSONReq(t, walletExtensionAddr, reqJSONMethodGetBalance, []string{dummyAccountAddr, "latest"})

	trimmedRespBody := strings.TrimSpace(string(respBody))
	expectedErr := fmt.Sprintf(errInsecure, reqJSONMethodGetBalance)
	if trimmedRespBody != expectedErr {
		t.Fatalf("Expected error message \"%s\", got \"%s\"", expectedErr, trimmedRespBody)
	}
}

func TestCanGetOwnBalanceAfterSubmittingViewingKey(t *testing.T) {
	t.Skip() // Skipping while support for viewing keys is being implemented.

	stopHandle, err := createObscuroNetwork(int(integration.StartPortWalletExtensionTest + 1))
	defer stopHandle()
	if err != nil {
		t.Fatalf("failed to create test Obscuro network. Cause: %s", err)
	}

	nodeRPCPort := integration.StartPortWalletExtensionTest + 1 + network.DefaultHostRPCOffset
	walletExtensionConfig := walletextension.Config{
		WalletExtensionPort: startPort,
		NodeRPCAddress:      fmt.Sprintf("%s:%d", network.Localhost, nodeRPCPort),
	}
	walletExtensionAddr := fmt.Sprintf("%s:%d", network.Localhost, integration.StartPortWalletExtensionTest)

	walletExtension := walletextension.NewWalletExtension(walletExtensionConfig)
	defer walletExtension.Shutdown()
	go walletExtension.Serve(walletExtensionAddr)
	waitForWalletExtension(t, walletExtensionAddr)

	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	accountAddr := crypto.PubkeyToAddress(privateKey.PublicKey).String()

	generateAndSubmitViewingKey(t, walletExtensionAddr, privateKey)

	getBalanceJSON := makeEthJSONReqAsJSON(t, walletExtensionAddr, reqJSONMethodGetBalance, []string{accountAddr, "latest"})

	if getBalanceJSON[respJSONKeyResult] != allocHex {
		t.Fatalf("Expected balance of %s, got %s", allocHex, getBalanceJSON[respJSONKeyResult])
	}
}

func TestCannotGetAnothersBalanceAfterSubmittingViewingKey(t *testing.T) {
	t.Skip() // Skipping while support for viewing keys is being implemented.

	stopHandle, err := createObscuroNetwork(int(integration.StartPortWalletExtensionTest + 1))
	defer stopHandle()
	if err != nil {
		t.Fatalf("failed to create test Obscuro network. Cause: %s", err)
	}

	nodeRPCPort := integration.StartPortWalletExtensionTest + 1 + network.DefaultHostRPCOffset
	walletExtensionConfig := walletextension.Config{
		WalletExtensionPort: startPort,
		NodeRPCAddress:      fmt.Sprintf("%s:%d", network.Localhost, nodeRPCPort),
	}
	walletExtensionAddr := fmt.Sprintf("%s:%d", network.Localhost, integration.StartPortWalletExtensionTest)

	walletExtension := walletextension.NewWalletExtension(walletExtensionConfig)
	defer walletExtension.Shutdown()
	go walletExtension.Serve(walletExtensionAddr)
	waitForWalletExtension(t, walletExtensionAddr)

	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}

	generateAndSubmitViewingKey(t, walletExtensionAddr, privateKey)

	respBody := makeEthJSONReq(t, walletExtensionAddr, reqJSONMethodGetBalance, []string{dummyAccountAddr, "latest"})

	trimmedRespBody := strings.TrimSpace(string(respBody))
	expectedErr := fmt.Sprintf(errInsecure, reqJSONMethodGetBalance)
	if trimmedRespBody != expectedErr {
		t.Fatalf("Expected error message\"%s\", got \"%s\"", expectedErr, trimmedRespBody)
	}
}

func TestCannotCallWithoutSubmittingViewingKey(t *testing.T) {
	t.Skip() // Skipping while support for viewing keys is being implemented.

	stopHandle, err := createObscuroNetwork(int(integration.StartPortWalletExtensionTest + 1))
	defer stopHandle()
	if err != nil {
		t.Fatalf("failed to create test Obscuro network. Cause: %s", err)
	}

	nodeRPCPort := integration.StartPortWalletExtensionTest + 1 + network.DefaultHostRPCOffset
	walletExtensionConfig := walletextension.Config{
		WalletExtensionPort: startPort,
		NodeRPCAddress:      fmt.Sprintf("%s:%d", network.Localhost, nodeRPCPort),
	}
	walletExtensionAddr := fmt.Sprintf("%s:%d", network.Localhost, integration.StartPortWalletExtensionTest)

	walletExtension := walletextension.NewWalletExtension(walletExtensionConfig)
	defer walletExtension.Shutdown()
	go walletExtension.Serve(walletExtensionAddr)
	waitForWalletExtension(t, walletExtensionAddr)

	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	accountAddr := crypto.PubkeyToAddress(privateKey.PublicKey).String()

	contractAddr := deployERC20Contract(t, walletExtensionAddr, privateKey)

	reqParams := map[string]interface{}{
		reqJSONKeyTo:   contractAddr,
		reqJSONKeyFrom: accountAddr,
	}
	respBody := makeEthJSONReq(t, walletExtensionAddr, reqJSONMethodCall, []interface{}{reqParams, "latest"})

	trimmedRespBody := strings.TrimSpace(string(respBody))
	expectedErr := fmt.Sprintf(errInsecure, reqJSONMethodCall)
	if trimmedRespBody != expectedErr {
		t.Fatalf("Expected error message \"%s\", got \"%s\"", expectedErr, trimmedRespBody)
	}
}

func TestCanCallAfterSubmittingViewingKey(t *testing.T) {
	t.Skip() // Skipping while support for viewing keys is being implemented.

	stopHandle, err := createObscuroNetwork(int(integration.StartPortWalletExtensionTest + 1))
	defer stopHandle()
	if err != nil {
		t.Fatalf("failed to create test Obscuro network. Cause: %s", err)
	}

	nodeRPCPort := integration.StartPortWalletExtensionTest + 1 + network.DefaultHostRPCOffset
	walletExtensionConfig := walletextension.Config{
		WalletExtensionPort: startPort,
		NodeRPCAddress:      fmt.Sprintf("%s:%d", network.Localhost, nodeRPCPort),
	}
	walletExtensionAddr := fmt.Sprintf("%s:%d", network.Localhost, integration.StartPortWalletExtensionTest)

	walletExtension := walletextension.NewWalletExtension(walletExtensionConfig)
	defer walletExtension.Shutdown()
	go walletExtension.Serve(walletExtensionAddr)
	waitForWalletExtension(t, walletExtensionAddr)

	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	accountAddr := crypto.PubkeyToAddress(privateKey.PublicKey).String()

	generateAndSubmitViewingKey(t, walletExtensionAddr, privateKey)

	contractAddr := deployERC20Contract(t, walletExtensionAddr, privateKey)

	reqParams := map[string]interface{}{
		reqJSONKeyTo:   contractAddr,
		reqJSONKeyFrom: accountAddr,
	}
	callJSON := makeEthJSONReqAsJSON(t, walletExtensionAddr, reqJSONMethodCall, []interface{}{reqParams, "latest"})

	// TODO - Consider executing an actual transaction, rather than passing an empty one.
	if callJSON[respJSONKeyResult] != emptyResult {
		t.Fatalf("Expected call result of %s, got %s", emptyResult, callJSON[respJSONKeyResult])
	}
}

func TestCannotCallForAnotherAddressAfterSubmittingViewingKey(t *testing.T) {
	t.Skip() // Skipping while support for viewing keys is being implemented.

	stopHandle, err := createObscuroNetwork(int(integration.StartPortWalletExtensionTest + 1))
	defer stopHandle()
	if err != nil {
		t.Fatalf("failed to create test Obscuro network. Cause: %s", err)
	}

	nodeRPCPort := integration.StartPortWalletExtensionTest + 1 + network.DefaultHostRPCOffset
	walletExtensionConfig := walletextension.Config{
		WalletExtensionPort: startPort,
		NodeRPCAddress:      fmt.Sprintf("%s:%d", network.Localhost, nodeRPCPort),
	}
	walletExtensionAddr := fmt.Sprintf("%s:%d", network.Localhost, integration.StartPortWalletExtensionTest)

	walletExtension := walletextension.NewWalletExtension(walletExtensionConfig)
	defer walletExtension.Shutdown()
	go walletExtension.Serve(walletExtensionAddr)
	waitForWalletExtension(t, walletExtensionAddr)

	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}

	generateAndSubmitViewingKey(t, walletExtensionAddr, privateKey)

	contractAddr := deployERC20Contract(t, walletExtensionAddr, privateKey)

	reqParams := map[string]interface{}{
		reqJSONKeyTo:   contractAddr,
		reqJSONKeyFrom: dummyAccountAddr,
	}
	respBody := makeEthJSONReq(t, walletExtensionAddr, reqJSONMethodCall, []interface{}{reqParams, "latest"})

	trimmedRespBody := strings.TrimSpace(string(respBody))
	expectedErr := fmt.Sprintf(errInsecure, reqJSONMethodCall)
	if trimmedRespBody != expectedErr {
		t.Fatalf("Expected error message \"%s\", got \"%s\"", expectedErr, trimmedRespBody)
	}
}

// Waits for wallet extension to be ready. Times out after three seconds.
func waitForWalletExtension(t *testing.T, walletExtensionAddr string) {
	retries := 30
	for i := 0; i < retries; i++ {
		resp, err := http.Get(httpProtocol + walletExtensionAddr) //nolint:noctx
		if resp != nil && resp.Body != nil {
			resp.Body.Close()
		}
		if err == nil {
			return
		}
		time.Sleep(300 * time.Millisecond)
	}
	t.Fatal("could not establish connection to wallet extension")
}

// Makes an Ethereum JSON RPC request and returns the response body.
func makeEthJSONReq(t *testing.T, walletExtensionAddr string, method string, params interface{}) []byte {
	reqBodyBytes, err := json.Marshal(map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  method,
		"params":  params,
		"id":      "1",
	})
	if err != nil {
		t.Fatal(err)
	}
	reqBody := bytes.NewBuffer(reqBodyBytes)

	var resp *http.Response
	// We retry for three seconds to handle node start-up time.
	timeout := time.Now().Add(3 * time.Second)
	for i := time.Now(); i.Before(timeout); i = time.Now() {
		resp, err = http.Post(httpProtocol+walletExtensionAddr, "text/html", reqBody) //nolint:noctx
		if err == nil {
			break
		}
		if resp != nil && resp.Body != nil {
			resp.Body.Close()
		}
	}

	if err != nil {
		t.Fatalf("received error response from wallet extension: %s", err)
	}
	if resp == nil {
		t.Fatal("did not receive a response from the wallet extension")
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	return respBody
}

// Makes an Ethereum JSON RPC request and returns the response body as JSON.
func makeEthJSONReqAsJSON(t *testing.T, walletExtensionAddr string, method string, params interface{}) map[string]interface{} {
	respBody := makeEthJSONReq(t, walletExtensionAddr, method, params)

	if respBody[0] != '{' {
		t.Fatalf("expected JSON response but received: %s", respBody)
	}

	var respBodyJSON map[string]interface{}
	err := json.Unmarshal(respBody, &respBodyJSON)
	if err != nil {
		t.Fatal(err)
	}

	return respBodyJSON
}

// Generates a signed viewing key and submits it to the wallet extension.
func generateAndSubmitViewingKey(t *testing.T, walletExtensionAddr string, accountPrivateKey *ecdsa.PrivateKey) {
	viewingKey := generateViewingKey(t, walletExtensionAddr)
	signature := signViewingKey(t, accountPrivateKey, viewingKey)

	submitViewingKeyBodyBytes, err := json.Marshal(map[string]interface{}{
		"signature": hex.EncodeToString(signature),
	})
	if err != nil {
		t.Fatal(err)
	}
	submitViewingKeyBody := bytes.NewBuffer(submitViewingKeyBodyBytes)
	resp, err := http.Post(httpProtocol+walletExtensionAddr+pathSubmitViewingKey, "application/json", submitViewingKeyBody) //nolint:noctx
	if err != nil {
		t.Fatal(err)
	}
	resp.Body.Close()
}

// Generates a viewing key.
func generateViewingKey(t *testing.T, walletExtensionAddr string) []byte {
	resp, err := http.Get(httpProtocol + walletExtensionAddr + pathGenerateViewingKey) //nolint:noctx
	if err != nil {
		t.Fatal(err)
	}
	viewingKey, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	resp.Body.Close()
	return viewingKey
}

// Signs a viewing key.
func signViewingKey(t *testing.T, privateKey *ecdsa.PrivateKey, viewingKey []byte) []byte {
	msgToSign := signedMsgPrefix + string(viewingKey)
	signature, err := crypto.Sign(accounts.TextHash([]byte(msgToSign)), privateKey)
	if err != nil {
		t.Fatal(err)
	}

	// We have to transform the V from 0/1 to 27/28, and add the leading "0".
	signature[64] += 27
	signatureWithLeadBytes := append([]byte("0"), signature...)

	return signatureWithLeadBytes
}

// Deploys an ERC20 contract and returns the contract's address and the block number it was deployed in.
func deployERC20Contract(t *testing.T, walletExtensionAddr string, signingKey *ecdsa.PrivateKey) string {
	tx := types.LegacyTx{
		Nonce:    0, // relies on a clean env
		GasPrice: big.NewInt(2000000000),
		Gas:      1025_000_000,
		Data:     common.Hex2Bytes(erc20Bytecode),
	}

	signedTx, err := types.SignNewTx(signingKey, types.NewEIP155Signer(big.NewInt(int64(chainID))), &tx)
	if err != nil {
		t.Fatal(err)
	}
	data, err := signedTx.MarshalBinary()
	if err != nil {
		t.Fatal(err)
	}
	encodedData := hexutil.Encode(data)

	respJSON := makeEthJSONReqAsJSON(t, walletExtensionAddr, "eth_sendRawTransaction", []string{encodedData})
	txHash := respJSON[respJSONKeyResult].(string)

	var txInfo map[string]interface{}
	for txInfo == nil {
		txReceipt := makeEthJSONReqAsJSON(t, walletExtensionAddr, "eth_getTransactionReceipt", []string{txHash})
		if txReceipt[respJSONKeyResult] != nil {
			txInfo = txReceipt[respJSONKeyResult].(map[string]interface{})
		}
		time.Sleep(100 * time.Millisecond)
	}

	return txInfo["contractAddress"].(string)
}

// Creates a single-node Obscuro network for testing.
func createObscuroNetwork(startPort int) (func(), error) {
	wallets := params.NewSimWallets(1, 1, 1)

	simParams := params.SimParams{
		NumberOfNodes:    1,
		AvgBlockDuration: 1 * time.Second,
		AvgGossipPeriod:  1 * time.Second / 3,
		MgmtContractLib:  ethereummock.NewMgmtContractLibMock(),
		ERC20ContractLib: ethereummock.NewERC20ContractLibMock(),
		Wallets:          wallets,
		StartPort:        startPort,
	}
	simStats := stats.NewStats(simParams.NumberOfNodes)

	obscuroNetwork := network.NewNetworkOfSocketNodes(wallets)
	_, _, _, err := obscuroNetwork.Create(&simParams, simStats) //nolint:dogsled
	if err != nil {
		return obscuroNetwork.TearDown, err
	}
	return obscuroNetwork.TearDown, nil
}