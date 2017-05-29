FROM golang:1.8

ADD ./ /go/src/tournament_server
# Govendor for vendoring
RUN go get -u github.com/kardianos/govendor
WORKDIR /go/src/tournament_server
RUN govendor sync

#EXPOSE 8000

#ENV PORT 8000
#ENV USER gopher
#ENV PASSWORD gopher
#ENV DATABASE social_tournament
#
#CMD ["go", "run", "server.go"]