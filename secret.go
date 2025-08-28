package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager/types"
)

type secretsManager struct {
	client *secretsmanager.Client
	config *configuration
}

func getKeyName(key string) (string, error) {
	parts := strings.SplitN(key, ":", 3)

	if len(parts) != 3 {
		return "", fmt.Errorf("unable to parse secret name, %s", key)
	}

	return parts[2], nil

}

func (s *secretsManager) Init(appConfig *configuration) {

	if s.client != nil {
		fmt.Println("Secrets Manager was already init")
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithSharedConfigProfile(appConfig.Profile))
	if err != nil {
		log.Fatal(err)
	}

	s.client = secretsmanager.NewFromConfig(cfg)
	s.config = appConfig

}

func (s *secretsManager) ListKeys(env string, secretsList *[]string, nextToken *string) error {
	input := &secretsmanager.ListSecretsInput{
		Filters: []types.Filter{
			{
				Key:    types.FilterNameStringTypeTagKey,
				Values: []string{"env"},
			},
			{
				Key:    types.FilterNameStringTypeTagValue,
				Values: []string{env},
			},
			{
				Key:    types.FilterNameStringTypeTagKey,
				Values: []string{"app"},
			},
			{
				Key:    types.FilterNameStringTypeTagValue,
				Values: []string{s.config.App},
			},
		},
	}

	if nextToken != nil {
		input.NextToken = nextToken
	}

	secrets, err := s.client.ListSecrets(context.TODO(),
		input,
	)

	if err != nil {
		return err

	}

	for _, v := range secrets.SecretList {

		name, err := getKeyName(*v.Name)

		if err != nil {
			return err
		}

		*secretsList = append(*secretsList, name)
	}

	// recurstion break case
	if secrets.NextToken == nil {
		return nil
	}

	return s.ListKeys(env, secretsList, secrets.NextToken)

}

func (s *secretsManager) GetValues(env string, secretsMap map[string]string, nextToken *string) error {

	input := &secretsmanager.BatchGetSecretValueInput{
		Filters: []types.Filter{
			{
				Key:    types.FilterNameStringTypeTagKey,
				Values: []string{"env"},
			},
			{
				Key:    types.FilterNameStringTypeTagValue,
				Values: []string{env},
			},
			{
				Key:    types.FilterNameStringTypeTagKey,
				Values: []string{"app"},
			},
			{
				Key:    types.FilterNameStringTypeTagValue,
				Values: []string{s.config.App},
			},
		},
	}

	secrets, err := s.client.BatchGetSecretValue(context.TODO(), input)

	if err != nil {
		return err
	}

	for _, secret := range secrets.SecretValues {
		name, err := getKeyName(*secret.Name)

		if err != nil {
			return err
		}

		secretsMap[name] = *secret.SecretString
	}

	// recurstion break case
	if secrets.NextToken == nil {
		return nil
	}

	return s.GetValues(env, secretsMap, secrets.NextToken)

}

func (s *secretsManager) Get(env string, key string) (string, error) {
	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(s.config.App + ":" + env + ":" + key),
	}

	secret, err := s.client.GetSecretValue(context.TODO(), input)

	if err != nil {

		return "", err
	}

	return *secret.SecretString, nil

}

func (s *secretsManager) Set(env string, key string, value string) error {
	app := s.config.App
	input := &secretsmanager.CreateSecretInput{
		Name:         aws.String(s.config.App + ":" + env + ":" + key),
		SecretString: aws.String(value),
		Tags: []types.Tag{
			{Key: aws.String("env"), Value: aws.String(env)},
			{Key: aws.String("app"), Value: aws.String(app)},
		},
	}

	_, err := s.client.CreateSecret(context.TODO(), input)

	return err

}
func (s *secretsManager) Update(env string, key string, value string) error {
	input := &secretsmanager.UpdateSecretInput{
		SecretId:     aws.String(s.config.App + ":" + env + ":" + key),
		SecretString: aws.String(value),
	}

	_, err := s.client.UpdateSecret(context.TODO(), input)

	return err

}
func (s *secretsManager) Remove(env string, key string) error {
	input := &secretsmanager.DeleteSecretInput{
		SecretId: aws.String(s.config.App + ":" + env + ":" + key),
	}

	_, err := s.client.DeleteSecret(context.TODO(), input)

	return err
}
