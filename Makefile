generate-dao:
	go run generator/gen_dao.go

generate-api:
	docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate -i  /local/api/openapi.yaml -g go-gin-server -o /local --additional-properties=interfaceOnly=true,packageName=generated,apiPath=generated