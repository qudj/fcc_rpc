regenerate:
	cd models/proto; protoc --gogofaster_out=plugins=grpc:. fcc.proto; mkdir ../fcc_serv; mv fcc.pb.go ../fcc_serv;