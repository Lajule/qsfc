package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/Lajule/qsfc/contact"
	"github.com/nimajalali/go-force/force"
	"log"
	"os"
)

type Configuration struct {
	UserName      string `json:"username"`
	Password      string `json:"password"`
	SecurityToken string `json:"securitytoken"`
	Environment   string `json:"environment"`
}

func NewConfiguration(path string) *Configuration {
	config := Configuration{}
	file, _ := os.Open(path)
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		log.Fatal(err)
	}
	return &config
}

func NewForceApi(config *Configuration) *force.ForceApi {
	forceApi, err := force.Create(
		"YOUR-API-VERSION",
		"YOUR-CLIENT-ID",
		"YOUR-CLIENT-SECRET",
		config.UserName,
		config.Password,
		config.SecurityToken,
		config.Environment,
	)
	if err != nil {
		log.Fatal(err)
	}
	return forceApi
}

var configFlag string
var countFlag bool

func init() {
	flag.StringVar(&configFlag, "config", "./config.json", "Path to config file")
	flag.BoolVar(&countFlag, "c", false, "Counts Salesforce contacts")
	flag.Usage = func() {
		fmt.Println("Usage: qsfc [flags] ...")
		flag.PrintDefaults()
	}
}

func print(forceApi *force.ForceApi, response *contact.ContactQueryResponse) {
	for {
		for _, c := range response.Records {
			fmt.Println(c.CSV())
		}
		if response.Done {
			break
		}
		response = contact.NextContacts(forceApi, response)
	}
}

func main() {
	flag.Parse()
	config := NewConfiguration(configFlag)
	forceApi := NewForceApi(config)
	response := contact.Contacts(forceApi)
	if countFlag {
		fmt.Printf("%v\n", response.TotalSize)
	} else {
		print(forceApi, response)
	}
}
