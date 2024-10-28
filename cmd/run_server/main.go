package main

import (
	"flag"
	"fmt"
	"html/template"
	"strconv"

	"github.com/ben-smyth/coolbox/web"
)

func main() {
	fmt.Println("Running server..")

	dev := flag.Bool("dev", true, "enable dev mode")
	localAssetPath := flag.String("localAssetPath", "web/static/", "location of static files")
	websiteAssetPath := flag.String("websiteAssetPath", "/static/", "web location of the static files")
	websiteUrl := flag.String("url", "localhost", "URL or IP of the website")
	websiteScheme := flag.String("scheme", "http", "scheme, http or https")
	port := flag.Int("port", 8080, "set port number")

	var app web.WebApp

	app.LocalAssetPath = *localAssetPath
	app.Dev = *dev
	app.Port = *port
	app.WebsiteUrl = fmt.Sprintf("%s:%s", *websiteUrl, strconv.Itoa(*port))
	app.WebsiteAssetUrl = template.URL(fmt.Sprint(*websiteScheme, "://", app.WebsiteUrl, *websiteAssetPath))

	fmt.Printf("Dev: \t\t\t\t\t %v\n", *dev)
	fmt.Printf("Local Static Files Location: \t\t %v\n", *localAssetPath)
	fmt.Printf("URL: \t\t\t\t\t %v\n", app.WebsiteUrl)
	fmt.Printf("Relative Website Asset Location:\t %v\n", app.WebsiteAssetUrl)
	fmt.Printf("Port: \t\t\t\t\t %v\n", *port)

	err := web.ServeWebsite(app)
	if err != nil {
		panic(err)
	}
}
