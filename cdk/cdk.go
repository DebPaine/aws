package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	// "github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	_ "github.com/joho/godotenv/autoload"
)

type AwsStackProps struct {
	awscdk.StackProps
}

func main() {
	createAwsStack()
}

func createAwsStack() {
	defer jsii.Close()
	app := awscdk.NewApp(nil)
	newAwsStack(app, "AwsStack", &AwsStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})
	app.Synth(nil)
}

func newAwsStack(scope constructs.Construct, id string, props *AwsStackProps) awscdk.Stack {
	var stackProps awscdk.StackProps
	if props != nil {
		stackProps = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &stackProps)

	// example resource
	// queue := awssqs.NewQueue(stack, jsii.String("AwsQueue"), &awssqs.QueueProps{
	// 	VisibilityTimeout: awscdk.Duration_Seconds(jsii.Number(300)),
	// })

	return stack
}

func env() *awscdk.Environment {
	return &awscdk.Environment{
		Account: jsii.String("CDK_DEFAULT_ACCOUNT"),
		Region:  jsii.String("CDK_DEFAULT_REGION"),
	}
}
