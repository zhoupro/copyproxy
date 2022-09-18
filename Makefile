VERSION=$(shell git describe --tags)

build:
	go build -ldflags "-s -w -X github.com/zhoupro/roperate.Version=$(VERSION)"

install:
	go install -ldflags "-s -w -X github.com/zhoupro/roperate.Version=$(VERSION)"

release:
	gox --arch 'amd64' --os 'windows linux darwin' --output "dist/{{.Dir}}_{{.OS}}_{{.Arch}}/{{.Dir}}" -ldflags "-s -w -X github.com/zhoupro/roperate.Version=$(VERSION)"
	zip      pkg/roperate_windows_amd64.zip   dist/roperate_windows_amd64/roperate.exe -j
	tar zcvf pkg/roperate_linux_amd64.tar.gz   dist/roperate_linux_amd64/roperate
	tar zcvf pkg/roperate_darwin_amd64.tar.gz  dist/roperate_darwin_amd64/roperate

clean:
	rm -rf dist/
	rm -f pkg/*.tar.gz pkg/*.zip
