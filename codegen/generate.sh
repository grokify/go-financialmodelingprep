export GO_POST_PROCESS_FILE="/usr/local/go/bin/gofmt -s -w"

curl -XGET 'https://raw.githubusercontent.com/openapis/api-specs/master/financialmodelingprep/v3/financialmodelingprep-v3_openapi-v3.1.0.yaml' > 'openapi-spec_v3.yaml'

java -jar openapi-generator-cli.jar generate -i openapi-spec_v3.json -g go -o client --package-name client

#java -jar openapi-generator-cli.jar generate -i partial-specs_v3.0.0/openapi-spec_teams.json -g go -o engagedigital --package-name engagedigital

#echo "\n\nfunc (apiClient *APIClient) HTTPClient() *http.Client { return apiClient.cfg.HTTPClient }" >> engagedigital/client.go
gofmt -s -w client/*.go
