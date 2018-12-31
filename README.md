# zmq-proxy[![Build Status](https://travis-ci.com/oxylbot/zmq-proxy.svg?branch=master)](https://travis-ci.com/oxylbot/zmq-proxy)
ZMQ device to redirect sharder message traffic to message-handlers. Mainly for binding purposes.


### Socket types

These are the potential values of `INCOMING_TYPE` and `OUTGOING_TYPE` which determine the behviour of the sockets.

| Socket type | Integer |
|-------------|---------|
| ZMQ_PAIR    | 0       |
| ZMQ_PUB     | 1       |
| ZMQ_SUB     | 2       |
| ZMQ_REQ     | 3       |
| ZMQ_REP     | 4       |
| ZMQ_DEALER  | 5       |
| ZMQ_ROUTER  | 6       |
| ZMQ_PULL    | 7       |
| ZMQ_PUSH    | 8       |
| ZMQ_XPUB    | 9       |
| ZMQ_XSUB    | 10      |
| ZMQ_STREAM  | 11      |
| ZMQ_SERVER  | 12      |
| ZMQ_CLIENT  | 13      |
| ZMQ_RADIO   | 14      |
| ZMQ_DISH    | 15      |

### Usage

1. Build project either using Docker or `go build`
2. Set all the required environment variables:
    - `INCOMING_ADDRESS`
    - `OUTGOING_ADDRESS`
    - `INCOMING_TYPE` (see above)
    - `OUTGOING_TYPE` (see above)
3. Execute the binary
4. Sockets will be listening until program is killed.

### Requirements
* libcmzq
* libzmq
* libsodium