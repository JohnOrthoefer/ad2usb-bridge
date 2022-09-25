SRC=build.go const.go main.go mqtt.go options.go status.go types.go debug.go
GOLANG=/usr/bin/go
CURL=/usr/bin/curl
GIT=/usr/bin/git
REPONAME=$(shell basename `git rev-parse --show-toplevel`)
DOCKERREPO=${REPONAME}
VERSION=v1.0-beta
SHA1=$(shell git rev-parse --short HEAD)
NOW=$(shell date +%Y-%m-%d_%T)

ad2usb: ${SRC}
	echo ${REPONAME}
	${GOLANG} build \
		-o ad2usb \
		-ldflags "-X main.sha1ver=${SHA1} \
		-X main.buildTime=${NOW} \
		-X main.repoName=${REPONAME}"

update-go:
	${GOLANG} get github.com/alecthomas/kong
	${GOLANG} get github.com/eclipse/paho.mqtt.golang
	${GOLANG} mod tidy

clean:
	rm -f ad2usb

