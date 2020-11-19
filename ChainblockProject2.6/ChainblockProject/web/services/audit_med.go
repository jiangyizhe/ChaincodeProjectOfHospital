package services

import "github.com/ChainblockProject/web/model"

func AddAuditMed(audit model.AuditMed)  (err error){

	if err = model.AddAuditMed(audit); err != nil {
		return err
	}
	return nil

}
func GetAuditMed() (audits []model.AuditMed, err error) {
	if audits, err = model.SelectAuditMed(); err != nil {
		return nil, err
	}
	return audits, nil
}
func IsData2(id int) bool{

	pat,_ :=model.GetAuditMedByID(id)
	if pat.Business_number==""{
		return true
	}else {
		return false
	}
}

func GetAuditMedByBus(bus string) (audits []model.AuditMed, err error) {
	if audits, err = model.GetAuditMedByBus(bus); err != nil {
		return nil, err
	}
	return audits, nil
}
