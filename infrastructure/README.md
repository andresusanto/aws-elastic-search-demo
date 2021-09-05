```
terraform init -backend-config "bucket=andres-tf-states" --backend-config "region=ap-southeast-2"

terraform apply -auto-approve -var-file="envs/prod.tfvars"

terraform plan -var-file="envs/prod.tfvars"
```
