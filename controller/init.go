package controller

import "github.com/gorilla/mux"

var Router *mux.Router

func init() {
	Router = mux.NewRouter()

	Router.HandleFunc("/activations", getActivations).Methods("GET")

	Router.HandleFunc("/postings/{user}", createPosting).Methods("GET")
	Router.HandleFunc("/posting", createPosting).Methods("POST")
	Router.HandleFunc("/posting", updatePosting).Methods("PUT")
	Router.HandleFunc("/posting", deletePosting).Methods("DELETE")

	Router.HandleFunc("/commit", createCommit).Methods("POST")
	Router.HandleFunc("/commit", updateCommit).Methods("PUT")
	Router.HandleFunc("/commit", deleteCommit).Methods("DELETE")
}
