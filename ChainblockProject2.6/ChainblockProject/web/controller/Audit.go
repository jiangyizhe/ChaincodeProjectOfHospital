package controller

import (
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/ChainblockProject/service2"
	"github.com/ChainblockProject/web/model"
	"github.com/ChainblockProject/web/services"
	"io"
	"os"
	"path"
	"strconv"
	"time"

	"log"
	"net/http"

)

const (
	upload_path string = "./files/"
)

type History struct {
	Historys []service2.HistoryItem
}

func fileXlsx(filePath string) {
	xlsx, err := excelize.OpenFile(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	rows := xlsx.GetRows("sheet1")
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}

	}
}


func (app *Application)AddAudit(w http.ResponseWriter, r *http.Request){
	//DB:=common.GetDB()
		file, head, err := r.FormFile("file")
		if err != nil {
			fmt.Println(err)
			return
		}

		defer file.Close()
		//当期时间格式化
		filename := time.Now().Format("20060102150405")
		//获取文件的后缀
		fileSuffix := path.Ext(head.Filename)

		filePath := upload_path + filename + fileSuffix
		//创建文件
		fW, err := os.Create(filePath)
		if err != nil {
			fmt.Println("文件创建失败")
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
	b:=0
	var med model.Audit
	rows := f.GetRows("Sheet1")

	var a [18]string
	for _, row := range rows {
		i:=0
		for _, value := range row {
			a[i]=value
			i++
		}
		med.Audit_id=b
		med.Business_number =a[0]
		med.Cost_Classes =a[1]
		med.Cost_Classes_Name =a[2]
		med.Main_Drugs_Name =a[3]
		med.Main_Drugs_Code =a[4]
		med.Hospital_Drug_Name =a[5]
		med.Drug_Dosage_Form =a[6]
		med.Manufactor =a[7]
		med.Specifications =a[8]
		med.Number=a[9]
		med.Baimonovalent =a[10]
		med.Money =a[11]
		med.Cost_Time=a[12]
		med.Entry_Time =a[13]
		med.Bookkeeper =a[14]
		med.Mypay =a[15]
		med.Proportion =a[16]
		med.Hospital_Drug_Code =a[17]

		services.AddAudit(med)
		b++
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

func(app *Application) AddFabric(){

	var record []model.Audit//定义结构体数组
	record, err := services.GetAudit()
	if err != nil {
		fmt.Println(err)
		return
	}

	var data service2.Audit
	i := 0
	id := 0
	for {
		result, _ := app.Setup2.FindMedInfoByID(strconv.Itoa(i))
		var med = service2.Audit{}
		json.Unmarshal(result, &med)
		if med.Audit_Id == "" {
			id=i
			break;
		}
		i++
	}

	for i:=0;i<len(record);i++{

			data.Number=record[i].Number
			data.Audit_Id= strconv.Itoa(id+i)
			data.Baimonovalent=record[i].Baimonovalent
			data.Bookkeeper=record[i].Bookkeeper
			data.Business_number=record[i].Business_number
			data.Cost_Classes=record[i].Cost_Classes
			data.Cost_Classes_Name=record[i].Cost_Classes_Name
			data.Cost_Time=record[i].Cost_Time
			data.Drug_Dosage_Form=record[i].Drug_Dosage_Form
			data.Entry_Time=record[i].Entry_Time
			data.Hospital_Drug_Code=record[i].Hospital_Drug_Code
			data.Hospital_Drug_Name=record[i].Hospital_Drug_Name
			data.Main_Drugs_Code=record[i].Main_Drugs_Code
			data.Main_Drugs_Name=record[i].Main_Drugs_Name
			data.Manufactor=record[i].Manufactor
			data.Money=record[i].Money
			data.Mypay=record[i].Mypay
			data.Proportion, _ =strconv.ParseFloat(record[i].Proportion,64)
			data.Specifications=record[i].Specifications

			app.Setup2.SaveMed(data)
	}
	_, err1 := model.DelAudit()
	if err1 != nil {
		fmt.Println(err)
		return
	}
}

//信息上传页面
func (app *Application) AddAuditShow(w http.ResponseWriter, r *http.Request)  {
	data := &struct {
		CurrentUser User
		Msg string
		Flag bool
	}{
		CurrentUser:cuser,
		Msg:"",
		Flag:false,
	}
	ShowView(w, r, "uploaddata.html", data)
}

