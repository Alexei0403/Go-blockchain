package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/elarsaks/Go-blockchain/pkg/utils"
	"github.com/elarsaks/Go-blockchain/wallet_server/handlers"
	"github.com/gorilla/mux"
)

type WalletServer struct {
	port    uint16
	gateway string
}

// Create a new instance of WalletServer
func NewWalletServer(port uint16, gateway string) *WalletServer {
	return &WalletServer{port, gateway}
}

// Get the port of the WalletServer
func (ws *WalletServer) Port() uint16 {
	return ws.port
}

// Get the gateway of the WalletServer
func (ws *WalletServer) Gateway() string {
	return ws.gateway
}

func apiDescription() map[string]string {
	return map[string]string{
		"/":               "index",
		"/wallet":         "Wallet description...",
		"/wallet/balance": "Wallet balance description...",
		"/transaction":    "Transaction description...",
		"/miner/blocks":   "Miner blocks description...",
	}
}

func init() {
	log.SetPrefix("Wallet Server: ")
}

func (ws *WalletServer) Run() {
	// Create router
	router := mux.NewRouter()
	router.Use(utils.CorsMiddleware())

	// Return API route descriptions
	router.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(apiDescription())
	})

	router.HandleFunc("/user/wallet", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetUserWallet(w, r, ws.gateway)
	})

	router.HandleFunc("/wallet/balance", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetWalletBalance(w, r, ws.gateway)
	})

	router.HandleFunc("/transaction", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateTransaction(w, r, ws.gateway)
	})

	router.HandleFunc("/miner/blocks", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetBlocks(w, r, ws.gateway)
	})

	router.HandleFunc("/miner/wallet", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetMinerWallet(w, r, ws.gateway)
	})

	// Start server
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(ws.Port())), router))
}

func main() {
	portStr := os.Getenv("PORT") // Retrieve port from environment variable
	port, err := strconv.Atoi(portStr)

	if err != nil || port <= 0 {
		port = 8080 // It defaults to 8080 in production
	}

	// Create and run the WalletServer with the configured ports and gateway
	app := NewWalletServer(uint16(port), "http://miner-1:5001")
	app.Run()
}
