# barc-adapter

## 项目依赖库

- [openwallet](https://github.com/blocktree/openwallet.git)

## 如何测试

openwtester包下的测试用例已经集成了openwallet钱包体系，创建conf目录，新建bar.ini文件，编辑如下内容：

```ini

#wallet api url
ServerAPI = "https://localhost:8080"
# Cache data file directory, default = "", current directory: ./data
dataDir = ""

```

## 浏览器
http://block.barc.io/#/

## Github
https://github.com/bar-chain/bar-core

## 相关文档
https://doc.barc.io/