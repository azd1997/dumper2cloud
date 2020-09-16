default:
	go build -o ./bin/d2c .;
	@echo 'ok';

test:
	go test .