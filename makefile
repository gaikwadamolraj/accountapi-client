# Set API_HOST env variable
export API_HOST=http://localhost:8080

# Set up go
goinstall:
	go mod download

gomod:
	GO111MODULE=on go mod tidy
	GO111MODULE=on go mod vendor

#  Testing
gotest: 
	go test  ./form3/...  -coverprofile=coverage.out

gobdd:
	cd integration && godog

gopact:
	go test -v ./contract-testing/pact_test.go

# Scan
gosecurityscan:
	gosec -exclude-dir=vendor ./...

gocodecov:
	make gotest
	go tool cover -html=coverage.out

#  local
godev:
	go run ./form3-client.go

#  PROD 
goprod:
	go build ./*.go

goprodrun:
	make build
	./form3-client

# Docker compose 
dcup:
	docker-compose -f docker-compose.yml up --build

dcstop:
	docker-compose -f docker-compose.yml stop

dcdown:
	docker-compose -f docker-compose.yml down

# Docker test compose 
dctestcompose:
	make dcdown
	docker-compose -f docker-compose-test.yml up --build
	docker-compose -f docker-compose-test.yml down --volumes