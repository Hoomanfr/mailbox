FROM golang:1.20-alpine as build
# Setup private git repo
ARG GITHUB_USERNAME
ARG GITHUB_TOKEN
RUN apk add --no-cache git
RUN git config --global url."https://$GITHUB_USERNAME:$GITHUB_TOKEN@github.com/".insteadOf "https://git@github.com/"
RUN echo "machine github.com login ${GITHUB_USERNAME} password ${GITHUB_TOKEN}" > ~/.netrc
RUN chmod 600 ~/.netrc
RUN export GOPRIVATE=github.com/Hoomanfr
# Build go project
WORKDIR /mailbox
COPY ./mailbox ./mailbox
COPY go.mod go.sum ./
RUN go mod download && go mod verify
WORKDIR /mailbox/mailbox/cmd
ENV GIN_MODE=release
RUN go build -o mailbox-srv
# Final image
FROM golang:1.20-alpine
COPY --from=build /mailbox /mailbox
WORKDIR /mailbox/mailbox/cmd
EXPOSE 8080
CMD [ "./mailbox-srv" ]

# Checkout https://vsupalov.com/build-docker-image-clone-private-repo-ssh-key/
# ENTRYPOINT ["tail", "-f", "/dev/null"] # for just keep the container up and running to be able to exec into it

# for building image run: 
# docker build -f <path-to-docker-file> --build-arg <arg1-name>=<arg1-value> --build-arg <arg2-name>=<arg2-value> -t <image-name> <path-to-building-context-directory>
# e.g.  docker build -f ./mailbox/builds/Dockerfile -t mailbox .
# e.g.  docker build -f ./mailbox/builds/Dockerfile --build-arg GITHUB_USERNAME=hoomanfr --build-arg GITHUB_TOKEN="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" -t mailbox .