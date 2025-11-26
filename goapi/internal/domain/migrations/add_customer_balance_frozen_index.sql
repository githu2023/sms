-- Add composite index on balance and frozen_amount to speed up reserve queries
ALTER TABLE `sms_customers`
    ADD INDEX `idx_customers_balance_frozen` (`balance`, `frozen_amount`);

