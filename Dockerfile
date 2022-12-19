FROM golang

WORKDIR /app

ARG CACHEBUST=1

RUN git clone -b develop https://PondXIV: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@github.com/PondXIV/BackEndCoach.git .

RUN go mod download

ENV GIN_MODE=release

COPY ./ ./

RUN go install github.com/cosmtrek/air@latest

RUN rm .air.toml
# Don't forget to add .air.toml .gitignore
RUN air init

# 9775-9780-DailyWorkout
EXPOSE 8080

CMD ["air"]