package model

type AuditMed struct {

	Audit_id string
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
	Proportion float64 //比例
	Hospital_Drug_Code string//医院药品编码
	IsRight string //是否合理
	Reason string//原因
}

func AddAuditMed(audit AuditMed)  (err error){
	err = db.Create(&audit).Error
	return
}

//查询所有
func SelectAuditMed() (audits []AuditMed, err error) {
	err = db.Debug().Model(&AuditMed{}).Find(&audits).Error
	return
}

func GetAuditMedByID(id int) (AuditMed, error) {
	var f AuditMed
	err := db.Model(&AuditMed{}).Where("audit_id = ?", id).Find(&f).Error
	if err != nil {
		return AuditMed{}, err
	}
	return f, nil
}

func GetAuditMedByBus(bus string) ([]AuditMed, error) {
	var f []AuditMed
	err := db.Model(&AuditMed{}).Where("business_number = ?", bus).Find(&f).Error
	if err != nil {
		return nil, err
	}
	return f, nil
}

func SelectAuditMedBack() (audits []AuditMed, err error) {

	err = db.Debug().Model(&AuditMed{}).Where("is_right = ?","否" ).Find(&audits).Error
	return
}

//查询所有
func DelAuditMed() (audits []AuditMed, err error) {
	err = db.Debug().Model(&AuditMed{}).Delete(&audits).Error
	return
}

func SavaAuditMed(audit AuditMed)  (err error){
	err = db.Model(&AuditMed{}).Where("audit_id = ?", audit.Audit_id).Update(&audit).Error
	return
}