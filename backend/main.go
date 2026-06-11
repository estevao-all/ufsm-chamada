package main

import (
	"log"
	"net/http"
	"os"

	"backend/routes"
	"backend/utils"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/user/login", routes.HandleLogin)
	mux.HandleFunc("/api/user/info", utils.WithCache(routes.HandleUserInfo))
	mux.HandleFunc("/api/user/teacher-schedule", utils.WithCache(routes.HandleTeacherSchedule))
	mux.HandleFunc("/api/user/disciplines/{classId}", utils.WithCache(routes.HandleDisciplineClass))
	mux.HandleFunc("/api/user/disciplines/{classId}/save-lesson", routes.HandleSaveLesson)

	frontend_static_files_dir := os.Getenv("FRONTEND_STATIC_FILES_DIR")
	if frontend_static_files_dir == "" {
		frontend_static_files_dir = "./frontend"
	}

	fs := http.FileServer(http.Dir(frontend_static_files_dir))
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := frontend_static_files_dir + r.URL.Path

		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			http.ServeFile(w, r, frontend_static_files_dir+"/index.html")
			return
		}

		fs.ServeHTTP(w, r)
	}))

	log.Println("Listening on :3030")
	log.Fatal(http.ListenAndServe(":3030", mux))
}
