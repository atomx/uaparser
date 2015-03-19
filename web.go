// +build ignore

// This is a simple web server that allowes you to parse and test user agent strings.
// Run using: go run web.go
// For a demo see: http://dubbelboer.com:8080
package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	. "."
)

type data struct {
	UserAgent              string
	DeviceType             string
	OperatingSystem        string
	OperatingSystemVersion string
	Browser                string
	BrowserVersion         string
}

var content = template.Must(template.New("content").Parse(`<doctype html>
<html>
<head>
<meta charset=utf-8>
<title>uaparser</title>
<style>
input {
width: 100%;
}
</style>
</head>
<body>
<form action="" method=get>
<p>
<input type=text name=useragent value="{{.UserAgent}}">
</p>
<p>
<input type=submit value=Lookup>
</p>
</form>
{{if .DeviceType}}
<p>
DeviceType: {{.DeviceType}}
</p>
{{end}}
{{if .OperatingSystem}}
<p>
OperatingSystem: {{.OperatingSystem}} {{.OperatingSystemVersion}}
</p>
{{end}}
{{if .Browser}}
<p>
Browser: {{.Browser}} {{.BrowserVersion}}
</p>
{{end}}
`))

func versionStr(version int) string {
	major, minor := Unversion(version)

	return strconv.FormatInt(int64(major), 10) + "." + strconv.FormatInt(int64(minor), 10)
}

func index(w http.ResponseWriter, r *http.Request) {
	d := data{
		UserAgent: r.Header.Get("User-Agent"),
	}

	if r.FormValue("useragent") != "" {
		d.UserAgent = r.FormValue("useragent")
	}

	d.DeviceType = DeviceTypes[DeviceType(d.UserAgent)]

	operatingSystemId, operatingSystemVersion := OperatingSystem(d.UserAgent)
	d.OperatingSystem = OperatingSystems[operatingSystemId]
	d.OperatingSystemVersion = versionStr(operatingSystemVersion)

	browserId, browserVersion := Browser(d.UserAgent)
	d.Browser = Browsers[browserId]
	d.BrowserVersion = versionStr(browserVersion)

	w.Header().Set("Content-Type", "text/html")

	if err := content.Execute(w, d); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", index)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Printf("[ERR] %v", err)
	}
}
