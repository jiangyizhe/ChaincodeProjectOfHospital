package controller

import (
	"encoding/json"
	"fmt"
	"github.com/ChainblockProject/service2"
	"github.com/ChainblockProject/web/model"
	"github.com/ChainblockProject/web/services"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type List struct {
	AuditMed []model.AuditMed
}

func (app *Application)AddAuditMed(){
	_, err1 := model.DelAuditMed()
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	i:=0
	var data model.AuditMed
	for {
		result, _ := app.Setup2.FindMedInfoByID(strconv.Itoa(i))
		var med = service2.Audit{}
		json.Unmarshal(result, &med)
		if med.Audit_Id=="" {
			break;
		}
		data.Number=med.Number
		data.Audit_id=med.Audit_Id
		data.Baimonovalent=med.Baimonovalent
		data.Bookkeeper=med.Bookkeeper
		data.Business_number=med.Business_number
		data.Cost_Classes=med.Cost_Classes
		data.Cost_Classes_Name=med.Cost_Classes_Name
		data.Cost_Time=med.Cost_Time
		data.Drug_Dosage_Form=med.Drug_Dosage_Form
		data.Entry_Time=med.Entry_Time
		data.Hospital_Drug_Code=med.Hospital_Drug_Code
		data.Hospital_Drug_Name=med.Hospital_Drug_Name
		data.Main_Drugs_Code=med.Main_Drugs_Code
		data.Main_Drugs_Name=med.Main_Drugs_Name
		data.Manufactor=med.Manufactor
		data.Money=med.Money
		data.Mypay=med.Mypay
		data.Proportion =med.Proportion
		data.Specifications=med.Specifications

		reasons := make([]string, 0)
		reasons = append(reasons,

			"符合要求",
		)
		rand.Seed(time.Now().Unix()) // initialize global pseudo random generator

		data.Reason=reasons[rand.Intn(len(reasons))]
		if data.Reason=="符合要求"{
			data.IsRight="是"
		}else {
			data.IsRight="否"
		}
		i++

		a,_:=strconv.Atoi(data.Audit_id)
		if services.IsData2(a){
			services.AddAuditMed(data)
		}

	}
}

func (app *Application)GetAuditMed(w http.ResponseWriter, r *http.Request){

	var record []model.AuditMed //定义结构体数组
	record, err := services.GetAuditMed()
	if err != nil {
		fmt.Println(err)
		return
	}




	data := &struct {
		AuditMed []model.AuditMed
		CurrentUser User
		Msg string
		Flag bool
		History bool
	}{
		AuditMed:record,
		CurrentUser:cuser,
		Msg:"",
		Flag:false,
		History:true,
	}

	ShowView(w, r, "meddata.html", data)
}

func (app *Application)GetAuditMedByBus(w http.ResponseWriter, r *http.Request){
	Business_number := r.FormValue("Business_number")
	var record []model.AuditMed //定义结构体数组
	record, err := services.GetAuditMedByBus(Business_number)
	if err != nil {
		fmt.Println(err)
		return
	}




	data := &struct {
		AuditMed []model.AuditMed
		CurrentUser User
		Msg string
		Flag bool
		History bool
	}{
		AuditMed:record,
		CurrentUser:cuser,
		Msg:"",
		Flag:false,
		History:true,
	}

	ShowView(w, r, "meddatabybus.html", data)
}

func (app *Application)IsRight(w http.ResponseWriter, r *http.Request){
	audit_id,_ := strconv.Atoi(r.FormValue("audit_id"))
	var record model.AuditMed //定义结构体数组

	record, _ = model.GetAuditMedByID(audit_id)
	record.IsRight="否"
	model.SavaAuditMed(record)
	app.GetAuditMed(w,r)
}