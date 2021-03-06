package main

import (
	"flag"
	"time"

	"github.com/ShiryaevAnton/ses/config"
	"github.com/ShiryaevAnton/ses/control"
	"github.com/ShiryaevAnton/ses/excel"
)

func main() {

	var configPath string
	var excelPath string

	flag.StringVar(&configPath, "cp", "config.toml", "Path to config file")
	flag.StringVar(&excelPath, "ep", "", "Path to excel file")

	flag.Parse()

	conf, err := config.GetConfig(configPath)
	if err != nil {
		panic(err)
	}

	r := excel.NewReader(
		excelPath,
		conf.SheetConfig.Name,
		conf.SheetConfig.StartRow,
		conf.SheetConfig.NameCol,
		conf.SheetConfig.CodeCol,
		conf.SheetConfig.PhoneCol)

	users, err := r.GetUsers()
	if err != nil {
		panic(err)
	}

	c := control.NewControl()

	PrintToSes(c, conf, users)

}

func PrintToSes(c *control.Control, config *config.Config, users []excel.User) {

	time.Sleep(10 * time.Second)

	for _, user := range users {

		c.TypeString(user.Code)
		c.Press(config.Seqs[0].Seq)

		c.TypeString(user.Name)
		c.Press(config.Seqs[1].Seq)

		c.TypeString(user.Phone)
		c.Press(config.Seqs[2].Seq)

	}
}
