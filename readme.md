#AWS-SECRETS
Generate your .secret (for serverless) with a CLI touch!

#Example
All the parameters need to have a name like this
--name /<<environment>>/<<product>>/<<key>>

aws ssm put-parameter --name /prod/database/DB_HOST --value "Username" --type SecureString --key-id "alias/aws/ssm" --region us-west-2