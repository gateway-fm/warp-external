APP:=warp
APP_ENTRY_POINT:=./cmd/waaagh.go

proxy-summon:
	MallocNanoZone=0 go run -race $(APP_ENTRY_POINT) proxy-summon $(ARGS)  #example: make proxy-summon ARGS="solana" !!! list of all proxy cfgs located in config/proxy/config_storage !!!

infra-summon:
	MallocNanoZone=0 go run -race $(APP_ENTRY_POINT) infra-summon

run-mkdir:
	MallocNanoZone=0 go run -race $(APP_ENTRY_POINT) run-mkdir $(ARGS) #example: make run-mkdir ARGS="solana"




