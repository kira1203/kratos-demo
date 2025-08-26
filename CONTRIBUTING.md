

## 包职责规范

| 包名 | 职责 | 禁止 |
|------|------|------|
| `biz/dto` | 数据传输对象：请求、响应、分页 | ❌ 不可有方法<br>❌ 不可放业务逻辑 |
| `biz/domain` | 领域实体：有状态、有行为 | ❌ 不可出现 `Req`/`Reply`<br>❌ 不可直接用于 HTTP 响应 |
| `biz` 接口 | 定义 `Repo`、`UseCase` 接口 | ❌ 不可依赖 `v1` |
| `service` | 协议适配：gRPC/HTTP → dto → biz → dto → v1 | ✅ 可做参数校验、转换 |



## 命名规范

- 请求：`{Verb}{Entity}Req` → `CreateBannerReq`, `QueryBannerReq`
- 响应：
    - 列表：`{Entity}ListReply` 或使用 `ListReply[T]`
    - 单个：`GetBannerReply` 或直接返回 `*domain.Banner`
    - 聚合：`DashboardSummary`（放 `domain`）
- 领域实体：`Banner`, `Order`, `User`（不要加 `VO`/`BO` 后缀）


## biz 接口返回值规范

| 场景 | 返回类型 | 示例 |
|------|----------|------|
| 查询单个实体 | `*domain.Xxx` | `GetBanner(ctx, id) (*domain.Banner, error)` |
| 查询列表 | `*dto.ListReply[domain.Xxx]` | `GetBannerList(...) (*dto.ListReply[domain.Banner], error)` |
| 创建 | `*domain.Xxx` | 返回完整实体（含 ID） |
| 更新 | `*domain.Xxx` 或 `error` | 可返回新状态 |
| 删除 | `error` | 成功无返回 |
| 聚合查询 | `*domain.XxxSummary` | 如 `DashboardSummary` |



// ✅ 正确示例
func (b *ProductBiz) GetBannerList(ctx context.Context, req *dto.BannerListReq) (*dto.ListReply[domain.Banner], error)

func (b *ProductBiz) CreateBanner(ctx context.Context, req *dto.CreateBannerReq) (*domain.Banner, error)

func (b *ProductBiz) GetDashboard(ctx context.Context) (*domain.DashboardSummary, error)