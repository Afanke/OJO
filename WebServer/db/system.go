package db

import (
	"github.com/afanke/OJO/WebServer/dto"
	"github.com/ilibs/gosql/v2"
)

type System struct {
}

func (System) GetAll() (*dto.SystemConfig, error) {
	var data dto.SystemConfig
	err := gosql.Get(&data, "select server, port, email, password, name, footer, allow_register from ojo.system_config limit 1")
	return &data, err
}

func (System) GetWebConfig() (*dto.SystemConfig, error) {
	var data dto.SystemConfig
	err := gosql.Get(&data, "select name, footer, allow_register from ojo.system_config limit 1")
	return &data, err
}

func (System) UpdateSMTP(cfg *dto.SystemConfig) error {
	_, err := gosql.Exec(`update ojo.system_config
					set server=?,
					port=?,
					email=?,
					password=?
					limit 1`,
		cfg.Server, cfg.Port, cfg.Email, cfg.Password)
	return err
}

func (System) UpdateWeb(cfg *dto.SystemConfig) error {
	_, err := gosql.Exec(`update ojo.system_config
					set name=?,
					footer=?,
					allow_register=?
					limit 1`,
		cfg.Name, cfg.Footer, cfg.AllowRegister)
	return err
}
