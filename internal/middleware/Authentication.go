package middleware

// func ForceSSL(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		if os.Getenv("MODE") != "development" {
// 			if r.Header.Get("x-forwarded-proto") != "https" {
// 				sslURL := "https://" + r.Host + r.RequestURI
// 				http.Redirect(w, r, sslURL, http.StatusTemporaryRedirect)
// 				return
// 			}
// 		}
// 		next.ServeHTTP(w, r)
// 	})
// }