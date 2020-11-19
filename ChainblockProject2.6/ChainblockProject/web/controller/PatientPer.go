package controller

import (
	"encoding/json"
	"fmt"
	"github.com/ChainblockProject/service3"
	"github.com/ChainblockProject/web/model"
	"github.com/ChainblockProject/web/services"
	"math/rand"
	"net/http"
	"strconv"
)

type List2 struct {
	PatientPer []model.PatientPer
}

func (app *Application)AddPatientPer(){
	_, err1 := model.DelPatientPer()
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	i:=0
	var data model.PatientPer

	for {
		result, _ := app.Setup3.FindPatientInfoByID("pat"+strconv.Itoa(i))
		var med = service3.PatientInfo{}
		json.Unmarshal(result, &med)
		if med.Patient_Id=="" {
			break;
		}
		data.Patient_Id=med.Patient_Id
		data.Sex=med.Sex
		data.Mypay=med.Mypay
		data.Entry_Time=med.Entry_Time
		data.Business_number=med.Business_number
		data.Department=med.Department
		data.EntityID=med.EntityID
		data.Name=med.Name
		data.BedNum=med.BedNum
		data.DiseasesName=med.DiseasesName
		data.Hosnum=med.Hosnum
		data.Hospital_Location=med.Hospital_Location
		data.HospitalizedDay=med.HospitalizedDay
		data.HospitalizedType=med.HospitalizedType
		data.Note=med.Note
		data.ObjectType=med.ObjectType
		data.ReAmount=med.ReAmount
		data.Reimbursement_Date=med.Reimbursement_Date
		data.Reimbursement_Amount=med.Reimbursement_Amount

		i++
		a,_:=strconv.Atoi(data.Patient_Id)
		if services.IsData(a){
			services.AddPatientPer(data)
		}

	}
}

func (app *Application)GetPatientPer(w http.ResponseWriter, r *http.Request){
	var sum float64

	var record []model.PatientPer//定义结构体数组
	record, err := services.GetPatientPer()
	if err != nil {
		fmt.Println(err)
		return
	}

	for i:=0;i<len(record);i++{
		sum = 0
		if record[i].Sex=="1"{
			record[i].Sex="男"
		}else {
			record[i].Sex="女"
		}

		pri, err := services.GetAuditMedByBus(record[i].Business_number)
		if err != nil {
			fmt.Println(err)
			return
		}
		if(len(pri)>1){
			for j:=0;j<len(pri);j++{
				float,_ := strconv.ParseFloat(pri[j].Money,64)
				sum+=float
			}
		}

		value := int64(sum)
		record[i].ReAmount=strconv.FormatInt(value,10)
		value =value/1000*9
		record[i].ReAmount=record[i].ReAmount+"."+strconv.FormatInt(value,10)
	}


	data := &struct {
		PatientPer []model.PatientPer
		CurrentUser User
		Msg string
		Flag bool
		History bool
	}{
		PatientPer:record,
		CurrentUser:cuser,
		Msg:"",
		Flag:false,
		History:true,
	}

	ShowView(w, r, "patdata.html", data)
}

func (app *Application)GetPatientPerResult(w http.ResponseWriter, r *http.Request){
	var sum float64

	var record []model.PatientPer//定义结构体数组
	record, err := services.GetPatientPer()
	if err != nil {
		fmt.Println(err)
		return
	}

	for i:=0;i<len(record);i++{
		sum = 0
		if record[i].Sex=="1"{
			record[i].Sex="男"
		}else {
			record[i].Sex="女"
		}

		pri, err := services.GetAuditMedByBus(record[i].Business_number)
		if err != nil {
			fmt.Println(err)
			return
		}
		if(len(pri)>1){
			for j:=0;j<len(pri);j++{
				float,_ := strconv.ParseFloat(pri[j].Money,64)
				sum+=float
			}
		}

		value := int64(sum)
		n:=rand.Intn(33)+56
		mon :=value*int64(n)/100
		record[i].ReAmount=strconv.FormatInt(value,10)
		value =value/1000*9

		record[i].ReAmount=record[i].ReAmount+"."+strconv.FormatInt(value,10)


		n2:=rand.Intn(98)+1

		if record[i].ReAmount=="0.0"{
			record[i].Reimbursement_Amount="0"
			record[i].Mypay="0"
		}else {
			record[i].Mypay="0."+strconv.FormatInt(int64(n),10)
			record[i].Reimbursement_Amount=strconv.FormatInt(int64(mon),10)+"."+strconv.FormatInt(int64(n2),10)
		}

	}



	data := &struct {

		PatientPer []model.PatientPer
		CurrentUser User
		Msg string
		Flag bool
		History bool
	}{
		PatientPer:record,
		CurrentUser:cuser,
		Msg:"",
		Flag:false,
		History:true,
	}

	ShowView(w, r, "patresult.html", data)
}

func (app *Application)GetPatientPerResultByEnd(w http.ResponseWriter, r *http.Request){
	var sum float64
	End := r.FormValue("end")
	var record []model.PatientPer//定义结构体数组
	record, err := services.GetPatientPerByEnd(End)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i:=0;i<len(record);i++{
		sum = 0
		if record[i].Sex=="1"{
			record[i].Sex="男"
		}else {
			record[i].Sex="女"
		}

		pri, err := services.GetAuditMedByBus(record[i].Business_number)
		if err != nil {
			fmt.Println(err)
			return
		}
		if(len(pri)>1){
			for j:=0;j<len(pri);j++{
				float,_ := strconv.ParseFloat(pri[j].Money,64)
				sum+=float
			}
		}

		value := int64(sum)
		n:=rand.Intn(33)+56
		mon :=value*int64(n)/100
		record[i].ReAmount=strconv.FormatInt(value,10)
		value =value/1000*9

		record[i].ReAmount=record[i].ReAmount+"."+strconv.FormatInt(value,10)


		n2:=rand.Intn(98)+1

		if record[i].ReAmount=="0.0"{
			record[i].Reimbursement_Amount="0"
			record[i].Mypay="0"
		}else {
			record[i].Mypay="0."+strconv.FormatInt(int64(n),10)
			record[i].Reimbursement_Amount=strconv.FormatInt(int64(mon),10)+"."+strconv.FormatInt(int64(n2),10)
		}

	}



	data := &struct {

		PatientPer []model.PatientPer
		CurrentUser User
		Msg string
		Flag bool
		History bool
	}{
		PatientPer:record,
		CurrentUser:cuser,
		Msg:"",
		Flag:false,
		History:true,
	}

	ShowView(w, r, "patresult.html", data)
}


func (app *Application)GetPatientPerByEnd(w http.ResponseWriter, r *http.Request){
	var sum float64
	End := r.FormValue("end")
	var record []model.PatientPer//定义结构体数组
	record, err := services.GetPatientPerByEnd(End)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i:=0;i<len(record);i++{
		sum = 0
		if record[i].Sex=="1"{
			record[i].Sex="男"
		}else {
			record[i].Sex="女"
		}

		pri, err := services.GetAuditMedByBus(record[i].Business_number)
		if err != nil {
			fmt.Println(err)
			return
		}
		if(len(pri)>1){
			for j:=0;j<len(pri);j++{
				float,_ := strconv.ParseFloat(pri[j].Money,64)
				sum+=float
			}
		}

		value := int64(sum)
		record[i].ReAmount=strconv.FormatInt(value,10)
		value =value/1000*9
		record[i].ReAmount=record[i].ReAmount+"."+strconv.FormatInt(value,10)
	}


	data := &struct {
		PatientPer []model.PatientPer
		CurrentUser User
		Msg string
		Flag bool
		History bool
	}{
		PatientPer:record,
		CurrentUser:cuser,
		Msg:"",
		Flag:false,
		History:true,
	}

	ShowView(w, r, "patdatabyend.html", data)
}