module github.com/temporalio/reference-app-orders-go

go 1.22.0

toolchain go1.22.2

require (
	buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go v1.35.2-20241127180247-a33202765966.1
	connectrpc.com/connect v1.17.0
	github.com/bufbuild/protovalidate-go v0.7.3
	github.com/cludden/protoc-gen-go-temporal v1.14.3
	github.com/glebarez/sqlite v1.11.0
	github.com/google/go-cmp v0.6.0
	github.com/mitchellh/go-homedir v1.1.0
	github.com/oklog/run v1.1.0
	github.com/prometheus/client_golang v1.11.0
	github.com/stretchr/testify v1.9.0
	github.com/uber-go/tally/v4 v4.1.1
	github.com/urfave/cli/v2 v2.27.5
	go.temporal.io/api v1.40.0
	go.temporal.io/sdk v1.30.1
	go.temporal.io/sdk/contrib/tally v0.2.0
	golang.org/x/net v0.28.0
	golang.org/x/sync v0.8.0
	google.golang.org/protobuf v1.35.2
	gorm.io/gorm v1.25.12
)

require (
	cel.dev/expr v0.18.0 // indirect
	github.com/Jeffail/gabs/v2 v2.7.0 // indirect
	github.com/OneOfOne/xxhash v1.2.8 // indirect
	github.com/PaesslerAG/gval v1.2.2 // indirect
	github.com/PaesslerAG/jsonpath v0.1.1 // indirect
	github.com/antlr4-go/antlr/v4 v4.13.0 // indirect
	github.com/aymerick/douceur v0.2.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bwmarrin/snowflake v0.3.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/clbanning/mxj/v2 v2.7.0 // indirect
	github.com/cludden/benthos/v4 v4.24.0-no-otel.2 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.5 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/facebookgo/clock v0.0.0-20150410010913-600d898af40a // indirect
	github.com/generikvault/gvalstrings v0.0.0-20180926130504-471f38f0112a // indirect
	github.com/glebarez/go-sqlite v1.21.2 // indirect
	github.com/go-faker/faker/v4 v4.2.0 // indirect
	github.com/gofrs/uuid v4.4.0+incompatible // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/mock v1.7.0-rc.1 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/cel-go v0.22.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/css v1.0.1 // indirect
	github.com/gosimple/slug v1.13.1 // indirect
	github.com/gosimple/unidecode v1.0.1 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.22.0 // indirect
	github.com/itchyny/timefmt-go v0.1.5 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/klauspost/compress v1.17.3 // indirect
	github.com/klauspost/pgzip v1.2.6 // indirect
	github.com/matoous/go-nanoid/v2 v2.0.0 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/microcosm-cc/bluemonday v1.0.26 // indirect
	github.com/ncruces/go-strftime v0.1.9 // indirect
	github.com/nexus-rpc/sdk-go v0.0.12 // indirect
	github.com/oklog/ulid v1.3.1 // indirect
	github.com/pborman/uuid v1.2.1 // indirect
	github.com/pierrec/lz4/v4 v4.1.18 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.26.0 // indirect
	github.com/prometheus/procfs v0.6.0 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect
	github.com/rickb777/date v1.20.5 // indirect
	github.com/rickb777/plural v1.4.1 // indirect
	github.com/robfig/cron v1.2.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/segmentio/ksuid v1.0.4 // indirect
	github.com/shopspring/decimal v1.3.1 // indirect
	github.com/stoewer/go-strcase v1.3.0 // indirect
	github.com/stretchr/objx v0.5.2 // indirect
	github.com/tilinna/z85 v1.0.0 // indirect
	github.com/twmb/murmur3 v1.1.5 // indirect
	github.com/vmihailenco/msgpack/v5 v5.4.1 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	github.com/xeipuuv/gojsonpointer v0.0.0-20190905194746-02993c407bfb // indirect
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415 // indirect
	github.com/xeipuuv/gojsonschema v1.2.0 // indirect
	github.com/xrash/smetrics v0.0.0-20240521201337-686a1a2994c1 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	golang.org/x/exp v0.0.0-20240325151524-a685a6edb6d8 // indirect
	golang.org/x/sys v0.24.0 // indirect
	golang.org/x/text v0.17.0 // indirect
	golang.org/x/time v0.4.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240827150818-7e3bb234dfed // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240827150818-7e3bb234dfed // indirect
	google.golang.org/grpc v1.66.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	modernc.org/libc v1.55.3 // indirect
	modernc.org/mathutil v1.6.0 // indirect
	modernc.org/memory v1.8.0 // indirect
	modernc.org/sqlite v1.31.1 // indirect
)
