# aws-docker-login

After the deprection of the command `get-login --no-include-email` in awscli version [1.7.10](https://github.com/aws/aws-cli/blob/8330afe7caef89a4012200a6cac7acdc53f5474e/CHANGELOG.rst#11710), I still miss this usfull command.

So, I wrote this short code to have this command in hand.

To use this:

```bash
# Install
go get -u github.com/omerh/aws-docker-login

# Run
$(aws-docker-login eu-west-2)
```
