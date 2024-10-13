package configs

import (
	"os"
	"strings"

	ut "github.com/naywin-programmer/RSS_Aggregator/utils"
)

func useDevTools() (devAssets string) {
	if appEnv := strings.ToLower(os.Getenv("APP_ENV")); appEnv == "dev" {
		devAssets = `<script src="/static/devtools/autorefreshbrowser.js" async></script>`
	}
	return devAssets
}

func SetUpView() {
	ut.SetUpAssets(`

		<link rel="preload" href="/static/build/style.css" as="style" onload="this.onload=null;this.rel='stylesheet'">
		<noscript><link rel="stylesheet" href="/static/build/style.css" type="text/css"></noscript>

		<script src="/static/js/htmx_v2_0_3.min.js" async></script>

	` + useDevTools())

	// use ViewNoCacheHTML instead if u don't want to add new template filenames in here
	ut.InitializeViewTemplatesCache(
		"404",
		"index",
		"profile/about_us",
	)
}
