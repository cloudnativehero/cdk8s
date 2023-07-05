package main

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

func NewRedis(scope constructs.Construct, id string) cdk8s.Chart {

	chart := cdk8s.NewChart(scope, jsii.String(id), nil)

	cdk8s.NewHelm(chart, jsii.String("mysql"), &cdk8s.HelmProps{
		Chart:   jsii.String("bitnami/mysql"),
		Version: jsii.String("9.10.4"),
		Values: &map[string]interface{}{
			"auth": map[string]string{
				"database": "mysqldb",
			},
		},
	})

	return chart
}

func main() {
	app := cdk8s.NewApp(nil)
	NewRedis(app, "mysql")
	app.Synth()
}
