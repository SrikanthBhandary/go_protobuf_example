Download the binaries from the following link

     https://github.com/protocolbuffers/protobuf/releases/tag/v3.11.4 
    
Add the bin folder to the path
            
     export PATH={path_to_protoc}:$PATH

Add GOPATH to the path aswell

    export PATH=$GOPATH:$PATH

Download the protoc-gen-go
  
    go get -u github.com/golang/protobuf/protoc-gen-go

Install the protoc-gen-go using the make command.

Example to generate the go output:
    
    protoc -I=/Users/apple/Desktop/go_understanding/src/github.com/srikanthbhandary/go_protobuf_example/proto --go_out=. /Users/apple/Desktop/go_understanding/src/github.com/srikanthbhandary/go_protobuf_example/proto/address.proto
 
 How to Run?
 
    go run main.go addressbook.data
    
 How to add data?
 
    go run add_person.go addressbook.data       

