module github.com/acidprime/go-zebra-scanner/devices

require (
	github.com/acidprime/go-zebra-scanner/snapi
	github.com/google/gousb v0.0.0
	github.com/sirupsen/logrus v1.1.1
)

replace github.com/acidprime/go-zebra-scanner/snapi => ../snapi

replace github.com/google/gousb => ../vendor/github.com/google/gousb
