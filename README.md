# mstock-api

## How to use

First of all, you need to have Golang installed. If you need help with this, you can check these [docs](https://golang.org/doc/).

## Running the server

Build the executable:

`go build main.go`

Run it:

`./main`

## Using the online service

Run:

`alias gcurl='curl --header "Authorization: Bearer $(gcloud auth print-identity-token)"'`

## To locally run and test docker image

Install [docker](https://docs.docker.com/install/) and run

PORT=8080 docker build --tag mstock .

## To build and a deploy a new container without having to trigger a new build

`gcloud builds submit . --config ./ci/cloudbuild.yaml --substitutions=COMMIT_SHA="[something here]"`
