package main

/*
#cgo LDFLAGS: -L. -lrustlib
#include <stdint.h>
#include <stdlib.h>

char* hash_identity(const char* email, uint64_t timestamp);
void free_str(char* s);
*/
import "C"

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"unsafe"
)

type HashRequest struct {
	Email     string `json:"email"`
	Timestamp uint64 `json:"timestamp"`
}

type HashResponse struct {
	Hash string `json:"hash"`
}

func hash(email string, timestamp uint64) string {
	cEmail := C.CString(email)
	defer C.free(unsafe.Pointer(cEmail))

	result := C.hash_identity(cEmail, C.uint64_t(timestamp))
	if result == nil {
		return "error"
	}
	defer C.free_str(result)

	return C.GoString(result)
}

func handler(w http.ResponseWriter, r *http.Request) {
	var req HashRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", 400)
		return
	}

	if req.Timestamp == 0 {
		req.Timestamp = uint64(time.Now().Unix())
	}

	hashed := hash(req.Email, req.Timestamp)

	json.NewEncoder(w).Encode(HashResponse{Hash: hashed})
}

func main() {
	http.HandleFunc("/hash", handler)
	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
