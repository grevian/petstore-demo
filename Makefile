.PHONY: generate

generate:
	docker run --rm -v  ${CURDIR}:/local openapitools/openapi-generator-cli generate -i /local/petstore-openapi.yaml -g go-server --additional-properties=addResponseHeaders=true,router=chi -o /local/api
	# The generated demo service and module configuration confuse the compiler, and we're not using them, so remove them
	del api\main.go api\go.mod

