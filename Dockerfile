FROM golang:alpine AS build

ADD . /go/src/mdNote
WORKDIR /go/src/mdNote

RUN apk add --no-cache git dep npm build-base
RUN rm -rf /var/cache/apk/*

RUN dep ensure -vendor-only
RUN go vet

WORKDIR /go/src/mdNote/MdNote

RUN npm install
RUN npm run build

WORKDIR /go/src/mdNote

RUN go build -a -o mdNote

FROM scratch
COPY --from=build /go/src/mdNote /mdNote
COPY --from=build /go/src/mdNote/MdNote/dist /MdNote/dist
ENV PORT $PORT
ENV DATABASE_URL $DATABASE_URL

CMD [ "/mdNote" ]