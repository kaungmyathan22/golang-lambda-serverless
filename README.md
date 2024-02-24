## Running go on aws lambda
### 1. Create iam role for the lambda function
```bash
aws iam create-role \
    --role-name lambda-ex \
    --assume-role-policy-document '{
        "Version": "2012-10-17",
        "Statement": [
            {
                "Effect": "Allow",
                "Principal": {
                    "Service": "lambda.amazonaws.com"
                },
                "Action": "sts:AssumeRole"
            }
        ]
    }'
```
### 2. Build go code
```bash
GOOS=linux GOARCH=amd64 go build -o bootstrap main.go
```
### 3. Create go lambda function
```bash
aws lambda create-function \
  --function-name  gosimple\
  --runtime go1.x \
  --role arn-no:role/your-created-role \
  --handler bootstrap \
  --zip-file fileb://./bootstrap.zip
```

### 3. In case you updated code and want to redeploy
```bash
aws lambda update-function-code --function-name  gosimple --zip-file fileb://./dist/bootstrap.zip
```

### 4. Function executing or invocation.
```bash
aws lambda invoke \
  --function-name gosimple \
  --payload file://testdata/event.json \
  --cli-binary-format raw-in-base64-out \
  output.txt
```