test:
	go test ./...
swagger:
	swagger generate spec -o spec.yaml
serve-docs:
	swagger serve spec.yaml
serve-swagger:
	swagger serve spec.yaml --flavor=swagger