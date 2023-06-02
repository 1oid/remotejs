rm -rf remotejs_windows_386.exe remotejs_windows_amd64.exe remotejs_darwin_amd64 remotejs_darwin_arm64
echo "gen windows"
CGO_ENABLED=1 GOOS=windows GOARCH=386 CC="i686-w64-mingw32-gcc" go build --trimpath -ldflags "-s -w"  -o remotejs_windows_386.exe main.go
CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC="x86_64-w64-mingw32-gcc" go build --trimpath -ldflags "-s -w"  -o remotejs_windows_amd64.exe main.go
echo "gen darwin"
go build --trimpath -ldflags "-s -w" -o remotejs_darwin_amd64 main.go
echo "gen darwin arm64"
CGO_ENABLED=1 GOOS=darwin GOARCH=arm64 SDKROOT=$(xcrun --sdk macosx --show-sdk-path)  go build --trimpath -ldflags "-s -w" -o remotejs_darwin_arm64 main.go
