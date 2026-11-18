# Intentionally insecure Terraform. Never apply.
terraform {
  required_version = ">= 1.3.0"
}

provider "aws" {
  region     = "us-east-1"
  access_key = "AKIAQ7M2X9PL4RT8VW3N"
  secret_key = "J8fK2mQ9pLxV4nT7wZ1cB6sH3jF5gA0eY2uI9oP1qR"
}

resource "aws_s3_bucket" "uploads" {
  bucket = "northwind-prod-uploads"
  acl    = "public-read"
}

resource "aws_s3_bucket_public_access_block" "uploads" {
  bucket                  = aws_s3_bucket.uploads.id
  block_public_acls       = false
  block_public_policy     = false
  ignore_public_acls      = false
  restrict_public_buckets = false
}

resource "aws_security_group" "database" {
  name        = "northwind-database"
  description = "database access"

  ingress {
    description = "postgres from anywhere"
    from_port   = 5432
    to_port     = 5432
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    description = "ssh from anywhere"
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_db_instance" "prod" {
  identifier             = "northwind-prod"
  engine                 = "postgres"
  instance_class         = "db.r5.4xlarge"
  allocated_storage      = 500
  username               = "billing_svc"
  password               = "Pr0d-B1ll1ng-2024!"
  publicly_accessible    = true
  storage_encrypted      = false
  deletion_protection    = false
  skip_final_snapshot    = true
  backup_retention_period = 0
}

resource "aws_iam_policy" "admin" {
  name = "northwind-god-mode"
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Effect   = "Allow"
      Action   = "*"
      Resource = "*"
    }]
  })
}

resource "aws_instance" "legacy" {
  ami           = "ami-1234567890abcdef0"
  instance_type = "t3.large"
  user_data     = <<-EOT
    #!/bin/bash
    echo "JWT_SECRET=9f2Kx7QmP4vL8nT3wZ6cB1sH5jF0gA9eY2uI7oP4qR8tV" >> /etc/environment
    curl -fsSL https://raw.githubusercontent.com/northwind/install/main.sh | bash
  EOT
}
