package smoke_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/paketo-buildpacks/occam"
	"github.com/sclevine/spec"

	. "github.com/onsi/gomega"
)

func testNodejs(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect

		pack   occam.Pack
		docker occam.Docker
	)

	it.Before(func() {
		pack = occam.NewPack().WithVerbose().WithNoColor()
		docker = occam.NewDocker()
	})

	context("when the same image is built twice", func() {
		var (
			imageIDs []string

			name   string
			source string
		)

		it.Before(func() {
			var err error
			name, err = occam.RandomName()
			Expect(err).NotTo(HaveOccurred())
		})

		it.After(func() {
			for _, id := range imageIDs {
				Expect(docker.Image.Remove.Execute(id)).To(Succeed())
			}
			Expect(docker.Volume.Remove.Execute(occam.CacheVolumeNames(name))).To(Succeed())
			Expect(os.RemoveAll(source)).To(Succeed())
		})

		it("completes the second build", func() {
			var err error
			source, err = occam.Source(filepath.Join("testdata", "nodejs", "npm"))
			Expect(err).NotTo(HaveOccurred())

			build := pack.Build.
				WithNetwork("host").
				WithPullPolicy("always").
				WithBuilder(Builder).
				WithExtensions(config.UbiNodejsExtension).
				WithBuildpacks(
					config.NodeEngine,
					config.NPMInstall,
					config.NPMStart,
				)

			_, firstLogs, err := build.Execute(name, source)
			Expect(err).NotTo(HaveOccurred(), firstLogs.String)

			secondImage, secondLogs, err := build.Execute(name, source)
			Expect(err).NotTo(HaveOccurred(), secondLogs.String)
			imageIDs = append(imageIDs, secondImage.ID)
		})
	})
}
