.PHONY: generate

generate:
	docker run --rm -v  ${CURDIR}:/local openapitools/openapi-generator-cli generate -i /local/petstore-openapi.yaml -g go-gin-server -o /local/api
	# The generated code includes an example service and Dockerfile which we don't need, and which confuse the compiler, so remove them
	del api\main.go api\Dockerfile
