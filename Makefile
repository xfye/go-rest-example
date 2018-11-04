default: src/rest-example/cmd/main.go
	@echo "开始编译"
	@rm -rf build/ && mkdir -p build/bin/ && \
	go build src/rest-example/cmd/main.go
	mv main build/bin/

.PHONY : clean
clean: 
	@echo "开始清理"
	rm -rf build
