package web

import (
	"context"
	"fmt"
	"github.com/elvismdnin/match_gateway/internal/web"
	"log"
	"net/http"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
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

			log.Println("Creating deployment...")
			deploymentsClient := clientset.AppsV1().Deployments("cchess")
			result, err := deploymentsClient.Create(context.TODO(), createDeployment(), metav1.CreateOptions{})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())
			deployReturn := fmt.Sprintf("Created deployment %s.\n", result.GetObjectMeta().GetName())

			_, _ = fmt.Fprint(w, deployReturn)
		},
	}
}

func createDeployment() *appsv1.Deployment {

	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "manager",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(2),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "manager",
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "manager",
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "manager",
							Image: "docker.pkg.github.com/elvismdnin/match_gateway/match_manager:0.1",
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 8000,
								},
							},
						},
					},
				},
			},
		},
	}
}

func int32Ptr(i int32) *int32 { return &i }