module github.com/acidprime/go-zebra-scanner/devices

require (
	github.com/acidprime/go-zebra-scanner/snapi v0.0.0-20210916191512-8fd01cee8ae9
	github.com/google/gousb/v2 v2.1.0
	github.com/sirupsen/logrus v1.1.1
)

replace github.com/acidprime/go-zebra-scanner/snapi => ../snapi

replace github.com/google/gousb => ../vendor/github.com/google/gousb
