SUBDIRS = \
	./conv \
	./dict \
	./inflator \
	./proto \
	./readutil

test:
	go test $(SUBDIRS)
