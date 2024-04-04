.PHONY: tailwindcss
tailwindcss:
	@echo "Building Tailwind CSS..."
	./tailwindcss -i server/ui/style/input.css -o server/ui/style/output.css -m
	@echo "Done."

.PHONY: templ
templ:
	@echo "Building templates..."
	templ generate
	@echo "Done."

.PHONY: generate-ui
generate-ui: tailwindcss templ
