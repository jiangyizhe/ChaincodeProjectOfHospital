package services

import "github.com/ChainblockProject/web/model"

func AddPatientInfo(info model.PatientInfo)  (err error){

	if err = model.AddPatient(info);err!=nil{
		return err
	}
	return nil

}

func GetPatientInfo()(info []model.PatientInfo,err error){
	if info, err = model.SelectPatient(); err != nil {
		return nil, err
	}
	return info, nil
}