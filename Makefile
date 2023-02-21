default: install

generate:
	$(info ************  Generating Provider Documentation  ************)
	go generate ./...

install:
	$(info ************  Building and Installing Binary  ************)
	go install .

test:
	$(info ************  Running Unit Tests  ************)
	go test -v -count=1 -parallel=4 ./...

testacc:
	$(info ************  Running Acceptance Tests  ************)
	$(info You must have the TF_VAR_slack_token env var set in your environment.)
	$(info You must also have a .env file in the git repo root with valid test data.)
	$(info )
	TF_ACC=1 go test -count=1 -parallel=4 -timeout 10m -v ./...
