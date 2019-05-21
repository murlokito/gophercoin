package gcd

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// BuildAndServeAPI is the function used to serve the API endpoints
func (s *GcdServer) BuildAndServeAPI() {
	log.Println("[GCDAPI] Building API endpoints.")
	s.Router = mux.NewRouter()
	s.Router.HandleFunc("/", s.Index).Methods("GET")
	s.Router.HandleFunc("/create_wallet", s.CreateWallet).Methods("POST")
	s.Router.HandleFunc("/create_blockchain", s.CreateBlockchain).Methods("POST")
	s.Router.HandleFunc("/get_balance", s.GetBalance).Methods("GET")
	s.Router.HandleFunc("/submit_tx", s.SubmitTx).Methods("POST")
	s.Router.HandleFunc("/generate_blocks", s.GenerateBlocks).Methods("POST")
	log.Println("[GCDAPI] Listening and Serving API. Port: 9000")
	log.Fatal(http.ListenAndServe(":9000", s.Router))
}
