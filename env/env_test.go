package env_test

import (
	. "github.com/EngineerBetter/cf-plex/env"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("env", func() {
	Describe("getCoordinate", func() {
		It("returns coordinate for an API", func() {
			cfEnv := "username:password@api.com"
			coords := GetCoordinate(cfEnv)
			Ω(coords.Username).Should(Equal("username"))
			Ω(coords.Password).Should(Equal("password"))
			Ω(coords.Api).Should(Equal("api.com"))
		})
	})

	Describe("getTriples", func() {
		It("returns coordinates for an API", func() {
			cfEnvs := "user1:pass1@api1.com;user2:pass2@api2.com"
			triples := GetTriples(cfEnvs)
			Ω(triples[0]).Should(Equal("user1:pass1@api1.com"))
			Ω(triples[1]).Should(Equal("user2:pass2@api2.com"))
		})
	})

	Describe("GetCoordinates", func() {
		It("returns coordinates for many APIs", func() {
			cfEnvs := "user1:pass1@api1.com;user2:pass2@api2.com"
			coords := GetCoordinates(cfEnvs)
			Ω(coords[0]).Should(Equal(Coord{Username: "user1", Password: "pass1", Api: "api1.com"}))
			Ω(coords[1]).Should(Equal(Coord{Username: "user2", Password: "pass2", Api: "api2.com"}))
		})
	})
})
