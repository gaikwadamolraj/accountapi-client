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
	req, _ := client.Create(ctx, accountData)
	response, _ := client.Do(req)
	if response.StatusCode == http.StatusCreated {
		log.Printf("Acoounts Fetched sucessfully %s", response.Status)
	} else {
		log.Println(fmt.Sprintf("Error occured status code %s", response.Status))
	}
	log.Println("---------------- create --------------------")

	log.Println("---------------- Fetch ---------------------")
	req, _ = client.Fetch(ctx, accountId)
	response, _ = client.Do(req)
	if response.StatusCode == http.StatusOK {
		log.Printf("Acoounts Fetched sucessfully %s", response.Status)
	} else {
		log.Println(fmt.Sprintf("Error occured status code %s", response.Status))
	}
	log.Println("---------------- Fetch ---------------------")

	log.Println("---------------- Delete --------------------")
	req, _ = client.Delete(ctx, accountId, version)
	response, _ = client.Do(req)
	if response.StatusCode == http.StatusNoContent {
		log.Printf("Account deleted %s", response.Status)
	} else {
		log.Println(fmt.Sprintf("Error occured status code  %s", response.Status))
	}

	log.Println("---------------- Delete --------------------")
}
