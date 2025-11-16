package templates

import (
	"html/template"
	"time"

	"github.com/gin-contrib/multitemplate"
)

func formatDate(layout string, dateString string) (string, error) {
	layoutShort := "2006-01-02T15:04"
	layoutLong := "2006-01-02T15:04:05"

	t, err := time.Parse(layoutLong, dateString)
	if err == nil {
		return t.Format(layout), nil
	}

	t, err = time.Parse(layoutShort, dateString)
	if err == nil {
		return t.Format(layout), nil
	}

	return "", err
}

func CreateMyRender() multitemplate.Renderer {
	router := multitemplate.NewRenderer()
	funcMap := template.FuncMap{
		"formatDate": formatDate,
	}
	router.AddFromFiles("login", "templates/pages/login.html")

	router.AddFromFilesFuncs("home", funcMap,
		"templates/layouts/base.html",
		"templates/partials/header.html",
		"templates/partials/menu.html",
		"templates/pages/home.html",
		"templates/partials/player.html",
		"templates/partials/ad.html",
		"templates/partials/footer.html",
	)

	router.AddFromFilesFuncs("category-frontend", funcMap,
		"templates/layouts/base.html",
		"templates/partials/header.html",
		"templates/partials/menu.html",
		"templates/pages/category-frontend.html",
		"templates/partials/ad.html",
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

	router.AddFromFilesFuncs("page", funcMap,
		"templates/layouts/baseAdmin.html",
		"templates/pages/admin.html",
		"templates/pages/page.html",
		"templates/partials/headerAdmin.html",
		"templates/partials/footer.html",
		"templates/partials/menuAdmin.html",
		"templates/partials/items.html",
	)

	router.AddFromFilesFuncs("page-edit", funcMap,
		"templates/layouts/baseAdmin.html",
		"templates/pages/admin.html",
		"templates/pages/page-edit.html",
		"templates/partials/headerAdmin.html",
		"templates/partials/footer.html",
		"templates/partials/menuAdmin.html",
		"templates/partials/items.html",
	)

	router.AddFromFilesFuncs("upload", funcMap,
		"templates/layouts/baseAdmin.html",
		"templates/pages/admin.html",
		"templates/pages/upload.html",
		"templates/partials/headerAdmin.html",
		"templates/partials/footer.html",
		"templates/partials/menuAdmin.html",
		"templates/partials/items.html",
	)

	router.AddFromFilesFuncs("user", funcMap,
		"templates/layouts/baseAdmin.html",
		"templates/pages/admin.html",
		"templates/pages/user.html",
		"templates/partials/headerAdmin.html",
		"templates/partials/footer.html",
		"templates/partials/menuAdmin.html",
		"templates/partials/items.html",
	)

	router.AddFromFilesFuncs("user-edit", funcMap,
		"templates/layouts/baseAdmin.html",
		"templates/pages/admin.html",
		"templates/pages/user-edit.html",
		"templates/partials/headerAdmin.html",
		"templates/partials/footer.html",
		"templates/partials/menuAdmin.html",
		"templates/partials/items.html",
	)

	router.AddFromFilesFuncs("setting", funcMap,
		"templates/layouts/baseAdmin.html",
		"templates/pages/admin.html",
		"templates/pages/setting.html",
		"templates/partials/headerAdmin.html",
		"templates/partials/footer.html",
		"templates/partials/menuAdmin.html",
		"templates/partials/items.html",
	)

	router.AddFromFilesFuncs("setting-edit", funcMap,
		"templates/layouts/baseAdmin.html",
		"templates/pages/admin.html",
		"templates/pages/setting-edit.html",
		"templates/partials/headerAdmin.html",
		"templates/partials/footer.html",
		"templates/partials/menuAdmin.html",
		"templates/partials/items.html",
	)

	router.AddFromFilesFuncs("search", funcMap,
		"templates/layouts/baseAdmin.html",
		"templates/pages/admin.html",
		"templates/pages/search.html",
		"templates/partials/headerAdmin.html",
		"templates/partials/footer.html",
		"templates/partials/menuAdmin.html",
		"templates/partials/items.html",
	)

	return router
}
