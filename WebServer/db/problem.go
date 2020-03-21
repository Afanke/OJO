package db

import (
	"fmt"
	"github.com/afanke/OJO/WebServer/dto"
)

type Problem struct {
}

var pb Problem

func (Problem) GetPbCase(pid int) ([]dto.ProblemCase, error) {
	var res []dto.ProblemCase
	fmt.Println(pid)
	err := db.Select(&res, "select * from problem_case where pid=?", pid)
	return res, err
}

func (Problem) GetProblem(id int) (*dto.Problem, error) {
	var res dto.Problem
	err := db.Get(&res, `select * from ojo.problem p where p.id=? limit 1`, id)
	return &res, err
}

func (Problem) GetProblemTag(pbid int) ([]dto.Tags, error) {
	var sql = `select t.id,t.name from tags t,problem_tags pt 
			where pt.pid=? and pt.tid=t.id`
	var tags []dto.Tags
	err := db.Select(&tags, sql, pbid)
	return tags, err
}

func (Problem) GetLanguage(pbid int) ([]dto.Language, error) {
	var sql = `select l.id,l.name from language l,problem_language pl 
			where pl.pid=? and pl.lid=l.id`
	var languages []dto.Language
	err := db.Select(&languages, sql, pbid)
	return languages, err
}

func (Problem) GetSample(pbid int) ([]dto.Sample, error) {
	var sql = `select s.id,s.input,s.output from sample s,problem_sample ps 
			where ps.pid=? and ps.sid=s.id`
	var samples []dto.Sample
	err := db.Select(&samples, sql, pbid)
	return samples, err
}

func (Problem) GetCreatorName(creatorId int) (string, error) {
	var s string
	err := db.Get(&s, "select a.name from administrator a where a.id=? limit 1", creatorId)
	return s, err
}

func (Problem) GetName(pbid int) (string, error) {
	var s string
	err := db.Get(&s, "select title from ojo.problem where id=? limit 1", pbid)
	return s, err
}
