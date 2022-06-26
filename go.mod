module rutvej/Splitpay_go

go 1.16

go env -w GO111MODULE=off

require (
	github.com/google/uuid v1.3.0
	github.com/gorilla/mux v1.8.0
	github.com/jinzhu/gorm v1.9.16
	github.com/tidwall/gjson v1.14.1
	github.com/tidwall/sjson v1.2.4
)
