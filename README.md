[comment]:  # ( Copyright Contributors to the Open Cluster Management project )

# EZRecruit

## Backend Hierarchy

### REST-API Server [/pkg/rest-api](pkg/rest-api/server.go)
The server uses [gin-gotonic](github.com/gin-gonic) pkg to run a RESTful API listener. 
The supported endpoints can be found in registerAPI function.

### DB [/pkg/db](pkg/db/db.go)
With the access restrictions in mind, we designed our backend to be DB-agnostic: the RESTAPI server does not care 
about the specifics of the DB, it only requires a client that implements the [db.go](pkg/db/db.go) interface.

For the time being we chose to go with a low-effort [in-memory database](pkg/db/in-memory/in_memory_db.go) 
implementation.

### Data-types [/datatypes](datatypes/datatypes.go)
The [datatypes.go](datatypes/datatypes.go) file contains the structures used in our endpoints (e.g. Job-Listing).

### Queries [/queries](queries)
The queries directory contains datatype-related query formats, e.g. URL query parameters for fetching job listings 
(in [job_listing.go](queries/job_listing.go)).

## Frontend
The [frontend](frontend) folder contains the code from Wix.

## Build

* [Makefile](Makefile)
* [Dockerfile](build/Dockerfile)
* [build scripts](build/scripts)

1.  Set the `REGISTRY` environment variable to hold the name of your docker registry:
    ```
    $ export REGISTRY=...
    ```

1.  Run the following command to build and push the image:
    ```
    make push-images
    ```
    
## Formatting

* [GCI](https://github.com/daixiang0/gci) for ordering imports.
* [gofumpt](https://github.com/mvdan/gofumpt) for formatting (a stricter tool than `go fmt`).
* `go fmt`

## Linting

* `go vet`
* [golangci-lint](https://github.com/golangci/golangci-lint), minimal version 1.43.0, the settings file is [.golangci.yaml](https://github.com/open-cluster-management/hub-of-hubs-spec-sync/blob/main/.golangci.yaml).
* [golint](https://github.com/golang/lint)

ℹ️ If you want to specify something as false-positive, use the [//nolint](https://golangci-lint.run/usage/false-positives/) comment.

ℹ️ If you see stale errors from [golangci-lint](https://github.com/golangci/golangci-lint), run `golangci-lint cache clean`.

## Tests

We did not implement any unit/e2e tests for this POC. 
At t is the developer's responsibility to build/test their code before merging.