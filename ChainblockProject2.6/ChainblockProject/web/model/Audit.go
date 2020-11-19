package model

type Audit struct {

	Audit_id int
	Business_number string //业务号
	Cost_Classes string //费用类别
	Cost_Classes_Name string //费用类别名称
	Main_Drugs_Name string //中心药品名称
	Main_Drugs_Code string //中心药品编码
	Hospital_Drug_Name string //医院药品名称
	Drug_Dosage_Form string //剂型
	Manufactor string //厂家
	Specifications string //规格
	Number string //数量
	Baimonovalent string //单价
	Money string //金额
	Cost_Time string //金额时间
	Entry_Time string//录入时间
	Bookkeeper string //记账人
	Mypay string //自付金额
	Proportion string //比例
	Hospital_Drug_Code string//医院药品编码

}

func AddAudit(audit Audit)  (err error){
	err = db.Create(&audit).Error
	return
}

//查询所有
func SelectAudit() (audits []Audit, err error) {
	err = db.Debug().Model(&Audit{}).Find(&audits).Error
	return
}

//查询所有
func DelAudit() (audits []Audit, err error) {
	err = db.Debug().Model(&Audit{}).Delete(&audits).Error
	return
}