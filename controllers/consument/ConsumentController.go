package consument

import (
	"kreditplus/entities"
	"kreditplus/models/consuments"
	RenderTemplate "kreditplus/services"
	"net/http"
	"strconv"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	consuments := consuments.GetAll()

	data := map[string]any {
		"consuments": consuments,
		"activePage": "consument",
	}

	temp := RenderTemplate.Render("views/pages/consument/index.html")

	temp.Execute(w,data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp := RenderTemplate.Render("views/pages/consument/create.html")

		data := map[string]any {
			"activePage": "consument",
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var consument entities.Consument

		consument.Nik = r.FormValue("nik")
		consument.FullName = r.FormValue("full_name")
		consument.LegalName = r.FormValue("legal_name")
		consument.PlaceOfBirth = r.FormValue("place_of_birth")
		consument.DateOfBirth, _ = time.Parse("2006-01-02", r.FormValue("date_of_birth"))
		consument.Salary, _ = strconv.Atoi(r.FormValue("salary"))
		consument.KtpPhoto = r.FormValue("ktp_photo")
		consument.SelfiePhoto = r.FormValue("selfie_photo")

		consument.CreatedAt = time.Now()
		consument.UpdatedAt = time.Now()

		consuments.Create(consument)

		http.Redirect(w, r, "/consuments", http.StatusSeeOther)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp := RenderTemplate.Render("views/pages/consument/edit.html")

		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)

		if err != nil {
			panic(err)
		}

		consument := consuments.Detail(id)
		data := map[string]any {
			"consument": consument,
			"dobformatted": consument.DateOfBirth.Format("2006-01-02"),
			"activePage": "consument",
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var consument entities.Consument

		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)

		if err != nil {
			panic(err)
		}

		consument.Nik = r.FormValue("nik")
		consument.FullName = r.FormValue("full_name")
		consument.LegalName = r.FormValue("legal_name")
		consument.PlaceOfBirth = r.FormValue("place_of_birth")
		consument.DateOfBirth, _ = time.Parse("2006-01-02", r.FormValue("date_of_birth"))
		consument.Salary, _ = strconv.Atoi(r.FormValue("salary"))
		consument.KtpPhoto = r.FormValue("ktp_photo")
		consument.SelfiePhoto = r.FormValue("selfie_photo")
		consument.UpdatedAt = time.Now()

		consuments.Update(id, consument)

		http.Redirect(w, r, "/consuments", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.FormValue("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		panic(err)
	}

	consuments.Delete(id)

	http.Redirect(w, r, "/consuments", http.StatusSeeOther)
}