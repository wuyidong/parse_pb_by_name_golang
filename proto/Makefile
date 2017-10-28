PROTO_DIR = ./
PROTOC = protoc --proto_path=$(PROTO_DIR)

RPC = rpc
RPC_PROTO = rpc.proto
ACCOUNT = account
ACCOUNT_PROTO = account.proto

all : $(RPC) $(ACCOUNT)

$(RPC) :
	@mkdir -p $(RPC); \
	$(PROTOC) --go_out=$(RPC) $(PROTO_DIR)/$(RPC_PROTO)

$(ACCOUNT):
	@mkdir -p $(ACCOUNT); \
	$(PROTOC) --go_out=$(ACCOUNT) $(PROTO_DIR)/$(ACCOUNT_PROTO); \
	sed -i "s/^.*import rpc.*$$/import rpc \"pbtest\/proto\/rpc\"/" $(ACCOUNT)/$(ACCOUNT).pb.go

clean :
	@rm -rf rpc/ account/
