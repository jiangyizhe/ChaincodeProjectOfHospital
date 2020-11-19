package model

type PatientPer struct {

	ObjectType string `json:"patType"`
	Patient_Id string
	Business_number string      //业务号
	Name                 string //姓名
	Sex                  string //性别
	Hosnum               string //hosnum
	HospitalizedType     string //住院类型
	HospitalizedDay      string //住院日期
	DiseasesName         string //病症名称
	Note                 string //备注
	Department           string //治疗科室
	BedNum               string //床号
	EntityID             string //身份证号码
	Reimbursement_Date   string //报销日期
	Hospital_Location    string //医院地点
	Entry_Time           string //录入时间
	ReAmount             string //原有金额
	Reimbursement_Amount string //报销金额
	Mypay string                //剩余金额
}



func AddPatientPer(patient PatientPer)(err error){
	err= db.Create(&patient).Error

	return
}
func SelectPatientPer()(patient []PatientPer,err error){
	err = db.Debug().Model(&PatientPer{}).Find(&patient).Error
	return
}
func SelectPatientPerByEnd(end string)(patient []PatientPer,err error){
	err = db.Debug().Model(&PatientPer{}).Where("entity_id  = ?", end).Find(&patient).Error
	return
}
func SelectPatientPerByBus(end string)(patient []PatientPer,err error){
	err = db.Debug().Model(&PatientPer{}).Where("business_number  = ?", end).Find(&patient).Error
	return
}
func SelectPatientPerByBusOne(end string)(patient PatientPer,err error){
	err = db.Debug().Model(&PatientPer{}).Where("business_number  = ?", end).Find(&patient).Error
	return
}

func GetPatientPerByID(id int) (PatientPer, error) {
	var f PatientPer
	err := db.Model(&PatientPer{}).Where("patient_id = ?", id).Find(&f).Error
	if err != nil {
		return PatientPer{}, err
	}
	return f, nil
}
//查询所有
func DelPatientPer() (audits []PatientPer, err error) {
	err = db.Debug().Model(&PatientPer{}).Delete(&audits).Error
	return
}