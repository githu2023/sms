-- Add frozen balance tracking columns to sms_transactions
ALTER TABLE `sms_transactions`
    ADD COLUMN `frozen_before` DECIMAL(10,2) NULL DEFAULT NULL COMMENT '变动前冻结金额';

ALTER TABLE `sms_transactions`
    ADD COLUMN `frozen_after` DECIMAL(10,2) NULL DEFAULT NULL COMMENT '变动后冻结金额';

