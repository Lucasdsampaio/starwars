build-qa:
	docker image build -t lucassampaio/starwars:qa .

push-qa:
	docker image push lucassampaio/starwars:qa

build-prod:
	docker image build -t lucassampaio/starwars:prod .

push-prod:
	docker image push lucassampaio/starwars:prod

