package services

import "github.com/ChainblockProject/web/model"

func AddPatientPer(audit model.PatientPer)  (err error){

	if err = model.AddPatientPer(audit); err != nil {
		return err
	}
	return nil

}
func GetPatientPer() (audits []model.PatientPer, err error) {
	if audits, err = model.SelectPatientPer(); err != nil {
		return nil, err
	}
	return audits, nil
}

func GetPatientPerByEnd(end string) (audits []model.PatientPer, err error) {
	if audits, err = model.SelectPatientPerByEnd(end); err != nil {
		return nil, err
	}
	return audits, nil
}

func IsData(id int) bool{

	pat,_ :=model.GetPatientPerByID(id)
	if pat.Name==""{
		return true
	}else {
		return false
	}
}