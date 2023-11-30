module bj38web/service/user

go 1.18

require (
	github.com/golang/protobuf v1.5.3
	github.com/micro/go-micro v1.18.0
	google.golang.org/protobuf v1.31.0
)

require (
	github.com/BurntSushi/toml v0.3.1 // indirect
	github.com/cheekybits/genny v1.0.0 // indirect
	github.com/coreos/etcd v3.3.17+incompatible // indirect
	github.com/coreos/go-systemd v0.0.0-20190719114852-fd7a80b32e1f // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/go-log/log v0.1.0 // indirect
	github.com/gogo/protobuf v1.2.1 // indirect
	github.com/golang/mock v1.6.0 // indirect
	github.com/google/btree v1.0.0 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/json-iterator/go v1.1.8 // indirect
	github.com/lucas-clemente/quic-go v0.13.1 // indirect
	github.com/marten-seemann/chacha20 v0.2.0 // indirect
	github.com/marten-seemann/qtls v0.4.1 // indirect
	github.com/micro/cli v0.2.0 // indirect
	github.com/micro/mdns v0.3.0 // indirect
	github.com/miekg/dns v1.1.43 // indirect
	github.com/mitchellh/hashstructure v1.0.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/nats-io/jwt v0.3.0 // indirect
	github.com/nats-io/nats.go v1.9.1 // indirect
	github.com/nats-io/nkeys v0.1.0 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/onsi/gomega v1.27.6 // indirect
	github.com/oxtoacart/bpool v0.0.0-20190530202638-03653db5a59c // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/stretchr/testify v1.8.1 // indirect
	go.uber.org/atomic v1.5.0 // indirect
	go.uber.org/multierr v1.3.0 // indirect
	go.uber.org/tools v0.0.0-20190618225709-2cfd321de3ee // indirect
	go.uber.org/zap v1.12.0 // indirect
	golang.org/x/crypto v0.4.0 // indirect
	golang.org/x/lint v0.0.0-20191125180803-fdd1cda4f05f // indirect
	golang.org/x/mod v0.10.0 // indirect
	golang.org/x/net v0.10.0 // indirect
	golang.org/x/sys v0.8.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0 // indirect
	golang.org/x/tools v0.9.1 // indirect
	google.golang.org/genproto v0.0.0-20230306155012-7f2fa6fef1f4 // indirect
	google.golang.org/grpc v1.54.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	honnef.co/go/tools v0.0.1-2019.2.3 // indirect
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.27.1
