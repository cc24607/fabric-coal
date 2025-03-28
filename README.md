# fabric-CoalMine
本文是基于hyperledger-fabric框架的Coal Mine Scheduling Speech Storage Based on Consortium Blockchain
，实现语音转文本后的数据上链存证、智能合约溯源码生成及提高煤矿企业监管透明性。
## 一、启动区块链网络

### 1.首先进入用于启动网络的脚本所在的目录
> cd trace-fabric/blockchain-trace-bcnetwork/basic-network

### 2.运行basic_network目录下的start.sh文件
> chmod -R 777 start.sh

>./start.sh

### 3.运行trace-fabric/blockchain-trace-bcnetwork/webapp目录下的./start.sh
先文件授权
```
chmod -R 777 startFarmerCC.sh  
```

### 4.执行npm install安装依赖  
> npm install  
 

### 5.注册用户
> node enrollAdmin.js  

> node registerUser.js  


### 6.启动node服务 
> node app.js  

如需常驻后台，使用pm2。
> 启动：pm2 start app.js

> 停止：pm2 stop app.js  

至此，完成区块链网络的部署  

## 二.PC端启动 blockchain-trace-pc

### 1.安装依赖
> npm install --registry=https://registry.npm.taobao.org

### 2.修改连接区块链网络地址
main.js， Vue.prototype.$httpUrl 修改为区块链网络所在服务器地址
```yaml
Vue.prototype.$httpUrl = "http://localhost:8080/route";
```
### 3.启动项目
> npm run dev


## 三、测试
1.首先进入测试工具Caliper CLI所在的目录：  
```
cd hyperledger/caliper/workspace
```

2.通过以下命令，利用 NPM 软件包安装 Caliper CLI。
根据"config.yaml,test-network.yaml"和  
"./hyperledger/caliper/workspace/benchmarks/samples/fabric/basic/"下的各种待测试的合约函数配置文件进行基准测试。
```
npm install --only=prod @hyperledger/caliper-cli@0.6.0
npx caliper bind --caliper-bind-sut fabric:2.5
npx caliper launch manager --caliper-workspace ./ --caliper-networkconfig networks/fabric/test-network.yaml \
--caliper-benchconfig benchmarks/samples/fabric/basic/config.yaml --caliper-flow-only-test --caliper-fabric-gateway-enabled
```



