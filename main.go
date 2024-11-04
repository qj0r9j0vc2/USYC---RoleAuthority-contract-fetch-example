package main

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"smart_contract/whitelist"
)

type Namespace struct {
	Denom           string
	WasmHook        string
	MintsPaused     bool
	SendsPaused     bool
	BurnsPaused     bool
	RolePermissions map[string]int
	AddressRoles    map[string][]string
}

var functionSigsMap = map[string]string{
	"a9059cbb": "transfer(address,uint256)",
	"40c10f19": "mint(address,uint256)",
	"7d40583d": "setRoleCapability(uint8,address,bytes4,bool)",
	"7641e6f3": "burn(uint256,string)",
	"23b872dd": "transferFrom(address,address,uint256)",
	"008fbe5e": "buyWithPermit(address,uint256,address,uint256,uint8,bytes32,bytes32)",
	"0323a8b0": "broadcast(bytes)",
	"742b5b74": "sellFor(uint256,address)",
	"b6b55f25": "deposit(uint256)",
	"d96a094a": "buy(uint256)",
	"e4849b32": "sell(uint256)",
	"a3a8573a": "buyFor(uint256,address)",
	"67aff484": "setUserRole(address,uint8,bool)",
	"c46c67fb": "batchExecute((address,(uint8,bytes)[])[])",
	"8e1b5c0e": "execute(address,(uint8,bytes)[])",
}

var contractAddrMap = map[string]string{
	"0x136471a34f6ef19fE571EFFC1CA711fdb8E49f2b": "ERC20 Proxy",
	"0x0dc09046f22EC756e633EcA91618E3c9a372699A": "ERC20",
	"0x470f3b37B9B20E13b0A2a5965Df6bD3f9640DFB4": "RoleAuthority Proxy",
	"0x3aF1730432f77437b3acb12C7FC35A7fF5fC4365": "RoleAuthority",
	"0x0a5EA26fdD38CF2Acb06Dc64198374C337879DAb": "PYUSD-YiledTokenTeller Proxy",
	"0x5C73E1cfdD85b7f1d608F7F7736fC8C653513B7A": "USDC-YiledTokenTeller Proxy",
	"0xf9e3430735de4b68ee41f652cc1394394ee6501b": "PYUSD-YieldTokenTeller",
	"0x62E2f30077907f2daa6d969F544908FAF1fa2bc6": "CrossChainTokenProxy - (hnSOL)",
	"0xe6ca5E409A68E81EACF62789198dFc2c4eF999F9": "CrossChainTokenProxy - (hnBTC_0x9b058e61)",
	"0x2BE3531ad3B5C8c3E74c128258Cd1e891d84CCda": "CrossChainTokenProxy - (hnBTC_0xd3147433)",
	"0xBAD11967371814054D505642a733ab6A83B7C2dB": "CrossChainTokenProxy - (hnBTC)",
	"0x57aFa1235a841c66641FE8be6a0CF6f1980E54cD": "CrossChainTokenProxy - (hnAVAX)",
	"0x88e6998CEc706E80CA3F38D635E9ee503F50D83d": "CrossChainTokenProxy - (hnETH)",
	"0xe1411AdCfF539eFdE126db8866e6306F50516a50": "CrossChainTokenProxy - (hnWSTETH)",
	"0x97D829C13a3cf9FD0615Ce687Ce9A2939234D218": "CrossChainTokenProxy - (hnMATIC)",
	"0x35d95e93E8aFCAaa3046B9E5926E92B46AAd9b11": "AxelarMessengerProxy",
	"0x2a6F3F4cD09106aD645A5d8b16FbE8D6Bc57E18F": "Opportunistic Stablecoin Yield Coin (OSYC)",
	"0x341C281d11677795F1192bF1f7438666909841c8": "CrossMarginPhysicalEngineProxy",
	"0x9c742Aef14CC875C49f52bBD4473B35beBAD26Ae": "CorssMarginEngineProxy",
}

func fetchHistory(contractAddress string) (*Namespace, error) {
	fmt.Printf("Fetching history for contract at address: %s...\n", contractAddress)

	client, err := ethclient.Dial(ETH_URL)
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress(contractAddress)

	contract, err := whitelist.NewWhitelist(
		address,
		client)

	target, _ := hex.DecodeString("0a5EA26fdD38CF2Acb06Dc64198374C337879DAb") // PYUSD-YiledTokenTeller Proxy
	iter, err := contract.FilterRoleCapabilityUpdated(
		&bind.FilterOpts{
			Start: 0,
			End:   nil, // set end to latest height
		}, nil, []common.Address{common.Address(target)}, nil)

	if err != nil {
		log.Fatal(err)
	}

	defer iter.Close()

	var lowestHeight uint64 = 9999999999
	//var whiteListUserRoleUpdated []*whitelist.WhitelistUserRoleUpdated
	for iter.Next() {
		contractName := contractAddrMap[iter.Event.Target.Hex()]
		if contractName == "" {
			contractName = iter.Event.Target.Hex()
		}

		fmt.Printf("%4d / %x / %70s / %50s (%v) / %t \n",
			iter.Event.Role,
			iter.Event.FunctionSig,
			functionSigsMap[fmt.Sprintf("%x", iter.Event.FunctionSig)],
			contractName,
			fmt.Sprintf("%v...%v", iter.Event.Target.Hex()[:5], iter.Event.Target.Hex()[35:]),
			iter.Event.Enabled)
		if iter.Event.Raw.BlockNumber < lowestHeight {
			lowestHeight = iter.Event.Raw.BlockNumber
		}
		//whiteListUserRoleUpdated = append(whiteListUserRoleUpdated, iter.Event)
	}

	fmt.Printf("most lowest height is %d\n", lowestHeight)
	return nil, nil
}

const ETH_URL = "https://mainnet.infura.io/v3/xxxxxxxxx"

func main() {
	// Example usage
	fetchHistory("0x470f3b37B9B20E13b0A2a5965Df6bD3f9640DFB4")
}
