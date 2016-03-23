// +build ignore

// This is a simple web server that allowes you to parse and test user agent strings.
// Run using: go run web.go
// For a demo see: http://dubbelboer.com:8080
package main

import (
	"html/template"
	"log"
	"net/http"

	. "."
)

type data struct {
	UserAgent                   string
	DeviceType                  string
	OperatingSystem             string
	OperatingSystemVersionMajor int
	OperatingSystemVersionMinor int
	Browser                     string
	BrowserVersionMajor         int
	BrowserVersionMinor         int
}

var content = template.Must(template.New("content").Parse(`<!doctype html>
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
OperatingSystem: {{.OperatingSystem}} {{.OperatingSystemVersionMajor}}.{{.OperatingSystemVersionMinor}}
</p>
{{end}}
{{if .Browser}}
<p>
Browser: {{.Browser}} {{.BrowserVersionMajor}}.{{.BrowserVersionMinor}}
</p>
{{end}}
`))

func index(w http.ResponseWriter, r *http.Request) {
	d := data{
		UserAgent: r.Header.Get("User-Agent"),
	}

	if r.FormValue("useragent") != "" {
		d.UserAgent = r.FormValue("useragent")
	}

	d.DeviceType = DeviceTypes[DeviceType(d.UserAgent)]

	operatingSystemId, operatingSystemVersionMajor, operatingSystemVersionMinor := OperatingSystem(d.UserAgent)
	d.OperatingSystem = OperatingSystems[operatingSystemId]
	d.OperatingSystemVersionMajor = operatingSystemVersionMajor
	d.OperatingSystemVersionMinor = operatingSystemVersionMinor

	browserId, browserVersionMajor, browserVersionMinor := Browser(d.UserAgent)
	d.Browser = Browsers[browserId]
	d.BrowserVersionMajor = browserVersionMajor
	d.BrowserVersionMinor = browserVersionMinor

	log.Printf("%22s | %10s | %14s %6d | %14s %6d | %s\n", r.RemoteAddr, d.DeviceType, d.OperatingSystem, d.OperatingSystemVersionMajor, d.Browser, d.BrowserVersionMajor, d.UserAgent)

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
