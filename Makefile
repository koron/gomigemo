SUBDIRS = \
	./conv \
	./dict \
	./inflator \
	./readutil \
	./migemo

test:
	go test $(SUBDIRS)
