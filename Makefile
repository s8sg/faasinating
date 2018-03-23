all: build deploy

build:        
	./build.sh

deploy:
	./deploy.sh $(APIKEY) 

clean:
	./remove.sh
