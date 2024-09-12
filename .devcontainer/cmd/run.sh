echo "building"
echo ""

wasm-pack build --target web --out-dir ./src/public/script

echo ""
echo "wasm builded"
echo "Starting server..."

go run .