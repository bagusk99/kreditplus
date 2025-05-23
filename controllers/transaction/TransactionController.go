package transaction

import (
	"kreditplus/entities"
	"kreditplus/models/transactions"
	RenderTemplate "kreditplus/services"
	"net/http"
	"strconv"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	transactions := transactions.GetAll()

	data := map[string]any {
		"transactions": transactions,
		"activePage": "transaction",
	}

	temp := RenderTemplate.Render("views/pages/transaction/index.html")

	temp.Execute(w,data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp := RenderTemplate.Render("views/pages/transaction/create.html")

		data := map[string]any {
			"activePage": "transaction",
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var transaction entities.Transaction

		transaction.ConsumentId, _ = strconv.Atoi(r.FormValue("consument_id"))
		transaction.ContractNumber = r.FormValue("contract_number")
		transaction.Otr, _ = strconv.Atoi(r.FormValue("otr"))
		transaction.AdminFee, _ = strconv.Atoi(r.FormValue("admin_fee"))
		transaction.Installment, _ = strconv.Atoi(r.FormValue("admin_fee"))
		transaction.Interest, _ = strconv.Atoi(r.FormValue("interest"))
		transaction.AssetName = r.FormValue("asset_name")

		transaction.CreatedAt = time.Now()
		transaction.UpdatedAt = time.Now()

		transactions.Create(transaction)

		http.Redirect(w, r, "/transactions", http.StatusSeeOther)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp := RenderTemplate.Render("views/pages/transaction/edit.html")

		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)

		if err != nil {
			panic(err)
		}

		transaction := transactions.Detail(id)
		data := map[string]any {
			"transaction": transaction,
			"activePage": "transaction",
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var transaction entities.Transaction

		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)

		if err != nil {
			panic(err)
		}

		transaction.ConsumentId, _ = strconv.Atoi(r.FormValue("consument_id"))
		transaction.ContractNumber = r.FormValue("contract_number")
		transaction.Otr, _ = strconv.Atoi(r.FormValue("otr"))
		transaction.AdminFee, _ = strconv.Atoi(r.FormValue("admin_fee"))
		transaction.Installment, _ = strconv.Atoi(r.FormValue("admin_fee"))
		transaction.Interest, _ = strconv.Atoi(r.FormValue("interest"))
		transaction.AssetName = r.FormValue("asset_name")
		transaction.UpdatedAt = time.Now()

		transactions.Update(id, transaction)

		http.Redirect(w, r, "/transactions", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.FormValue("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		panic(err)
	}

	transactions.Delete(id)

	http.Redirect(w, r, "/transactions", http.StatusSeeOther)
}