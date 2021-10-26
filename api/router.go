package api

import mux "github.com/julienschmidt/httprouter"

/*The code creates a new router instance and iterate through all the Routes
to get each’s Route’s Method, Pattern and Handle and registers a new request handle
with the given path and method.*/

func NewRouter() *mux.Router {
        
        router := mux.New()
        
        for _, route := range routes {
                
                router.Handle(route.Method, route.Path, route.Handle)
                
        }
        
        return router
}