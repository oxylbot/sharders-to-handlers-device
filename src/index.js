const zmq = require("zeromq");

// PULL from sharders, continue to PUSH to handlers
const pull = zmq.socket("pull");
const push = zmq.socket("push");
pull.connect(process.env.FROM_SHARDERS);
push.bind(process.env.TO_HANDLERS);

pull.on("message", msg => push.send(msg));
