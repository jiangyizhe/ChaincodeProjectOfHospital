package service3

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"time"
)


type PatientInfo struct {

	ObjectType string `json:"patType"`
	Patient_Id string
	Business_number string      //业务号
	Name                 string //姓名
	Sex                  string //性别
	Hosnum               string //hosnum
	HospitalizedType     string //住院类型
	HospitalizedDay      string //住院日期
	DiseasesName         string //病症名称
	Note                 string //备注
	Department           string //治疗科室
	BedNum               string //床号
	EntityID             string //身份证号码
	Reimbursement_Date   string //报销日期
	Hospital_Location    string //医院地点
	Entry_Time           string //录入时间
	ReAmount             string //原有金额
	Reimbursement_Amount string //报销金额
	Mypay string                //剩余金额

	Historys   []HistoryItem  // 当前的历史记录
}

type HistoryItem struct {
	TxId   string
	PatientInfo    PatientInfo
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
