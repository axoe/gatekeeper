// Copyright © 2018 The GΛTEKEEPER Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command

var uRegion string
var uSecret string
var uValue string
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a secret in AWS Secrets Manager",
	Long: `Update a secret in AWS Secrets Manager.

	Access to secrets is granted via fine-grained 
	AWS Identity and Access Management (IAM) policies.`,
	Run: func(cmd *cobra.Command, args []string) {

		updateSecret(cmd)

	},
}

func updateSecret(cmd *cobra.Command) {

	region := cmd.Flag("region").Value.String()
	secret := cmd.Flag("secret").Value.String()
	value := cmd.Flag("value").Value.String()
	svc := secretsmanager.New(session.New(), aws.NewConfig().WithRegion(region))
	input := &secretsmanager.PutSecretValueInput{
		SecretId:     aws.String(secret),
		SecretString: aws.String(value),
	}

	result, err := svc.PutSecretValue(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case secretsmanager.ErrCodeInvalidParameterException:
				fmt.Println(secretsmanager.ErrCodeInvalidParameterException, aerr.Error())
			case secretsmanager.ErrCodeInvalidRequestException:
				fmt.Println(secretsmanager.ErrCodeInvalidRequestException, aerr.Error())
			case secretsmanager.ErrCodeLimitExceededException:
				fmt.Println(secretsmanager.ErrCodeLimitExceededException, aerr.Error())
			case secretsmanager.ErrCodeEncryptionFailure:
				fmt.Println(secretsmanager.ErrCodeEncryptionFailure, aerr.Error())
			case secretsmanager.ErrCodeResourceExistsException:
				fmt.Println(secretsmanager.ErrCodeResourceExistsException, aerr.Error())
			case secretsmanager.ErrCodeResourceNotFoundException:
				fmt.Println(secretsmanager.ErrCodeResourceNotFoundException, aerr.Error())
			case secretsmanager.ErrCodeInternalServiceError:
				fmt.Println(secretsmanager.ErrCodeInternalServiceError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

func init() {

	updateCmd.Flags().StringVarP(&uRegion, "region", "r", "", "The region the secret is stored in")
	updateCmd.Flags().StringVarP(&uSecret, "secret", "s", "", "The name of the secret stored in AWS Secrets Manager")
	updateCmd.Flags().StringVarP(&uValue, "value", "v", "", "The value of the secret")
	rootCmd.AddCommand(updateCmd)
}
