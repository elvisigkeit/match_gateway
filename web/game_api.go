package web

import (
	"context"
	"fmt"
	"github.com/elvismdnin/match_gateway/internal/web"
	"net/http"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"github.com/gorilla/mux"
)

func CreateGame() web.Handler {
	return web.Handler {
		Route: func(r *mux.Route) {
			r.Path("/new").Methods("GET")
		},
		Func: func(w http.ResponseWriter, r *http.Request) {
			config, err := rest.InClusterConfig()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			clientset, err := kubernetes.NewForConfig(config)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			for _, pod := range pods.Items {
				println("Pod: ", pod.Spec.NodeName, " - ", pod.Status.StartTime)
			}
			fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

			_, _ = fmt.Fprint(w, "{success}")
		},
	}
}