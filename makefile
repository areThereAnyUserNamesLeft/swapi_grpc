# protos
mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
project_root := $(patsubst %/,%,$(dir $(mkfile_path)))
proto_src_root:=$(project_root)/protos

clean-protos:
	rm -rf $(proto_src_root)/*

check-protoc:
ifeq (, $(shell which protoc))
	error "protoc not installed"
else
	@echo protoc √
endif

proto-src-remove:
	@rm -rf $(proto_src_root)

define clone_proto
	@echo "Cloning $(1)"
	@-git clone --quiet --depth 1 https://github.com/$(1) \
		$(proto_src_root)/github.com/$(1)/$(2)

	@if [ -n "$2" ]; then \
		git -C $(proto_src_root)/github.com/$(1)/$(2) filter-branch --subdirectory-filter $(2); \
	fi
endef

define clone_uw_proto
	$(call clone_proto,"areThereAnyUserNamesLeft/$(1)",$(2))
endef

download-protos: check-protoc proto-src-remove
	@-mkdir -p $(proto_src_root)
	$(call clone_uw_proto,swapi_contract)

generate: 
	for V in films people planets species starships vehicles ; do \
		mkdir -p generated/$$V/ ; \
		protoc --plugin=protoc-gen-go=$(HOME)/go/bin/protoc-gen-go \
		-I ./protos/github.com/areThereAnyUserNamesLeft/swapi_contract/ \
		--go_out=plugins=grpc:./generated/$$V/  $$V.proto  ; \
	done

