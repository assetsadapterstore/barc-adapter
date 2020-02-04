/*
 * Copyright 2018 The OpenWallet Authors
 * This file is part of the OpenWallet library.
 *
 * The OpenWallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The OpenWallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package barc

import (
	"github.com/blocktree/bitshares-adapter/bitshares"
	"github.com/blocktree/openwallet/log"
	bts "github.com/denkhaus/bitshares"
	"github.com/denkhaus/bitshares/config"
)

const (
	ChainIDBAR = "8ac5756ab6f78c1ae92d9654072eecf7fc81b418a05073ad067095061067e3c9"
)

type WalletManager struct {
	*bitshares.WalletManager
}

func NewWalletManager() *WalletManager {
	wm := WalletManager{}
	wm.WalletManager = bitshares.NewWalletManager(nil)
	wm.Config = bitshares.NewConfig(Symbol)
	wm.Decoder = NewAddressDecoder(&wm)
	wm.DecoderV2 = NewAddressDecoder(&wm)
	wm.Log = log.NewOWLogger(wm.Symbol())
	wm.Api = bitshares.NewWalletClient(wm.Config.ServerAPI, wm.Config.WalletAPI, false)
	wm.WebsocketAPI = NewWebsocketAPI(wm.Config.ServerWS)
	return &wm
}

func NewWebsocketAPI(api string) bts.WebsocketAPI {
	config.Add(config.ChainConfig{
		Name:      "BAR",
		CoreAsset: "BAR",
		Prefix:    "BAR",
		ID:        ChainIDBAR,
	})
	config.SetCurrent(ChainIDBAR)
	return bts.NewWebsocketAPI(api)
}
