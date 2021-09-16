module github.com/acidprime/go-zebra-scanner/devices

require (
	github.com/acidprime/go-zebra-scanner/snapi v0.0.0-20210916191512-8fd01cee8ae9
	github.com/google/gousb v1.1.1
	github.com/sirupsen/logrus v1.1.1
)

replace github.com/acidprime/go-zebra-scanner/snapi => ../snapi

replace github.com/google/gousb => github.com/Elemecca/gousb v0.0.0-20181019140747-a9eea035a91c
