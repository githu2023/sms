#!/bin/bash

# SMS平台API自动化测试脚本
# 作者: AI Assistant
# 日期: 2025-11-24

# 配置
BASE_URL="http://38.60.203.212:6060/v1"
CLIENT_API_PREFIX="/client/v1"
API_PREFIX="/api/v1"
TIMESTAMP=$(date +%s)
TEST_USER="testuser_${TIMESTAMP}"
TEST_EMAIL="testuser_${TIMESTAMP}@example.com"
TEST_PASSWORD="TestPassword123!"

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# 全局变量
TOTAL_TESTS=0
PASSED_TESTS=0
FAILED_TESTS=0
CLIENT_TOKEN=""
API_TOKEN=""
USER_ID=""
PHONE_NUMBER=""
ASSIGNMENT_ID=""

# 日志函数
log_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

log_success() {
    echo -e "${GREEN}[PASS]${NC} $1"
    ((PASSED_TESTS++))
}

log_error() {
    echo -e "${RED}[FAIL]${NC} $1"
    ((FAILED_TESTS++))
}

log_warn() {
    echo -e "${YELLOW}[WARN]${NC} $1"
}

# 测试函数
run_test() {
    local test_name="$1"
    local expected_status="$2"
    local response_file="/tmp/api_test_response.json"
    
    ((TOTAL_TESTS++))
    log_info "运行测试: $test_name"
    
    # 执行curl命令并保存响应
    shift 2
    local curl_cmd="$@"
    
    # 添加静态header和响应保存
    curl_cmd="curl -s -w '%{http_code}' -o $response_file $curl_cmd"
    
    local actual_status=$(eval $curl_cmd)
    local response_body=$(cat $response_file 2>/dev/null || echo "{}")
    
    # 验证状态码
    if [[ "$actual_status" == "$expected_status" ]]; then
        log_success "$test_name - 状态码: $actual_status"
        echo "响应: $(echo $response_body | jq . 2>/dev/null || echo $response_body)"
        return 0
    else
        log_error "$test_name - 期望状态码: $expected_status, 实际状态码: $actual_status"
        echo "响应: $(echo $response_body | jq . 2>/dev/null || echo $response_body)"
        return 1
    fi
}

# 提取JSON字段值
extract_json_field() {
    local json="$1"
    local field="$2"
    echo "$json" | jq -r ".data.$field" 2>/dev/null || echo ""
}

# 健康检查
test_health() {
    log_info "=== 健康检查测试 ==="
    run_test "健康检查" "200" \
        "-X GET $BASE_URL/health"
}

# 用户注册测试
test_register() {
    log_info "=== 用户注册测试 ==="
    
    # 正常注册
    local response_file="/tmp/register_response.json"
    local status=$(curl -s -w '%{http_code}' -o $response_file \
        -X POST "$BASE_URL$CLIENT_API_PREFIX/register" \
        -H "Content-Type: application/json" \
        -d "{\"username\":\"$TEST_USER\",\"email\":\"$TEST_EMAIL\",\"password\":\"$TEST_PASSWORD\"}")
    
    ((TOTAL_TESTS++))
    if [[ "$status" == "201" ]]; then
        local response=$(cat $response_file)
        USER_ID=$(extract_json_field "$response" "id")
        log_success "用户注册成功 - 用户ID: $USER_ID"
        echo "响应: $(echo $response | jq .)"
    else
        log_error "用户注册失败 - 状态码: $status"
        echo "响应: $(cat $response_file)"
    fi
    
    # 重复注册测试
    run_test "重复注册" "400" \
        "-X POST $BASE_URL$CLIENT_API_PREFIX/register" \
        "-H 'Content-Type: application/json'" \
        "-d '{\"username\":\"$TEST_USER\",\"email\":\"$TEST_EMAIL\",\"password\":\"$TEST_PASSWORD\"}'"
    
    # 无效参数测试
    run_test "注册-缺少密码" "400" \
        "-X POST $BASE_URL$CLIENT_API_PREFIX/register" \
        "-H 'Content-Type: application/json'" \
        "-d '{\"username\":\"testuser2\",\"email\":\"test2@example.com\"}'"
}

# 用户登录测试
test_login() {
    log_info "=== 用户登录测试 ==="
    
    # 正常登录
    local response_file="/tmp/login_response.json"
    local status=$(curl -s -w '%{http_code}' -o $response_file \
        -X POST "$BASE_URL$CLIENT_API_PREFIX/login" \
        -H "Content-Type: application/json" \
        -d "{\"username\":\"$TEST_USER\",\"password\":\"$TEST_PASSWORD\"}")
    
    ((TOTAL_TESTS++))
    if [[ "$status" == "200" ]]; then
        local response=$(cat $response_file)
        CLIENT_TOKEN=$(extract_json_field "$response" "token")
        log_success "用户登录成功"
        echo "响应: $(echo $response | jq .)"
        log_info "Client Token: ${CLIENT_TOKEN:0:50}..."
    else
        log_error "用户登录失败 - 状态码: $status"
        echo "响应: $(cat $response_file)"
    fi
    
    # 错误密码测试
    run_test "登录-错误密码" "400" \
        "-X POST $BASE_URL$CLIENT_API_PREFIX/login" \
        "-H 'Content-Type: application/json'" \
        "-d '{\"username\":\"$TEST_USER\",\"password\":\"wrongpassword\"}'"
    
    # 不存在的用户
    run_test "登录-用户不存在" "400" \
        "-X POST $BASE_URL$CLIENT_API_PREFIX/login" \
        "-H 'Content-Type: application/json'" \
        "-d '{\"username\":\"nonexistentuser\",\"password\":\"$TEST_PASSWORD\"}'"
}

# API Token获取测试
test_get_api_token() {
    log_info "=== API Token获取测试 ==="
    
    # 使用真实的API密钥
    local api_secret="431915928ea2df20f0a3af17921dac2efee5c63d46191927ac781e528e3fd0a5"
    
    # 正常获取API Token
    local response_file="/tmp/api_token_response.json"
    local status=$(curl -s -w '%{http_code}' -o $response_file \
        -X POST "$BASE_URL$API_PREFIX/get_token" \
        -H "Content-Type: application/json" \
        -d "{\"secret\":\"$api_secret\"}")
    
    ((TOTAL_TESTS++))
    if [[ "$status" == "200" ]]; then
        local response=$(cat $response_file)
        API_TOKEN=$(extract_json_field "$response" "token")
        log_success "API Token获取成功"
        echo "响应: $(echo $response | jq .)"
        log_info "API Token: ${API_TOKEN:0:50}..."
    else
        log_error "API Token获取失败 - 状态码: $status"
        echo "响应: $(cat $response_file)"
    fi
    
    # 无效API密钥
    run_test "API Token-无效密钥" "400" \
        "-X POST $BASE_URL$API_PREFIX/get_token" \
        "-H 'Content-Type: application/json'" \
        "-d '{\"secret\":\"invalid_secret\"}'"
}

# 业务类型测试
test_business_types() {
    log_info "=== 业务类型测试 ==="
    
    # 客户端API
    if [[ -n "$CLIENT_TOKEN" ]]; then
        run_test "客户端-获取业务类型" "200" \
            "-X GET $BASE_URL$CLIENT_API_PREFIX/business_types" \
            "-H 'Authorization: Bearer $CLIENT_TOKEN'"
    fi
    
    # 编程API
    if [[ -n "$API_TOKEN" ]]; then
        run_test "编程API-获取业务类型" "200" \
            "-X GET $BASE_URL$API_PREFIX/business_types" \
            "-H 'Authorization: Bearer $API_TOKEN'"
    fi
    
    # 无Token访问
    run_test "业务类型-无Token" "401" \
        "-X GET $BASE_URL$CLIENT_API_PREFIX/business_types"
}

# 余额查询测试
test_balance() {
    log_info "=== 余额查询测试 ==="
    
    # 客户端API
    if [[ -n "$CLIENT_TOKEN" ]]; then
        run_test "客户端-查询余额" "200" \
            "-X GET $BASE_URL$CLIENT_API_PREFIX/balance" \
            "-H 'Authorization: Bearer $CLIENT_TOKEN'"
    fi
    
    # 编程API
    if [[ -n "$API_TOKEN" ]]; then
        run_test "编程API-查询余额" "200" \
            "-X GET $BASE_URL$API_PREFIX/balance" \
            "-H 'Authorization: Bearer $API_TOKEN'"
    fi
    
    # 无Token访问
    run_test "余额查询-无Token" "401" \
        "-X GET $BASE_URL$CLIENT_API_PREFIX/balance"
}

# 手机号服务测试
test_phone_services() {
    log_info "=== 手机号服务测试 ==="
    
    # 获取手机号测试
    if [[ -n "$CLIENT_TOKEN" ]]; then
        local response_file="/tmp/get_phone_response.json"
        local status=$(curl -s -w '%{http_code}' -o $response_file \
            -X POST "$BASE_URL$CLIENT_API_PREFIX/get_phone" \
            -H "Authorization: Bearer $CLIENT_TOKEN" \
            -H "Content-Type: application/json" \
            -d '{"business_type":"verification","card_type":"physical"}')
        
        ((TOTAL_TESTS++))
        if [[ "$status" == "200" ]]; then
            local response=$(cat $response_file)
            PHONE_NUMBER=$(extract_json_field "$response" "phone_number")
            log_success "获取手机号成功 - 手机号: $PHONE_NUMBER"
            echo "响应: $(echo $response | jq .)"
        else
            log_error "获取手机号失败 - 状态码: $status"
            echo "响应: $(cat $response_file)"
        fi
    fi
    
    # 获取验证码测试
    if [[ -n "$CLIENT_TOKEN" && -n "$PHONE_NUMBER" ]]; then
        run_test "客户端-获取验证码" "200" \
            "-X POST $BASE_URL$CLIENT_API_PREFIX/get_code" \
            "-H 'Authorization: Bearer $CLIENT_TOKEN'" \
            "-H 'Content-Type: application/json'" \
            "-d '{\"phone_number\":\"$PHONE_NUMBER\",\"timeout\":30}'"
    fi
    
    # 手机状态查询
    if [[ -n "$CLIENT_TOKEN" && -n "$PHONE_NUMBER" ]]; then
        run_test "客户端-手机状态查询" "200" \
            "-X GET '$BASE_URL$CLIENT_API_PREFIX/phone_status?phone_number=$PHONE_NUMBER'" \
            "-H 'Authorization: Bearer $CLIENT_TOKEN'"
    fi
    
    # 无效参数测试
    if [[ -n "$CLIENT_TOKEN" ]]; then
        run_test "获取手机号-无效业务类型" "400" \
            "-X POST $BASE_URL$CLIENT_API_PREFIX/get_phone" \
            "-H 'Authorization: Bearer $CLIENT_TOKEN'" \
            "-H 'Content-Type: application/json'" \
            "-d '{\"business_type\":\"invalid_type\",\"card_type\":\"physical\"}'"
    fi
}

# 分配记录测试
test_assignments() {
    log_info "=== 分配记录测试 ==="
    
    # 获取分配记录
    if [[ -n "$CLIENT_TOKEN" ]]; then
        run_test "客户端-获取分配记录" "200" \
            "-X GET $BASE_URL$CLIENT_API_PREFIX/assignments" \
            "-H 'Authorization: Bearer $CLIENT_TOKEN'"
    fi
    
    # 获取成本统计
    if [[ -n "$CLIENT_TOKEN" ]]; then
        local today=$(date +%Y-%m-%d)
        run_test "客户端-成本统计" "200" \
            "-X GET '$BASE_URL$CLIENT_API_PREFIX/assignments/statistics?start_date=$today&end_date=$today'" \
            "-H 'Authorization: Bearer $CLIENT_TOKEN'"
    fi
    
    # 编程API测试
    if [[ -n "$API_TOKEN" ]]; then
        run_test "编程API-获取分配记录" "200" \
            "-X GET $BASE_URL$API_PREFIX/assignments" \
            "-H 'Authorization: Bearer $API_TOKEN'"
    fi
}

# 白名单测试
test_whitelist() {
    log_info "=== 白名单测试 ==="
    
    local test_ip="192.168.1.100"
    local test_description="测试IP地址"
    
    # 添加白名单
    if [[ -n "$CLIENT_TOKEN" ]]; then
        run_test "客户端-添加白名单" "201" \
            "-X POST $BASE_URL$CLIENT_API_PREFIX/whitelist" \
            "-H 'Authorization: Bearer $CLIENT_TOKEN'" \
            "-H 'Content-Type: application/json'" \
            "-d '{\"ip_address\":\"$test_ip\",\"description\":\"$test_description\"}'"
    fi
    
    # 获取白名单
    if [[ -n "$CLIENT_TOKEN" ]]; then
        run_test "客户端-获取白名单" "200" \
            "-X GET $BASE_URL$CLIENT_API_PREFIX/whitelist" \
            "-H 'Authorization: Bearer $CLIENT_TOKEN'"
    fi
    
    # 删除白名单
    if [[ -n "$CLIENT_TOKEN" ]]; then
        run_test "客户端-删除白名单" "200" \
            "-X DELETE '$BASE_URL$CLIENT_API_PREFIX/whitelist?ip_address=$test_ip'" \
            "-H 'Authorization: Bearer $CLIENT_TOKEN'"
    fi
    
    # 无效IP测试
    if [[ -n "$CLIENT_TOKEN" ]]; then
        run_test "添加白名单-无效IP" "400" \
            "-X POST $BASE_URL$CLIENT_API_PREFIX/whitelist" \
            "-H 'Authorization: Bearer $CLIENT_TOKEN'" \
            "-H 'Content-Type: application/json'" \
            "-d '{\"ip_address\":\"invalid_ip\",\"description\":\"测试\"}'"
    fi
}

# 安全测试
test_security() {
    log_info "=== 安全性测试 ==="
    
    # SQL注入测试
    run_test "SQL注入防护" "400" \
        "-X POST $BASE_URL$CLIENT_API_PREFIX/login" \
        "-H 'Content-Type: application/json'" \
        "-d '{\"username\":\"admin'\''or 1=1--\",\"password\":\"test\"}'"
    
    # XSS测试
    run_test "XSS防护" "400" \
        "-X POST $BASE_URL$CLIENT_API_PREFIX/register" \
        "-H 'Content-Type: application/json'" \
        "-d '{\"username\":\"<script>alert(1)</script>\",\"email\":\"test@test.com\",\"password\":\"test123\"}'"
    
    # 无效Token测试
    run_test "无效Token" "401" \
        "-X GET $BASE_URL$CLIENT_API_PREFIX/profile" \
        "-H 'Authorization: Bearer invalid_token_here'"
}

# 边界条件测试
test_edge_cases() {
    log_info "=== 边界条件测试 ==="
    
    # 超长用户名
    local long_username=$(printf 'a%.0s' {1..1000})
    run_test "超长用户名" "400" \
        "-X POST $BASE_URL$CLIENT_API_PREFIX/register" \
        "-H 'Content-Type: application/json'" \
        "-d '{\"username\":\"$long_username\",\"email\":\"test@test.com\",\"password\":\"test123\"}'"
    
    # 空JSON
    run_test "空JSON请求" "400" \
        "-X POST $BASE_URL$CLIENT_API_PREFIX/login" \
        "-H 'Content-Type: application/json'" \
        "-d '{}'"
    
    # 无效JSON格式
    run_test "无效JSON格式" "400" \
        "-X POST $BASE_URL$CLIENT_API_PREFIX/login" \
        "-H 'Content-Type: application/json'" \
        "-d '{invalid json}'"
}

# 生成测试报告
generate_report() {
    log_info "=== 测试报告 ==="
    echo "=========================================="
    echo "SMS平台API测试结果"
    echo "=========================================="
    echo "总测试数: $TOTAL_TESTS"
    echo "通过: $PASSED_TESTS"
    echo "失败: $FAILED_TESTS"
    echo "成功率: $(( PASSED_TESTS * 100 / TOTAL_TESTS ))%"
    echo "=========================================="
    
    if [[ $FAILED_TESTS -gt 0 ]]; then
        log_error "发现 $FAILED_TESTS 个测试失败，请检查日志"
        return 1
    else
        log_success "所有测试通过！"
        return 0
    fi
}

# 主函数
main() {
    log_info "开始SMS平台API自动化测试"
    log_info "测试用户: $TEST_USER"
    log_info "测试邮箱: $TEST_EMAIL"
    log_info "时间戳: $TIMESTAMP"
    echo "=========================================="
    
    # 执行测试
    test_health
    test_register
    test_login
    test_get_api_token
    test_business_types
    test_balance
    test_phone_services
    test_assignments
    test_whitelist
    test_security
    test_edge_cases
    
    # 生成报告
    echo ""
    generate_report
}

# 运行测试
main "$@"