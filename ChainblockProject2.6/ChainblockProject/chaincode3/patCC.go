
package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"encoding/json"
	"fmt"
	"bytes"
)

const DOC_TYPE = "patObj"

// 保存pat
// args: patcation
func PutPat(stub shim.ChaincodeStubInterface, pat PatientInfo) ([]byte, bool) {

	pat.ObjectType = DOC_TYPE

	b, err := json.Marshal(pat)
	if err != nil {
		return nil, false
	}

	// 保存pat状态
	err = stub.PutState(pat.Patient_Id, b)
	if err != nil {
		return nil, false
	}

	return b, true
}

// 根据业务号查询信息状态
// args: Business_number
func GetPatInfo(stub shim.ChaincodeStubInterface, Audit_Id string) (PatientInfo, bool)  {
	var pat PatientInfo
	// 根据业务号号码查询信息状态
	b, err := stub.GetState(Audit_Id)
	if err != nil {
		return pat, false
	}

	if b == nil {
		return pat, false
	}

	// 对查询到的状态进行反序列化
	err = json.Unmarshal(b, &pat)
	if err != nil {
		return pat, false
	}

	// 返回结果
	return pat, true
}

// 根据指定的查询字符串实现富查询
func getPatByQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer  resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}

		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		bArrayMemberAlreadyWritten = true
	}

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil

}

// 添加信息
// args: patObject
// 身份证号为 key, Pat为 value
func (t *PatientChaincode) addPat(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 2{
		return shim.Error("给定的参数个数不符合要求")
	}

	var pat PatientInfo
	err := json.Unmarshal([]byte(args[0]), &pat)
	if err != nil {
		return shim.Error("反序列化信息时发生错误")
	}

	// 查重: 身份证号码必须唯一
	_, exist := GetPatInfo(stub, pat.Patient_Id)
	if exist {
		return shim.Error("要添加的身份证号码已存在")
	}

	_, bl := PutPat(stub, pat)
	if !bl {
		return shim.Error("保存信息时发生错误")
	}

	err = stub.SetEvent(args[1], []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("信息添加成功"))
}

// 根据证书编号及姓名查询信息
// args: CertNo, name
func (t *PatientChaincode) queryPatByNoAndName(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 2 {
		return shim.Error("给定的参数个数不符合要求")
	}
	Business_number := args[0]
	Name := args[1]

	// 拼装CouchDB所需要的查询字符串(是标准的一个JSON串)
	// queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"patObj\", \"CertNo\":\"%s\"}}", CertNo)
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"%s\", \"Business_number\":\"%s\", \"Name\":\"%s\"}}", DOC_TYPE, Business_number, Name)

	// 查询数据
	result, err := getPatByQueryString(stub, queryString)
	if err != nil {
		return shim.Error("根据证书编号及姓名查询信息时发生错误")
	}
	if result == nil {
		return shim.Error("根据指定的证书编号及姓名没有查询到相关的信息")
	}
	return shim.Success(result)
}

// 根据身份证号码查询详情（溯源）
// args: entityID
func (t *PatientChaincode) queryPatInfoByID(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("给定的参数个数不符合要求")
	}

	// 根据身份证号码查询pat状态
	b, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("根据身份证号码查询信息失败")
	}

	if b == nil {
		return shim.Error("根据身份证号码没有查询到相关的信息")
	}

	// 对查询到的状态进行反序列化
	var pat PatientInfo
	err = json.Unmarshal(b, &pat)
	if err != nil {
		return  shim.Error("反序列化pat信息失败")
	}

	// 获取历史变更数据
	iterator, err := stub.GetHistoryForKey(pat.Patient_Id)
	if err != nil {
		return shim.Error("根据指定的身份证号码查询对应的历史变更数据失败")
	}
	defer iterator.Close()

	// 迭代处理
	var historys []HistoryItem
	var hisPat PatientInfo
	for iterator.HasNext() {
		hisData, err := iterator.Next()
		if err != nil {
			return shim.Error("获取pat的历史变更数据失败")
		}

		var historyItem HistoryItem
		historyItem.TxId = hisData.TxId
		json.Unmarshal(hisData.Value, &hisPat)

		if hisData.Value == nil {
			var empty PatientInfo
			historyItem.PatientInfo= empty
		}else {
			historyItem.PatientInfo = hisPat
		}

		historys = append(historys, historyItem)

	}

	pat.Historys = historys

	// 返回
	result, err := json.Marshal(pat)
	if err != nil {
		return shim.Error("序列化pat信息时发生错误")
	}
	return shim.Success(result)
}

