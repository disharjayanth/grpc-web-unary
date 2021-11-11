<h1>Run envoy proxy (to convert http1 client request to http2 grpc request):</h1>

docker run -d -v "$(pwd)"/envoy.yaml:/etc/envoy/envoy.yaml:ro  -p 8080:8080 -p 9901:9901 envoyproxy/envoy:v1.17.0

<h1>gRPC client code generation:</h1>
protoc -I=$DIR echo.proto \
    --js_out=import_style=commonjs:$OUT_DIR \
    --grpc-web_out=import_style=commonjs,mode=grpcwebtext:$OUT_DIR

$DIR => directory in which .proto file is present
$OUT_DIR => where to put grpc-web-client generated code

Above command creates two files _pb.js and _grpc_web_pb.js

<h1>Webpack command:</h1>
npx webpack client/client.js 


