package api

import (
	"bytes"
	"database/sql"
	"fmt"
	"go-api/cmd/routes/user"
	"log"
	"net/http"
	"os/exec"
	"strings"

	"github.com/gorilla/mux"
)

type APIServer struct {
	db   *sql.DB
	addr string
}

func NewAPIServer(db *sql.DB, addr string) *APIServer {
	return &APIServer{
		db:   db,
		addr: addr,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userHandler := user.NewHandler()

	userHandler.RegisterRoutes(subrouter)

	log.Println("Starting server on ", s.addr)

	return http.ListenAndServe(s.addr, router)
}

func HandleOpenConnections() {
	cmd := exec.Command("lsof", "-i", ":8080")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error finding processes:", err)
		return
	}

	lines := strings.Split(out.String(), "\n")
	var pids []string
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) > 1 && fields[1] != "PID" {
			pids = append(pids, fields[1])
		}
	}

	for _, pid := range pids {
		killCmd := exec.Command("kill", "-9", pid)
		err := killCmd.Run()
		if err != nil {
			fmt.Println("Error killing process", pid, ":", err)
		} else {
			fmt.Println("Killed process", pid)
		}
	}
}
