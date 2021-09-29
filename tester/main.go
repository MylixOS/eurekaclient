package main

import (
	"encoding/json"
	eureka "eurekaclient/client"
	"fmt"
	"net/http"
	"time"
)

func main() {
	// create eureka client
	client := eureka.NewClient(&eureka.Config{
		//DefaultZone:           "http://localhost:8761/eureka/",
		DefaultZone: "http://admin:123456@118.118.118.11:9090/eureka/",
		//DefaultZone:           "http://admin:123456@eureka.t.dacube.cn/eureka/",
		App:                   "go-example",
		Port:                  10000,
		RenewalIntervalInSecs: 10,
		DurationInSecs:        30,
		Metadata: map[string]interface{}{
			"VERSION":              "0.1.0",
			"NODE_GROUP_ID":        0,
			"PRODUCT_CODE":         "DEFAULT",
			"PRODUCT_VERSION_CODE": "DEFAULT",
			"PRODUCT_ENV_CODE":     "DEFAULT",
			"SERVICE_VERSION_CODE": "DEFAULT",
		},
	})
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
		fmt.Println("app: ", "COMPANY-ADMIN-NEW ", app, b)
		sleep := time.Duration(client.Config.RegistryFetchIntervalSeconds)
		fmt.Println("sleep seconds", sleep*time.Second, " seconds")
		time.Sleep(sleep * time.Second)
	}
}
