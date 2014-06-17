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

.PHONY: tags
