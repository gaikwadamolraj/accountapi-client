gosec -exclude-dir=vendor ./...  
cd integration
godog
cd ../form3 
export API_HOST="http://accountapi:8080"
go test ./... -coverprofile=coverage.out
# cd ..
# cwd=$(pwd)
# cd $GOPATH
# curl -fsSL https://raw.githubusercontent.com/pact-foundation/pact-ruby-standalone/master/install.sh | bash
# export PATH=$PATH:$GOPATH/pact/bin
# go install github.com/pact-foundation/pact-go@v1
# cd $cwd
# go test -v ./contract-testing/pact_test.go