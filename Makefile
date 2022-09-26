all: ad2usb ad2logs install

ad2usb:
	$(MAKE) -C cmd/ad2usb

ad2logs:
	$(MAKE) -C cmd/ad2logs

install:
	-mkdir bin
	cp cmd/ad2usb/ad2usb bin/
	cp cmd/ad2logs/ad2logs bin/

clean:
	$(MAKE) -C cmd/ad2usb clean
	$(MAKE) -C cmd/ad2logs clean
	rm bin/*

