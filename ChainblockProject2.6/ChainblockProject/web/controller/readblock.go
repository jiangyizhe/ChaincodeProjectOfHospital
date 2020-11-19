package controller

import "net/http"

func (app *Application) ReadBlockShow(w http.ResponseWriter, r *http.Request)  {

	app.AddAuditMed()
	app.AddPatientPer()
	data := &struct {
		CurrentUser User
		Msg string
		Flag bool
	}{
		CurrentUser:cuser,
		Msg:"OK",
		Flag:false,
	}
	ShowView(w, r, "help.html", data)
}
func (app *Application) AddBlock1Show(w http.ResponseWriter, r *http.Request)  {

	app.AddFabric()

	data := &struct {
		CurrentUser User
		Msg string
		Flag bool
	}{
		CurrentUser:cuser,
		Msg:"OK",
		Flag:false,
	}
	ShowView(w, r, "help.html", data)
}

func (app *Application) AddBlock2Show(w http.ResponseWriter, r *http.Request)  {

	app.AddFabric2()

	data := &struct {
		CurrentUser User
		Msg string
		Flag bool
	}{
		CurrentUser:cuser,
		Msg:"OK",
		Flag:false,
	}
	ShowView(w, r, "help.html", data)
}
