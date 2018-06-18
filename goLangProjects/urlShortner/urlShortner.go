package urlShortner

import "net/http"

func MapHandler(pathstoUrls map[string]string, fallback http.Handler) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := pathstoUrls[r.RequestURI]
		if res == "" {
			fallback.ServeHTTP(w, r)
		} else {
			http.Redirect(w, r, res, 301)
		}
	})
}
