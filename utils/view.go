package utils

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type PageData struct {
	HeadMetaTag  string
	HeadAssetTag string
	FooterTag    string
	Data         any
}

var (
	viewTemplates *template.Template
	viewAssets    string
)

func InitializeViewTemplatesCache(filenames ...string) {
	templates := make([]string, 0)
	for _, v := range filenames {
		templates = append(templates, "views/"+v+".html")
	}

	viewTemplates = template.Must(template.New("html_template_cache").Funcs(tagFuncs()).ParseFiles(templates...))
}

func SetUpAssets(assets string) {
	viewAssets = assets
}

func ViewHtml(w http.ResponseWriter, data any, page string) {
	w.Header().Set("Content-Type", "text/html")
	err := viewTemplates.ExecuteTemplate(w, page+".html", getPageData(data))
	if err != nil {
		View404(w)
	}
}

func ViewNoCacheHTML(w http.ResponseWriter, data any, page string) {
	templ := template.Must(template.New("no_cache_html").Funcs(tagFuncs()).ParseFiles("views/" + page + ".html"))
	w.Header().Set("Content-Type", "text/html")
	err := templ.Execute(w, getPageData(data))
	if err != nil {
		View404(w)
	}
}

func View404(w http.ResponseWriter) {
	type Page404 struct {
		Title string
	}

	data404 := Page404{
		Title: os.Getenv("APP_NAME"),
	}

	w.WriteHeader(http.StatusNotFound)
	err := viewTemplates.ExecuteTemplate(w, "404.html", getPageData(data404))
	if err != nil {
		fmt.Fprint(w, "404 Page Not Found")
	}
}

func tagFuncs() template.FuncMap {
	return template.FuncMap{
		"headMetaTag": func(s PageData) template.HTML {
			return template.HTML(s.HeadMetaTag)
		},
		"headAssetTag": func(s PageData) template.HTML {
			return template.HTML(s.HeadAssetTag)
		},
		"footerTag": func(s PageData) template.HTML {
			return template.HTML(s.FooterTag)
		},
	}
}

func getPageData(data any) PageData {
	return PageData{
		HeadMetaTag: `<!DOCTYPE html><html lang="en"><head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">`,
		HeadAssetTag: viewAssets + `</head><body>`,
		FooterTag:    `</body></html>`,
		Data:         data,
	}
}
