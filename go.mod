module github.com/nordix/meridio-operator

go 1.16

require (
	github.com/go-logr/logr v0.4.0
	github.com/nordix/meridio v0.0.0-20210820133400-65b3add684a6
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.15.0
	golang.org/x/net v0.0.0-20210614182718-04defd469f4e
	gopkg.in/yaml.v2 v2.4.0
	k8s.io/api v0.21.3
	k8s.io/apiextensions-apiserver v0.21.3
	k8s.io/apimachinery v0.21.3
	k8s.io/client-go v0.21.3
	sigs.k8s.io/controller-runtime v0.9.6
)
