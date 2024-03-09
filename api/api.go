package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rulanugrh/aitne/api/apps"
	"github.com/rulanugrh/aitne/internal/config"
	srv "github.com/rulanugrh/aitne/internal/service/apps"
)

var (
	kubeconfig = flag.String("kubeconfig", "$HOME/.kube/config/", "flag for k8s config file")
)

type API struct {
	deployment apps.DaemonEndpoint
	daemon     apps.DaemonEndpoint
	replica    apps.ReplicaEndpoint
	statefull  apps.StatefullEndpoint
}

func main() {
	client, err := config.GetConfig(kubeconfig)
	if err != nil {
		log.Printf("error cannot connect k8s: %s", err.Error())
	}

	deployment := srv.NewDeployment(client)
	daemon := srv.NewDaemonSet(client)
	replica := srv.NewReplicaSet(client)
	statefull := srv.NewStatefulSet(client)

	api := API{
		statefull:  apps.NewStatefullEndpoint(statefull),
		daemon:     apps.NewDameonEndpoint(daemon),
		replica:    apps.NewReplicaEndpoint(replica),
		deployment: apps.NewDeploymentEndpoint(deployment),
	}

	router := mux.NewRouter()
	api.DaemonRouter(router)
	api.DeploymentRoute(router)
	api.ReplicaRouter(router)
	api.StatefullRouter(router)

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
