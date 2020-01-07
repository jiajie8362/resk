package resk

import (
	_ "github.com/vonnwang/account/core/accounts"
	"github.com/vonnwang/infra"
	"github.com/vonnwang/infra/base"
	"github.com/vonnwang/resk/apis/gorpc"
	_ "github.com/vonnwang/resk/apis/gorpc"
	_ "github.com/vonnwang/resk/apis/web"
	_ "github.com/vonnwang/resk/core/envelopes"
	"github.com/vonnwang/resk/jobs"
)

func init() {
	infra.Register(&base.PropsStarter{})
	infra.Register(&base.DbxDatabaseStarter{})
	infra.Register(&base.ValidatorStarter{})
	infra.Register(&base.GoRPCStarter{})
	infra.Register(&gorpc.GoRpcApiStarter{})
	infra.Register(&jobs.RefundExpiredJobStarter{})
	infra.Register(&base.IrisServerStarter{})
	infra.Register(&infra.WebApiStarter{})
	infra.Register(&base.EurekaStarter{})
	infra.Register(&base.HookStarter{})
}
