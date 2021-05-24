package web

import (
	"context"
	"fmt"
	"github.com/elvismdnin/match_gateway/internal/web"
	"log"
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

			cchessSSets := clientset.AppsV1().StatefulSets("cchess")

			log.Println("\nUpdating manager statefulSet")
			log.Println(cchessSSets)

			managerSSet, err := cchessSSets.Get(context.TODO(), "manager", metav1.GetOptions{})
			if err != nil {
				log.Println("Error getting managerSSet")
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			log.Println("After got managerSSet:")
			log.Println(managerSSet)
			log.Println(managerSSet.Name)
			log.Println(managerSSet.Size())

			replicas := *managerSSet.Spec.Replicas
			managerSSet.Spec.Replicas = int32Ptr(replicas + 1)

			_, err = cchessSSets.Update(context.TODO(), managerSSet, metav1.UpdateOptions{})
			if err != nil {
				log.Println("Error updating managerSSet")
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			log.Println("Finished update")
			fmt.Printf("Updated statefulSet %q.\nIt was %d replicas, but now there are %d.", managerSSet.Name, replicas, replicas+1)
			updateRet := fmt.Sprintf("{replicaNum: %d}", replicas + 1)

			_, _ = fmt.Fprint(w, updateRet)
		},
	}
}

func int32Ptr(i int32) *int32 { return &i }
