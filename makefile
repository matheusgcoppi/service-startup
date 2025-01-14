# Check to see if we can use ash, in Alpine images, or default to BASH.
SHELL_PATH = /bin/ash
SHELL = $(if $(wildcard $(SHELL_PATH)),/bin/ash,/bin/bash)

run:
	go run api/service/sales/main.go | go run api/tooling/logfmt/main.go

# ======================================================================================================================
# Define dependencies

GOLANG          := golang:1.23
ALPINE          := alpine:3.21
KIND            := kindest/node:v1.32.0
POSTGRES        := postgres:17.2
GRAFANA         := grafana/grafana:11.4.0
PROMETHEUS      := prom/prometheus:v3.0.0
TEMPO           := grafana/tempo:2.6.0
LOKI            := grafana/loki:3.3.0
PROMTAIL        := grafana/promtail:3.3.0

KIND_CLUSTER    := ardan-starter-cluster
NAMESPACE       := sales-system
SALES_APP       := sales
AUTH_APP        := auth
BASE_IMAGE_NAME := localhost/ardanlabs
VERSION         := 0.0.1
SALES_IMAGE     := $(BASE_IMAGE_NAME)/$(SALES_APP):$(VERSION)
METRICS_IMAGE   := $(BASE_IMAGE_NAME)/metrics:$(VERSION)
AUTH_IMAGE      := $(BASE_IMAGE_NAME)/$(AUTH_APP):$(VERSION)

# ======================================================================================================================
# Building containers

build: sales

sales:
	docker build \
		-f zarf/docker/Dockerfile.sales \
		-t ${SALES_IMAGE} \
		--build-arg BUILD_REF=$(VERSION) \
		--build-arg BUILD_DATE=$(date -u +"%Y-%m-%dT%H:%M:%SZ") \
		.



# ======================================================================================================================
# Running from within k8s/kind
dev-up:
	kind create cluster \
		--image $(KIND) \
		--name $(KIND_CLUSTER) \
		--config zarf/k8s/dev/kind-config.yaml

	kubectl wait --timeout=120s --namespace=local-path-storage --for=condition=Available deployment/local-path-provisioner

dev-down:
	kind delete cluster --name $(KIND_CLUSTER)

dev-status-all:
	kubectl get nodes -o wide
	kubectl get svc -o wide
	kubectl get pods -o wide --watch --all-namespaces

dev-status:
	watch -n 2 kubectl get pods -o wide --all-namespaces

# ======================================================================================================================

dev-load:
	kind load docker-image $(SALES_IMAGE) --name $(KIND_CLUSTER)

dev-apply:
	kustomize build zarf/k8s/dev/sales | kubectl apply -f -
	kubectl wait pods --namespace=$(ANMESPACE) --selector app=$(SALES_APP) --timeout=120s --for=condition=Ready

# ======================================================================================================================
# Modules support
tidy:
	go mod tidy
	go mod vendor
