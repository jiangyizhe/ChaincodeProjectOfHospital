package main

import (
	_ "encoding/json"
	"fmt"
	"github.com/ChainblockProject/sdkInit"
	_ "github.com/ChainblockProject/sdkInit"
	_ "github.com/ChainblockProject/web"
	_ "github.com/ChainblockProject/web/controller"
	"os"
	_ "os"
)


func main() {

	////初始化
	//if err := model.InitSQLite(); err != nil {
	//	panic("数据库初始化失败")
	//}
	//defer model.Close()
	//sdk, err := fabsdk.New(config.FromFile(configFile))
	////channelClient, err := sdkInit.InstallAndInstantiateCC(sdk, initInfo)
	//clientChannelContext := sdk.ChannelContext("shuibianhospital", fabsdk.WithUser("User1"), fabsdk.WithOrg("Org1"))
	//// returns a Client instance. Channel client can query chaincode1, execute chaincode1 and register/unregister for chaincode1 events on specific channel.
	//fmt.Println("通道客户端创建成功，可以利用此客户端调用链码进行查询或执行事务.")
	//channelClient, err := channel.New(clientChannelContext)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//fmt.Println(channelClient)

	//===========================================//
	initInfo := &sdkInit.InitInfo{

		ChannelID: "shuibianhospital",
		ChannelConfig: os.Getenv("GOPATH") + "/src/github.com/ChainblockProject/fixtures/artifacts/channel.tx",

		OrgAdmin:"Admin",
		OrgName:"Org1",
		OrdererOrgName: "orderer.shuibian.hospital.com",

		ChaincodeID: EduCC,
		ChaincodeID2: MedCC,
		ChaincodeID3: PatCC,
		ChaincodeGoPath: os.Getenv("GOPATH"),
		ChaincodePath: "github.com/ChainblockProject/chaincode1/",
		ChaincodePath2: "github.com/ChainblockProject/chaincode2/",
		ChaincodePath3: "github.com/ChainblockProject/chaincode3/",
		UserName:"User1",
	}

	sdk, err := sdkInit.SetupSDK(configFile, initialized)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	defer sdk.Close()
	err = sdkInit.CreateChannel(sdk, initInfo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}//


	channelClient, err := sdkInit.InstallAndInstantiateCC(sdk, initInfo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(channelClient)

	//serviceSetup := service.ServiceSetup{
	//	ChaincodeID:EduCC,
	//	Client:channelClient,
	//}
	//serviceSetup2 := service2.ServiceSetup{
	//	ChaincodeID:MedCC,
	//	Client:channelClient,
	//}
	//serviceSetup3 := service3.ServiceSetup{
	//	ChaincodeID:PatCC,
	//	Client:channelClient,
	//}
	//
	//
	//data := service.Education{
	//	Name         :"11",
	//	EntityID    :"11",
	//}

	//var med service.Education
	//serviceSetup.SaveEdu(data)
	//result, err := serviceSetup.FindEduInfoByEntityID("11")
	//json.Unmarshal(result, &med)
	//fmt.Println(med.Name)
	//
	//fmt.Println(result)
	//
	//data2 := service2.Audit{
	//	Audit_Id         :"10",
	//	Number   :"11",
	//}
	//
	//
	//serviceSetup2.SaveMed(data2)
	//result2, err := serviceSetup2.FindMedInfoByID("10")
	//fmt.Println(result2)
	//
	//data3 := service3.PatientInfo{
	//	Patient_Id         :"11",
	//	Sex   :"11",
	//}
	//
	//
	//var med model.PatientPer
	//serviceSetup3.SavePatient(data3)
	//result3, err := serviceSetup3.FindPatientInfoByID("11")
	//json.Unmarshal(result3, &med)
	//fmt.Println(med)



	//app := controller.Application{
	//	Setup: &serviceSetup,
	//	Setup2: &serviceSetup2,
	//	Setup3: &serviceSetup3,
	//}
	//定时器
	//ticker := time.NewTicker(time.Hour * 24)
	//go func() {
	//	for _ = range ticker.C {
	//		app.AddPatientPer()
	//		fmt.Printf("ticked at %v", time.Now())
	//	}
	//}()

	//web.WebStart(app)

}
