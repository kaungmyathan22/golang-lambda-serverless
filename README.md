### Use this to create iam role for that lambda function
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



{
    "Role": {
        "Path": "/",
        "RoleName": "lambda-ex",
        "RoleId": "AROARTTEOM2DXVKHTXX4B",
        "Arn": "arn:aws:iam::110806066823:role/lambda-ex",
        "CreateDate": "2024-02-24T05:03:58+00:00",
        "AssumeRolePolicyDocument": {
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
        }
    }
}

aws lambda create-function \
  --function-name github-com-kaungmyathan22-golang-lambda-serverless \
  --runtime go1.x \
  --role arn:aws:iam::110806066823:role/lambda-ex \
  --handler main \
  --zip-file fileb://./dist/bootstrap.zip

aws lambda update-function-code --function-name  gosimple --zip-file fileb://./dist/bootstrap.zip

------------
```
{
    "FunctionName": "go-lambda-serverless",
    "FunctionArn": "arn:aws:lambda:eu-west-2:110806066823:function:go-lambda-serverless",
    "Runtime": "go1.x",
    "Role": "arn:aws:iam::110806066823:role/lambda-ex",
    "Handler": "main",
    "CodeSize": 5350565,
    "Description": "",
    "Timeout": 3,
    "MemorySize": 128,
    "LastModified": "2024-02-24T05:15:20.195+0000",
    "CodeSha256": "2kJNk19uSRrAs7sMZhi5IEh0j09UW8nPh96LPlKAorg=",
    "Version": "$LATEST",
    "TracingConfig": {
        "Mode": "PassThrough"
    },
    "RevisionId": "8098f2c6-82d4-4fcb-9722-f8a9d1d77a2e",
    "State": "Pending",
    "StateReason": "The function is being created.",
    "StateReasonCode": "Creating",
    "PackageType": "Zip",
    "Architectures": [
        "x86_64"
    ],
    "EphemeralStorage": {
        "Size": 512
    },
    "SnapStart": {
        "ApplyOn": "None",
        "OptimizationStatus": "Off"
    },
    "RuntimeVersionConfig": {
        "RuntimeVersionArn": "arn:aws:lambda:eu-west-2::runtime:30052276b0b7733e82eddf1f0942de1022c7dfbc0ca93cfc121c868194868dec"
    },
    "LoggingConfig": {
        "LogFormat": "Text",
        "LogGroup": "/aws/lambda/go-lambda-serverless"
    }
}
```

aws lambda invoke \
  --function-name github-com-kaungmyathan22-golang-lambda-serverless \
  --payload '{"What is your name?":"Kaung Myat Han","How old are you?":22}' \
  --cli-binary-format raw-in-base64-out \
  output.txt