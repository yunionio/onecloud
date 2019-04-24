// Copyright 2019 Yunion
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tasks

import (
	"context"
	"fmt"

	"yunion.io/x/jsonutils"

	api "yunion.io/x/onecloud/pkg/apis/compute"
	"yunion.io/x/onecloud/pkg/cloudcommon/db"
	"yunion.io/x/onecloud/pkg/cloudcommon/db/taskman"
	"yunion.io/x/onecloud/pkg/cloudcommon/notifyclient"
	"yunion.io/x/onecloud/pkg/compute/models"
	"yunion.io/x/onecloud/pkg/util/logclient"
)

type LoadbalancerListenerRuleCreateTask struct {
	taskman.STask
}

func init() {
	taskman.RegisterTask(LoadbalancerListenerRuleCreateTask{})
}

func (self *LoadbalancerListenerRuleCreateTask) taskFail(ctx context.Context, lbr *models.SLoadbalancerListenerRule, reason string) {
	lbr.SetStatus(self.GetUserCred(), api.LB_CREATE_FAILED, reason)
	db.OpsLog.LogEvent(lbr, db.ACT_ALLOCATE_FAIL, reason, self.UserCred)
	logclient.AddActionLogWithStartable(self, lbr, logclient.ACT_CREATE, reason, self.UserCred, false)
	notifyclient.NotifySystemError(lbr.Id, lbr.Name, api.LB_CREATE_FAILED, reason)
	self.SetStageFailed(ctx, reason)
}

func (self *LoadbalancerListenerRuleCreateTask) OnInit(ctx context.Context, obj db.IStandaloneModel, data jsonutils.JSONObject) {
	lbr := obj.(*models.SLoadbalancerListenerRule)
	region := lbr.GetRegion()
	if region == nil {
		self.taskFail(ctx, lbr, fmt.Sprintf("failed to find region for lbr %s", lbr.Name))
		return
	}
	self.SetStage("OnLoadbalancerListenerRuleCreateComplete", nil)
	if err := region.GetDriver().RequestCreateLoadbalancerListenerRule(ctx, self.GetUserCred(), lbr, self); err != nil {
		self.taskFail(ctx, lbr, err.Error())
	}
}

func (self *LoadbalancerListenerRuleCreateTask) OnLoadbalancerListenerRuleCreateComplete(ctx context.Context, lbr *models.SLoadbalancerListenerRule, data jsonutils.JSONObject) {
	lbr.SetStatus(self.GetUserCred(), api.LB_STATUS_ENABLED, "")
	db.OpsLog.LogEvent(lbr, db.ACT_ALLOCATE, lbr.GetShortDesc(ctx), self.UserCred)
	logclient.AddActionLogWithStartable(self, lbr, logclient.ACT_CREATE, nil, self.UserCred, true)
	self.SetStageComplete(ctx, nil)
}

func (self *LoadbalancerListenerRuleCreateTask) OnLoadbalancerListenerRuleCreateCompleteFailed(ctx context.Context, lbr *models.SLoadbalancerListenerRule, reason jsonutils.JSONObject) {
	self.taskFail(ctx, lbr, reason.String())
}
