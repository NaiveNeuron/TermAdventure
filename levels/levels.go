package levels

import (
	"flag"
	"fmt"
	"github.com/rakyll/globalconf"
	"log"
	"strconv"
)

type flagValue struct {
	str string
}

func (f *flagValue) String() string {
	return f.str
}

func (f *flagValue) Set(value string) error {
	f.str = value
	return nil
}

type Level struct {
	Name        string
	PreTextCmd  string
	Text        string
	PostTextCmd string
	TestCmd     string
}

func (level *Level) Print() {
	terminalized_text := MarkdownToTerminal(level.Text)
	PrintText(terminalized_text)
}

type Challenge struct {
	Name         string
	Levels       []Level
	conf         *globalconf.GlobalConf
	CurrentLevel *int
}

func NewChallenge(name string) Challenge {
	cfg, err := globalconf.New(name)
	if err != nil {
		log.Fatal(err)
	}

	c := Challenge{
		Name:         name,
		conf:         cfg,
		CurrentLevel: flag.Int("level", 0, "Current Level"),
	}
	return c
}

func (c *Challenge) AddLevel(level Level) {
	c.Levels = append(c.Levels, level)
}

func (c *Challenge) CheckCurrentLevel() bool {
	return CmdOK(c.Levels[*c.CurrentLevel].TestCmd)
}

func (c *Challenge) PrintCurrentLevel() {
	c.Levels[*c.CurrentLevel].Print()
}

func (c *Challenge) IncreaseLevel() {
	*c.CurrentLevel += 1
	fint := &flagValue{str: strconv.Itoa(*c.CurrentLevel)}
	f := &flag.Flag{Name: "level", Value: fint}
	c.conf.Set("", f)
}

func (c *Challenge) LoadCfg() {
	c.conf.ParseAll()
}

func (c *Challenge) PrintIdentifier() {
	fmt.Printf("[%s %s]", c.Name, c.Levels[*c.CurrentLevel].Name)
}
