package ctrl

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gogotime/OJO/WebServer/dto"
	jsp "github.com/gogotime/OJO/WebServer/judge"
	"github.com/gogotime/OJO/utils/log"
	"github.com/kataras/iris/v12"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type Problem struct {
}

var pb Problem

func (Problem) LocalTest(c iris.Context) {
	var form dto.JudgeForm
	err := c.ReadJSON(&form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	_, err = isAdmin(c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	addr, err := jsp.GetAddr()
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	client := &http.Client{
		Timeout: time.Duration(form.MaxRealTime*form.SPJMp*form.CompMp) * time.Second * 5,
	}
	buff, err := json.Marshal(&form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	res, err := client.Post("http://"+addr+"/judge", "application/json", bytes.NewBuffer(buff))
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: form})
}

func (Problem) AddProblem(c iris.Context) {
	var p dto.Problem
	err := c.ReadJSON(&p)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	userId, err := isAdmin(c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	p.Cid = userId
	err = pbdb.InsertProblem(&p)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: "save successfully"})
}

func (Problem) UpdateProblem(c iris.Context) {
	var p dto.Problem
	err := c.ReadJSON(&p)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	err = pb.isPermitted(c, p.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	err = pbdb.UpdateProblem(&p)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: "save successfully"})
}

func (Problem) DeleteProblem(c iris.Context) {
	var id dto.Id
	err := c.ReadJSON(&id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	err = pb.isCreator(c, id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	count, err := pbdb.IsDepended(id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	if count != 0 {
		c.JSON(&dto.Res{Error: errors.New("can't delete problem: the problem is depended by " + strconv.Itoa(count) + " contests").Error(), Data: nil})
		return
	}
	err = pbdb.DeleteProblem(id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: "delete successfully"})
}

func (Problem) GetAll(c iris.Context) {
	var form dto.ProblemForm
	err := c.ReadJSON(&form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	userId, err := isAdmin(c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	form.Cid = userId
	data, err := pbdb.GetAll(&form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: data})
}

func (Problem) GetAllShared(c iris.Context) {
	var form dto.ProblemForm
	err := c.ReadJSON(&form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	userId, err := isAdmin(c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	form.Cid = userId
	data, err := pbdb.GetAll(&form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: data})
}

func (Problem) GetSharedCount(c iris.Context) {
	userId, err := isAdmin(c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	data, err := pbdb.GetSharedCount(userId)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: data})
}

func (Problem) GetDetail(c iris.Context) {
	var id dto.Id
	err := c.ReadJSON(&id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	err = pb.isPermitted(c, id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	data, err := pbdb.GetDetail(id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: data})
}

func (Problem) GetCount(c iris.Context) {
	var form dto.ProblemForm
	err := c.ReadJSON(&form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	userId, err := isAdmin(c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	form.Cid = userId
	tags, err := pbdb.GetCount(&form)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: tags})
}

func (Problem) SetVisibleTrue(c iris.Context) {
	var id dto.Id
	err := c.ReadJSON(&id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	err = pb.isPermitted(c, id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	err = pbdb.SetVisibleTrue(id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: "set up successfully"})
}

func (Problem) SetVisibleFalse(c iris.Context) {
	var id dto.Id
	err := c.ReadJSON(&id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	err = pb.isPermitted(c, id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	err = pbdb.SetVisibleFalse(id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: "set up successfully"})
}

func (Problem) TryEdit(c iris.Context) {
	var id dto.Id
	err := c.ReadJSON(&id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	err = pb.isPermitted(c, id.Id)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	c.JSON(&dto.Res{Error: "", Data: "ok"})
}

func (Problem) isCreator(c iris.Context, id int64) error {
	adminId, err := isSuperAdmin(c)
	if err == nil {
		return nil
	}
	creatorId, err := pbdb.GetCreatorId(id)
	if err != nil {
		return err
	}
	if adminId != creatorId {
		return errors.New("not allowed")
	}
	return nil
}

func (Problem) isPermitted(c iris.Context, id int64) error {
	err := pb.isCreator(c, id)
	if err == nil {
		return nil
	}
	err = pbdb.IsShared(id)
	return err
}

// ------------------------------------

func (Problem) ImportFromQDUOJ(c iris.Context) {
	adminId, err := isAdmin(c)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	file, info, err := c.FormFile("file")
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	defer file.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	path := "./dist/file/" + strconv.Itoa(int(adminId)) + "_" +
		time.Now().Format("2006_01_02_15_04_05") + "_" + info.Filename
	isZip := isZip(fileBytes[:4])
	if !isZip {
		c.JSON(&dto.Res{Error: "not zip", Data: nil})
		return
	}
	err = ioutil.WriteFile(path, fileBytes, 0666)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	problems, err := readPbFromZip(path, adminId)
	if err != nil {
		c.JSON(&dto.Res{Error: err.Error(), Data: nil})
		return
	}
	for i, j := 0, len(problems); i < j; i++ {
		err := pbdb.InsertProblem(&problems[i])
		if err != nil {
			c.JSON(&dto.Res{Error: err.Error(), Data: nil})
			return
		}
	}
	c.JSON(&dto.Res{Error: "", Data: "upload file successfully"})
}

func subString(str string, start, end int) (string, error) {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		return "", errors.New("start is wrong")
	}

	if end < start || end > length {
		return "", errors.New("end is wrong")
	}
	return string(rs[start:end]), nil
}

func getDir(path string) (string, error) {
	return subString(path, 0, strings.LastIndex(path, "/"))
}

func unzip(archive, target string) error {
	reader, err := zip.OpenReader(archive)
	if err != nil {
		log.Error("%v\n", err)
		return err
	}

	if err := os.MkdirAll(target, 0755); err != nil {
		log.Error("%v\n", err)
		return err
	}

	for _, file := range reader.File {
		err := func() error {
			path := filepath.Join(target, file.Name)
			if runtime.GOOS == "windows" {
				path = strings.ReplaceAll(path, "\\", "/")
			}
			dir, err := getDir(path)
			if err != nil {
				log.Error("%v\n", err)
				return err
			}
			err = os.MkdirAll(dir, 0755)
			if err != nil {
				log.Error("%v\n", err)
				return err
			}
			o, err := os.Create(path)
			if err != nil {
				log.Error("%v\n", err)
				return err
			}
			defer o.Close()
			i, err := file.Open()
			if err != nil {
				log.Error("%v\n", err)
				return err
			}
			defer i.Close()
			_, err = io.Copy(o, i)
			if err != nil {
				log.Error("%v\n", err)
				return err
			}
			return nil
		}()
		if err != nil {
			log.Error("%v\n", err)
			return err
		}
	}

	return nil
}

func isZip(buf []byte) bool {
	return bytes.Equal(buf, []byte("PK\x03\x04"))
}

func readPbFromZip(path string, cid int64) ([]dto.Problem, error) {
	unzipPath := strings.ReplaceAll(path, ".zip", "")
	err := unzip(path, unzipPath)
	if err != nil {
		log.Error("%v\n", err)
	}
	i := 1
	res := make([]dto.Problem, 0)
	for {
		pbDir := unzipPath + "/" + strconv.Itoa(i)
		_, err := os.Stat(pbDir)
		if os.IsNotExist(err) {
			break
		}
		problem, err := readPbFromDir(pbDir, cid)
		if err != nil {
			log.Error("%v\n", err)
			return nil, err
		}
		res = append(res, problem)
		i++
	}
	return res, nil
}

func readPbFromDir(path string, cid int64) (dto.Problem, error) {
	file, err := ioutil.ReadFile(path + "/problem.json")
	if err != nil {
		log.Error("%v\n", err)
		return dto.Problem{}, err
	}
	var data dto.QDUOJProblem
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Error("%v\n", err)
		return dto.Problem{}, err
	}
	var res dto.Problem
	res.Cid = cid
	res.Ref = data.Ref
	res.Title = data.Title
	res.Description = data.Description.Value
	res.InputDescription = data.InputDescription.Value
	res.OutputDescription = data.OutputDescription.Value
	res.Hint = data.Hint.Value
	res.Difficulty = "Normal"
	res.Visible = false
	res.UseSPJ = false
	res.Shared = false
	res.ProblemCase = make([]dto.ProblemCase, len(data.ProblemCase))
	for i, j := 0, len(data.ProblemCase); i < j; i++ {
		in, err := ioutil.ReadFile(path + "/testcase/" + data.ProblemCase[i].InputName)
		if err != nil {
			log.Error("%v\n", err)
			return dto.Problem{}, err
		}
		out, err := ioutil.ReadFile(path + "/testcase/" + data.ProblemCase[i].OutputName)
		if err != nil {
			log.Error("%v\n", err)
			return dto.Problem{}, err
		}
		res.ProblemCase[i].Input = string(in)
		res.ProblemCase[i].Output = string(out)
		res.ProblemCase[i].Score = data.ProblemCase[i].Score
	}
	res.Language = []dto.Language{{Id: 1, Name: "C"}, {Id: 2, Name: "Cpp"}, {Id: 3, Name: "Java"}, {Id: 4, Name: "Python"}, {Id: 5, Name: "Go"}}
	res.Sample = data.Sample
	res.Tag = nil
	res.Template = make([]dto.Template, 0)
	if data.Template.C.Prepend != "" || data.Template.C.Template != "" || data.Template.C.Append != "" {
		res.Template = append(res.Template, dto.Template{Id: 0, Pid: 0, Lid: 1, Prepend: data.Template.C.Prepend, Content: data.Template.C.Template, Append: data.Template.C.Append})
	}
	if data.Template.Cpp.Prepend != "" || data.Template.Cpp.Template != "" || data.Template.Cpp.Append != "" {
		res.Template = append(res.Template, dto.Template{Id: 0, Pid: 0, Lid: 2, Prepend: data.Template.Cpp.Prepend, Content: data.Template.Cpp.Template, Append: data.Template.Cpp.Append})
	}
	if data.Template.Java.Prepend != "" || data.Template.Java.Template != "" || data.Template.Java.Append != "" {
		res.Template = append(res.Template, dto.Template{Id: 0, Pid: 0, Lid: 3, Prepend: data.Template.Java.Prepend, Content: data.Template.Java.Template, Append: data.Template.Java.Append})
	}
	if data.Template.Python3.Prepend != "" || data.Template.Python3.Template != "" || data.Template.Python3.Append != "" {
		res.Template = append(res.Template, dto.Template{Id: 0, Pid: 0, Lid: 4, Prepend: data.Template.Python3.Prepend, Content: data.Template.Python3.Template, Append: data.Template.Python3.Append})
	}
	res.Limit = []dto.ProblemLimit{
		{Id: 0, Pid: 0, Lid: 1, MaxCpuTime: data.TimeLimit, MaxRealTime: data.TimeLimit, MaxMemory: data.MemoryLimit * 1024 * 1024, CompMp: 2, SPJMp: 2},
		{Id: 0, Pid: 0, Lid: 2, MaxCpuTime: data.TimeLimit, MaxRealTime: data.TimeLimit, MaxMemory: data.MemoryLimit * 1024 * 1024, CompMp: 2, SPJMp: 2},
		{Id: 0, Pid: 0, Lid: 3, MaxCpuTime: data.TimeLimit, MaxRealTime: data.TimeLimit, MaxMemory: data.MemoryLimit * 1024 * 1024, CompMp: 3, SPJMp: 2},
		{Id: 0, Pid: 0, Lid: 4, MaxCpuTime: data.TimeLimit, MaxRealTime: data.TimeLimit, MaxMemory: data.MemoryLimit * 1024 * 1024, CompMp: 2, SPJMp: 2},
		{Id: 0, Pid: 0, Lid: 5, MaxCpuTime: data.TimeLimit, MaxRealTime: data.TimeLimit, MaxMemory: data.MemoryLimit * 1024 * 1024, CompMp: 2, SPJMp: 2},
	}

	fmt.Printf("%v", StringStruct(&res))
	return res, nil
}

func StringStruct(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return fmt.Sprintf("%+v", v)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "    ")
	if err != nil {
		return fmt.Sprintf("%+v", v)
	}
	return out.String()
}

// ------------------------------------
