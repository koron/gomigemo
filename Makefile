SUBDIRS = \
	./conv \
	./dict \
	./embedict \
	./inflator \
	./readutil \
	./migemo

DICTDIR = ./_dict

test:
	go test $(SUBDIRS)

bench:
	go test -bench . $(SUBDIRS)

tags:
	ctags -R $(SUBDIRS)

bindata: embedict/bindata.go

embedict/bindata.go: $(DICTDIR)/*
	go-bindata -o $@ -nomemcopy -pkg="embedict" -prefix="$(DICTDIR)" $(DICTDIR)

.PHONY: tags
