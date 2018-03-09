package main

import (
	"flag"
	"github.com/albertocubeddu/aws-secrets/strategies"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"log"
	"strings"
)

func main() {
	var container = map[string]map[string]string{}

	modulesArg := flag.String("modules", "database", "All the modules to import separated by a ',' e.g. database,elasticsearch")
	strategyToUse := flag.String("output", "screen,file", "The various output method ['screen', 'file']")
	//environmnetArg := flag.String("environment", "testing", "The environment that you want to call [production/testing/etc.]")
	flag.Parse()

	modules := strings.Split(*modulesArg, ",")
	strategiesToUse := strings.Split(*strategyToUse, ",")
	for _, value := range modules {
		ExportVariables("/"+value+"/", "", container)
	}

	for _, value := range strategiesToUse {
		operation := fakeFactoryOperator(value)
		apply := StrategyOperation{operation}
		apply.Operate(container)
	}
}

//The idea in the future is to have a proper factory that allow to select multiple strategy to use!
func fakeFactoryOperator(strategyToUse string) Operator {
	switch strategyToUse {
	case "screen":
		return new(strategies.OutputScreen)
	case "file":
		return new(strategies.OutputFile)
	default:
		return new(strategies.OutputScreen)
	}
}

func ExportVariables(path string, nextToken string, container map[string]map[string]string) {

	session := session.Must(session.NewSession(&aws.Config{Region: aws.String("ap-southeast-2")}))
	client := ssm.New(session)

	input := &ssm.GetParametersByPathInput{
		Path:           &path,
		WithDecryption: aws.Bool(true),
		Recursive:      aws.Bool(true),
	}

	if nextToken != "" {
		input.SetNextToken(nextToken)
	}

	output, err := client.GetParametersByPath(input)

	if err != nil {
		log.Panic(err)
	}

	for _, element := range output.Parameters {
		PrintExportParameter(path, element, container)
	}

	if output.NextToken != nil {
		ExportVariables(path, *output.NextToken, container)
	}
}

func PrintExportParameter(path string, parameter *ssm.Parameter, container map[string]map[string]string) {
	name := *parameter.Name
	value := *parameter.Value

	key := strings.Trim(name[len(path):], "/")
	splitted := strings.Split(key, "/")

	if len(splitted) != 2 {
		panic("[ERROR] Wrong Format: please use: /name/environment/key")
	}

	if _, present := container[splitted[0]]; !present {
		container[splitted[0]] = make(map[string]string)
	}

	//Create the container with all the variables
	container[splitted[0]][splitted[1]] = value
}
