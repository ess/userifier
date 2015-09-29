package config_test

import (
	. "github.com/ess/userifier/config"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config", func() {
  Describe(".New()", func() {
    var (
      new_config *Config
      err error
    )

    Context("with a good config file", func() {
      BeforeEach(func() {
        new_config, err = New("../config.toml.example")
      })

      It("returns no error", func() {
        Expect(err).To(BeNil())
      })

      It("populates the Config instance's ApiId", func() {
        Expect(new_config.ApiId).To(Equal("USERIFY API ID"))
      })

      It("populates the Config instance's ApiKey", func() {
        Expect(new_config.ApiKey).To(Equal("USERIFY API KEY"))
      })
    })

    Context("with a missing config file", func() {
      BeforeEach(func() {
        new_config, err = New("lolnoconfig.toml")
      })

      It("returns an error", func() {
        Expect(err).NotTo(BeNil())
      })
    })
  })
})
