package web

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/ben-smyth/coolbox/pkg/core_tools/json2yaml"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type WebApp struct {
	Dev             bool
	Port            int
	LocalAssetPath  string
	WebsiteUrl      string
	WebsiteAssetUrl template.URL
}

func ServeWebsite(app WebApp) error {
	r := mux.NewRouter()

	fs := http.FileServer(http.Dir(app.LocalAssetPath))

	// hot reloading for dev environment
	if app.Dev {
		r.HandleFunc("/dev", handleWebSocket)
	}

	r.PathPrefix("/static/").Handler(cacheControlMiddleware(http.StripPrefix("/static/", fs)))
	r.HandleFunc("/", app.IndexHandler)

	r.HandleFunc("/plugin/json2yaml", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			r.ParseForm()
			jsonData := r.FormValue("json")
			yaml, err := json2yaml.ConvertJson2Yaml(jsonData)
			if err != nil {
				w.Write([]byte(err.Error()))
			}
			w.Write([]byte(yaml))
		}
	})

	return http.ListenAndServe(app.WebsiteUrl, r)
}

func (a *WebApp) IndexHandler(w http.ResponseWriter, r *http.Request) {
	Templ.ExecuteTemplate(w, "home", a)
}

var Templ = func() *template.Template {
	t := template.New("")
	err := filepath.Walk("web/templates", func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".html") {
			fmt.Println(path)
			_, err = t.ParseFiles(path)
			if err != nil {
				fmt.Println(err)
			}
		}
		return err
	})

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	return t
}()

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Upgrade failed:", err)
		return
	}
	defer conn.Close()

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Read failed:", err)
			break
		}

		fmt.Printf("Received: %s\n", message)

		if err := conn.WriteMessage(messageType, message); err != nil {
			fmt.Println("Write failed:", err)
			break
		}
	}
}

// cacheControlMiddleware - prevent caching static files for too long
func cacheControlMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "max-age=7200, must-revalidate")

		next.ServeHTTP(w, r)
	})
}
