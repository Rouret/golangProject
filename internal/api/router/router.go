package router

import (
	"log"

	Models "github.com/Rouret/golangProject/internal/models"
	mux "github.com/julienschmidt/httprouter"
)


func NewRouter(routes Models.Routes) *mux.Router {
        router := mux.New()
        
        for _, route := range routes {
                router.Handle(route.Method, route.Path, route.Handle)
                log.Println("Route registered " + route.Method + " " + route.Path)
        }

        log.Println("Router OK")
        return router
}

