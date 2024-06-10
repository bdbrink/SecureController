package main

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
    var kubeconfig *string
    if home := homedir.HomeDir(); home != "" {
        kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
    } else {
        kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
    }
    flag.Parse()

    // Build the config from the kubeconfig file
    config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig)
    if err != nil {
        panic(err.Error())
    }

    // Create the clientset
    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        panic(err.Error())
    }

    // Get the admission controllers
    admissionControllers, err := clientset.AdmissionregistrationV1().ValidatingWebhookConfigurations().List(context.TODO(), metav1.ListOptions{})
    if err != nil {
        panic(err.Error())
    }

    // Print the admission controllers
    fmt.Println("Validating Webhook Configurations:")
    for _, webhook := range admissionControllers.Items {
        fmt.Printf("- Name: %s\n", webhook.Name)
        for _, rule := range webhook.Webhooks {
            fmt.Printf("  Webhook: %s\n", rule.Name)
        }
    }
}
