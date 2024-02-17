.PHONY: tailwindcss
tailwindcss:
	@echo "Building Tailwind CSS..."
	npx tailwindcss -i ui/style/input.css -o ui/style/output.css
	@echo "Done."

.PHONY: templ
templ:
	@echo "Building templates..."
	templ generate
	@echo "Done."

.PHONY: generate-ui
generate-ui: tailwindcss templ
