PROTOBF = $(shell find ./ -name '*.proto')
#OBJECTS = $(subst proto,pb.go,$(PROTOBF))
OBJECTS = center.pb.go

TARGET_PATH=./

default: $(OBJECTS)
	@echo -n

%.pb.go: %.proto
	protoc --go_out=$(TARGET_PATH) $<
