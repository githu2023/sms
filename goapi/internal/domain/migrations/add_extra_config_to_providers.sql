-- 为 sms_providers 表添加 extra_config 字段
-- 用于存储运营商的额外配置（JSON格式）

-- 检查字段是否已存在，如果不存在则添加
SET @dbname = DATABASE();
SET @tablename = "sms_providers";
SET @columnname = "extra_config";
SET @preparedStatement = (SELECT IF(
  (
    SELECT COUNT(*) FROM INFORMATION_SCHEMA.COLUMNS
    WHERE
      (TABLE_SCHEMA = @dbname)
      AND (TABLE_NAME = @tablename)
      AND (COLUMN_NAME = @columnname)
  ) > 0,
  "SELECT 'Column extra_config already exists.' AS result;",
  CONCAT("ALTER TABLE ", @tablename, " ADD COLUMN ", @columnname, " JSON COMMENT '运营商额外配置(JSON格式，用于存储特殊配置如projectName等)' AFTER api_config;")
));
PREPARE alterIfNotExists FROM @preparedStatement;
EXECUTE alterIfNotExists;
DEALLOCATE PREPARE alterIfNotExists;

-- 示例：为 BigBus666 运营商添加配置
-- UPDATE sms_providers 
-- SET extra_config = JSON_OBJECT('projectName', 'hema')
-- WHERE code = 'bigbus666' AND extra_config IS NULL;

