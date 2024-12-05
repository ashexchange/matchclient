package spotapi

import (
	"context"

	"github.com/ashexchange/matchclient/v2"
	"github.com/ashexchange/matchclient/v2/types"
)

type AssetClient interface {
	// AssetList 列出资产的数据库保存精度及显示精度
	AssetList(ctx context.Context) (AssetListResponse, error)

	// AssetUpdate 更新(加或减)单个资产
	AssetUpdate(ctx context.Context, req *AssetUpdateRequest) (*types.Result, error)

	// AssetUpdateBatch 批量更新(加或减)资产
	AssetUpdateBatch(ctx context.Context, req *AssetUpdateBatchRequest) ([]*types.Result, error)

	// AssetLock 锁定资产
	AssetLock(ctx context.Context, req *AssetLockRequest) (*types.Result, error)

	// AssetUnlock 解锁资产
	AssetUnlock(ctx context.Context, req *AssetUnlockRequest) (*types.Result, error)

	// AssetQuery 查询资产(因为matchengine读写分离，该接口可能会有延迟)
	AssetQuery(ctx context.Context, req *AssetQueryRequest) (AssetQueryResponse, error)

	// AssetQueryInTime 实时查询资产
	AssetQueryInTime(ctx context.Context, req *AssetQueryInTimeRequest) (AssetQueryResponse, error)

	// AssetQueryUsers 查询多个用户的资产
	AssetQueryUsers(ctx context.Context, req *AssetQueryUsersRequest) (AssetQueryUsersResponse, error)

	// AssetQueryUsersInTime 实时查询多个用户的资产
	AssetQueryUsersInTime(ctx context.Context, req *AssetQueryUsersInTimeRequest) (AssetQueryUsersResponse, error)

	// AssetSummary 根据资产 账户id进行资产统计
	AssetSummary(ctx context.Context, req *AssetSummaryRequest) (*AssetSummaryResponse, error)

	// AssetQueryAll 查询某个用户的全部资产
	AssetQueryAll(ctx context.Context, req *AssetQueryAllRequest) (AssetQueryAllResponse, error)

	// AssetQueryAllInTime 实时查询某个用户的全部资产
	AssetQueryAllInTime(ctx context.Context, req *AssetQueryAllInTimeRequest) (AssetQueryAllResponse, error)

	// AssetQueryLock 查询某个用户被锁定的资产
	AssetQueryLock(ctx context.Context, req *AssetQueryLockRequest) (AssetQueryLockResponse, error)

	// AssetQueryLockInTime 实时查询某个用户被锁定的资产
	AssetQueryLockInTime(ctx context.Context, req *AssetQueryLockInTimeRequest) (AssetQueryLockResponse, error)

	// AssetBackup 将用户资产备份到数据库中
	AssetBackup(ctx context.Context) (*AssetBackupResponse, error)

	// AssetHistory 查询某个用户的资产流水
	AssetHistory(ctx context.Context, req *AssetHistoryRequest) (*AssetHistoryResponse, error)

	// AssetAllUsers 查询币种下所有的用户资产
	AssetAllUsers(ctx context.Context, req *AssetAllUsersRequest) ([]*UsersAsset, error)
}

type Asset struct {
	PrecShow int `json:"prec_show"`
	PrecSave int `json:"prec_save"`
}

type AssetListResponse map[string]map[string]*Asset

type (
	AssetBalance struct {
		Available string `json:"available"`
		Frozen    string `json:"frozen"`
	}

	// the key is asset name
	AssetAccount map[string]*AssetBalance
)

type (
	AssetUpdateRequest struct {
		UserId     types.UserID
		Account    int
		Asset      string
		Business   string
		BusinessId uint64
		Change     string
		Detail     map[string]interface{}
	}

	AssetUpdateBatchRequest struct {
		UpdateList []*AssetUpdateRequest
	}
)

type (
	AssetLockRequest struct {
		UserId     types.UserID
		Account    int
		Asset      string
		Business   string
		BusinessId uint64
		Amount     string
	}

	AssetUnlockRequest = AssetLockRequest
)

type (
	AssetQueryRequest struct {
		UserId    types.UserID
		Account   int
		AssetList []string
	}

	AssetQueryInTimeRequest struct {
		UserId    types.UserID
		Account   int
		AssetList []string
	}

	// the key is asset name
	AssetQueryResponse = AssetAccount
)

type (
	AssetQueryUsersRequest struct {
		Account int
		UserIds []types.UserID
	}

	AssetQueryUsersInTimeRequest struct {
		Account int // must be grater than 0
		UserIds []types.UserID
	}

	// map key is user_id
	AssetQueryUsersResponse map[string]AssetAccount
)

type (
	AssetSummaryRequest struct {
		Asset   string
		Account *int // optional
	}

	AssetSummaryResponse struct {
		TotalUsers     uint32 `json:"total_users"`
		AvailableUsers uint32 `json:"available_users"`
		LockUsers      uint32 `json:"lock_users"`
		FrozenUsers    uint32 `json:"frozen_users"`
		Total          string `json:"total"`
		Available      string `json:"available"`
		Lock           string `json:"lock"`
		Frozen         string `json:"frozen"`
	}
)

type (
	AssetQueryAllRequest struct {
		UserId types.UserID
	}

	AssetQueryAllInTimeRequest = AssetQueryAllRequest

	// AssetQueryAllResponse map key is account
	AssetQueryAllResponse map[string]AssetAccount
)

type (
	AssetQueryLockRequest struct {
		UserId    types.UserID
		Account   int
		AssetList []string // optional
	}

	AssetQueryLockInTimeRequest = AssetQueryLockRequest

	// AssetQueryLockResponse key: asset name, value: lock amount
	AssetQueryLockResponse map[string]string
)

type (
	AssetHistoryRequest struct {
		UserId    types.UserID
		Account   int
		Asset     string
		Business  string
		StartTime int64
		EndTime   int64
		Offset    int64
		Limit     int64
	}

	AssetHistory struct {
		UserId   types.UserID           `json:"user"`
		Account  int                    `json:"account"`
		Time     float64                `json:"time"`
		Asset    string                 `json:"asset"`
		Business string                 `json:"business"`
		Change   string                 `json:"change"`
		Balance  string                 `json:"balance"`
		Detail   map[string]interface{} `json:"detail"`
	}

	AssetHistoryResponse struct {
		Offset  uint            `json:"offset"`
		Limit   uint            `json:"limit"`
		Records []*AssetHistory `json:"records"`
	}
)

type AssetBackupResponse struct {
	Table string  `json:"table"`
	Time  float64 `json:"time"`
}

type AssetAllUsersRequest struct {
	Account int    `json:"account"`
	Asset   string `json:"asset"`
}

type UsersAsset struct {
	UserId    types.UserID `json:"user_id"`
	Account   int          `json:"account"`
	Frozen    string       `json:"frozen"`
	Available string       `json:"available"`
}

type assetClient struct {
	cli matchclient.Invoker
}

func NewAssetClient(c matchclient.Invoker) AssetClient {
	return &assetClient{cli: c}
}

func (c *assetClient) AssetList(ctx context.Context) (reply AssetListResponse, err error) {
	err = c.cli.Invoke(ctx, "asset.list", &reply)
	return
}

var noDetails = make(map[string]interface{})

func (c *assetClient) AssetUpdate(ctx context.Context, req *AssetUpdateRequest) (reply *types.Result, err error) {
	if req.Detail == nil {
		req.Detail = noDetails
	}
	err = c.cli.Invoke(ctx, "asset.update", &reply,
		req.UserId,
		req.Account,
		req.Asset,
		req.Business,
		req.BusinessId,
		req.Change,
		req.Detail,
	)
	return
}

func (c *assetClient) AssetUpdateBatch(ctx context.Context, req *AssetUpdateBatchRequest) (reply []*types.Result, err error) {
	params := make([]interface{}, 0, len(req.UpdateList))
	for _, req := range req.UpdateList {
		params = append(params, []interface{}{
			req.UserId,
			req.Account,
			req.Asset,
			req.Business,
			req.BusinessId,
			req.Change,
			req.Detail,
		})
	}
	err = c.cli.Invoke(ctx, "asset.update", &reply, params...)
	return
}

func (c *assetClient) AssetLock(ctx context.Context, req *AssetLockRequest) (reply *types.Result, err error) {
	err = c.cli.Invoke(ctx, "asset.lock", &reply,
		req.UserId, req.Account, req.Asset, req.Business, req.BusinessId, req.Amount,
	)
	return
}

func (c *assetClient) AssetUnlock(ctx context.Context, req *AssetUnlockRequest) (reply *types.Result, err error) {
	err = c.cli.Invoke(ctx, "asset.unlock", &reply,
		req.UserId, req.Account, req.Asset, req.Business, req.BusinessId, req.Amount,
	)
	return
}

func (c *assetClient) AssetQuery(ctx context.Context, req *AssetQueryRequest) (reply AssetQueryResponse, err error) {
	params := []interface{}{req.UserId, req.Account}
	for _, asset := range req.AssetList {
		params = append(params, asset)
	}
	err = c.cli.Invoke(ctx, "asset.query", &reply, params...)
	return
}

func (c *assetClient) AssetQueryInTime(ctx context.Context, req *AssetQueryInTimeRequest) (reply AssetQueryResponse, err error) {
	params := []interface{}{req.UserId, req.Account}
	for _, asset := range req.AssetList {
		params = append(params, asset)
	}
	err = c.cli.Invoke(ctx, "asset.query_intime", &reply, params...)
	return
}

func (c *assetClient) AssetQueryUsers(ctx context.Context, req *AssetQueryUsersRequest) (reply AssetQueryUsersResponse, err error) {
	err = c.cli.Invoke(ctx, "asset.query_users", &reply, req.Account, req.UserIds)
	return
}

func (c *assetClient) AssetQueryUsersInTime(ctx context.Context, req *AssetQueryUsersInTimeRequest) (reply AssetQueryUsersResponse, err error) {
	err = c.cli.Invoke(ctx, "asset.query_users_intime", &reply, req.Account, req.UserIds)
	return
}

func (c *assetClient) AssetSummary(ctx context.Context, req *AssetSummaryRequest) (reply *AssetSummaryResponse, err error) {
	params := []interface{}{req.Asset}
	if req.Account != nil {
		params = append(params, *req.Account)
	}
	err = c.cli.Invoke(ctx, "asset.summary", &reply, params...)
	return
}

func (c *assetClient) AssetQueryAll(ctx context.Context, req *AssetQueryAllRequest) (reply AssetQueryAllResponse, err error) {
	err = c.cli.Invoke(ctx, "asset.query_all", &reply, req.UserId)
	return
}

func (c *assetClient) AssetQueryAllInTime(ctx context.Context, req *AssetQueryAllInTimeRequest) (reply AssetQueryAllResponse, err error) {
	err = c.cli.Invoke(ctx, "asset.query_all_intime", &reply, req.UserId)
	return
}

func (c *assetClient) AssetQueryLock(ctx context.Context, req *AssetQueryLockRequest) (reply AssetQueryLockResponse, err error) {
	params := []interface{}{req.UserId, req.Account}
	for _, asset := range req.AssetList {
		params = append(params, asset)
	}
	err = c.cli.Invoke(ctx, "asset.query_lock", &reply)
	return
}

func (c *assetClient) AssetQueryLockInTime(ctx context.Context, req *AssetQueryLockInTimeRequest) (reply AssetQueryLockResponse, err error) {
	params := []interface{}{req.UserId, req.Account}
	for _, asset := range req.AssetList {
		params = append(params, asset)
	}
	err = c.cli.Invoke(ctx, "asset.query_lock_intime", &reply, params...)
	return
}

func (c *assetClient) AssetBackup(ctx context.Context) (reply *AssetBackupResponse, err error) {
	err = c.cli.Invoke(ctx, "asset.backup", &reply)
	return
}

func (c *assetClient) AssetHistory(ctx context.Context, req *AssetHistoryRequest) (reply *AssetHistoryResponse, err error) {
	err = c.cli.Invoke(ctx, "asset.history", &reply,
		req.UserId,
		req.Account,
		req.Asset,
		req.Business,
		req.StartTime,
		req.EndTime,
		req.Offset,
		req.Limit,
	)
	return
}

func (c *assetClient) AssetAllUsers(ctx context.Context, req *AssetAllUsersRequest) (reply []*UsersAsset, err error) {
	err = c.cli.Invoke(ctx, "asset.query_asset_users", &reply,
		req.Account,
		req.Asset,
	)
	return
}
