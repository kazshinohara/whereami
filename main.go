package main

import (
	"cloud.google.com/go/compute/metadata"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	port    = os.Getenv("PORT")
	version = os.Getenv("VERSION")
	kind    = os.Getenv("KIND")
)

type rootResponse struct {
	Kind     string `json:"kind"`    // backend, backend-b, backend-c
	Version  string `json:"version"` // v1, v2, v3
	Region   string `json:"region"`
	Cluster  string `json:"cluster"`
	Hostname string `json:"hostname"`
}

func resolveRegion() string {
	if !metadata.OnGCE() {
		log.Println("This app is not running on GCE")
	} else {
		zone, err := metadata.Zone()
		if err != nil {
			log.Printf("could not get zone info: %v", err)
			return "unknown"
		}
		region := zone[:strings.LastIndex(zone, "-")]
		return region
	}
	return "unknown"
}

func resolveCluster() string {
	if !metadata.OnGCE() {
		log.Println("This app is not running on GCE")
	} else {
		cluster, err := metadata.Get("/instance/attributes/cluster-name")
		if err != nil {
			log.Printf("could not get cluster name: %v", err)
			return "unknown"
		}
		return cluster
	}
	return "unknown"
}

func resolveHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Printf("could not get hostname: %v", err)
		return "unknown"
	}
	return hostname
}

func fetchRootResponse(w http.ResponseWriter, r *http.Request) {
	responseBody, err := json.Marshal(&rootResponse{
		Kind:     kind,
		Version:  version,
		Region:   resolveRegion(),
		Cluster:  resolveCluster(),
		Hostname: resolveHostname(),
	})
	if err != nil {
		log.Printf("could not json.Marshal: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.Write(responseBody)
}

func main() {
	// Set up Routing and Server
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", fetchRootResponse).Methods("GET")
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}
