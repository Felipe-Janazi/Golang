package middleware

import "net/http"

// Recebe os controllers que estamos usando e devolvemos o mesmo mas com o setup atualizado
func ContentTypeMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-type", "application/json")
		// Colocamos em todos os handlers
		next.ServeHTTP(w, r)
	})
}