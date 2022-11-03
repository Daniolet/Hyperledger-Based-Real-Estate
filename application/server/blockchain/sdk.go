package blockchain

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

// Configuration information
var (
	sdk           *fabsdk.FabricSDK                              // Fabric SDK
	configPath    = "config.yaml"                                // Configuration file path
	channelName   = "appchannel"                                 // Channel name
	user          = "Admin"                                      // user
	chainCodeName = "fabric-realty"                              // Chain code name
	endpoints     = []string{"peer0.jd.com", "peer0.taobao.com"} // To send the node of the transaction
)

// Init initialization
func Init() {
	var err error
	// Initialize SDK by configuration file
	sdk, err = fabsdk.New(config.FromFile(configPath))
	if err != nil {
		panic(err)
	}
}

// ChannelExecute Blockchain interconnection
func ChannelExecute(fcn string, args [][]byte) (channel.Response, error) {
	// Create a client to indicate the identity in the channel
	ctx := sdk.ChannelContext(channelName, fabsdk.WithUser(user))
	cli, err := channel.New(ctx)
	if err != nil {
		return channel.Response{}, err
	}
	// Write operation of blockchain ledger (invoke of the chain code)
	resp, err := cli.Execute(channel.Request{
		ChaincodeID: chainCodeName,
		Fcn:         fcn,
		Args:        args,
	}, channel.WithTargetEndpoints(endpoints...))
	if err != nil {
		return channel.Response{}, err
	}
	//The result after the chain code is executed
	return resp, nil
}

// ChannelQuery Blockchain query
func ChannelQuery(fcn string, args [][]byte) (channel.Response, error) {
	// Create a client to indicate the identity in the channel
	ctx := sdk.ChannelContext(channelName, fabsdk.WithUser(user))
	cli, err := channel.New(ctx)
	if err != nil {
		return channel.Response{}, err
	}
	// The operation of the blockchain account query (calling invoke of the chain code), only the result returns the result
	resp, err := cli.Query(channel.Request{
		ChaincodeID: chainCodeName,
		Fcn:         fcn,
		Args:        args,
	}, channel.WithTargetEndpoints(endpoints...))
	if err != nil {
		return channel.Response{}, err
	}
	//The result after the chain code is executed
	return resp, nil
}
