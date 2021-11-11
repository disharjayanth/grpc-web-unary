const {EchoRequest, EchoResponse} = require('./echo_pb.js');
const {EchoServiceClient} = require('./echo_grpc_web_pb.js');

var echoService = new EchoServiceClient('http://localhost:8080');

var request = new EchoRequest();
request.setMessage('Hello World!');

echoService.echo(request, {}, function(err, response) {
  if (err) {
      console.log("Error code:",err.code)
      console.log("Error message",err.message )
  } else {
    response.array.forEach((res) => {
      console.log("reponse:",res)
    })
  }
});