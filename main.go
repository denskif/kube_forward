package main

import (
	"fmt"
	"os/exec"
	"strings"
	"sync"
)


func main() {
	cmd := "kubectl"
	args := [] string {
		"port-forward --namespace game-service svc/casino-content-provider-sv 9079:9090",
		"port-forward --namespace game-service svc/game-session-sv 9080:9090",
		"port-forward --namespace payment-gateway svc/payment-service-sv 9081:9000",
		"port-forward --namespace payment-gateway svc/payment-service-sv 9082:9090",
		"port-forward --namespace retail svc/retail-service-headless-sv 9083:9090",
		"port-forward --namespace retail svc/retail-service-headless-sv 9084:9000",
		"port-forward --namespace retail svc/retail-location-headless-sv 9085:9090",
		"port-forward --namespace analytics svc/query-service-sv 9086:9090",
		"port-forward --namespace retail svc/retail-location-headless-sv 9087:9090",
		"port-forward --namespace crm svc/bonus-service-sv 9088:9000",
		"port-forward --namespace crm svc/bonus-service-sv 9089:9090",
		"port-forward --namespace game-service svc/game-manager-sv 9090:9090",
		"port-forward --namespace billing-service svc/billing-sv 9091:9090",
		"port-forward --namespace message-service svc/template-service-sv 9092:9090",
		"port-forward --namespace language-service svc/language-service-sv 9093:9090",
		"port-forward --namespace user-management svc/registration-service-sv 9094:9090",
		"port-forward --namespace cms svc/cms-sv 9095:9090",
		"port-forward --namespace user-management svc/auth-service-sv 9096:9090",
		"port-forward --namespace user-management svc/client-service-sv 9097:9000",
		"port-forward --namespace user-management svc/client-service-sv 9098:9090",
		"port-forward --namespace user-management svc/bouser-service-sv 9099:9090",
	}

	//Common Channel for the goroutines
	tasks := make(chan *exec.Cmd, len(args))

	//Spawning 4 goroutines
	var wg sync.WaitGroup
	for i := 0; i < len(args); i++ {
		wg.Add(1)
		go func(num int, w *sync.WaitGroup) {
			defer w.Done()
			var (
				out []byte
				err error
			)
			    cmd := <- tasks  // this will exit the loop when the channel closes
				out, err = cmd.Output()
				if err != nil {
					fmt.Printf("can't get stdout:", err)
				}
				fmt.Printf("goroutine %d command output:%s", num, string(out))

		}(i, &wg)
	}
	//Generate Tasks
	for i := 0; i < len(args); i++ {
		tasks <- exec.Command(cmd, strings.Split(args[i]," ")...)
	}
	close(tasks)

	// wait for the workers to finish
	wg.Wait()

	fmt.Println("Done")
}