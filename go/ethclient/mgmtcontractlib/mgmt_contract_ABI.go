package mgmtcontractlib

const (
	AddRollupMethod     = "AddRollup"
	StoreSecretMethod   = "StoreSecret"
	RequestSecretMethod = "RequestSecret"

	MgmtContractByteCode = `608060405234801561001057600080fd5b50610c58806100206000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80637887aa4f1161005b5780637887aa4f146100d85780638353ffca146100f45780638ef74f8914610110578063e0643dfc146101405761007d565b806347f03a84146100825780635c26b626146100a057806370342156146100bc575b600080fd5b61008a610170565b604051610097919061095d565b60405180910390f35b6100ba60048036038101906100b59190610719565b610259565b005b6100d660048036038101906100d19190610719565b6102ab565b005b6100f260048036038101906100ed9190610766565b6102fe565b005b61010e600480360381019061010991906107e2565b61031a565b005b61012a600480360381019061012591906106ec565b610365565b604051610137919061097f565b60405180910390f35b61015a60048036038101906101559190610822565b610405565b604051610167919061097f565b60405180910390f35b6060600080438152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b828210156102505783829060005260206000200180546101c390610aed565b80601f01602080910402602001604051908101604052809291908181526020018280546101ef90610aed565b801561023c5780601f106102115761010080835404028352916020019161023c565b820191906000526020600020905b81548152906001019060200180831161021f57829003601f168201915b5050505050815260200190600101906101a4565b50505050905090565b6000804381526020019081526020016000208282909180600181540180825580915050600190039060005260206000200160009091929091929091929091925091906102a69291906104be565b505050565b8181600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002091906102f99291906104be565b505050565b8260029080519060200190610314929190610544565b50505050565b8073ffffffffffffffffffffffffffffffffffffffff166108fc839081150290604051600060405180830381858888f19350505050158015610360573d6000803e3d6000fd5b505050565b6001602052806000526040600020600091509050805461038490610aed565b80601f01602080910402602001604051908101604052809291908181526020018280546103b090610aed565b80156103fd5780601f106103d2576101008083540402835291602001916103fd565b820191906000526020600020905b8154815290600101906020018083116103e057829003601f168201915b505050505081565b6000602052816000526040600020818154811061042157600080fd5b9060005260206000200160009150915050805461043d90610aed565b80601f016020809104026020016040519081016040528092919081815260200182805461046990610aed565b80156104b65780601f1061048b576101008083540402835291602001916104b6565b820191906000526020600020905b81548152906001019060200180831161049957829003601f168201915b505050505081565b8280546104ca90610aed565b90600052602060002090601f0160209004810192826104ec5760008555610533565b82601f1061050557803560ff1916838001178555610533565b82800160010185558215610533579182015b82811115610532578235825591602001919060010190610517565b5b50905061054091906105ca565b5090565b82805461055090610aed565b90600052602060002090601f01602090048101928261057257600085556105b9565b82601f1061058b57805160ff19168380011785556105b9565b828001600101855582156105b9579182015b828111156105b857825182559160200191906001019061059d565b5b5090506105c691906105ca565b5090565b5b808211156105e35760008160009055506001016105cb565b5090565b60006105fa6105f5846109c6565b6109a1565b90508281526020810184848401111561061657610615610bbd565b5b610621848285610aab565b509392505050565b60008135905061063881610bdd565b92915050565b60008135905061064d81610bf4565b92915050565b60008083601f84011261066957610668610bb3565b5b8235905067ffffffffffffffff81111561068657610685610bae565b5b6020830191508360018202830111156106a2576106a1610bb8565b5b9250929050565b600082601f8301126106be576106bd610bb3565b5b81356106ce8482602086016105e7565b91505092915050565b6000813590506106e681610c0b565b92915050565b60006020828403121561070257610701610bc7565b5b600061071084828501610629565b91505092915050565b600080602083850312156107305761072f610bc7565b5b600083013567ffffffffffffffff81111561074e5761074d610bc2565b5b61075a85828601610653565b92509250509250929050565b60008060006040848603121561077f5761077e610bc7565b5b600084013567ffffffffffffffff81111561079d5761079c610bc2565b5b6107a9868287016106a9565b935050602084013567ffffffffffffffff8111156107ca576107c9610bc2565b5b6107d686828701610653565b92509250509250925092565b600080604083850312156107f9576107f8610bc7565b5b6000610807858286016106d7565b92505060206108188582860161063e565b9150509250929050565b6000806040838503121561083957610838610bc7565b5b6000610847858286016106d7565b9250506020610858858286016106d7565b9150509250929050565b600061086e83836108eb565b905092915050565b600061088182610a07565b61088b8185610a2a565b93508360208202850161089d856109f7565b8060005b858110156108d957848403895281516108ba8582610862565b94506108c583610a1d565b925060208a019950506001810190506108a1565b50829750879550505050505092915050565b60006108f682610a12565b6109008185610a3b565b9350610910818560208601610aba565b61091981610bcc565b840191505092915050565b600061092f82610a12565b6109398185610a4c565b9350610949818560208601610aba565b61095281610bcc565b840191505092915050565b600060208201905081810360008301526109778184610876565b905092915050565b600060208201905081810360008301526109998184610924565b905092915050565b60006109ab6109bc565b90506109b78282610b1f565b919050565b6000604051905090565b600067ffffffffffffffff8211156109e1576109e0610b7f565b5b6109ea82610bcc565b9050602081019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b6000610a6882610a81565b9050919050565b6000610a7a82610a81565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b83811015610ad8578082015181840152602081019050610abd565b83811115610ae7576000848401525b50505050565b60006002820490506001821680610b0557607f821691505b60208210811415610b1957610b18610b50565b5b50919050565b610b2882610bcc565b810181811067ffffffffffffffff82111715610b4757610b46610b7f565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b610be681610a5d565b8114610bf157600080fd5b50565b610bfd81610a6f565b8114610c0857600080fd5b50565b610c1481610aa1565b8114610c1f57600080fd5b5056fea2646970667358221220f6fb6eff95ba2f4befdd5b2c555eb3ade256d9f6a77215380df4c9670ef8d2e264736f6c63430008070033`
	MgmtContractABI      = ` [
	{
		"inputs": [
			{
				"internalType": "string",
				"name": "rollupData",
				"type": "string"
			}
		],
		"name": "AddRollup",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "string",
				"name": "requestReport",
				"type": "string"
			}
		],
		"name": "RequestSecret",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "Rollup",
		"outputs": [
			{
				"internalType": "string[]",
				"name": "",
				"type": "string[]"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "string",
				"name": "inputSecret",
				"type": "string"
			},
			{
				"internalType": "string",
				"name": "requestReport",
				"type": "string"
			}
		],
		"name": "StoreSecret",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "withdrawAmount",
				"type": "uint256"
			},
			{
				"internalType": "address payable",
				"name": "destination",
				"type": "address"
			}
		],
		"name": "Withdraw",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"name": "attestationRequests",
		"outputs": [
			{
				"internalType": "string",
				"name": "",
				"type": "string"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"name": "rollups",
		"outputs": [
			{
				"internalType": "string",
				"name": "",
				"type": "string"
			}
		],
		"stateMutability": "view",
		"type": "function"
	}
]`
)