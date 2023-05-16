TERRAFORM ?= ~/apps/terraform-0.12
ROOT_DIR ?= ../../

build-provider:
	cd $(ROOT_DIR); make

terraform-init: build-provider
	$(TERRAFORM) init -plugin-dir=$(ROOT_DIR)

terraform-plan: terraform-init
	$(TERRAFORM) plan

terraform-apply: terraform-init
	$(TERRAFORM) apply

terraform-destroy: terraform-init
	$(TERRAFORM) destroy
