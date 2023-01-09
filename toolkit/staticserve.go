package toolkit

import "net/http"

func StaticServe(dir string) {
	fs := http.FileServer(http.Dir(dir))
	http.Handle("/"+dir, http.StripPrefix("/"+dir, fs))
}
