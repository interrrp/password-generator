package main

import (
	"fmt"
	"log/slog"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

const (
	addr = ":8080"

	chars      = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	defaultLen = 16
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /passwords", func(w http.ResponseWriter, r *http.Request) {
		length := getIntQueryParam(r, "length", defaultLen)
		pass := randomString(length)

		fmt.Fprintf(w, "%s", pass)
	})

	slog.Info("listening", "addr", addr)
	http.ListenAndServe(addr, mux)
}

func randomString(length int) string {
	var sb strings.Builder

	for range length {
		idx := rand.Intn(len(chars))
		char := chars[idx]
		sb.WriteByte(char)
	}

	return sb.String()
}

func getIntQueryParam(r *http.Request, key string, fallback int) int {
	s := r.URL.Query().Get(key)
	if s == "" {
		return fallback
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		return fallback
	}
	return i
}
