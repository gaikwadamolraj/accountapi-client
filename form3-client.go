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
	accountId := utils.GetUUID()
	version := 0

	accountData := model.GetAccountModel()

	accountData.SetAccountID(accountId)
	accountData.SetCountry("GB")
	accountData.SetStatus("confirmed")

	log.Println("---------------- create --------------------")
	response, err := form3.Create(ctx, accountData)
	if err != nil {
		log.Println(err)
	} else {
		log.Printf("Acoounts Fetched sucessfully  %s", response.Status)
	}
	log.Println("---------------- create --------------------")

	log.Println("---------------- Fetch ---------------------")
	response, err = form3.FetchById(ctx, accountId)
	if err != nil {
		log.Println(err)
	} else {
		log.Printf("Acoounts Fetched sucessfully %s", response.Status)
	}
	log.Println("---------------- Fetch ---------------------")
	log.Println("---------------- Delete --------------------")

	response, err = form3.DeleteByIdAndVer(ctx, accountId, version)
	if err != nil {
		log.Println(err)
	} else {
		if response.StatusCode == http.StatusNoContent {
			log.Printf("Account deleted %s", response.Status)
		} else {
			log.Println(fmt.Sprintf("Some error got while deleting accout %d", response.StatusCode))
		}
	}
	log.Println("---------------- Delete --------------------")
}
