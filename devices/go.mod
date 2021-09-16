module github.com/acidprime/go-zebra-scanner/devices

require (
	github.com/acidprime/go-zebra-scanner/snapi v0.0.0-20210916190706-c85d127cea5a
	github.com/google/gousb a9eea03
	github.com/sirupsen/logrus v1.1.1
)

replace github.com/acidprime/go-zebra-scanner/snapi => ../snapi

replace github.com/google/gousb => ../vendor/github.com/google/gousb
