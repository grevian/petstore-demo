.PHONY: generate

generate:
	docker run --rm -v  ${CURDIR}:/local openapitools/openapi-generator-cli generate -i /local/petstore-openapi.yaml -g go-server --additional-properties=addResponseHeaders=true,outputAsLibrary=true,onlyInterfaces=true,router=chi -o /local/api


