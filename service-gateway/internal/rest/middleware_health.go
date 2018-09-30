package rest

import (
	"encoding/json"
	"github.com/1ambda/domain-driven-design-go/service-gateway/internal/config"
	"net/http"
)

type healthCheckResponse struct {
	Version  string `json:"version"`
	Commit   string `json:"commit"`
	PodName  string `json:"pod_name"`
	PodIp    string `json:"pod_ip"`
	NodeName string `json:"node_name"`
}

func InjectHealthCheckMiddleware(h http.Handler) http.Handler {
	env := config.Env

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// if GCE LB (GLBC Ingress) health check request
		if r.Method == http.MethodGet && r.URL.Path == "/" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)

			resp := healthCheckResponse{
				Version:  env.Version,
				Commit:   env.GitCommit,
				PodName:  env.PodName,
				PodIp:    env.PodIP,
				NodeName: env.NodeName,
			}

			json.NewEncoder(w).Encode(&resp)
			return
		}

		h.ServeHTTP(w, r)
	})
}
