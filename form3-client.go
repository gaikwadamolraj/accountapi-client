package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gaikwadamolraj/form3"
	"github.com/gaikwadamolraj/form3/model"
	"github.com/gaikwadamolraj/form3/utils"
)

func main() {
	accountId := utils.GetUUID()
	version := 0

	accountData := model.GetAccountModel()

	accountData.SetAccountID(accountId)
	accountData.SetCountry("GB")
	accountData.SetStatus("confirmed")

	log.Println("---------------- create --------------------")
	createResp, err := form3.Create(accountData)
	if err != nil {
		log.Println(err)
	} else {
		log.Printf("AcoountId %s with status \"%s\"", createResp.GetAccountID(), createResp.GetStatus())
	}
	log.Println("---------------- create --------------------")

	log.Println("---------------- Fetch ---------------------")
	response, err := form3.FetchById(accountId)
	if err != nil {
		log.Println(err)
	} else {
		log.Printf("Account found with ID %s", response.ID)
	}
	log.Println("---------------- Fetch ---------------------")
	log.Println("---------------- Delete --------------------")

	err = form3.DeleteByIdAndVer(accountId, version)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(fmt.Sprintf("Account with id %s and version \"%d\" got deleted", accountId, version))
	}
	log.Println("---------------- Delete --------------------")
	http.ListenAndServe("localhost:5000", nil)
}
