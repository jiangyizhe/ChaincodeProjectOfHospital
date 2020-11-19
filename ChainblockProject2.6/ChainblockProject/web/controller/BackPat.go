package controller

import (
	"fmt"
	"github.com/ChainblockProject/web/model"
	"github.com/ChainblockProject/web/services"
	"math/rand"
	"net/http"
	"strconv"
)

func(app *Application) AddBackPat(){

	var rr []model.AuditMed
	rr,_ = model.SelectAuditMedBack()
	for w:=0;w<len(rr);w++{

		var record []model.PatientPer//定义结构体数组
		record, err := model.SelectPatientPerByBus(rr[w].Business_number)
		if err != nil {
			fmt.Println(err)
			return
		}
		var data model.BackPat

		if len(record)>0 {
			i := 0
			data.Patient_Id = record[i].Patient_Id
			data.Sex = record[i].Sex
			data.Mypay = record[i].Mypay
			data.Entry_Time = record[i].Entry_Time
			data.Business_number = record[i].Business_number
			data.Department = record[i].Department
			data.EntityID = record[i].EntityID
			data.Name = record[i].Name
			data.BedNum = record[i].BedNum
			data.DiseasesName = record[i].DiseasesName
			data.Hosnum = record[i].Hosnum
			data.Hospital_Location = record[i].Hospital_Location
			data.HospitalizedDay = record[i].HospitalizedDay
			data.HospitalizedType = record[i].HospitalizedType
			data.Note = record[i].Note
			data.ObjectType = record[i].ObjectType
			data.ReAmount = record[i].ReAmount
			data.Reimbursement_Date = record[i].Reimbursement_Date
			data.Reimbursement_Amount = record[i].Reimbursement_Amount
			data.Status = "已反馈"
			data.Message = "药品不合理"

			a, _ := model.SelectBackPatByBus(data.Business_number)
			if len(a) < 1 {
				model.AddBackPat(data)
			}
		}


	}

}
func (app *Application) AddBackPatShow(w http.ResponseWriter, r *http.Request)  {

	app.AddBackPat()
	data := &struct {
		CurrentUser User
		Msg string
		Flag bool
	}{
		CurrentUser:cuser,
		Msg:"OK",
		Flag:false,
	}
	ShowView(w, r, "help.html", data)
}

func (app *Application)GetBackPat(w http.ResponseWriter, r *http.Request){
	var sum float64


	var record []model.BackPat//定义结构体数组

	record, err := model.SelectBackPat()
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

		PatientPer []model.BackPat
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
	fmt.Println(data)
	ShowView(w, r, "backpat.html", data)
}