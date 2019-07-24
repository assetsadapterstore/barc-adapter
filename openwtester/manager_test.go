package openwtester

import (
	"path/filepath"
	"testing"

	"github.com/blocktree/openwallet/log"
	"github.com/blocktree/openwallet/openw"
	"github.com/blocktree/openwallet/openwallet"
)

var (
	testApp        = "assets-adapter"
	configFilePath = filepath.Join("conf")
)

func testInitWalletManager() *openw.WalletManager {
	log.SetLogFuncCall(true)
	tc := openw.NewConfig()

	tc.ConfigDir = configFilePath
	tc.EnableBlockScan = false
	tc.SupportAssets = []string{
		"BAR",
	}

	return openw.NewWalletManager(tc)
}

func TestWalletManager_CreateWallet(t *testing.T) {
	tm := testInitWalletManager()
	w := &openwallet.Wallet{Alias: "HELLO BAR", IsTrust: true, Password: "12345678"}
	nw, key, err := tm.CreateWallet(testApp, w)
	if err != nil {
		log.Error(err)
		return
	}

	log.Info("wallet:", nw)
	log.Info("key:", key)

}

func TestWalletManager_GetWalletInfo(t *testing.T) {

	tm := testInitWalletManager()

	wallet, err := tm.GetWalletInfo(testApp, "WF3Bq2FJ444BH1ZHfjuDRC72bhqH8NRNLb")
	if err != nil {
		log.Error("unexpected error:", err)
		return
	}
	log.Info("wallet:", wallet)
}

func TestWalletManager_GetWalletList(t *testing.T) {

	tm := testInitWalletManager()

	list, err := tm.GetWalletList(testApp, 0, 10000000)
	if err != nil {
		log.Error("unexpected error:", err)
		return
	}
	for i, w := range list {
		log.Info("wallet[", i, "] :", w)
	}
	log.Info("wallet count:", len(list))

	tm.CloseDB(testApp)
}

func TestWalletManager_CreateAssetsAccount(t *testing.T) {

	tm := testInitWalletManager()

	walletID := "WKMowUwix8Eo6Y1rLcnEcA8oszVRL2C1tj"
	account := &openwallet.AssetsAccount{Alias: "zbcat111", WalletID: walletID, Required: 1, Symbol: "BAR", IsTrust: true}
	account, address, err := tm.CreateAssetsAccount(testApp, walletID, "12345678", account, nil)
	if err != nil {
		log.Error(err)
		return
	}

	log.Info("account:", account)
	log.Info("address:", address)

	tm.CloseDB(testApp)
}

func TestWalletManager_GetAssetsAccountList(t *testing.T) {

	tm := testInitWalletManager()

	walletID := "WHx7NV9em582AXU5tu8r5xdpcuBjrtMTUZ"
	list, err := tm.GetAssetsAccountList(testApp, walletID, 0, 10000000)
	if err != nil {
		log.Error("unexpected error:", err)
		return
	}
	for i, w := range list {
		log.Info("account[", i, "] :", w)
	}
	log.Info("account count:", len(list))

	tm.CloseDB(testApp)

}

func TestWalletManager_CreateAddress(t *testing.T) {

	tm := testInitWalletManager()

	walletID := "WHx7NV9em582AXU5tu8r5xdpcuBjrtMTUZ"
	accountID := "WHx7NV9em582AXU5tu8r5xdpcuBjrtMTUZ"
	address, err := tm.CreateAddress(testApp, walletID, accountID, 1)
	if err != nil {
		log.Error(err)
		return
	}

	log.Info("address:", address)

	tm.CloseDB(testApp)
}

func TestWalletManager_GetAddressList(t *testing.T) {

	tm := testInitWalletManager()

	// walletID := "W4oCSSCw9YrCSy4M959ChKM7BRphgZVeMw"
	// accountID := "DZazoeronbtQRkm6gEyqDubvYwCNGXUWP2fzBx3bSBrJ" //zbalice111 BAR6SsHffQqnnBMM3wg8LKQdRc7rMJwkvye2JB8LTihCJXAJGLs6z

	// walletID := "WKMowUwix8Eo6Y1rLcnEcA8oszVRL2C1tj"
	// accountID := "AszEboNMoJRbrLZtuSdci312StDrtrLuxwbrc38o6moH" //zbbob111 BAR7VZ6vU5T4Yfcd1oQnwk8nVWTGg1cT8CRB7cswhmKNXUCTc4KtU

	walletID := "WKMowUwix8Eo6Y1rLcnEcA8oszVRL2C1tj"
	accountID := "Fo8dJUXcHA1n3Q6Cegr5ou2pppe9fpsd7kqSh4GsyXDc" //zbcat111 BAR5s3e7wthf4BQ61JjDCegZvt9LDEtmUJJKsFfgRnrrgjJG66BLr

	list, err := tm.GetAddressList(testApp, walletID, accountID, 0, -1, false)
	if err != nil {
		log.Error("unexpected error:", err)
		return
	}
	for i, w := range list {
		log.Info("address[", i, "] :", w.Address)
	}
	log.Info("address count:", len(list))

	tm.CloseDB(testApp)
}
