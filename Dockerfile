FROM golang:1.23 AS builder

WORKDIR /workdir

COPY go.mod ./
COPY ./main.go ./
COPY ./functions.go ./

RUN go build -o gha-summarizer github.com/VOID404/gha-summarizer

FROM scratch

WORKDIR /github/workspace

COPY --from=builder /workdir/gha-summarizer /gha-summarizer
ENTRYPOINT [ "/gha-summarizer" ]
