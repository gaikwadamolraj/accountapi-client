<p align="left"> <a href="https://golang.org" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/go/go-original.svg" alt="go"  width="320" height="320"/> </a> </p>

# Client for accountapi
## Project Status
![Codecov](https://img.shields.io/badge/codecoverage-100%25-green)
![Seccan](https://img.shields.io/badge/goscan-passing-green)
![Bdd](https://img.shields.io/badge/bddtests-100%25-green)
![Container](https://img.shields.io/badge/Containerisation-Yes-green)

 > These badges are as per local run

## This go client is created for accountapi.

To run this app, you'll need:

- Go 1.17

## How to use  ##
 - Instructions
Use of this client library access account API, which support Create, Fetch and delete methood. Please refer to the
[Form3 documentation](http://api-docs.form3.tech/api.html#organisation-accounts) for information for each API deatils

   #### Set API host using env variable
   ```sh
    API_HOST
    // ex: export API_HOST=http://localhost:8080
   ```
   #### Import below modules
   ```sh
    github.com/gaikwadamolraj/form3
    github.com/gaikwadamolraj/form3/model
	github.com/gaikwadamolraj/form3/utils
   ```

   #### Sample Module created with Set and Get methods
   ```sh
    accountId := utils.GetUUID()
	version := 0

	accountData := model.GetAccountModel()

	accountData.SetAccountID(accountId)
	accountData.SetCountry("GB")
	accountData.SetStatus("confirmed")

    // Create and fetch response will be 
    // type AccountData struct
   ```

   #### Create
   ```sh
    response, err := form3.Create(accountData)
     // err is any errors found
     // Sucess will result the of AccountData module
     // You can use GET as like createResp.GetAccountID()
   ```

   #### Fetch
   ```sh
    response, err := form3.FetchById(accountId)
     // err is any errors found
     // Sucess will result the of AccountData module
   ```

   #### Delete
   ```sh
    err = form3.DeleteByIdAndVer(accountId, version)
    // If any errors then err
    // If no reponse means your account deleted
   ```

## Run example application in development
- ### Start dependency containers
    ```sh
    make dcup
    ```
- ### Run local sample example
    ```sh
    make godev
    ```
## Application Lifecycle

   - ### Install modules & build
     ```sh
      make goinstall
      ```

   - ### Build application
     ```sh
     make goprod
     ```
## Testing Lifecycle
   > ### Make sure your docker compose should up before running below commands

   - ### Unit tests
     ```sh
     make gotest
     ```
   - ### BDD tests(Cucumber)
     ```sh
     make gobdd
     ```
     > Need [GoDog](https://github.com/cucumber/godog) set up on local.
  - ### Pact testing
     ```sh
     make gopact
     ```
     > Need [Pact](https://github.com/pact-foundation/pact-go) set up on local.
   - ### Code coverage
     ```sh
     make gocodecov
     ```

   - ### Security Scan
     ```sh
      make gosecurityscan
     ```
     > Need [Gosec](https://github.com/securego/gosec) set up on local.
## Docker Lifecycle
   - ### Docker Compose ups
     ```sh
      make dcup
     ```

   - ### Docker Compose stop
     ```sh
      make dcstop
     ```

   - ### Docker Compose down
     ```sh
     make dcdown
     ```
   - ### Docker compose for test application
     ```sh
      make dctestcompose
     ```
## Features

:white_check_mark: Containerize Application Using Docker

:white_check_mark: Unit Testing

:white_check_mark: Integration testing(BDD)

:white_check_mark: Contract testing(Pact)

:white_check_mark: Security scan

:white_check_mark: Code coverage

:white_check_mark: CI/CD

:white_check_mark: K8s Deployment files

:white_check_mark: Prototype Design pattern

### Future plan

:ballot_box_with_check: Authentication & Authorisation

:ballot_box_with_check: List api

:ballot_box_with_check: Rate limit autoretry

:ballot_box_with_check: Bulk fetch/delete/create

:ballot_box_with_check: Context with timeout for api

:ballot_box_with_check: Context for godog bdd

:ballot_box_with_check: Audit Trail logs

:ballot_box_with_check: Go routines and channels

## Author
[Amol Gaikwad - Linkedin](https://www.linkedin.com/in/gaikwadamolraj)

**I am new in GO** (Learnt basic go lang firstly and then completed assignment. Still learning process is going on)

## Screenshots
  > Added test result screenshots in /screenshots folder
  
   ![Bdd](https://github.com/gaikwadamolraj/accountapi-client/screenshots/Bdd.png)

   ![CodCov](https://github.com/gaikwadamolraj/accountapi-client/screenshots/Codecov.png)

   ![Pact](https://github.com/gaikwadamolraj/accountapi-client/screenshots/Pact.png)

   ![Security](https://github.com/gaikwadamolraj/accountapi-client/screenshots/Security.png)
