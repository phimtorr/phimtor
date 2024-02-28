.PHONY: tailwindcss
tailwindcss:
	@echo "Building Tailwind CSS..."
	npx tailwindcss -i server/ui/style/input.css -o server/ui/style/output.css
	@echo "Done."

.PHONY: templ
templ:
	@echo "Building templates..."
	templ generate
	@echo "Done."

.PHONY: generate-ui
generate-ui: tailwindcss templ
