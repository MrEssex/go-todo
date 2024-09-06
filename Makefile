templ:
	@templ generate --watch --proxy="http://localhost:3030" --open-browser=false

# run air to detect any go file changes to re-build and re-run the server.
server:
	@air \
	--build.exclude_dir "node_modules" \
	--build.include_ext "go" \
	--build.stop_on_error "false" \
	--misc.clean_on_exit true \
	--screen.clear_on_rebuild true \
	--log.main_only true


# start the application in development
dev:
	@make -j5 templ server
