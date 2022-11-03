#!/bin/bash

if [[ `uname` == 'Darwin' ]]; then
    echo "Mac OS"
    export PATH=${PWD}/hyperledger-fabric-darwin-amd64-1.4.12/bin:$PATH
fi
if [[ `uname` == 'Linux' ]]; then
    echo "Linux"
    export PATH=${PWD}/hyperledger-fabric-linux-amd64-1.4.12/bin:$PATH
fi

echo "1. Clean up the environment"
./stop.sh

echo "2. Generate certificates and keys（ MSP Materials), the resulting result will be preserved in crypto-config In the folder"
Echo "2. Generate certificate and key (MSP material), and the result will be saved in the Crypto-config folder"
cryptogen generate --config=./crypto-config.yaml

echo "3. Create a Sorting Channel Creation Block"
configtxgen -profile TwoOrgsOrdererGenesis -outputBlock ./config/genesis.block -channelID firstchannel

echo "Fourth, generate channel configuration transactions'appchannel.tx'"
configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ./config/appchannel.tx -channelID appchannel

echo "Five, for TDefine anchor node Define anchor node"
configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./config/TaobaoAnchor.tx -channelID appchannel -asOrg Taobao

echo "Six, for the sake of JD Define anchor node"
configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./config/JDAnchor.tx -channelID appchannel -asOrg JD

echo "Blockchain: Start"
docker-compose up -d
echo "Waiting for the start of the node, wait for 10 seconds"
sleep 10

TaobaoPeer0Cli="CORE_PEER_ADDRESS=peer0.taobao.com:7051 CORE_PEER_LOCALMSPID=TaobaoMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/taobao.com/users/Admin@taobao.com/msp"
TaobaoPeer1Cli="CORE_PEER_ADDRESS=peer1.taobao.com:7051 CORE_PEER_LOCALMSPID=TaobaoMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/taobao.com/users/Admin@taobao.com/msp"
JDPeer0Cli="CORE_PEER_ADDRESS=peer0.jd.com:7051 CORE_PEER_LOCALMSPID=JDMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/jd.com/users/Admin@jd.com/msp"
JDPeer1Cli="CORE_PEER_ADDRESS=peer1.jd.com:7051 CORE_PEER_LOCALMSPID=JDMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/jd.com/users/Admin@jd.com/msp"

echo "Seven, Create Channel"
docker exec cli bash -c "$TaobaoPeer0Cli peer channel create -o orderer.qq.com:7050 -c appchannel -f /etc/hyperledger/config/appchannel.tx"

echo "8. Add all nodes to the channel"
docker exec cli bash -c "$TaobaoPeer0Cli peer channel join -b appchannel.block"
docker exec cli bash -c "$TaobaoPeer1Cli peer channel join -b appchannel.block"
docker exec cli bash -c "$JDPeer0Cli peer channel join -b appchannel.block"
docker exec cli bash -c "$JDPeer1Cli peer channel join -b appchannel.block"

echo "9. Update anchor node"
docker exec cli bash -c "$TaobaoPeer0Cli peer channel update -o orderer.qq.com:7050 -c appchannel -f /etc/hyperledger/config/TaobaoAnchor.tx"
docker exec cli bash -c "$JDPeer0Cli peer channel update -o orderer.qq.com:7050 -c appchannel -f /etc/hyperledger/config/JDAnchor.tx"

# -N chain code name, you can set it by yourself
# -V version number
# -P chain code directory, under/opt/gopath/src/directory
echo "Ten, install chain code"
docker exec cli bash -c "$TaobaoPeer0Cli peer chaincode install -n fabric-realty -v 1.0.0 -l golang -p chaincode"
docker exec cli bash -c "$JDPeer0Cli peer chaincode install -n fabric-realty -v 1.0.0 -l golang -p chaincode"

# Only need one of the nodes instantiated
# -N Corresponding to the name of the chain code in the previous step
# -V version number
# -C is a channel, in the world of Fabric, a channel is a different chain
# -c is a parameter, and passed into the init parameter
echo "Eleven, instantiated chain codeinstantiated chain codeinstantiated chain code"
docker exec cli bash -c "$TaobaoPeer0Cli peer chaincode instantiate -o orderer.qq.com:7050 -C appchannel -n fabric-realty -l golang -v 1.0.0 -c '{\"Args\":[\"init\"]}' -P \"AND ('TaobaoMSP.member','JDMSP.member')\""

echo "Waiting for the chain code instantiated completion and waiting for 5 seconds chain code instantiated completion and waiting for 5 seconds"
sleep 5

# Perform the chain code interaction, verify whether the chain code is correctly installed and the blockchain network can work properlyction, verify whether the chain code is correctly installed and the blockchain network can work properlyction, verify whether the chain code is correctly installed and the blockchain network can work properly
echo "Twelve, verification chain code"
docker exec cli bash -c "$TaobaoPeer0Cli peer chaincode invoke -C appchannel -n fabric-realty -c '{\"Args\":[\"hello\"]}'"
docker exec cli bash -c "$JDPeer0Cli peer chaincode invoke -C appchannel -n fabric-realty -c '{\"Args\":[\"hello\"]}'"

