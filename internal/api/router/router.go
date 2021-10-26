package router

import (
	Models "github.com/Rouret/golangProject/internal/models"
	mux "github.com/julienschmidt/httprouter"
)


func NewRouter(routes Models.Routes) *mux.Router {
        router := mux.New()
        
        for _, route := range routes {
                router.Handle(route.Method, route.Path, route.Handle)
        }
        
        return router
}

