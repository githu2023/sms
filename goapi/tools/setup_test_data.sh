#!/bin/bash

# 为测试用户设置业务类型和配置数据

BASE_URL="http://localhost:6060"
TEST_USER="testuser_1763969884"
API_SECRET="431915928ea2df20f0a3af17921dac2efee5c63d46191927ac781e528e3fd0a5"
MERCHANT_ID="368570"

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

log_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

log_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# 登录获取token
log_info "登录获取token..."
LOGIN_RESPONSE=$(curl -s -X POST "$BASE_URL/client/v1/login" \
    -H "Content-Type: application/json" \
    -d "{\"username\":\"$TEST_USER\",\"password\":\"TestPassword123!\"}")

CLIENT_TOKEN=$(echo "$LOGIN_RESPONSE" | jq -r '.data.token')
USER_ID=$(echo "$LOGIN_RESPONSE" | jq -r '.data.user_id' 2>/dev/null || echo "4")

if [[ "$CLIENT_TOKEN" == "null" || -z "$CLIENT_TOKEN" ]]; then
    log_error "登录失败"
    echo "$LOGIN_RESPONSE" | jq .
    exit 1
fi

log_success "登录成功，用户ID: $USER_ID, Token: ${CLIENT_TOKEN:0:20}..."

# 直接使用MySQL命令行插入测试数据
log_info "插入业务类型和客户配置数据到数据库..."

# 注意：你需要根据实际的数据库配置调整这个命令
# 这里假设数据库名为 sms_platform，用户为 root，密码为空
mysql -u root sms_platform << EOF
-- 插入业务类型（如果不存在）
INSERT INTO sms_business_types (name, code, is_enabled) 
VALUES 
    ('腾讯QQ', 'qq', 1),
    ('微信', 'wechat', 1),
    ('支付宝', 'alipay', 1),
    ('抖音', 'douyin', 1),
    ('淘宝', 'taobao', 1)
ON DUPLICATE KEY UPDATE name=VALUES(name);

-- 为测试用户插入业务配置
INSERT INTO sms_customer_business_config 
    (customer_id, platform_business_type_id, business_code, business_name, weight, status, created_at, updated_at)
VALUES
    ($USER_ID, 1, 'qq', '腾讯QQ', 10, 1, NOW(), NOW()),
    ($USER_ID, 2, 'wechat', '微信', 8, 1, NOW(), NOW()),
    ($USER_ID, 3, 'alipay', '支付宝', 5, 1, NOW(), NOW()),
    ($USER_ID, 4, 'douyin', '抖音', 3, 1, NOW(), NOW()),
    ($USER_ID, 5, 'taobao', '淘宝', 2, 1, NOW(), NOW())
ON DUPLICATE KEY UPDATE 
    business_name=VALUES(business_name),
    weight=VALUES(weight),
    status=VALUES(status),
    updated_at=NOW();

SELECT '数据插入完成' as status;
SELECT * FROM sms_business_types;
SELECT * FROM sms_customer_business_config WHERE customer_id = $USER_ID;
EOF

log_success "数据库数据设置完成"

# 测试获取业务类型
log_info "测试获取业务类型..."
BUSINESS_TYPES=$(curl -s -X GET "$BASE_URL/client/v1/business_types" \
    -H "Authorization: Bearer $CLIENT_TOKEN")

echo "业务类型列表:"
echo "$BUSINESS_TYPES" | jq .

# 测试获取手机号（使用qq业务类型）
log_info "测试获取手机号（业务类型：qq）..."
GET_PHONE_RESPONSE=$(curl -s -X POST "$BASE_URL/client/v1/get_phone" \
    -H "Authorization: Bearer $CLIENT_TOKEN" \
    -H "Content-Type: application/json" \
    -d '{"business_type":"qq","card_type":"physical"}')

echo "获取手机号响应:"
echo "$GET_PHONE_RESPONSE" | jq .

log_success "测试完成"
