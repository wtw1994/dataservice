module github.com/DataWorkbench/dataservice

go 1.15

replace github.com/DataWorkbench/gproto => D:\github\gproto
replace github.com/DataWorkbench/common => D:\github\common

require (
	github.com/DataWorkbench/common v0.0.0-20211025072126-7ab59e5fdb2b
	github.com/DataWorkbench/glog v0.0.0-20210809050640-4960fd6de6ab
	github.com/DataWorkbench/gproto v0.0.0-20211124042600-7a50397fa864
	github.com/DataWorkbench/loader v0.0.0-20210801212447-9ef10fa3f297
	github.com/go-playground/validator/v10 v10.4.1
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/spf13/cobra v0.0.3
	github.com/spf13/pflag v1.0.5 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
	gorm.io/gorm v1.21.12
)
