# QH

https://grpc.io/docs/languages/go/quickstart/

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=require_unimplemented_servers=false:. --go-grpc_opt=paths=source_relative internal/prompt/proto/teamcubot.proto


$ cat ./* | xclip -selection clipboard

$ go build -o ./cmd/crawler/crawler  -v ./cmd/crawler/

$ ./cmd/crawler/crawler crawl http://quotes.toscrape.com

$ go run ./cmd/crawler/crawler.go ./cmd/crawler/crawler_launcher.go crawl http://quotes.toscrape.com

$ go run main.go crawl http://quotes.toscrape.com


$ go generate mockgen -source=./../service.go -destination=../../../mocks/service_mock.go -package=mocks

$ go generate mockgen -source=./internal/url-lister/crawler/htmlparser_adapter.go -destination=./internal/url-lister/crawler/mocks/htlmparser_adapeter_mock.go -package=mocks

