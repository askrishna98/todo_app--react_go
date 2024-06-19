package router

// this provides the routing and request handling for the web server application.
import (
	"github.com/aswin/go-react-todo/middleware"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	// Creates instance of new router
	router := mux.NewRouter()

	// Define the routes and their corresponding handlers from the middleware package

	router.HandleFunc("/api/task", middleware.GetAllTasks).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/tasks", middleware.CreateTask).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/tasks/{id}", middleware.TaskComplete).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/undoTask/{id}", middleware.UndoTask).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/deleteTask/{id}", middleware.DeleteTask).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/deleteAllTask", middleware.DeleteAllTask).Methods("DELETE", "OPTIONS")
	return router
}
