# NATS experimentations

## NATS server

### Install the server to localhost
See https://github.com/nats-io/nats-server
1) `wget https://github.com/nats-io/nats-server/releases/download/v2.8.4/nats-server-v2.8.4-amd64.deb`
2) `sudo dpkg --install nats-server-v2.8.4-amd64.deb`

### Install the cli tool nats

See https://github.com/nats-io/natscli
1) `wget https://github.com/nats-io/natscli/releases/download/v0.0.33/nats-0.0.33-386.deb`
2) `sudo dpkg --install nats-0.0.33-386.deb`

### Run
`nats-server -DV -m 8222 -user foo -pass bar`