package main

import (
	"encoding/json"
	"fmt"
	eureka "github.com/mylixos/eurekaclient/client"
	"net/http"
	"time"
)

func main() {

	// create eureka client
	client := eureka.NewClient(eureka.NewConfig(
		"", // eureka server endpoint
		"go-client-example",
		10000,
	))
	go printApp(client)
	// start client, register、heartbeat、refresh
	client.Start()

	// http server
	http.HandleFunc("/v1/services", func(writer http.ResponseWriter, request *http.Request) {
		// full applications from eureka server
		apps := client.Applications

		b, _ := json.Marshal(apps)
		_, _ = writer.Write(b)
	})

	// start http server
	if err := http.ListenAndServe(":10000", nil); err != nil {
		fmt.Println(err)
	}

}

func printApp(client *eureka.Client) {
	for true {
		app, b := client.GetAppByName("COMPANY-ADMIN-NEW")
		//if client.Applications != nil {
		//	appBytes, e := json.Marshal(client.Applications.Applications)
		//	if e != nil {
		//		fmt.Println("Marshal apps Error")
		//	}
		//	fmt.Println("Apps:\n\t", string(appBytes), "\n\n")
		//}
		if b {
			bytes, err := json.Marshal(app)
			if err != nil {
				fmt.Println("Marshal apps Error")
			}
			fmt.Println("COMPANY-ADMIN-NEW App:\n\t", string(bytes))
		}
		sleep := time.Duration(client.Config.RegistryFetchIntervalSeconds)
		fmt.Println("printApp sleep seconds", sleep*time.Second, " seconds")
		time.Sleep(sleep * time.Second)
	}
}
