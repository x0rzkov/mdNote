FROM golang:latest
MAINTAINER KimMachineGun <geon0250@gmail.com>

RUN apt-get update && apt-get upgrade -y
RUN apt-get install curl -y
RUN curl -sL https://deb.nodesource.com/setup_6.x | bash -
RUN apt-get install -y nodejs
# RUN apt-get install -y npm

RUN go get -u github.com/golang/dep/cmd/dep

ADD . /go/src/mdNote

WORKDIR /go/src/mdNote

RUN dep ensure
RUN go build -i

WORKDIR /go/src/mdNote/MdNote

RUN npm install
RUN npm run build

WORKDIR /go/src/mdNote

ENV PORT $PORT
ENV DATABASE_URL $DATABASE_URL

CMD [ "./mdNote" ]