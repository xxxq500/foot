package service

import (
	"strings"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-api/module/odds/pojo"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
)

type AsiaHisService struct {
	mysql.BaseService
}

func (this *AsiaHisService) FindExists(v *pojo.AsiaHis) bool {
	sql_build := strings.Builder{}
	sql_build.WriteString(" MatchId = '" + v.MatchId + "' AND CompId = '" + v.CompId + "' AND OddDate = '" + v.OddDate + "'")
	result, err := mysql.GetEngine().Where(sql_build.String()).Exist(new(pojo.AsiaHis))
	if err != nil {
		base.Log.Error("FindExists:", err)
	}
	return result
}

//根据比赛ID查找亚赔
func (this *AsiaHisService) FindByMatchId(matchId string) []*pojo.AsiaHis {
	dataList := make([]*pojo.AsiaHis, 0)
	err := mysql.GetEngine().Where(" MatchId = ? ", matchId).Find(dataList)
	if err != nil {
		base.Log.Error("FindByMatchId:", err)
	}
	return dataList
}

//根据比赛ID和波菜公司ID查找亚赔
func (this *AsiaHisService) FindByMatchIdCompId(matchId string, compIds ...string) []*pojo.AsiaHis {
	dataList := make([]*pojo.AsiaHis, 0)
	sql_build := strings.Builder{}
	sql_build.WriteString(" MatchId = '" + matchId + "' AND CompId in ( '0' ")
	for _, v := range compIds {
		sql_build.WriteString(" ,'")
		sql_build.WriteString(v)
		sql_build.WriteString("'")
	}
	sql_build.WriteString(")")
	err := mysql.GetEngine().Where(sql_build.String()).Find(&dataList)
	if err != nil {
		base.Log.Error("FindByMatchIdCompId:", err)
	}
	return dataList
}
