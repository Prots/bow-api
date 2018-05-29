# defining environment variables

.PHONY: local
local: dependencies localbuild test code-quality

.PHONY: localbuild
localbuild:
	go build .

.PHONY: dependencies
dependencies:
	echo "Installing dependencies"
	glide install

packages = \
	./services \
	./models \
	./config

.PHONY: test
test:
	@$(foreach package,$(packages), \
		set -e; \
		go test -coverprofile $(package)/cover.out -covermode=count $(package);)

.PHONY: cover
cover: test
	echo "mode: count" > cover-all.out
	@$(foreach package,$(packages), \
		tail -n +2 $(package)/cover.out >> cover-all.out;)
	gocover-cobertura < cover-all.out > cover-cobertura.xml

.PHONY: show
show:
	echo "Launching web browser to show overall coverage..."
	go tool cover -html=cover-all.out

.PHONY: code-quality
code-quality:
	gometalinter --vendor --tests --skip=mock \
		--exclude='_gen.go' --exclude='docs.go' --exclude='vendor/*'\
		--disable=gotype --disable=errcheck --disable=gas \
		--deadline=1500s --checkstyle --sort=linter ./... > static-analysis.xml



