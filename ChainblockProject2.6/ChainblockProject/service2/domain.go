
package service2

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"time"
)

type Audit struct {

	ObjectType	string	`json:"medType"`
	Audit_Id string
	Business_number string //业务号
	Cost_Classes string //费用类别
	Cost_Classes_Name string //费用类别名称
	Main_Drugs_Name string //中心药品名称
	Main_Drugs_Code string //中心药品编码
	Hospital_Drug_Name string //医院药品名称
	Drug_Dosage_Form string //剂型
	Manufactor string //厂家
	Specifications string //规格
	Number string //数量
	Baimonovalent string //单价
	Money string //金额
	Cost_Time string //金额时间
	Entry_Time string//录入时间
	Bookkeeper string //记账人
	Mypay string //自付金额
	Proportion float64 //比例
	Hospital_Drug_Code string//医院药品编码

	Historys	[]HistoryItem	// 当前的历史记录
}

type HistoryItem struct {
	TxId	string
	Audit	Audit
}

type ServiceSetup struct {
	ChaincodeID	string
	Client	*channel.Client
}

func regitserEvent(client *channel.Client, chaincodeID, eventID string) (fab.Registration, <-chan *fab.CCEvent) {

	reg, notifier, err := client.RegisterChaincodeEvent(chaincodeID, eventID)
	if err != nil {
		fmt.Println("注册链码事件失败: %s", err)
	}
	return reg, notifier
}

func eventResult(notifier <-chan *fab.CCEvent, eventID string) error {
	select {
	case ccEvent := <-notifier:
		fmt.Printf("接收到链码事件: %v\n", ccEvent)
	case <-time.After(time.Second * 20):
		return fmt.Errorf("不能根据指定的事件ID接收到相应的链码事件(%s)", eventID)
	}
	return nil
}
