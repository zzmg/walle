list:
	@echo "sync-codes: 同步代码"
	@echo "image: 制作镜像"
	@echo "build: 编译可执行文件"
	@echo "upload: 上传本地镜像"
	@echo "publish: 同步最新代码，制作镜像并上传"
sync-codes:
	git pull
build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o server ./main.go
image: build
	@echo "start to make docker image..."
	docker build --tag="d-hub.wallstcn.com:5000/wallstreetcn/autoDelUser:latest" .
upload:
	@echo "start to upload..."
	docker push d-hub.wallstcn.com:5000/wallstreetcn/autoDelUser:latest
upload-qcloud:
	@echo "upload to qcloud"
	docker login --username=3518936217 ccr.ccs.tencentyun.com
	docker build --tag ccr.ccs.tencentyun.com/dhub.wallstcn.com/autoDelUser:latest .
	docker push ccr.ccs.tencentyun.com/dhub.wallstcn.com/autoDelUser:latest
publish: sync-codes image upload
deps:
	mkdir -p /tmp/govendor/src/gitlab.wallstcn.com/wscnbackend
	@if [ "$(CI_COMMIT_REF_NAME)" = "master" ]; then\
		echo "checkout ivankastd:master";\
		git clone -b sit git@gitlab.wallstcn.com:wscnbackend/ivankastd.git /tmp/govendor/src/gitlab.wallstcn.com/wscnbackend/ivankastd;\
		git clone git@gitlab.wallstcn.com:wscnbackend/govendor.git /tmp/govendor_temp;\
        git clone -b sit git@gitlab.wallstcn.com:wscnbackend/ivankaprotocol.git /tmp/govendor/src/gitlab.wallstcn.com/wscnbackend/ivankaprotocol;\
	else\
		echo "checkout ivankastd:sit";\
		git clone -b sit git@gitlab.wallstcn.com:wscnbackend/ivankastd.git /tmp/govendor/src/gitlab.wallstcn.com/wscnbackend/ivankastd;\
		git clone  git@gitlab.wallstcn.com:wscnbackend/govendor.git /tmp/govendor_temp;\
	    git clone -b sit git@gitlab.wallstcn.com:wscnbackend/ivankaprotocol.git /tmp/govendor/src/gitlab.wallstcn.com/wscnbackend/ivankaprotocol;\
        fi
	cp -r /tmp/govendor_temp/vendor/* /tmp/govendor/src
	mkdir -p /tmp/govendor/bin
	mkdir -p /go/src/cradle/
	cp -R "/builds/operation/$(SERVICE_NAME)" "/go/src/cradle/$(SERVICE_NAME)/"
test:
	go test
