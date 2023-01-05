package main

import (
	"bot_huan/core/server"
	pb_base "bot_huan/proto/steam/base"
	base "bot_huan/service/steam/service/base"
	"google.golang.org/grpc/reflection"
)

func main() {
	s, lis := server.NewServer(8089)
	pb_base.RegisterBaseServiceServer(s, &base.BaseService{})
	reflection.Register(s)
	s.Serve(lis)
}
