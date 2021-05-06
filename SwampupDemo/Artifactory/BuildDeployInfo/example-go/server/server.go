package server

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	_ "jfrog.com/jfrog-partnership/example-go/docs"
	logger "jfrog.com/jfrog-partnership/example-go/log"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	muxtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"github.com/swaggo/http-swagger"
)

// loads the oAuth server and enables routes w/ oAuth authentication
func StartServer() {
	tracer.Start()
	defer tracer.Stop()

	// logger
	log := logger.NewLogger()

	// oAuth2 Server
	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

	// token store
	manager.MustTokenStorage(store.NewMemoryTokenStore())
	manager.MapAccessGenerate(generates.NewAccessGenerate())

	// client memory store
	clientStore := store.NewClientStore()
	clientStore.Set("example-go", &models.Client{
		ID:     "example-go",
		Secret: "example-go",
		Domain: "http://localhost",
	})

	manager.MapClientStorage(clientStore)

	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	// Create a traced mux router.
	myRouter := muxtrace.NewRouter().StrictSlash(true)

	// Render swagger ui
	myRouter.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), //The url pointing to API definition"
	))

	// common middleware
	myRouter.Use(commonMiddleware)

	// declare the routes
	myRouter.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// handles new authentication requests for an oAuth token by clientId+Secret
	// for additional security pass the clientId+Secret as an http basic auth base64 encoded string
	myRouter.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		// support basic auth for swagger ui
		authHeader := r.Header.Get("Authorization")
		if len(authHeader) > 0 {
			user, pass, ok := r.BasicAuth()
			if !ok {
				http.Error(w, "Invalid authorization", http.StatusUnauthorized)
			}
			form, _ := url.ParseQuery(r.URL.RawQuery)
			form.Add("client_id", user)
			form.Add("client_secret", pass)
			r.URL.RawQuery = form.Encode()
		}
		srv.HandleTokenRequest(w, r)
	})

	myRouter.HandleFunc("/good", func(w http.ResponseWriter, r *http.Request) {
		_, err := srv.ValidationBearerToken(r)
		if err != nil {
			errJson := Error{Error: err.Error()}
			errBody, _ := json.Marshal(errJson)
			errString := string(errBody[:])
			http.Error(w, errString, http.StatusBadRequest)
			return
		}
		if r.Method == http.MethodGet {
			GoodHandler(w, r)
		} else {
			errJson := Error{Error: "HTTP method not allowed"}
			errBody, _ := json.Marshal(errJson)
			errString := string(errBody[:])
			http.Error(w, errString, http.StatusMethodNotAllowed)
		}
	})

	myRouter.HandleFunc("/bad", func(w http.ResponseWriter, r *http.Request) {
		_, err := srv.ValidationBearerToken(r)
		if err != nil {
			errJson := Error{Error: err.Error()}
			errBody, _ := json.Marshal(errJson)
			errString := string(errBody[:])
			http.Error(w, errString, http.StatusBadRequest)
			return
		}
		if r.Method == http.MethodGet {
			BadHandler(w, r)
		} else {
			errJson := Error{Error: "HTTP method not allowed"}
			errBody, _ := json.Marshal(errJson)
			errString := string(errBody[:])
			http.Error(w, errString, http.StatusMethodNotAllowed)
		}
	})

	// Serve the HTTP server
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}


func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.URL.String(),"swagger") {
			w.Header().Set("Cache-Control", "no-cache, no-store")
			w.Header().Set("Referrer-Policy", "strict-origin")
			w.Header().Add("Content-Type", "application/json")
		}
		next.ServeHTTP(w, r)
	})
}