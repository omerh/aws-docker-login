package main

import (
	"encoding/base64"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Missing region argument")
		os.Exit(1)
	}
	region := os.Args[1]
	token, endpoint := getEcrToken(region)
	token = cleanToken(token)
	endpoint = cleanURL(endpoint)
	fmt.Printf("docker login -u AWS -p %v %v", token, endpoint)
}

func getEcrToken(region string) (string, string) {
	config := aws.NewConfig().WithRegion(region)
	session, _ := session.NewSession(config)
	svc := ecr.New(session, config)
	input := &ecr.GetAuthorizationTokenInput{}
	// retrive result from aws api
	result, _ := svc.GetAuthorizationToken(input)
	// get token from result
	token := *result.AuthorizationData[0].AuthorizationToken
	t, _ := base64.StdEncoding.DecodeString(token)
	// get endpoint from result
	endpoint := *result.AuthorizationData[0].ProxyEndpoint

	return string(t), endpoint
}

func cleanURL(address string) string {
	// get the registry host from url
	t, _ := url.Parse(address)
	return t.Host
}

func cleanToken(token string) string {
	// clean the token
	if strings.HasPrefix(token, "AWS:") {
		token = strings.Trim(token, "AWS:")
	}
	return token
}
