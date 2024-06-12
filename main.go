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
    config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
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
    for _, webhookConfig := range admissionControllers.Items {
        fmt.Printf("- Name: %s\n", webhookConfig.Name)
        for _, webhook := range webhookConfig.Webhooks {
            fmt.Printf("  Webhook: %s\n", webhook.Name)
            fmt.Printf("    Client Config: %v\n", webhook.ClientConfig)
            fmt.Printf("    Admission Review Versions: %v\n", webhook.AdmissionReviewVersions)
            fmt.Printf("    Rules:\n")
            for _, rule := range webhook.Rules {
                fmt.Printf("      Operations: %v\n", rule.Operations)
                fmt.Printf("      APIGroups: %v\n", rule.APIGroups)
                fmt.Printf("      APIVersions: %v\n", rule.APIVersions)
                fmt.Printf("      Resources: %v\n", rule.Resources)
                fmt.Printf("      Scope: %v\n", rule.Scope)
            }
            fmt.Printf("    Failure Policy: %v\n", webhook.FailurePolicy)
            fmt.Printf("    Match Policy: %v\n", webhook.MatchPolicy)
            fmt.Printf("    Namespace Selector: %v\n", webhook.NamespaceSelector)
            fmt.Printf("    Object Selector: %v\n", webhook.ObjectSelector)
            fmt.Printf("    Side Effects: %v\n", webhook.SideEffects)
            fmt.Printf("    Timeout Seconds: %v\n", webhook.TimeoutSeconds)
        }
    }
}
