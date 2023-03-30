package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/argoproj/argo-cd/v2/pkg/clientset/versioned"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	masterURL string
)

func main() {
	fmt.Println(nil)
	flag.Parse()
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return
	}

	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, fmt.Sprintf("%s/Downloads/civo-argocd-sync-kubeconfig", homeDir))
	if err != nil {
		fmt.Println("Error building kubeconfig: %v", err)
	}

	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		fmt.Println("Error building kube clientset: %v", err)
	}
	fmt.Println(kubeClient)

	config, err := argocd.GetClientConfig()
	if err != nil {
		fmt.Printf("Failed to get ArgoCD client config: %v\n", err)
		os.Exit(1)
	}

	argocdClient := versioned.NewClient(config)

	// appName := "ingress-nginx"
	// namespace := "argocd"

	// Create a new informer factory for the Argo CD application resource
	// factory := externalversions.NewSharedInformerFactoryWithOptions(argocdClient, time.Second*30, externalversions.WithNamespace(namespace))

	// Create an informer for the Argo CD application resource
	// informer := factory.Argoproj().V1alpha1().Applications().Informer()

	// // Add a handler function for when the informer receives an update
	// informer.AddEventHandler(
	// 	cache.ResourceEventHandlerFuncs{
	// 		UpdateFunc: func(oldObj, newObj interface{}) {
	// 			newApp := newObj.(*v1alpha1.Application)
	// 			if newApp.Name == appName {
	// 				fmt.Println("Argo CD application %s updated, status: %s", appName, newApp.Status.Sync.Status)
	// 			}
	// 		},
	// 	},
	// )

	// Start the informer
	// stopCh := make(chan struct{})
	// defer close(stopCh)
	// go informer.Run(stopCh)

	// Wait for the informer to sync with the API server
	// if !cache.WaitForCacheSync(stopCh, informer.HasSynced) {
	// 	fmt.Println("Error syncing informer cache")
	// }

	// Wait indefinitely
	select {}
}
