BUILDPATH=$(CURDIR)/output
GO=$(shell which go)
GOINSTALL=$(GO) install
GOCLEAN=$(GO) clean
GOGET=$(GO) get
EXENAME=cmd

export GOPATH=$(CURDIR)/output

tree:
	@echo "start building tree..."
	@if [ ! -d $(BUILDPATH)/bin ] ; then mkdir -p $(BUILDPATH)/bin ; fi
	@if [ ! -d $(BUILDPATH)/pkg ] ; then mkdir -p $(BUILDPATH)/pkg ; fi

get:
	@echo "getting packages..."
	@$(GOGET) "github.com/aws/aws-sdk-go/aws"
	@$(GOGET) "github.com/aws/aws-sdk-go/aws/awserr"
	@$(GOGET) "github.com/aws/aws-sdk-go/aws/session"
	@$(GOGET) "github.com/aws/aws-sdk-go/service/secretsmanager"
	@$(GOGET) "github.com/spf13/cobra"

build:
	@echo "start building..."
	$(GOINSTALL) $(EXENAME)
	@echo "Yay! all DONE!"

clean:
	@rm -rf $(BUILDPATH)
	@echo "looks fresh as a daisy!"

all: tree get build
