package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gaikwadamolraj/form3"
	"github.com/gaikwadamolraj/form3/model"
	"github.com/gaikwadamolraj/form3/utils"
)

func main() {
	ctx := context.Background()
	client := form3.NewClient()

	accountId := utils.GetUUID()
	version := 0

	accountData := model.GetAccountModel()

	accountData.SetAccountID(accountId)

	log.Println("---------------- create --------------------")
	req, err := client.Create(ctx, accountData)
	if err == nil {
		log.Println("Error occured")
	}
	response, err := client.Do(req)
	if err != nil {
		log.Println("Error occured")
	}
	if response.StatusCode == http.StatusCreated {
		log.Printf("Acoounts Fetched sucessfully %s", response.Status)
	} else {
		log.Println(fmt.Sprintf("Error occured status code %s", response.Status))
	}

	log.Println("---------------- create --------------------")

	log.Println("---------------- Fetch ---------------------")
	req, err = client.Fetch(ctx, accountId)
	if err != nil {
		log.Println("Error occured")
	}

	response, err = client.Do(req)
	if err != nil {
		log.Println("Error occured")
	}
	if response.StatusCode == http.StatusOK {
		log.Printf("Acoounts Fetched sucessfully %s", response.Status)
	} else {
		log.Println(fmt.Sprintf("Error occured status code %s", response.Status))
	}
	log.Println("---------------- Fetch ---------------------")

	log.Println("---------------- Delete --------------------")
	req, err = client.Delete(ctx, accountId, version)
	if err != nil {
		log.Println("Error occured")
	}

	response, err = client.Do(req)
	if err != nil {
		log.Println("Error occured")
	}

	if response.StatusCode == http.StatusNoContent {
		log.Printf("Account deleted %s", response.Status)
	} else {
		log.Println(fmt.Sprintf("Error occured status code  %s", response.Status))
	}

	log.Println("---------------- Delete --------------------")
}
