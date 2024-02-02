package main

import (
	"context"
	"flag"
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

func main() {

	var kubeconfig *string

	home := homedir.HomeDir()

	kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "actual path to kubeconfigure file")

	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)

	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	//pod list return
	list, err := clientset.CoreV1().Pods("").List(context.Background(), v1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}

	//printing pods
	fmt.Printf("total Pods are: %d\n", len(list.Items))

	fmt.Println("Podlist with corresponding namespaces are shown here:\n")
	for _, pod := range list.Items {
		fmt.Printf("%s  -- %s\n", pod.Name, pod.Namespace)
	}

	pod, err := clientset.CoreV1().Pods("kube-system").Get(context.Background(), "kube-apiserver-kind-control-plane", v1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(pod.Labels) //no usage

	//running deployments
	fmt.Println("Deployments are: ")
	deployments, err := clientset.AppsV1().Deployments("default").List(context.Background(), v1.ListOptions{})
	for _, d := range deployments.Items {
		fmt.Println(d.Name)
	}
	if err != nil {
		panic(err.Error())
	}

}
