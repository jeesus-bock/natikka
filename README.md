# NATS experimentations


## What's this about?
[NATS.io](https://nats.io/) – Cloud Native, Open Source, High-performance Messaging is a messaging and persistence system I've heard a lot good things about. This is a repo with some basic client functionality implemented in Go.

## NATS server

### Install the server to localhost
See https://github.com/nats-io/nats-server
1) `wget https://github.com/nats-io/nats-server/releases/download/v2.8.4/nats-server-v2.8.4-amd64.deb`
2) `sudo dpkg --install nats-server-v2.8.4-amd64.deb`

### Install the cli tool nats

See https://github.com/nats-io/natscli
1) `wget https://github.com/nats-io/natscli/releases/download/v0.0.33/nats-0.0.33-386.deb`
2) `sudo dpkg --install nats-0.0.33-386.deb`

### Run the server
Opens the monitoring at 127.0.0.1:8222
`nats-server -DV -m 8222 -user foo -pass bar`

### Run natikka
`natikka publish` and `natikka consume`, tmux recommended :)
