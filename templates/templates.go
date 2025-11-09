package templates

import (
	"html/template"
	"time"

	"github.com/gin-contrib/multitemplate"
)

func formatDate(layout string, dateString string) (string, error) {
	t, err := time.Parse("2006-01-02T15:04:05", dateString)
	if err != nil {
		return "", err
	}
	return t.Format(layout), nil
}

func CreateMyRender() multitemplate.Renderer {
	router := multitemplate.NewRenderer()
	funcMap := template.FuncMap{
		"formatDate": formatDate,
	}
	router.AddFromFiles("login", "templates/pages/login.html")

	router.AddFromFilesFuncs("home", funcMap,
		"templates/layouts/base.html",
		"templates/pages/home.html",
		"templates/partials/footer.html",
	)

	router.AddFromFilesFuncs("admin", funcMap,
		"templates/layouts/baseAdmin.html",
		"templates/pages/admin.html",
		"templates/partials/headerAdmin.html",
		"templates/partials/footer.html",
		"templates/partials/menuAdmin.html",
		"templates/partials/items.html",
	)

	router.AddFromFilesFuncs("admin-edit", funcMap,
		"templates/layouts/baseAdmin.html",
		"templates/pages/admin-edit.html",
		"templates/partials/headerAdmin.html",
		"templates/partials/footer.html",
		"templates/partials/menuAdmin.html",
		"templates/partials/items.html",
	)

	router.AddFromFilesFuncs("category", funcMap,
		"templates/layouts/baseAdmin.html",
		"templates/pages/admin.html",
		"templates/pages/category.html",
		"templates/partials/headerAdmin.html",
		"templates/partials/footer.html",
		"templates/partials/menuAdmin.html",
		"templates/partials/items.html",
	)

	router.AddFromFilesFuncs("category-edit", funcMap,
		"templates/layouts/baseAdmin.html",
		"templates/pages/admin.html",
		"templates/pages/category-edit.html",
		"templates/partials/headerAdmin.html",
		"templates/partials/footer.html",
		"templates/partials/menuAdmin.html",
		"templates/partials/items.html",
	)

	return router
}
