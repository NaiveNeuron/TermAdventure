package levels

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"github.com/rakyll/globalconf"
	"log"
	"os/user"
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
	CurrentLevel *string
}

func NewChallenge(name string) Challenge {
	cfg, err := globalconf.New(name)
	if err != nil {
		log.Fatal(err)
	}

	c := Challenge{
		Name:         name,
		conf:         cfg,
		CurrentLevel: flag.String("level", IndexToID(0, name), "Current Level"),
	}
	return c
}

func (c *Challenge) AddLevel(level Level) {
	c.Levels = append(c.Levels, level)
}

func (c *Challenge) CheckCurrentLevel() bool {
	level := c.IDToIndex(*c.CurrentLevel)
	return CmdOK(c.Levels[level].TestCmd)
}

func (c *Challenge) PrintCurrentLevel() {
	c.Levels[c.IDToIndex(*c.CurrentLevel)].Print()
}

func (c *Challenge) IncreaseLevel() {
	index := c.IDToIndex(*c.CurrentLevel)
	index += 1
	fint := &flagValue{str: IndexToID(index, c.Name)}
	f := &flag.Flag{Name: "level", Value: fint}
	c.conf.Set("", f)
	c.conf.ParseAll()
}

func (c *Challenge) LoadCfg() {
	c.conf.ParseAll()
}

func (c *Challenge) PrintIdentifier() {
	index := c.IDToIndex(*c.CurrentLevel)
	fmt.Printf("[%s %s]", c.Name, c.Levels[index].Name)
}

func (c *Challenge) IDToIndex(id string) int {
	for i := 0; i < len(c.Levels); i++ {
		if id == IndexToID(i, c.Name) {
			return i
		}
	}
	return -1
}

func IndexToID(index int, challenge_name string) string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return GetMD5Hash(fmt.Sprintf("i%sj%dk%sl", challenge_name, index, usr.HomeDir))
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
