package middleware

import "net/http"

func (m *Middleware) AuthToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		_, err := m.Models.Tokens.AuthenticateToken(request)
		if err != nil {
			var payload struct {
				Error   bool   `json:"error"`
				Message string `json:"message"`
			}

			payload.Error = true
			payload.Message = "invalid authentication credentials"

			_ = m.App.WriteJSON(writer, http.StatusUnauthorized, payload)
		}
	})
}
