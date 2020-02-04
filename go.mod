module github.com/assetsadapterstore/barc-adapter

go 1.13

require (
	github.com/astaxie/beego v1.12.0
	github.com/blocktree/bitshares-adapter v1.2.0
	github.com/blocktree/go-owcdrivers v1.2.0
	github.com/blocktree/go-owcrypt v1.1.1
	github.com/blocktree/openwallet v1.7.0
	github.com/denkhaus/bitshares v0.6.1-0.20190502142618-5ae8c00cb394
)

replace github.com/imroc/req => github.com/blocktree/req v0.2.5

//replace github.com/blocktree/bitshares-adapter => ../bitshares-adapter
