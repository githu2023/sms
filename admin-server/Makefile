.PHONY: linux web clean help

# 构建输出目录
BUILD_DIR := build

# 默认目标
help:
	@echo "可用命令:"
	@echo "  make linux  - 编译 Go 程序为 Linux 可执行文件（输出到 build/ 目录）"
	@echo "  make web     - 编译打包前端项目（输出到 build/ 目录）"
	@echo "  make clean   - 清理编译产物"

# 编译 Go 程序为 Linux 可执行文件
linux:
	@echo "正在编译 Linux 版本..."
	@mkdir -p $(BUILD_DIR)
	@cd server && \
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o ../$(BUILD_DIR)/server main.go
	@echo "编译完成！可执行文件位于: $(BUILD_DIR)/server"

# 编译打包前端项目
web:
	@echo "正在编译前端项目..."
	@mkdir -p $(BUILD_DIR)
	@cd web && \
	if [ ! -d "node_modules" ]; then \
		echo "正在安装依赖..."; \
		npm install; \
	fi && \
	npm run build
	@echo "正在移动前端构建产物到 build 目录..."
	@if [ -d "web/dist" ]; then \
		mv web/dist $(BUILD_DIR)/; \
	fi
	@echo "前端编译完成！输出目录: $(BUILD_DIR)/"

# 清理编译产物
clean:
	@echo "正在清理编译产物..."
	@rm -rf $(BUILD_DIR)
	@rm -rf web/dist
	@echo "清理完成！"

