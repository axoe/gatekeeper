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

// delCmd represents the del command

var dSecret string
var dRegion string
var delCmd = &cobra.Command{
	Use:   "del",
	Short: "Delete a secret from AWS Secrets Manager",
	Long: `Delete a secret from AWS Secrets Manager.

	Access to secrets is granted via fine-grained 
	AWS Identity and Access Management (IAM) policies.`,
	Run: func(cmd *cobra.Command, args []string) {

		delSecret(cmd)

	},
}

func delSecret(cmd *cobra.Command) {

	region := cmd.Flag("region").Value.String()
	secret := cmd.Flag("secret").Value.String()
	svc := secretsmanager.New(session.New(), aws.NewConfig().WithRegion(region))
	input := &secretsmanager.DeleteSecretInput{
		RecoveryWindowInDays: aws.Int64(7),
		SecretId:             aws.String(secret),
	}

	result, err := svc.DeleteSecret(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case secretsmanager.ErrCodeResourceNotFoundException:
				fmt.Println(secretsmanager.ErrCodeResourceNotFoundException, aerr.Error())
			case secretsmanager.ErrCodeInvalidParameterException:
				fmt.Println(secretsmanager.ErrCodeInvalidParameterException, aerr.Error())
			case secretsmanager.ErrCodeInvalidRequestException:
				fmt.Println(secretsmanager.ErrCodeInvalidRequestException, aerr.Error())
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
	getCmd.Flags().StringVarP(&dRegion, "region", "r", "", "The region the secret is stored in")
	getCmd.Flags().StringVarP(&dSecret, "secret", "s", "", "The name of the secret stored in AWS Secrets Manager")
	rootCmd.AddCommand(delCmd)
}
