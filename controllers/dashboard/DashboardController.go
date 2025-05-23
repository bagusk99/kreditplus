package dashboard

import (
	RenderTemplate "kreditplus/services"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	temp := RenderTemplate.Render("views/pages/dashboard.html")

	data := map[string]any {
		"activePage": "dashboard",
	}

	temp.Execute(w, data);
}