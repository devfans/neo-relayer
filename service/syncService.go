package service

import (
	neoRpc "github.com/joeqian10/neo-gogogo/rpc"
	"github.com/joeqian10/neo-gogogo/wallet"
	"github.com/polynetwork/neo-relayer/config"
	"github.com/polynetwork/neo-relayer/db"
	"github.com/polynetwork/neo-relayer/log"
	poly_bridge_sdk "github.com/polynetwork/poly-bridge/bridgesdk"
	rsdk "github.com/polynetwork/poly-go-sdk"
	"os"
)

// SyncService ...
type SyncService struct {
	relayAccount    *rsdk.Account
	relaySdk        *rsdk.PolySdk
	relaySyncHeight uint32

	neoAccount       *wallet.Account
	neoSdk           *neoRpc.RpcClient
	neoSdk4Listen    *neoRpc.RpcClient
	neoSyncHeight    uint32
	neoNextConsensus string

	bridgeSdk *poly_bridge_sdk.BridgeSdk

	db     *db.BoltDB
	config *config.Config
}

// NewSyncService ...
func NewSyncService(acct *rsdk.Account, relaySdk *rsdk.PolySdk, neoAccount *wallet.Account, neoSdk *neoRpc.RpcClient, neoSdk4Listen *neoRpc.RpcClient) *SyncService {
	if !checkIfExist(config.DefConfig.DBPath) {
		os.Mkdir(config.DefConfig.DBPath, os.ModePerm)
	}
	boltDB, err := db.NewBoltDB(config.DefConfig.DBPath)
	if err != nil {
		log.Errorf("db.NewWaitingDB error:%s", err)
		os.Exit(1)
	}
	sdk := poly_bridge_sdk.NewBridgeSdk(config.DefConfig.BridgeUrl[0][0])
	syncSvr := &SyncService{
		relayAccount: acct,
		relaySdk:     relaySdk,

		neoAccount:    neoAccount,
		neoSdk:        neoSdk,
		neoSdk4Listen: neoSdk4Listen,
		db:            boltDB,
		config:        config.DefConfig,
		bridgeSdk:     sdk,
	}
	return syncSvr
}

// Run ...
func (this *SyncService) Run() {
	go this.RelayToNeo()
	go this.RelayToNeoRetry()
	go this.NeoToRelay()
	go this.NeoToRelayCheckAndRetry()
}

func checkIfExist(dir string) bool {
	_, err := os.Stat(dir)
	if err != nil && !os.IsExist(err) {
		return false
	}
	return true
}
