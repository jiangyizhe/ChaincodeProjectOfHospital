
package sdkInit

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
)

type InitInfo struct {
	ChannelID     string
	ChannelConfig string
	OrgAdmin      string
	OrgName       string
	OrdererOrgName	string
	OrgResMgmt *resmgmt.Client

	ChaincodeID	string
	ChaincodeID2	string
	ChaincodeID3	string
	ChaincodeGoPath	string
	ChaincodePath	string
	ChaincodePath2	string
	ChaincodePath3	string
	UserName	string
}
