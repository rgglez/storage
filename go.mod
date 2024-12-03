module github.com/rgglez/storage

go 1.22

toolchain go1.23.2

require (
	github.com/kr/pretty v0.3.1
	github.com/rgglez/go-storage/services/oss/v3 v3.0.3
	github.com/rgglez/go-storage/v5 v5.0.0
	github.com/ztrue/tracerr v0.4.0
)

require (
	github.com/aliyun/aliyun-oss-go-sdk v3.0.2+incompatible // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/rgglez/go-storage/credential v1.0.0 // indirect
	github.com/rgglez/go-storage/endpoint v1.2.1 // indirect
	github.com/rogpeppe/go-internal v1.13.1 // indirect
	golang.org/x/time v0.8.0 // indirect
)

//replace (
//	github.com/rgglez/go-storage/services/oss/v3 => ../../Portafolios/go-storage/services/oss/v3
//	github.com/rgglez/go-storage/v5 => ../../Portafolios/go-storage/v5
//)
