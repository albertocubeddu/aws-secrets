# AWS-SECRETS
Generate your .secret (for serverless) with a CLI touch using the AWS's parameter store.

## Usage
./aws-secrets --help

* environment string
        The environment that you want to call [production/testing/etc.] (default "") -> All environments
* modules string
        All the modules to import separated by a ',' e.g. database,elasticsearch (default "database")
* output string
        The various output method ['screen', 'file'] (default "screen,file")

## Example
All the parameters need to have a name like this
--name /\<\<environment>>/\<\<product>>/\<\<key>>

aws ssm put-parameter --name /prod/database/DB_HOST --value "Username" --type SecureString --key-id "alias/aws/ssm" --region us-west-2
aws ssm put-parameter --name /demo/database/DB_HOST --value "Demo-Username" --type SecureString --key-id "alias/aws/ssm" --region us-west-2