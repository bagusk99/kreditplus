package DashboardController

import (
	RenderTemplate "kreditplus/services"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	temp := RenderTemplate.Render("views/pages/dashboard.html")

	temp.Execute(w, nil);
}