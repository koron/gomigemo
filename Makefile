SUBDIRS = \
	./conv \
	./dict \
	./inflator \
	./readutil \
	./migemo

test:
	go test $(SUBDIRS)

tags:
	ctags -R $(SUBDIRS)

bindata: migemo/bindata.go

migemo/bindata.go: _dict/*
	go-bindata -o $@ -pkg="migemo" -prefix="./_dict" ./_dict

.PHONY: tags
