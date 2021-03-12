module github.com/dmorgan81/btt-expl

go 1.16

require (
	github.com/dmorgan81/go-btt v0.0.0-20210312205432-54fc360b4d0a
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.1.3
	github.com/spf13/viper v1.7.1
)

replace github.com/dmorgan81/go-btt => ../
