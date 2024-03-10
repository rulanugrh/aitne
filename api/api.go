package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rulanugrh/aitne/api/apps"
	"github.com/rulanugrh/aitne/api/core"
	"github.com/rulanugrh/aitne/internal/config"
	srv "github.com/rulanugrh/aitne/internal/service/apps"
	srv2 "github.com/rulanugrh/aitne/internal/service/core"
)

var (
	kubeconfig = os.Getenv("KUBECONFIG_PATH")
)

type API struct {
	deployment apps.DaemonEndpoint
	daemon     apps.DaemonEndpoint
	replica    apps.ReplicaEndpoint
	statefull  apps.StatefullEndpoint

	namespace core.NamespaceEndpoint
	pod       core.PodEndpoint
	service   core.ServiceEndpoint
}

func main() {
	client, err := config.GetConfig(&kubeconfig)
	if err != nil {
		log.Printf("error cannot connect k8s: %s", err.Error())
	}

	deployment := srv.NewDeployment(client)
	daemon := srv.NewDaemonSet(client)
	replica := srv.NewReplicaSet(client)
	statefull := srv.NewStatefulSet(client)

	pod := srv2.NewPod(client)
	namespace := srv2.NewNamespace(client)
	service := srv2.NewServiceKurbenetes(client)

	api := API{
		statefull:  apps.NewStatefullEndpoint(statefull),
		daemon:     apps.NewDameonEndpoint(daemon),
		replica:    apps.NewReplicaEndpoint(replica),
		deployment: apps.NewDeploymentEndpoint(deployment),

		pod:       core.NewPodEndpoint(pod),
		namespace: core.NewNamespaceEndpoint(namespace),
		service:   core.NewServiceEndpoint(service),
	}

	router := mux.NewRouter()
	api.DaemonRouter(router)
	api.DeploymentRoute(router)
	api.ReplicaRouter(router)
	api.StatefullRouter(router)

	api.PodRouter(router)
	api.NamespaceRouter(router)
	api.ServiceRouter(router)

	server := http.Server{
		Addr:    "0.0.0.0:3000",
		Handler: router,
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Printf("cannot running http service : %s", err.Error())
	}

	log.Println("Server runnint at :3000")

}

func (api *API) DeploymentRoute(r *mux.Router) {
	app := r.PathPrefix("/api/deployment").Subrouter()
	app.HandleFunc("/create/", api.deployment.Create).Methods("POST")
	app.HandleFunc("/getAll/", api.deployment.Get).Methods("GET")
	app.HandleFunc("/get/{name}", api.deployment.GetByName).Methods("GET")
	app.HandleFunc("/delete/{name}", api.deployment.Delete).Methods("DELETE")
}

func (api *API) DaemonRouter(r *mux.Router) {
	app := r.PathPrefix("/api/daemon").Subrouter()
	app.HandleFunc("/create/", api.daemon.Create).Methods("POST")
	app.HandleFunc("/getAll/", api.daemon.Get).Methods("GET")
	app.HandleFunc("/get/{name}", api.daemon.GetByName).Methods("GET")
	app.HandleFunc("/delete/{name}", api.daemon.Delete).Methods("DELETE")
}

func (api *API) ReplicaRouter(r *mux.Router) {
	app := r.PathPrefix("/api/replica").Subrouter()
	app.HandleFunc("/create/", api.replica.Create).Methods("POST")
	app.HandleFunc("/getAll/", api.replica.Get).Methods("GET")
	app.HandleFunc("/get/{name}", api.replica.GetByName).Methods("GET")
	app.HandleFunc("/delete/{name}", api.replica.Delete).Methods("DELETE")
}

func (api *API) StatefullRouter(r *mux.Router) {
	app := r.PathPrefix("/api/statefull").Subrouter()
	app.HandleFunc("/create/", api.statefull.Create).Methods("POST")
	app.HandleFunc("/getAll/", api.statefull.Get).Methods("GET")
	app.HandleFunc("/get/{name}", api.statefull.GetByName).Methods("GET")
	app.HandleFunc("/delete/{name}", api.statefull.Delete).Methods("DELETE")
}

func (api *API) PodRouter(r *mux.Router) {
	app := r.PathPrefix("/api/pod").Subrouter()
	app.HandleFunc("/create/", api.pod.Create).Methods("POST")
	app.HandleFunc("/getAll/", api.pod.Get).Methods("GET")
	app.HandleFunc("/get/{name}", api.pod.GetByName).Methods("GET")
	app.HandleFunc("/delete/{name}", api.pod.Delete).Methods("DELETE")
}

func (api *API) NamespaceRouter(r *mux.Router) {
	app := r.PathPrefix("/api/namespace").Subrouter()
	app.HandleFunc("/create/", api.namespace.Create).Methods("POST")
	app.HandleFunc("/getAll/", api.namespace.Get).Methods("GET")
	app.HandleFunc("/get/{name}", api.namespace.GetByName).Methods("GET")
	app.HandleFunc("/delete/{name}", api.namespace.Delete).Methods("DELETE")
}

func (api *API) ServiceRouter(r *mux.Router) {
	app := r.PathPrefix("/api/service").Subrouter()
	app.HandleFunc("/create/", api.service.Create).Methods("POST")
	app.HandleFunc("/getAll/", api.service.Get).Methods("GET")
	app.HandleFunc("/get/{name}", api.service.GetByName).Methods("GET")
	app.HandleFunc("/delete/{name}", api.service.Delete).Methods("DELETE")
}
