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

	Describe(".UsersFromJSON()", func() {
		var (
			json string
		)

		BeforeEach(func() {
			json = `{"shim-delay": 24, "users": {"user1": {"ssh_public_key": "ssh-key keyforuser1 user1@example.com", "name": "User 1", "perm": "ALL=NOPASSWD: ALL", "preferred_shell": "/bin/bash"}, "user2": {"ssh_public_key": "ssh-key keyforuser2 user2@example.com", "name": "User 2", "preferred_shell": "", "perm": ""}}}`

		})

		It("is a slice of User instances", func() {
			var expected []*User

			Expect(UsersFromJSON(json)).To(BeAssignableToTypeOf(expected))
		})

		It("contains a User for each represented in the provided JSON", func() {
			Expect(len(UsersFromJSON(json))).To(Equal(2))
		})
	})
})
