package middleware

import (
	"SistemManagementResto/model/web"
	"SistemManagementResto/service"
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func BasicAuth(userService service.UserService, next httprouter.Handle) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		// Ambil header Authorization
		authHeader := request.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(writer, "Authorization header is missing", http.StatusUnauthorized)
			return
		}

		// Pastikan header menggunakan Basic Auth
		if !strings.HasPrefix(authHeader, "Basic ") {
			http.Error(writer, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}

		// Decode token Basic Auth
		encodedCredentials := strings.TrimPrefix(authHeader, "Basic ")
		decodedBytes, err := base64.StdEncoding.DecodeString(encodedCredentials)
		if err != nil {
			http.Error(writer, "Invalid Basic Auth token", http.StatusUnauthorized)
			return
		}

		credentials := strings.SplitN(string(decodedBytes), ":", 2)
		if len(credentials) != 2 {
			http.Error(writer, "Invalid Basic Auth format", http.StatusUnauthorized)
			return
		}

		email, password := credentials[0], credentials[1]

		// Validasi kredensial melalui service.Login
		_, err = userService.Login(request.Context(), web.UserLoginRequest{
			Email:    email,
			Password: password,
		})
		if err != nil {
			http.Error(writer, "Invalid email or password", http.StatusUnauthorized)
			return
		}

		// Jika validasi berhasil, teruskan ke handler berikutnya
		next(writer, request, params)
	}
}
