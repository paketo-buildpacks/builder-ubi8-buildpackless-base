package smoke_test

import (
	"encoding/json"
	"flag"
	"os"
	"testing"
	"time"

	"github.com/onsi/gomega/format"
	"github.com/paketo-buildpacks/occam"
	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"

	. "github.com/onsi/gomega"
)

var (
	Builder string

	config struct {
		Procfile           string `json:"procfile"`
		GoDist             string `json:"go-dist"`
		BuildPlan          string `json:"build-plan"`
		NodeEngine         string `json:"node-engine"`
		NPMInstall         string `json:"npm-install"`
		NPMStart           string `json:"npm-start"`
		UbiNodejsExtension string `json:"ubi-nodejs-extension"`
	}
)

func init() {
	flag.StringVar(&Builder, "name", "", "")
}

func TestSmoke(t *testing.T) {
	format.MaxLength = 0
	Expect := NewWithT(t).Expect

	flag.Parse()

	Expect(Builder).NotTo(Equal(""))

	SetDefaultEventuallyTimeout(60 * time.Second)

	file, err := os.Open("../smoke.json")
	Expect(err).NotTo(HaveOccurred())

	Expect(json.NewDecoder(file).Decode(&config)).To(Succeed())
	Expect(file.Close()).To(Succeed())

	Expect(occam.NewDocker().Pull.Execute(config.Procfile))
	Expect(occam.NewDocker().Pull.Execute(config.GoDist))
	Expect(occam.NewDocker().Pull.Execute(config.BuildPlan))
	Expect(occam.NewDocker().Pull.Execute(config.NodeEngine))
	Expect(occam.NewDocker().Pull.Execute(config.NPMInstall))
	Expect(occam.NewDocker().Pull.Execute(config.NPMStart))
	Expect(occam.NewDocker().Pull.Execute(config.UbiNodejsExtension))

	suite := spec.New("Buildpackless Smoke", spec.Parallel(), spec.Report(report.Terminal{}))
	suite("Procfile", testProcfile)
	suite("Node.js", testNodejs)
	suite.Run(t)
}
