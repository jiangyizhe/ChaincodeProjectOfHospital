package main

import (

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"fmt"
	"github.com/hyperledger/fabric/protos/peer"
)

type AuditChaincode struct {

}



func (t *AuditChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response{

	return shim.Success(nil)
}

func (t *AuditChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response{
	// 获取用户意图
	fun, args := stub.GetFunctionAndParameters()
	if fun == "addMed"{
		return t.addMed(stub, args)		// 添加信息
	}else if fun == "queryMedByNoAndName" {
		return t.queryMedByNoAndName(stub, args)		// 根据证书编号及姓名查询信息
	}else if fun == "queryMedInfoByID" {
		return t.queryMedInfoByID(stub, args)	// 根据身份证号码及姓名查询详情
	}

	return shim.Error("指定的函数名称错误")

}

//***********************************************************************************************************


func main(){
	err := shim.Start(new(AuditChaincode))
	if err != nil{
		fmt.Printf("启动EducationChaincode时发生错误: %s", err)
	}
}
