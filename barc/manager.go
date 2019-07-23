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
	"github.com/blocktree/openwallet/openwallet"
)

type WalletManager struct {
	*bitshares.WalletManager
}

func NewWalletManager(cacheManager openwallet.ICacheManager) *WalletManager {
	wm := WalletManager{}
	wm.WalletManager = bitshares.NewWalletManager(cacheManager)
	wm.Config = bitshares.NewConfig(Symbol)
	wm.Decoder = NewAddressDecoder(&wm)
	wm.Log = log.NewOWLogger(wm.Symbol())
	return &wm
}
