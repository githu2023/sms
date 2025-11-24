#!/bin/bash

# SMS平台测试准备脚本
# 确保服务运行并准备测试环境

BASE_URL="http://localhost:6060"
SERVICE_PID_FILE="/tmp/goapi_service.pid"

log_info() {
    echo -e "\033[0;34m[INFO]\033[0m $1"
}

log_success() {
    echo -e "\033[0;32m[SUCCESS]\033[0m $1"
}

log_error() {
    echo -e "\033[0;31m[ERROR]\033[0m $1"
}

# 检查服务是否运行
check_service() {
    log_info "检查服务状态..."
    local response=$(curl -s -o /dev/null -w "%{http_code}" "$BASE_URL/health" 2>/dev/null)
    
    if [[ "$response" == "200" ]]; then
        log_success "服务正在运行，健康检查通过"
        return 0
    else
        log_error "服务未运行或健康检查失败 (状态码: $response)"
        return 1
    fi
}

# 启动服务
start_service() {
    log_info "启动SMS平台服务..."
    
    cd /Users/jarvis/work/tools/sms/goapi
    
    # 检查Go模块
    if [[ ! -f "go.mod" ]]; then
        log_error "未找到go.mod文件"
        return 1
    fi
    
    # 启动服务（后台运行）
    nohup go run ./cmd/server/main.go > /tmp/goapi_service.log 2>&1 &
    local pid=$!
    echo $pid > $SERVICE_PID_FILE
    
    # 等待服务启动
    log_info "等待服务启动..."
    sleep 3
    
    # 验证服务是否成功启动
    for i in {1..10}; do
        if check_service; then
            log_success "服务启动成功 (PID: $pid)"
            return 0
        fi
        sleep 1
    done
    
    log_error "服务启动失败"
    return 1
}

# 停止服务
stop_service() {
    if [[ -f $SERVICE_PID_FILE ]]; then
        local pid=$(cat $SERVICE_PID_FILE)
        if kill $pid 2>/dev/null; then
            log_info "停止服务 (PID: $pid)"
            rm -f $SERVICE_PID_FILE
        fi
    fi
    
    # 强制杀死可能残留的进程
    pkill -f "go run.*cmd/server/main.go" 2>/dev/null || true
}

# 准备测试环境
prepare_test_env() {
    log_info "准备测试环境..."
    
    # 检查必要的工具
    for tool in curl jq; do
        if ! command -v $tool &> /dev/null; then
            log_error "缺少必要工具: $tool"
            echo "请安装: brew install $tool"
            return 1
        fi
    done
    
    # 确保临时目录存在
    mkdir -p /tmp
    
    log_success "测试环境准备完成"
    return 0
}

# 主函数
main() {
    echo "=========================================="
    echo "SMS平台测试准备"
    echo "=========================================="
    
    case "${1:-prepare}" in
        "start")
            stop_service
            prepare_test_env && start_service
            ;;
        "stop")
            stop_service
            ;;
        "status")
            check_service
            ;;
        "prepare")
            prepare_test_env
            if ! check_service; then
                log_info "服务未运行，正在启动..."
                start_service
            fi
            ;;
        *)
            echo "用法: $0 {start|stop|status|prepare}"
            echo "  start   - 启动服务"
            echo "  stop    - 停止服务"
            echo "  status  - 检查状态"
            echo "  prepare - 准备环境（默认）"
            exit 1
            ;;
    esac
}

main "$@"