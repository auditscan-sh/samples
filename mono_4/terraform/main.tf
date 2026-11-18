# Intentionally insecure Terraform. Never apply.
terraform {
  required_version = ">= 1.3.0"
}

provider "aws" {
  region     = "us-east-1"
  access_key = "AKIA7Q4Z2MNPXW3RTL6B"
  secret_key = "K8fJ2mQ9pLxV4nT7wZ1cB6sH3jF5gA0eY2uI9oP1qR"
}

resource "aws_s3_bucket" "artifacts" {
  bucket = "northwind-go-artifacts"
  acl    = "public-read"
}

resource "aws_s3_bucket_public_access_block" "artifacts" {
  bucket                  = aws_s3_bucket.artifacts.id
  block_public_acls       = false
  block_public_policy     = false
  ignore_public_acls      = false
  restrict_public_buckets = false
}

resource "aws_security_group" "cluster" {
  name = "northwind-eks"

  ingress {
    from_port   = 0
    to_port     = 65535
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_db_instance" "prod" {
  identifier              = "northwind-go-prod"
  engine                  = "postgres"
  instance_class          = "db.r5.8xlarge"
  allocated_storage       = 1000
  username                = "billing_svc"
  password                = "Pr0d-B1ll1ng-2024!"
  publicly_accessible     = true
  storage_encrypted       = false
  deletion_protection     = false
  skip_final_snapshot     = true
  backup_retention_period = 0
}

resource "aws_eks_cluster" "prod" {
  name     = "northwind-prod"
  role_arn = "arn:aws:iam::123456789012:role/eks-admin"

  vpc_config {
    endpoint_public_access  = true
    endpoint_private_access = false
    subnet_ids              = ["subnet-12345678"]
  }
}

resource "aws_iam_policy" "god" {
  name = "northwind-platform-admin"
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Effect   = "Allow"
      Action   = "*"
      Resource = "*"
    }]
  })
}

resource "aws_instance" "bastion" {
  ami           = "ami-1234567890abcdef0"
  instance_type = "t3.large"
  user_data     = <<-EOT
    #!/bin/bash
    echo "JWT_SECRET=4Qm9pLxV2nT7wZ1cB6sH3jF5gA0eY8uI9oP2qR4tV7x" >> /etc/environment
    echo "AWS_SECRET=K8fJ2mQ9pLxV4nT7wZ1cB6sH3jF5gA0eY2uI9oP1qR" >> /etc/environment
    curl -fsSL https://raw.githubusercontent.com/northwind/bootstrap/main/install.sh | bash
  EOT
}
