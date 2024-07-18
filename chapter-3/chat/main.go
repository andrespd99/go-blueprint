package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/andrespd99/chat-rooms/trace"
	"github.com/joho/godotenv"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
	"github.com/stretchr/objx"
)

// set the active Avatar implementation
var avatars Avatar = TryAvatars{
	UseFileSystemAvatar,
	UseAuthAvatar,
	UseGravatar,
}

// templ represents a single template
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServeHTTP handles the HTTP request.
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates",
			t.filename)))
	})
	data := map[string]interface{}{
		"Host": r.Host,
	}
	if authCookie, err := r.Cookie("auth"); err == nil {
		data["UserData"] = objx.MustFromBase64(authCookie.Value)
	}

	t.templ.Execute(w, data)

}

func main() {
	var addr = flag.String("addr", ":8080", "The addr of the  application.")
	flag.Parse() // parse the flags

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	gomniSecKey := os.Getenv("GOMNIAUTH_SEC_KEY")
	googleClientId := os.Getenv("GOOGLE_OAUTH2_CLIENT_ID")
	googleSecret := os.Getenv("GOOGLE_OAUTH2_SECRET")

	// setup gomniauth
	gomniauth.SetSecurityKey(gomniSecKey)
	gomniauth.WithProviders(
		google.New(
			googleClientId,
			googleSecret,
			"http://localhost:8080/auth/callback/google",
		),
	)

	r := newRoom()
	r.tracer = trace.New(os.Stdout)

	http.Handle("/", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)
	http.Handle("/room", r)
	http.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:   "auth",
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		})
		w.Header().Set("Location", "/chat")
		w.WriteHeader(http.StatusTemporaryRedirect)
	})
	http.Handle("/upload", &templateHandler{filename: "upload.html"})
	http.HandleFunc("/uploader", uploaderHandler)
	http.Handle("/avatars/",
		http.StripPrefix("/avatars/",
			http.FileServer(http.Dir("./avatars"))))

	// get the room going
	go r.run()
	// start the web server
	log.Println("Starting web server on", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
