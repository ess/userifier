package main_test

import (
	. "github.com/ess/userifier/user"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("User", func() {
	var (
		user     *User
		username string
		realname string
		pubkey   string
		shell    string
		perms    string
	)

	BeforeEach(func() {
		username = "test_user"
		realname = "Test User"
		pubkey = "ssh-key abc123 user@example.com"
		shell = "/bin/bash"
		perms = ""

		user = New(username, realname, pubkey, shell, perms)
	})

	Describe("#String()", func() {
		It("includes the User's user name", func() {
			Expect(user.String()).To(ContainSubstring(user.UserName))
		})

		It("includes the User's real name", func() {
			Expect(user.String()).To(
				ContainSubstring("real name: '" + user.RealName + "'"))
		})

		It("includes the User's preferred shell", func() {
			Expect(user.String()).To(
				ContainSubstring("preferred shell: '" + user.PreferredShell + "'"))
		})
	})
})
