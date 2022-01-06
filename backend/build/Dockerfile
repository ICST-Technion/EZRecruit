# Stage 1: Use image builder to build the target binaries
FROM golang:1.17 AS builder

ARG COMPONENT
WORKDIR /workspace/${COMPONENT}
COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY Makefile Makefile
COPY cmd/ cmd/
COPY pkg/ pkg/
COPY datatypes/ datatypes/
COPY queries/ queries/
RUN make build

# Stage 2: Copy the binaries from the image builder to the base image
FROM registry.access.redhat.com/ubi8/ubi-minimal:latest
ARG COMPONENT
ENV USER_UID=1001
ENV USER_NAME=${COMPONENT}

# install operator binary
COPY --from=builder /workspace/${COMPONENT}/bin/${COMPONENT} /usr/local/bin/manager

# allow user to fully control the /usr/local/bin/resume directory (save resume files)
COPY web-templates/upload_resume.html /usr/local/bin/upload_resume.html
RUN mkdir /usr/local/bin/resume
RUN chown ${USER_UID}:0 /usr/local/bin/resume
RUN chmod 0755 /usr/local/bin/resume

COPY build/scripts/user_setup /usr/local/scripts/user_setup
RUN  /usr/local/scripts/user_setup

RUN microdnf update && \
    microdnf clean all

USER ${USER_UID}
ENTRYPOINT ["/usr/local/bin/manager"]
