# Project info
APP := sctx
VERSION := 0.3.0
DIST := dist

# Default build
build:
	@echo "🏗️  Building $(APP)..."
	@go build -o bin/$(APP) main.go
	@echo "✅ Build done: bin/$(APP)"

# Clean build artifacts
clean:
	@echo "🧹 Cleaning..."
	@rm -rf bin $(DIST)
	@echo "✅ Clean done."

# Release build for multiple platforms
release:
	@echo "🚀 Building release binaries for version $(VERSION)..."
	@mkdir -p $(DIST)
	GOOS=darwin GOARCH=arm64 go build -o $(DIST)/$(APP)-$(VERSION)-darwin-arm64 main.go
	GOOS=darwin GOARCH=amd64 go build -o $(DIST)/$(APP)-$(VERSION)-darwin-amd64 main.go
	GOOS=linux GOARCH=amd64 go build -o $(DIST)/$(APP)-$(VERSION)-linux-amd64 main.go
	GOOS=linux GOARCH=arm64 go build -o $(DIST)/$(APP)-$(VERSION)-linux-arm64 main.go
	cd $(DIST) && tar -czf $(APP)-$(VERSION)-darwin-arm64.tar.gz $(APP)-$(VERSION)-darwin-arm64
	cd $(DIST) && tar -czf $(APP)-$(VERSION)-darwin-amd64.tar.gz $(APP)-$(VERSION)-darwin-amd64
	cd $(DIST) && tar -czf $(APP)-$(VERSION)-linux-amd64.tar.gz $(APP)-$(VERSION)-linux-amd64
	cd $(DIST) && tar -czf $(APP)-$(VERSION)-linux-arm64.tar.gz $(APP)-$(VERSION)-linux-arm64
	@echo "✅ All binaries built and compressed in $(DIST)/"

# GitHub Release (requires GitHub CLI)
# Usage: make gh-release VERSION=0.3.0
gh-release: release
	@echo "📦 Creating GitHub release v$(VERSION)..."
	gh release create v$(VERSION) \
		$(DIST)/$(APP)-$(VERSION)-darwin-arm64.tar.gz \
		$(DIST)/$(APP)-$(VERSION)-darwin-amd64.tar.gz \
		$(DIST)/$(APP)-$(VERSION)-linux-amd64.tar.gz \
		$(DIST)/$(APP)-$(VERSION)-linux-arm64.tar.gz \
		bin/install.sh \
		--title "$(APP) v$(VERSION)" \
		--notes "- First public release\n- Colored prompt support\n- Zsh/Bash compatible\n- macOS + Linux binaries (ARM & AMD)"
	@echo "✅ GitHub release v$(VERSION) created successfully!"

# Tag and push
tag:
	@git tag v$(VERSION)
	@git push origin v$(VERSION)
	@echo "🏷️  Tag v$(VERSION) pushed."

# Convenience target: tag + release + publish
publish: tag gh-release
	@echo "🚀 Published $(APP) v$(VERSION)!"