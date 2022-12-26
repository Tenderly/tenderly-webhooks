// server.go
//
// Use this sample code to handle webhook events
// 
// Run the server on http://localhost:8008
// go run server.go

package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type TransactionLog struct {
	Address string   `json:"address"`
	Topics  []string `json:"topics"`
	Data    string   `json:"data"`
}

type TransactionPayload struct {
	Network           string           `json:"network"`
	BlockHash         string           `json:"block_hash"`
	BlockNumber       int64            `json:"block_number"`
	Hash              string           `json:"hash"`
	From              string           `json:"from"`
	To                *string          `json:"to"`
	Logs              []TransactionLog `json:"logs"`
	Input             *string          `json:"input"`
	Value             *string          `json:"value"`
	Nonce             *string          `json:"nonce"`
	Gas               *string          `json:"gas"`
	GasUsed           *string          `json:"gas_used"`
	CumulativeGasUsed *string          `json:"cumulative_gas_used"`
	GasPrice          *string          `json:"gas_price"`
	GasTipCap         *string          `json:"gas_tip_cap"`
	GasFeeCap         *string          `json:"gas_fee_cap"`
}

type WebhookRequestBody struct {
	ID          string              `json:"id"` // idempotency key
	EventType   string              `json:"event_type"`
	Transaction *TransactionPayload `json:"transaction"`
}

func main() {
	http.HandleFunc("/webhook", handleWebhook)
	address := "localhost:8008"
	log.Printf("Listening on %s", address)
	log.Fatal(http.ListenAndServe(address, nil))
}

// replace this with key taken from dashboard for specific webhook
// if you are testing webhook before creation set it on empty string
const signingKey = ""

func handleWebhook(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	//set req method get to verify that a webhook endpoint exists.
	case "GET":
		w.WriteHeader(http.StatusOK)
	case "POST":
		signature := req.Header.Get("X-Tenderly-Signature")
		timestamp := req.Header.Get("Date")

		var body WebhookRequestBody
		bytes, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error read body: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if !isValidSignature(signature, bytes, timestamp) {
			fmt.Fprintf(os.Stderr, "Error signature not valid \n")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if err := json.Unmarshal(bytes, &body); err != nil {
			fmt.Fprintf(os.Stderr, "Error unmarshal body: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		switch body.EventType {
		case "TEST":
			// Then define and call a function to handle the test event
		case "ALERT":
			// Then define and call a function to handle the alert event
		// ... handle other event types
		default:
			fmt.Fprintf(os.Stderr, "Unhandled event type: %s\n", body.EventType)
		}

		w.WriteHeader(http.StatusOK)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.\n")
		w.WriteHeader(http.StatusBadRequest)
	}
}

func isValidSignature(
	signature string,
	body []byte,
	timestamp string,
) bool {
	h := hmac.New(sha256.New, []byte(signingKey))
	h.Write(body)
	h.Write([]byte(timestamp))
	digest := hex.EncodeToString(h.Sum(nil))
	return digest == signature
}