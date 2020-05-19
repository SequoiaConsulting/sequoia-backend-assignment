FROM golang:latest 

RUN mkdir /trello
ADD . /trello/ 
WORKDIR /trello 
RUN go build -o main .
EXPOSE 8000
CMD ["/main"]