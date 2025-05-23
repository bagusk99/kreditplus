package RenderTemplate

import "html/template"

func Render(page string) *template.Template {
	temp := template.Must(template.ParseFiles(
		"views/app.html",
		"views/layouts/sidebar.html",
		"views/layouts/topbar.html",
		page,
	))

	return temp
}