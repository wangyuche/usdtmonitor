build_protoc:
		rm -rf ./proto/tronproto && \
        cp -R ~/go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis/google ./proto/tron && \
		mkdir proto/tronproto && \
		cd proto/tron && \
		protoc **/*.proto  --go_out=../tronproto \
		--go_opt=paths=source_relative --proto_path=. \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative && \
		protoc **/**/*.proto  --go_out=../tronproto \
		--go_opt=paths=source_relative --proto_path=. \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative 