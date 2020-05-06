package db

import (
	"github.com/afanke/OJO/WebServer/dto"
	"github.com/afanke/OJO/utils/log"
	"github.com/ilibs/gosql/v2"
)

type Rank struct {
}

var rankPageSize = 10

func (Rank) GetACMTop10() ([]dto.ACMRank2, error) {
	sql := `select uid,sum(ac) as ac,sum(total) as total
			from ojo.contest_acm_overall
			group by uid
			order by ac desc ,total
			limit 10`
	var data []dto.ACMRank2
	err := gosql.Select(&data, sql)
	if err != nil {
		log.Warn("error:%v\n", err)
		return nil, err
	}
	err = user.SelectUserName(len(data), func(i int) (target int64) {
		return data[i].Uid
	}, func(i int, res string) {
		data[i].Username = res
	})
	return data, err
}

func (Rank) GetACMRank(form dto.RankForm) ([]dto.ACMRank2, error) {
	if form.Page < 1 {
		form.Page = 1
	}
	form.Page -= 1
	form.Limit = rankPageSize
	form.Offset = form.Page * rankPageSize
	sql := `select uid,sum(ac) as ac,sum(total) as total
			from ojo.contest_acm_overall
			group by uid
			order by ac desc ,total
			limit ?,?`
	var data []dto.ACMRank2
	err := gosql.Select(&data, sql, form.Offset, form.Limit)
	if err != nil {
		log.Warn("error:%v\n", err)
		return nil, err
	}
	err = user.SelectUserNameAndSig(len(data), func(i int) (target int64) {
		return data[i].Uid
	}, func(i int, res *dto.UsernameAndSig) {
		data[i].Username = res.Username
		data[i].Signature = res.Signature
	})
	return data, err
}

func (Rank) GetACMRankCount() (int, error) {
	sql := `select count(distinct uid) as total
			from ojo.contest_acm_overall`
	var count int
	err := gosql.Get(&count, sql)
	return count, err
}
