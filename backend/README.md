protoc -I. -I/home/user/workspace/go//src/mairpc/proto 
--go_out=
Mproto/account/account.proto=mairpc/proto/proto/account,
Mproto/account/service.proto=mairpc/proto/proto/account,
Mproto/account/user.proto=mairpc/proto/proto/account,
Mvendor/github.com/grpc-ecosystem/grpc-gateway/internal/errors.proto=mairpc/proto/vendor/github.com/grpc-ecosystem/grpc-gateway/internal,
Mvendor/google.golang.org/grpc/reflection/grpc_reflection_v1alpha/reflection.proto=mairpc/proto/vendor/google.golang.org/grpc/reflection/grpc_reflection_v1alpha,
plugins=grpc+tag:. 
./proto/account/account.proto 
./proto/account/service.proto 
./proto/account/user.proto n

protoc -I. -I/home/user/workspace/go//src/mairpc/proto 
--go_out=
Mproto/account/account.proto=mairpc/proto/proto/account,
Mproto/account/service.proto=mairpc/proto/proto/account,
Mproto/account/user.proto=mairpc/proto/proto/account,
Mvendor/github.com/grpc-ecosystem/grpc-gateway/internal/errors.proto=mairpc/proto/vendor/github.com/grpc-ecosystem/grpc-gateway/internal,
Mvendor/google.golang.org/grpc/reflection/grpc_reflection_v1alpha/reflection.proto=mairpc/proto/vendor/google.golang.org/grpc/reflection/grpc_reflection_v1alpha,
plugins=grpc+tag:. 
./proto/account/account.proto 
./proto/account/service.proto 
./proto/account/user.proto n




protoc -I. -I=/Users/alomerry/workspace/bot-huan/backend/proto/ --go_out=Mproto/account/user.proto=bot_huan/proto/account ./proto/account/account.proto ./proto/account/service.proto 

protoc -I. -I/Users/alomerry/workspace/bot-huan/backend/proto/ --go_out=Mproto/account/user.proto=bot_huan/proto/account ./proto/account/account.proto ./proto/account/service.proto 

protoc -I. --go_out=Mproto/account/user.proto=bot_huan/proto/account,Mproto/account/service.proto=bot_huan/proto/account ./proto/account/user.proto ./proto/account/service.proto 

protoc --proto_path=src \
  --go_opt=Mprotos/buzz.proto=example.com/project/protos/fizz \
  --go_opt=Mprotos/bar.proto=example.com/project/protos/foo \
  protos/buzz.proto protos/bar.proto

protoc --proto_path=src --go_out=out --go_opt=paths=source_relative foo.proto bar/baz.proto

protoc -I. --go_out=Mproto/account/user.proto=bot_huan/proto/account,Mproto/account/service.proto=bot_huan/proto/account ./proto/account/user.proto ./proto/account/service.proto 

protoc -I. --go_out=Mproto/account/service.proto=bot_huan/proto/account ./proto/account/service.proto 