package consument

import (
	"fmt"
	"io"
	"kreditplus/entities"
	"kreditplus/models/consuments"
	RenderTemplate "kreditplus/services"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/google/uuid"
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
		err := r.ParseMultipartForm(10 << 20) // 10MB
		if err != nil {
			http.Error(w, "Gagal parsing form", http.StatusBadRequest)
			return
		}

		var consument entities.Consument

		consument.Nik = r.FormValue("nik")
		consument.FullName = r.FormValue("full_name")
		consument.LegalName = r.FormValue("legal_name")
		consument.PlaceOfBirth = r.FormValue("place_of_birth")
		consument.DateOfBirth, _ = time.Parse("2006-01-02", r.FormValue("date_of_birth"))
		consument.Salary, _ = strconv.Atoi(r.FormValue("salary"))

		consument.CreatedAt = time.Now()
		consument.UpdatedAt = time.Now()

		// === Upload KTP Photo ===
		ktpFile, ktpHeader, err := r.FormFile("ktp_photo")
		if err == nil {
			defer ktpFile.Close()

			originalFilename := ktpHeader.Filename
			ext := filepath.Ext(originalFilename)
			nameOnly := filepath.Base(originalFilename[:len(originalFilename)-len(ext)])
			newFileName := fmt.Sprintf("%s-%s%s", nameOnly, uuid.New().String(), ext)

			ktpPath := "public/uploads/" + newFileName
			dst, err := os.Create(ktpPath)
			if err != nil {
				log.Println("Gagal buat file:", err)
				return
			}
			defer dst.Close()
			io.Copy(dst, ktpFile)
			consument.KtpPhoto = newFileName
		}

		// === Upload Selfie Photo ===
		selfieFile, selfieHeader, err := r.FormFile("selfie_photo")
		if err == nil {
			defer selfieFile.Close()

			originalFilename := selfieHeader.Filename
			ext := filepath.Ext(originalFilename)
			nameOnly := filepath.Base(originalFilename[:len(originalFilename)-len(ext)])
			newFileName := fmt.Sprintf("%s-%s%s", nameOnly, uuid.New().String(), ext)

			selfiePath := "public/uploads/" + newFileName
			dst, err := os.Create(selfiePath)
			if err != nil {
				log.Println("Gagal buat file:", err)
				return
			}
			defer dst.Close()
			io.Copy(dst, selfieFile)
			consument.SelfiePhoto = newFileName
		}

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