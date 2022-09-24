all: ad2usb install

ad2usb:
	$(MAKE) -C cmd/ad2usb

install:
	-mkdir bin
	cp cmd/ad2usb/ad2usb bin/

clean:
	$(MAKE) -C cmd/ad2usb clean
	rm bin/*

