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
	}

	temp := RenderTemplate.Render("views/pages/consument/index.html")

	temp.Execute(w,data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp := RenderTemplate.Render("views/pages/consument/create.html")

		temp.Execute(w, nil)
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

		http.Redirect(w, r, "/categories", http.StatusSeeOther)
	}
}