.PHONY: linux clean test run help

# 构建输出目录
BUILD_DIR := build

# 默认目标
help:
	@echo "可用命令:"
	@echo "  make linux  - 编译 Go 程序为 Linux 可执行文件（输出到 build/ 目录）"
	@echo "  make test    - 运行测试"
	@echo "  make run     - 本地运行程序"
	@echo "  make clean   - 清理编译产物"

# 编译 Go 程序为 Linux 可执行文件
linux:
	@echo "正在编译 Linux 版本..."
	@mkdir -p $(BUILD_DIR)
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o $(BUILD_DIR)/goapi ./cmd/server
	@echo "编译完成！可执行文件位于: $(BUILD_DIR)/goapi"

# 运行测试
test:
	@echo "正在运行测试..."
	@go test -v ./...

# 本地运行程序
run:
	@echo "正在启动程序..."
	@go run ./cmd/server

# 清理编译产物
clean:
	@echo "正在清理编译产物..."
	@rm -rf $(BUILD_DIR)
	@rm -f coverage.out
	@echo "清理完成！"

