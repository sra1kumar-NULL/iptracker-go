package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/spf13/cobra"
)

var (
	traceCmd = &cobra.Command{
		Use:   "trace",
		Short: "trace the ip",
		Long:  `Trace the ip address`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				for _, ip := range args {
					showCommand(ip)
				}
			} else {
				fmt.Println("Please provide an ip to trace")
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(traceCmd)
}
type IP struct{

	STATUS string`json::"status"`
	COUNTRY string`json::"country"`
	CITY string`json::"city"`
	lat float64`json::"lat"`
	lon float64`json::"lon"`
	TIMEZONE string`json::"timezone"`
	REGIONNAME string`json::"regionName"`
}

func showCommand(ip string) {
	api := `http://ip-api.com/json/` + ip
	data := getData(api)
	dataGrepped :=IP{}
	err := json.Unmarshal(data,&dataGrepped)
	if  err!=nil{
		log.Println("Unable to marshall the response")
	}
	fmt.Println("Data Found : ")
	fmt.Printf("Ip : %s\nStatus : %s\nCountry : %s\ncity : %s\nlat : %f\nlong : %f\ntimezone : %s \nregion : %s",ip,dataGrepped.STATUS,dataGrepped.COUNTRY,
	dataGrepped.CITY,dataGrepped.lat,dataGrepped.lon,dataGrepped.TIMEZONE,
	dataGrepped.REGIONNAME)
}

func getData(api string) []byte {
	response, err := http.Get(api)
	if err != nil {
		fmt.Print(err.Error())
		fmt.Println("Failed to fetch the ip address data")
	}
	responseByte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return responseByte
}
