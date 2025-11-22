-- SMS Platform Database Schema
-- Version: 1.1 (Added Chinese Comments)
-- Author: Gemini

-- ----------------------------
-- 表结构: sms_customers (客户信息表)
-- ----------------------------
CREATE TABLE `sms_customers` (
    `id` BIGINT PRIMARY KEY AUTO_INCREMENT,
    `username` VARCHAR(255) UNIQUE COMMENT '客户端登录用户名',
    `email` VARCHAR(255) UNIQUE COMMENT '客户端登录邮箱',
    `password_hash` VARCHAR(255) COMMENT '客户端登录用的密码哈希',
    `api_secret_key` VARCHAR(255) UNIQUE NOT NULL COMMENT '用于生成API Token的唯一密钥',
    `balance` DECIMAL(10, 4) NOT NULL DEFAULT 0.00 COMMENT '客户余额',
    `status` TINYINT NOT NULL DEFAULT 1 COMMENT '客户状态 (1:正常, 2:冻结, 0:已删除)',
    `registration_ip` VARCHAR(45) COMMENT '注册时的IP地址',
    `last_login_ip` VARCHAR(45) COMMENT '最后一次登录的IP地址',
    `last_login_at` TIMESTAMP NULL COMMENT '最后一次登录的时间',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) COMMENT='客户信息表';

-- ----------------------------
-- 表结构: sms_providers (第三方服务商表)
-- ----------------------------
CREATE TABLE `sms_providers` (
    `id` INT PRIMARY KEY AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL COMMENT '服务商名称',
    `api_config` JSON COMMENT '服务商的API配置 (如URL, key等)',
    `is_enabled` BOOLEAN NOT NULL DEFAULT TRUE COMMENT '是否启用该服务商'
) COMMENT='第三方服务商表';

-- ----------------------------
-- 表结构: sms_business_types (业务类型表)
-- ----------------------------
CREATE TABLE `sms_business_types` (
    `id` INT PRIMARY KEY AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL COMMENT '业务名称, 例如 "腾讯QQ"',
    `code` VARCHAR(50) UNIQUE NOT NULL COMMENT '业务代码, 例如 "qq"',
    `is_enabled` BOOLEAN NOT NULL DEFAULT TRUE COMMENT '是否开放该业务'
) COMMENT='业务类型表';

-- ----------------------------
-- 表结构: sms_phone_assignments (手机号分配记录表)
-- ----------------------------
CREATE TABLE `sms_phone_assignments` (
    `id` BIGINT PRIMARY KEY AUTO_INCREMENT,
    `customer_id` BIGINT NOT NULL COMMENT '客户ID, 关联到sms_customers.id',
    `provider_id` INT NOT NULL COMMENT '服务商ID, 关联到sms_providers.id',
    `business_type_id` INT NOT NULL COMMENT '业务类型ID, 关联到sms_business_types.id',
    `card_type` VARCHAR(50) NOT NULL DEFAULT 'unknown' COMMENT '卡类型 (例如: physical, virtual)',
    `phone_number` VARCHAR(50) NOT NULL COMMENT '获取到的手机号',
    `verification_code` VARCHAR(50) COMMENT '获取到的验证码',
    `cost` DECIMAL(10, 4) NOT NULL COMMENT '本次操作的费用',
    `status` TINYINT NOT NULL COMMENT '状态 (1:待取码, 2:已完成, 3:已过期, 4:失败)',
    `expires_at` TIMESTAMP NULL COMMENT '手机号锁定的过期时间',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    INDEX `idx_customer_id` (`customer_id`),
    INDEX `idx_phone_number` (`phone_number`)
) COMMENT='手机号分配记录表';

-- ----------------------------
-- 表结构: sms_transactions (交易记录表)
-- ----------------------------
CREATE TABLE `sms_transactions` (
    `id` BIGINT PRIMARY KEY AUTO_INCREMENT,
    `customer_id` BIGINT NOT NULL COMMENT '客户ID',
    `amount` DECIMAL(10, 4) NOT NULL COMMENT '变动金额 (正数为充值, 负数为消费)',
    `balance_before` DECIMAL(10, 4) NOT NULL COMMENT '变动前余额',
    `balance_after` DECIMAL(10, 4) NOT NULL COMMENT '变动后余额',
    `type` TINYINT NOT NULL COMMENT '交易类型 (1:充值, 2:API消费)',
    `reference_id` BIGINT COMMENT '关联的业务ID, 例如sms_phone_assignments.id',
    `notes` VARCHAR(255) COMMENT '备注',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    INDEX `idx_customer_id` (`customer_id`)
) COMMENT='客户余额交易记录表';

-- ----------------------------
-- 表结构: sms_api_logs (API请求日志表)
-- ----------------------------
CREATE TABLE `sms_api_logs` (
    `id` BIGINT PRIMARY KEY AUTO_INCREMENT,
    `customer_id` BIGINT COMMENT '客户ID',
    `request_ip` VARCHAR(45) NOT NULL COMMENT '请求来源IP',
    `request_path` VARCHAR(255) NOT NULL COMMENT '请求的API路径',
    `request_body` TEXT COMMENT '请求体内容',
    `response_code` INT COMMENT 'HTTP响应状态码',
    `duration_ms` INT COMMENT '请求处理耗时(毫秒)',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    INDEX `idx_customer_id` (`customer_id`),
    INDEX `idx_created_at` (`created_at`)
) COMMENT='API请求日志表';

-- ----------------------------
-- 表结构: sms_ip_whitelist (IP白名单表)
-- ----------------------------
CREATE TABLE `sms_ip_whitelist` (
    `id` BIGINT PRIMARY KEY AUTO_INCREMENT,
    `customer_id` BIGINT NOT NULL COMMENT '客户ID',
    `ip_address` VARCHAR(45) NOT NULL COMMENT '白名单IP或IP段',
    `notes` VARCHAR(255) COMMENT '备注, 例如 "办公室IP"',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    UNIQUE `uniq_customer_ip` (`customer_id`, `ip_address`)
) COMMENT='API IP白名单表';