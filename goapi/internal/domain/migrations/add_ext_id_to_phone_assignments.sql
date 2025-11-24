-- 添加 ext_id 字段到 sms_phone_assignments 表
-- ext_id 用于存储运营商返回的外部ID（如 BigBus666 的 extId），用于后续获取验证码和释放手机号

ALTER TABLE `sms_phone_assignments`
ADD COLUMN `ext_id` VARCHAR(100) NULL COMMENT '运营商返回的外部ID(extId)，用于后续获取验证码和释放手机号' AFTER `phone_number`;

-- 添加索引（可选，如果经常根据 ext_id 查询）
-- CREATE INDEX `idx_ext_id` ON `sms_phone_assignments` (`ext_id`);

-- 号码(id,与KEY请找管理员拿取)
-- http://szb.jczl70.com:6086/mqtt/msg/getNumber?id=%d&key=%s

-- {"number":"1888888888","extId":"2025111215435966402734","id":1}

-- 拿验证码
-- http://szb.jczl70.com:6086/mqtt/msg/getCode?extId=%s
-- {
--   code,                   // 0 表示接口操作成功
--   message,                // 接口返回消息提示
--   data:{
--     receiveStatus: 1,     //0, 短信接收失败 1, 短信接收成功 
--     message: "xxx"        //短信信息，status为1时读取
--     error: "xxx"          //错误信息，status为0时读取  
--   }
-- }
-- eg:
-- {
--   "code" : 0,
--   "message" : "成功",
--   "data" : {
--     "receiveStatus" : 1,
--     "message" : "456039"
--   }
-- }

-- 释放号码返回状态：*如果注册失败，必须要执行**
-- http://szb.jczl70.com:6086/mqtt/msg/release?extId=%s&status=%d
-- status : 1,注册成功 2，超时 3，已注册 4，其它问题
