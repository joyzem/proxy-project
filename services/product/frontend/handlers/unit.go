package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	sharedUtils "github.com/joyzem/proxy-project/services/utils"

	"github.com/joyzem/proxy-project/services/product/domain"
	"github.com/joyzem/proxy-project/services/product/frontend/transport"
	"github.com/joyzem/proxy-project/services/product/frontend/utils"
	"github.com/levigross/grequests"
)

func UnitsHandler(w http.ResponseWriter, r *http.Request) {
	units, err := utils.GetUnitsFromBackend()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	unitsPage, _ := template.ParseFiles("../static/html/units.html")
	unitsPage.Execute(w, units)
}

func DeleteUnitHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		http.Redirect(w, r, "/product/units/", http.StatusBadRequest)
		return
	}
	body := transport.DeleteUnitRequest{Id: id}
	options := sharedUtils.CreateJsonRequestOption(body)
	unitsUrl := fmt.Sprintf("%s/units", utils.GetBackendAddress())
	resp, err := grequests.Delete(unitsUrl, options)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var deleteResponse transport.DeleteUnitResponse
	err = json.Unmarshal(resp.Bytes(), &deleteResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/product/units/", http.StatusSeeOther)
}

func CreateUnitGetHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../static/html/create-unit.html")
}

func CreateUnitPostHandler(w http.ResponseWriter, r *http.Request) {
	unitName := r.FormValue("name")
	if len(unitName) == 0 {
		http.Error(w, errors.New(sharedUtils.FIELDS_VALIDATION_ERROR).Error(), http.StatusUnprocessableEntity)
	}
	request := transport.CreateUnitRequest{
		Unit: unitName,
	}
	unitsUrl := fmt.Sprintf("%s/units", utils.GetBackendAddress())
	options := sharedUtils.CreateJsonRequestOption(request)
	resp, err := grequests.Post(unitsUrl, options)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var data transport.CreateUnitResponse
	err = json.Unmarshal(resp.Bytes(), &data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if data.Err == nil {
		http.Redirect(w, r, "/product/units/", http.StatusSeeOther)
	} else {
		http.Error(w, data.Err.Error(), http.StatusInternalServerError)
	}
}

func UpdateUnitGetHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Redirect(w, r, "/product/units/", http.StatusBadRequest)
		return
	}
	units, err := utils.GetUnitsFromBackend()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var requestedUnit *domain.Unit
	for i, unit := range units {
		if units[i].Id == id {
			requestedUnit = &unit
		}
	}
	if requestedUnit == nil {
		http.Error(w, "the unit does not exist", http.StatusBadRequest)
		return
	}
	updateUnitPage, err := template.ParseFiles("../static/html/update-unit.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	updateUnitPage.Execute(w, requestedUnit)
}

func UpdateUnitPostHandler(w http.ResponseWriter, r *http.Request) {
	unitId, err := strconv.Atoi(r.FormValue("id"))
	unitName := r.FormValue("name")
	if len(unitName) == 0 || err != nil {
		http.Error(w, errors.New(sharedUtils.FIELDS_VALIDATION_ERROR).Error(), http.StatusUnprocessableEntity)
		return
	}
	unit := &domain.Unit{
		Id:   unitId,
		Name: unitName,
	}

	unitsUrl := fmt.Sprintf("%s/units", utils.GetBackendAddress())
	request := transport.UpdateUnitRequest{
		Unit: unit,
	}
	options := &grequests.RequestOptions{
		JSON: request,
	}
	_, err = grequests.Put(unitsUrl, options)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/product/units/", http.StatusSeeOther)
}
