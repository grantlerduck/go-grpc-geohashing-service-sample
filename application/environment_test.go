package application_test

import (
	"os"

	"github.com/grntlduck-cloud/go-grpc-geohasing-service-sample/application"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test application environment load", func ()  {
	When("table name is set", func ()  {
		tableName := "some-table"
		os.Setenv("DYNAMO_TABLE_NAME", tableName)
		It("does not panic", func ()  {
			appEnv := application.NewServiceEnv()
			Expect(appEnv).Should(Not(BeNil()))
			Expect(appEnv.TableName).Should(Equal(tableName))
		})
	})
	
})