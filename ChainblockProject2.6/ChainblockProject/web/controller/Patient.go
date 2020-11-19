package controller

import (
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/ChainblockProject/service"
	"github.com/ChainblockProject/service3"
	"github.com/ChainblockProject/web/model"
	"github.com/ChainblockProject/web/services"
	_ "github.com/ChainblockProject/web/services"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)
const (
	patientInfo_path string = "./files/"
)



func (app *Application)AddPatientShow(w http.ResponseWriter, r *http.Request){
	data:=&struct {
		CurrentUser User
		Msg string
		Flag bool
	}{
		CurrentUser:cuser,
		Msg:"",
		Flag:false,
	}
	ShowView(w,r,"uploadpatdata.html",data)
}
func (app *Application)AddPatientInfo(w http.ResponseWriter,r *http.Request){
	file,head,err:=r.FormFile("file")
	if err!=nil{
		log.Println(err)
	}
	defer file.Close()
	filename := time.Now().Format("20060102150405")
	//获取文件后缀
	fileSuffix := path.Ext(head.Filename)
	filePath := patientInfo_path + filename + fileSuffix
	fW, err := os.Create(filePath)
	if err!=nil{
		log.Println("文件创建失败",err)
		return
	}
	defer fW.Close()
	_, err = io.Copy(fW, file)
	if err != nil {
		fmt.Println("文件保存失败")
		return
	}
	if fileSuffix == ".xlsx" {
		fileXlsx(filePath)
	}
	println(filePath)
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	var info model.PatientInfo
	rows := f.GetRows("Sheet1")
	b:=0

	var a [17]string
	for _, row := range rows {
		i:=0
		for _, value := range row {
			a[i]=value
			i++
		}
		info.Patient_Id=b
		info.Business_number =a[0]
		info.Name =a[1]
		info.Sex =a[2]
		info.Hosnum =a[3]
		info.HospitalizedType =a[4]
		info.HospitalizedDay =a[5]
		info.DiseasesName =a[6]
		info.Note =a[7]
		info.Department =a[8]
		info.BedNum =a[9]
		info.EntityID =a[10]
		info.Reimbursement_Date =a[11]
		info.Hospital_Location =a[12]
		info.Entry_Time =a[13]
		info.ReAmount =a[14]
		info.Reimbursement_Amount =a[15]
		info.Mypay =a[16]

		var edu service.Education
		info.EntityID = strings.Replace(info.EntityID, " ", "", -1)
		println(info.EntityID)
		result, _ := app.Setup.FindEduInfoByEntityID(info.EntityID)
		json.Unmarshal(result, &edu)

		if edu.Name==""{
			fmt.Println("该"+info.Name+"用户不是医保人")
		}else{
			services.AddPatientInfo(info)
			b++
		}

	}



	data := &struct {
		CurrentUser User
		Msg string
		Flag bool
	}{
		CurrentUser:cuser,
		Msg:"",
		Flag:false,
	}
	ShowView(w, r, "help.html", data)
}

func(app *Application) AddFabric2(){

	var record []model.PatientInfo//定义结构体数组
	record, err := services.GetPatientInfo()
	if err != nil {
		fmt.Println(err)
		return
	}

	var data service3.PatientInfo
	i := 0
	id := 0
	for {
		result, _ := app.Setup3.FindPatientInfoByID("pat"+strconv.Itoa(i))
		var med = service3.PatientInfo{}
		json.Unmarshal(result, &med)
		if med.Patient_Id == "" {
			id=i
			break;
		}
		i++
	}



	for i:=0;i<len(record);i++{

		data.Patient_Id="pat"+strconv.Itoa(id+i)
		data.Sex=record[i].Sex
		data.Mypay=record[i].Mypay
		data.Entry_Time=record[i].Entry_Time
		data.Business_number=record[i].Business_number
		data.Department=record[i].Department
		data.EntityID=record[i].EntityID
		data.Name=record[i].Name
		data.BedNum=record[i].BedNum
		data.DiseasesName=record[i].DiseasesName
		data.Hosnum=record[i].Hosnum
		data.Hospital_Location=record[i].Hospital_Location
		data.HospitalizedDay=record[i].HospitalizedDay
		data.HospitalizedType=record[i].HospitalizedType
		data.Note=record[i].Note
		data.ObjectType=record[i].ObjectType
		data.ReAmount=record[i].ReAmount
		data.Reimbursement_Date=record[i].Reimbursement_Date
		data.Reimbursement_Amount=record[i].Reimbursement_Amount
		patient, _ := model.SelectPatientPerByEnd(record[i].EntityID)
		if len(patient)<1{

		app.Setup3.SavePatient(data)
		}
	}

	_, err1 := model.DelPatient()
	if err1 != nil {
		fmt.Println(err)
		return
	}
}
