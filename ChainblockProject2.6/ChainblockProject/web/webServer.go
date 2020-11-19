package web

import (
	"fmt"
	"github.com/ChainblockProject/web/controller"
	"net/http"
)


// 启动Web服务并指定路由信息
func WebStart(app controller.Application)  {

	fs:= http.FileServer(http.Dir("web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// 指定路由信息(匹配请求)
	http.HandleFunc("/", app.LoginView)
	http.HandleFunc("/login", app.Login)
	http.HandleFunc("/loginout", app.LoginOut)

	http.HandleFunc("/index", app.Index)
	http.HandleFunc("/help", app.Help)

	http.HandleFunc("/addEduInfo", app.AddEduShow)	// 显示添加信息页面
	http.HandleFunc("/addEdu", app.AddEdu)	// 提交信息请求

	http.HandleFunc("/queryPage", app.QueryPage)	// 转至根据证书编号与姓名查询信息页面
	http.HandleFunc("/query", app.FindCertByNoAndName)	// 根据证书编号与姓名查询信息

	http.HandleFunc("/queryPage2", app.QueryPage2)	// 转至根据身份证号码查询信息页面
	http.HandleFunc("/query2", app.FindByID)	// 根据身份证号码查询信息


	http.HandleFunc("/modifyPage", app.ModifyShow)	// 修改信息页面
	http.HandleFunc("/modify", app.Modify)	//  修改信息

	http.HandleFunc("/upload", app.UploadFile)

	http.HandleFunc("/uploadfile",app.AddAudit)//添加信息
	http.HandleFunc("/uploadpatfile",app.AddPatientInfo)//添加信息

	http.HandleFunc("/uploaddata",app.AddAuditShow)//

	http.HandleFunc("/uploadpatdata",app.AddPatientShow)//添加信息页面显示

	http.HandleFunc("/readblock",app.ReadBlockShow)//链上信息一键更新
	http.HandleFunc("/addblock1",app.AddBlock1Show)//链上信息一键更新
	http.HandleFunc("/addblock2",app.AddBlock2Show)//链上信息一键更新

	http.HandleFunc("/meddata",app.GetAuditMed)//获取信息页面显示
	http.HandleFunc("/meddatabybus",app.GetAuditMedByBus)//获取信息页面显示
	http.HandleFunc("/isright",app.IsRight)//获取信息页面显示
	http.HandleFunc("/patdata",app.GetPatientPer)//获取信息页面显示
	http.HandleFunc("/patresult",app.GetPatientPerResult)//获取信息页面显示
	http.HandleFunc("/patresultend",app.GetPatientPerResultByEnd)//获取信息页面显示
	http.HandleFunc("/querypat",app.GetPatientPerByEnd)//获取信息页面显示
	http.HandleFunc("/addback",app.AddBackPatShow)//获取信息页面显示
	http.HandleFunc("/backshow",app.GetBackPat)//获取信息页面显示


	fmt.Println("启动Web服务, 监听端口号为: 9002")
	err := http.ListenAndServe(":9002", nil)
	if err != nil {
		fmt.Printf("Web服务启动失败: %v", err)
	}

}



