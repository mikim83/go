package main

import (
	"github.com/hudl/fargo"
	"fmt"
	"strings"
	"strconv"
	"net/http"
)


func main() {
	app_status_map := make(map[string]map[string]string)
	c , _ := fargo.NewConnFromConfigFile("config/config.gcfg")
	app,_ := c.GetApps()
	for _, value := range app {
	 	for _, ins := range value.Instances {
			if !strings.Contains(ins.App,"EUREKA") {
				resp, err := http.Get(ins.HealthCheckUrl)
				app_status_map[ins.App] = make(map[string]string)
				if err != nil {
					app_status_map[ins.App]["RESPONSE"] = "500"
					app_status_map[ins.App]["IPADDR"] = ins.IPAddr
					app_status_map[ins.App]["HOMEPAGE"] = ins.HomePageUrl
				}else{
					app_status_map[ins.App]["RESPONSE"] = strconv.Itoa(resp.StatusCode)
 					app_status_map[ins.App]["IPADDR"] = ins.IPAddr 
					app_status_map[ins.App]["HOMEPAGE"] = ins.HomePageUrl 
				}
			}
		}
	}
	for key, value := range app_status_map {
		fmt.Printf("APP: %s\n", key)
		fmt.Printf("RESPONSE: %v\n", value["RESPONSE"])
		fmt.Printf("HOMEPAGE: %v\n", value["HOMEPAGE"])
		fmt.Printf("IPADDR: %v\n", value["IPADDR"])
	} 
}
