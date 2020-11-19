package main

import (

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"fmt"
	"github.com/hyperledger/fabric/protos/peer"
)

type PatientChaincode struct {

}



func (t *PatientChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response{

	return shim.Success(nil)
}

func (t *PatientChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response{
	// 获取用户意图
	fun, args := stub.GetFunctionAndParameters()
	if fun == "addPat"{
		return t.addPat(stub, args)		// 添加信息
	}else if fun == "queryPatByNoAndNames" {
		return t.queryPatByNoAndName(stub, args)		// 根据证书编号及姓名查询信息
	}else if fun == "queryPatByID" {
		return t.queryPatInfoByID(stub, args)	// 根据身份证号码及姓名查询详情
	}

	return shim.Error("指定的函数名称错误")

}

//***********************************************************************************************************


func main(){
	err := shim.Start(new(PatientChaincode))
	if err != nil{
		fmt.Printf("启动Chaincode时发生错误: %s", err)
	}
}
