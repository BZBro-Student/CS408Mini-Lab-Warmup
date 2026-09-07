# Variables
APP_NAME := app
MAIN_FILE := cmd/app/main.go
CSS_INPUT := assets/css/input.css
CSS_OUTPUT := static/css/styles.css

.PHONY: all templ tailwind build run dev watch-tailwind watch-air clean

# Default target
all: build

# Generate Templ Go files from .templ components
templ:
	templ generate

# Build Tailwind CSS (one-off)
tailwind:
	tailwindcss -i $(CSS_INPUT) -o $(CSS_OUTPUT)

# Build the complete application for production
build: templ tailwind
	go build -o bin/$(APP_NAME) ./$(MAIN_FILE)

# Run the application directly (no live reload)
run: build
	./bin/$(APP_NAME)

# ==========================================
# Development & Live Reload
# ==========================================

# Run Tailwind watcher and Air concurrently
dev:
	@make -j2 watch-tailwind watch-air

# Watch for Tailwind changes and rebuild CSS
watch-tailwind:
	tailwindcss -i $(CSS_INPUT) -o $(CSS_OUTPUT) --watch

# Run Air for Go and Templ live reloading
watch-air:
	air

# ==========================================
# Cleanup
# ==========================================

# Clean build artifacts and generated files
clean:
	rm -rf bin/
	rm -f $(CSS_OUTPUT)
	find . -type f -name '*_templ.go' -delete