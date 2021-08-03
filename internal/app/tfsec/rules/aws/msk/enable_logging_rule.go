package msk

// ATTENTION!
// This rule was autogenerated!
// Before making changes, consider updating the generator.

// generator-locked

import (
	"github.com/aquasecurity/tfsec/internal/app/tfsec/block"
	"github.com/aquasecurity/tfsec/internal/app/tfsec/hclcontext"
	"github.com/aquasecurity/tfsec/internal/app/tfsec/scanner"
	"github.com/aquasecurity/tfsec/pkg/provider"
	"github.com/aquasecurity/tfsec/pkg/result"
	"github.com/aquasecurity/tfsec/pkg/rule"
	"github.com/aquasecurity/tfsec/pkg/severity"
)

func init() {
	scanner.RegisterCheckRule(rule.Rule{
		Provider:  provider.AWSProvider,
		Service:   "msk",
		ShortCode: "enable-logging",
		Documentation: rule.RuleDocumentation{
			Summary:     "Ensure MSK Cluster logging is enabled",
			Explanation: `Managed streaming for Kafka can log to Cloud Watch, Kinesis Firehose and S3, at least one of these locations should be logged to`,
			Impact:      "Without logging it is difficult to trace issues",
			Resolution:  "Enable logging",
			BadExample: []string{`
resource "aws_msk_cluster" "example" {
  cluster_name           = "example"
  kafka_version          = "2.4.1"
  number_of_broker_nodes = 3

  broker_node_group_info {
    instance_type   = "kafka.m5.large"
    ebs_volume_size = 1000
    client_subnets = [
      aws_subnet.subnet_az1.id,
      aws_subnet.subnet_az2.id,
      aws_subnet.subnet_az3.id,
    ]
    security_groups = [aws_security_group.sg.id]
  }
  tags = {
    foo = "bar"
  }
}
`},
			GoodExample: []string{`
resource "aws_msk_cluster" "example" {
  cluster_name           = "example"
  kafka_version          = "2.4.1"
  number_of_broker_nodes = 3

  broker_node_group_info {
    instance_type   = "kafka.m5.large"
    ebs_volume_size = 1000
    client_subnets = [
      aws_subnet.subnet_az1.id,
      aws_subnet.subnet_az2.id,
      aws_subnet.subnet_az3.id,
    ]
    security_groups = [aws_security_group.sg.id]
  }

  logging_info {
    broker_logs {
      firehose {
        enabled         = false
        delivery_stream = aws_kinesis_firehose_delivery_stream.test_stream.name
      }
      s3 {
        enabled = true
        bucket  = aws_s3_bucket.bucket.id
        prefix  = "logs/msk-"
      }
    }
  }

  tags = {
    foo = "bar"
  }
}
`, `
resource "aws_msk_cluster" "example" {
  cluster_name           = "example"
  kafka_version          = "2.4.1"
  number_of_broker_nodes = 3

  broker_node_group_info {
    instance_type   = "kafka.m5.large"
    ebs_volume_size = 1000
    client_subnets = [
      aws_subnet.subnet_az1.id,
      aws_subnet.subnet_az2.id,
      aws_subnet.subnet_az3.id,
    ]
    security_groups = [aws_security_group.sg.id]
  }

  logging_info {
    broker_logs {
      cloudwatch_logs {
        enabled   = false
        log_group = aws_cloudwatch_log_group.test.name
      }
      firehose {
        enabled         = true
        delivery_stream = aws_kinesis_firehose_delivery_stream.test_stream.name
      }
    }
  }

  tags = {
    foo = "bar"
  }
}
`, `
resource "aws_msk_cluster" "example" {
  cluster_name           = "example"
  kafka_version          = "2.4.1"
  number_of_broker_nodes = 3

  broker_node_group_info {
    instance_type   = "kafka.m5.large"
    ebs_volume_size = 1000
    client_subnets = [
      aws_subnet.subnet_az1.id,
      aws_subnet.subnet_az2.id,
      aws_subnet.subnet_az3.id,
    ]
    security_groups = [aws_security_group.sg.id]
  }

  logging_info {
    broker_logs {
      cloudwatch_logs {
        enabled   = true
        log_group = aws_cloudwatch_log_group.test.name
      }
      firehose {
        enabled         = false
        delivery_stream = aws_kinesis_firehose_delivery_stream.test_stream.name
      }
      s3 {
        enabled = true
        bucket  = aws_s3_bucket.bucket.id
        prefix  = "logs/msk-"
      }
    }
  }

  tags = {
    foo = "bar"
  }
}
`},
			Links: []string{
				"https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/msk_cluster#",
			},
		},
		RequiredTypes: []string{
			"resource",
		},
		RequiredLabels: []string{
			"aws_msk_cluster",
		},
		DefaultSeverity: severity.Medium,
		CheckFunc: func(set result.Set, resourceBlock block.Block, _ *hclcontext.Context) {

			if cwLogAttr := resourceBlock.GetNestedAttribute("logging_info.broker_logs.cloudwatch_logs.enabled"); cwLogAttr.IsNotNil() && cwLogAttr.IsTrue() {
				return
			}

			if firehoseLogAttr := resourceBlock.GetNestedAttribute("logging_info.broker_logs.firehose.enabled"); firehoseLogAttr.IsNotNil() && firehoseLogAttr.IsTrue() {
				return
			}

			if s3LogAttr := resourceBlock.GetNestedAttribute("logging_info.broker_logs.s3.enabled"); s3LogAttr.IsNotNil() && s3LogAttr.IsTrue() {
				return
			}

			set.AddResult().
				WithDescription("Resource '%s' does not have any broker logging enabled", resourceBlock.FullName())
		},
	})
}