echo "building\n\n"

wasm-pack build --target web --out-dir ./public/script

echo "ok, starting server\n\n"

go run .