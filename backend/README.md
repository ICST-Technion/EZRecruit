[comment]:  # ( Copyright Contributors to the Open Cluster Management project )

# EZRecruit - REST-API Backend

## Hierarchy

### REST-API Server [/pkg/rest-api](pkg/rest-api/server.go)

The server uses [gin-gotonic](github.com/gin-gonic) pkg to run a RESTful API listener. 
The supported API can be found in registerAPI function.

### DB [/pkg/db](pkg/db/db.go)

With the access restrictions in mind, we designed our backend to be DB-agnostic: the RESTAPI server does not care 
about the specifics of the DB, it only requires a client that implements the [db.go](pkg/db/db.go) interface.

For the time being we chose to go with a low-effort [in-memory database](pkg/db/in-memory/in_memory_db.go) 
implementation.

## Building and deploying

Refer to the [Makefile](Makefile)

1.  Set the `REGISTRY` environment variable to hold the name of your docker registry:
    ```
    $ export REGISTRY=...
    ```

1.  Run the following command to build and push the image:
    ```
    make push-images
    ```