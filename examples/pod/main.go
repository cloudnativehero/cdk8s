package main

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/cloudnativehero/cdk8s/examples/pod/imports/k8s"
)

type MyChartProps struct {
	cdk8s.ChartProps
}

func NewMyChart(scope constructs.Construct, id string, props *MyChartProps) cdk8s.Chart {
	var cprops cdk8s.ChartProps
	if props != nil {
		cprops = props.ChartProps
	}
	chart := cdk8s.NewChart(scope, jsii.String(id), &cprops)

	// define resources here

	k8s.NewKubePod(chart, jsii.String("my-pod"), &k8s.KubePodProps{
		Spec: &k8s.PodSpec{
			Containers: &[]*k8s.Container{{
				Name:  jsii.String("mypod"),
				Image: jsii.String("nginx"),
				Ports: &[]*k8s.ContainerPort{{ContainerPort: jsii.Number(80)}},
			}},
		},
	})
	return chart
}

func main() {
	app := cdk8s.NewApp(nil)
	NewMyChart(app, "pod", nil)
	app.Synth()
}
