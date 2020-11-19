package services

import "github.com/ChainblockProject/web/model"

func AddAudit(audit model.Audit)  (err error){

	if err = model.AddAudit(audit); err != nil {
		return err
	}
	return nil

}
func GetAudit() (audits []model.Audit, err error) {
	if audits, err = model.SelectAudit(); err != nil {
		return nil, err
	}
	return audits, nil
}


