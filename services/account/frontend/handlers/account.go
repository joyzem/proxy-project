package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/joyzem/proxy-project/services/account/domain"
	"github.com/joyzem/proxy-project/services/account/dto"
	"github.com/joyzem/proxy-project/services/account/frontend/utils"
	"github.com/joyzem/proxy-project/services/base"
	"github.com/levigross/grequests"
)

func AccountsHandler(w http.ResponseWriter, r *http.Request) {
	accountsResp, err := utils.GetAccountsFromBackend()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	accountsPage, err := template.ParseFiles("../static/html/accounts.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if accountsResp.Err != "" {
		http.Error(w, accountsResp.Err, http.StatusInternalServerError)
	}
	accountsPage.Execute(w, accountsResp.Accounts)
}

func DeleteAccountHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		http.Redirect(w, r, "/account/accounts", http.StatusBadRequest)
		return
	}
	body := dto.DeleteAccountRequest{Id: id}
	accountsUrl := fmt.Sprintf("%s/accounts", utils.GetBackendAddress())
	resp, err := grequests.Delete(accountsUrl, &grequests.RequestOptions{
		JSON: body,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var deleteResponse dto.DeleteAccountResponse
	err = json.Unmarshal(resp.Bytes(), &deleteResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/account/accounts", http.StatusSeeOther)
}

func CreateAccountGetHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../static/html/create-account.html")
}

func CreateAccountPostHandler(w http.ResponseWriter, r *http.Request) {
	bankName := r.FormValue("bank_name")
	bankIdentityNumber := r.FormValue("bin")
	if len(bankName) == 0 || len(bankIdentityNumber) != 9 {
		http.Error(w, errors.New(base.FIELDS_VALIDATION_ERROR).Error(), http.StatusUnprocessableEntity)
		return
	}
	request := dto.CreateAccountRequest{
		BankName:           bankName,
		BankIdentityNumber: bankIdentityNumber,
	}
	accountsUrl := fmt.Sprintf("%s/accounts", utils.GetBackendAddress())
	resp, err := grequests.Post(accountsUrl, &grequests.RequestOptions{
		JSON: request,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var data dto.CreateAccountResponse
	err = json.Unmarshal(resp.Bytes(), &data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if data.Err == "" {
		http.Redirect(w, r, "/account/accounts", http.StatusSeeOther)
	} else {
		http.Error(w, data.Err, http.StatusInternalServerError)
	}
}

func UpdateAccountGetHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Redirect(w, r, "/account/accounts", http.StatusBadRequest)
		return
	}
	accountsResp, err := utils.GetAccountsFromBackend()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if accountsResp.Err != "" {
		http.Error(w, accountsResp.Err, http.StatusInternalServerError)
		return
	}
	accounts := accountsResp.Accounts
	var requestedAccount *domain.Account
	for i, account := range accounts {
		if accounts[i].Id == id {
			requestedAccount = &account
		}
	}
	if requestedAccount == nil {
		http.Error(w, "the account does not exist", http.StatusBadRequest)
		return
	}
	updateAccountPage, err := template.ParseFiles("../static/html/update-account.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	updateAccountPage.Execute(w, requestedAccount)

}

func UpdateAccountPostHandler(w http.ResponseWriter, r *http.Request) {
	accountId, _ := strconv.Atoi(r.FormValue("id"))
	bankName := r.FormValue("bank_name")
	bankIdentityCode := r.FormValue("bin")
	account := domain.Account{
		Id:                 accountId,
		BankName:           bankName,
		BankIdentityNumber: bankIdentityCode,
	}
	accountsUrl := fmt.Sprintf("%s/accounts", utils.GetBackendAddress())
	request := dto.UpdateAccountRequest{
		Account: account,
	}
	options := &grequests.RequestOptions{
		JSON: request,
	}
	response, err := grequests.Put(accountsUrl, options)
	var updateResponse dto.UpdateAccountResponse
	json.Unmarshal(response.Bytes(), &updateResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if updateResponse.Err != "" {
		http.Error(w, updateResponse.Err, http.StatusInternalServerError)
	}
	http.Redirect(w, r, "/account/accounts", http.StatusSeeOther)
}
