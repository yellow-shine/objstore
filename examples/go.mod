module example.com/m

go 1.22.7

toolchain go1.23.7

require (
	github.com/go-kit/log v0.2.1
	github.com/thanos-io/objstore v0.0.0-20250317105316-a0136a6f898d

)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/clbanning/mxj v1.8.4 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/efficientgo/core v1.0.0-rc.0.0.20221201130417-ba593f67d2a4 // indirect
	github.com/go-logfmt/logfmt v0.5.1 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/mitchellh/mapstructure v1.4.3 // indirect
	github.com/mozillazg/go-httpheader v0.2.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/prometheus/client_golang v1.17.0 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/prometheus/common v0.44.0 // indirect
	github.com/prometheus/procfs v0.11.1 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common v1.0.1129 // indirect
	github.com/tencentyun/cos-go-sdk-v5 v0.7.40 // indirect
	golang.org/x/sync v0.11.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	google.golang.org/protobuf v1.36.4 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace github.com/thanos-io/objstore => github.com/yellow-shine/objstore v0.0.0-20250325120849-1036980b7796

replace github.com/tencentyun/cos-go-sdk-v5 => github.com/yellow-shine/cos-go-sdk-v5 v0.0.0-20250324084202-fd0081c12590
