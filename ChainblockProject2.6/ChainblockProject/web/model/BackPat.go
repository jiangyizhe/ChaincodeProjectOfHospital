package model

type BackPat struct {

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
	Status string
	Message string

}



func AddBackPat(patient BackPat)(err error){
	err= db.Create(&patient).Error

	return
}
func SelectBackPat()(patient []BackPat,err error){
	err = db.Debug().Model(&BackPat{}).Find(&patient).Error
	return
}

func GetBackPatByID(id string) (int64, error) {
	var f BackPat
	row := db.Model(&PatientPer{}).Where("patient_id = ?", id).Find(&f).RowsAffected

	return row, nil
}

	func SelectBackPatByBus(end string)(patient []BackPat,err error){
	err = db.Debug().Model(&BackPat{}).Where("business_number  = ?", end).Find(&patient).Error
	return
}