all: dist

build:
	go mod download && CGO_ENADLED=0 GOOS=linux go build -o votesapp-0.1
dist: build
	mkdir votesapp
	cp votesapp-0.1 votesapp/
	cp file.txt votesapp/
	rm -rf votesapp-0.1
	
