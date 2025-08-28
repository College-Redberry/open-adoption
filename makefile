# Carrega variáveis do .env
ifneq (,$(wildcard .env))
	include .env
	export $(shell sed 's/=.*//' .env)
endif

container_run:
	podman compose down
	podman compose build
	podman compose up
