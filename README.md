# GoAWS

Simple Go programme to display ec2 attributes. It uses the AWS Go SDK to query AWS account and then renders html table to display the results.
I use it at work when i am asked to provide a summary of deployed ec2s.

## Prerequisite

You must have your AWS credentials (access keys) present in your credentials file (refer to AWS documentations).

## To run

Install Go

Install AWS Go SDK:

```
go get github.com/aws/aws-sdk-go
```

Clone this repo:

```
go get github.com/Abdul2/GoAWS
```

run:

```
go run listinstances.go

```

Don't forget to update HTTML to reflect your organisation.

## Screenshot

![screenshot](https://github.com/Abdul2/GoAWS/raw/master/documentation/screenshot.png)
