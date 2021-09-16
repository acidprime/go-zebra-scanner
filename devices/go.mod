module github.com/acidprime/go-zebra-scanner/devices

require (
	github.com/acidprime/go-zebra-scanner/snapi v0.0.0-20210916191512-8fd01cee8ae9
	github.com/elemecca/gousb v0.0.0-20210916192740-a9eea035a91c
	github.com/sirupsen/logrus v1.1.1
)

replace github.com/acidprime/go-zebra-scanner/snapi => ../snapi

replace github.com/google/gousb => ../vendor/github.com/google/gousb
