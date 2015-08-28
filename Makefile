
build: jarvis

jarvis: deps main.go
	go build github.com/mhoc/jarvis

deps:
	go get gopkg.in/yaml.v2
	go get github.com/gorilla/websocket
	go get gopkg.in/redis.v3

clean:
	rm jarvis

run: build
	./jarvis
