## Transaction & Balance Refactor - TODO

1. **模型调整**
   - [x] 在 `sms_transactions` 表结构中增加 `frozen_before`、`frozen_after` 字段（或单独的变动字段）以记录冻结金额变化。
   - [x] 确认 `sms_customers` 中已有 `frozen_amount`（已存在），必要时为 `balance + frozen_amount` 加联合索引以支持新查询。

2. **交易类型扩展**
   - [x] 定义新的交易类型常量/文档说明：
     - `6`：预冻结（可用余额 -> 冻结）
     - `7`：解冻/释放（冻结 -> 可用）
     - `8`（如需要）记录冻结转实扣或完成状态
   - [x] 更新 `go_api.md` / 开发文档描述新的交易类型及含义。

3. **服务层流程改造**
   - [x] 在 `phone_service.GetPhone` 中将流程拆分为：
     1. **预扣阶段**：事务内调用 `transactionRepo.ReserveBalance` (balance 减、frozen 增、写 type=6 交易)。
     2. **调用运营商/创建 assignment 等操作**（不持锁）。
     3. **结算阶段**：
        - 成功：标记交易/assignment 关联，并更新 `frozen_amount -> 扣减`（type=2 或更新 type=6 状态）。
        - 失败：调用 `transactionRepo.ReleaseBalance`（frozen 减、balance 加、写 type=7）。
   - [x] 调整 `scheduler_service` 中的过期退款逻辑：按“释放冻结 -> balance 增”的方式处理。

4. **Repository 层支持**
   - [x] 在 `TransactionRepository` 中新增：
     - `ReserveBalance(ctx, tx, customerID, amount)`：原子更新 balance & frozen。
     - `CommitReservedBalance(ctx, tx, customerID, amount)`：将冻结金额转为正式扣减（若需要）。
     - `ReleaseReservedBalance(ctx, tx, customerID, amount)`：冻结金额释放回余额。
   - [x] 更新 `TransactionService` 对应方法，保证调用方可以获得一致的 API。

5. **API / DTO 调整**
   - [ ] 在 `GET /client/v1/profile`、`/balance` 等返回中加入 `frozen_amount` 字段，让前端可见剩余冻结金额。
   - [ ] 更新 `go_api.md` 对返回字段的说明。

6. **测试与工具**
   - [ ] 更新/新增单元测试覆盖预扣、释放、并发场景（PhoneService、TransactionService）。
   - [ ] 调整 `tools/balance_stress.go` 脚本，验证新逻辑下的并发扣费正确性。

7. **日志与监控**
   - [ ] 在预扣/释放/结算的关键路径打印结构化日志，便于排查。
   - [ ] 可选：添加指标（成功扣费数、预扣失败数等）便于监控。

