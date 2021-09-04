package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func routesServe() error {
	router := mux.NewRouter()

	subdir := pathStrip(os.Getenv("ORIGIN"))
	if subdir != "" {
		router = router.PathPrefix(subdir).Subrouter()
	}

	if err := apiRouterInit(router); err != nil {
		return err
	}

	if err := staticRouterInit(router); err != nil {
		return err
	}

	origins := handlers.AllowedOrigins([]string{"*"})
	headers := handlers.AllowedHeaders([]string{"X-Requested-With"})
	methods := handlers.AllowedMethods([]string{"GET", "POST"})

	addrPort := os.Getenv("BIND_ADDRESS") + ":" + os.Getenv("PORT")
	ssl := os.Getenv("SSL")
	fmt.Println(ssl)
	if ssl == "true" {
		cert := os.Getenv("SSL_CERT")
		fmt.Println(cert)
		key := os.Getenv("SSL_KEY")
		fmt.Println(key)
		if cert == "" || key == "" {
			return fmt.Errorf("missing cert %s or key %s file", cert, key)
		}
		logger.Infof("starting SSL server on %s\n", addrPort)
		if err := http.ListenAndServeTLS(addrPort, cert, key, handlers.CORS(origins, headers, methods)(router)); err != nil {
			logger.Errorf("cannot start server: %v", err)
		return err
		}
	}
	logger.Infof("starting server on %s\n", addrPort)
	if err := http.ListenAndServe(addrPort, handlers.CORS(origins, headers, methods)(router)); err != nil {
		logger.Errorf("cannot start server: %v", err)
		return err
	}

	return nil
}
