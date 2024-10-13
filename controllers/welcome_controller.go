package controllers

import (
	"net/http"
	"os"

	ut "github.com/naywin-programmer/RSS_Aggregator/utils"
)

type welcomeController struct {
}

var WelcomeController = welcomeController{}

func (c welcomeController) Welcome(w http.ResponseWriter, r *http.Request) {
	type WelcomePage struct {
		Title       string
		Description string
		Name        string
	}

	data := WelcomePage{
		Title:       os.Getenv("APP_NAME"),
		Description: os.Getenv("APP_NAME"),
		Name:        "Perfect Hotreload is here",
	}

	ut.ViewHtml(w, data, "index")
}

func (c welcomeController) AboutUs(w http.ResponseWriter, r *http.Request) {
	type AboutUsPage struct {
		Title       string
		Description string
		Injection   string
	}

	data := AboutUsPage{
		Title:       "About Us - " + os.Getenv("APP_NAME"),
		Description: "About Us - This is just for fun",
		Injection: `<p>
        <script>alert("XSS Injection")</script>
    </p>`,
	}

	ut.ViewHtml(w, data, "about_us")
}
