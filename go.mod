module alert

go 1.16

require (
	github.com/apache/thrift v0.16.0
	github.com/cloudwego/kitex v0.1.4
	gorm.io/driver/mysql v1.3.2
	gorm.io/gorm v1.23.1
)

replace github.com/apache/thrift => github.com/apache/thrift v0.13.0
