package main

import (
	"database/src/config"
	"database/src/middlewares"
	"database/src/router"
	"fmt"
	"log"
	"net/http"
)

/* func init() {
	key := make([]byte, 64)

	if _, error := rand.Read(key); error != nil {
		log.Fatal(error)
	}
	stringBase64 := base64.StdEncoding.EncodeToString(key)
	fmt.Println(stringBase64)
} */

func main() {
	config.EnveriormentsVariable()

	routers := router.Generate()

	handler := middlewares.CorsMiddleware(routers)

	fmt.Printf("\nRodando API na porta :%d\n", config.Port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), handler))
}
