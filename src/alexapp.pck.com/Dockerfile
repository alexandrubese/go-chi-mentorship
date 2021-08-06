FROM golang:latest

COPY . /go/src/alexapp.pck.com/

WORKDIR /go/src/alexapp.pck.com/

RUN go get gopkg.in/mgo.v2
#RUN go get github.com/oxequa/realize
RUN go get github.com/go-chi/chi
RUN go get github.com/go-chi/cors
RUN go get github.com/go-chi/jwtauth
RUN go get github.com/go-chi/render
RUN go get github.com/dgrijalva/jwt-go
#RUN PATH=/go/bin/linux_arm:$PATH
# this means go/app/src/alexaap.pck.com/app/app.go

WORKDIR /go/src/alexapp.pck.com/app/
RUN env GOOS=linux go install

#WORKDIR /go/src/alexapp.pck.com/
WORKDIR /go/bin/

EXPOSE 1323

CMD ["./go/bin/"]
