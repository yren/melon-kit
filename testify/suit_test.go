package testify

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

// ExampleTestSuit, define suite
type ExampleTestSuit struct{
	suite.Suite
	VarShould5 int
}

// SetupTest, make sure VarShould5 set 5
func (s *ExampleTestSuit) SetupTest() {
	s.VarShould5 = 5
}

// TestExample, all method begin with "Test", run as tests within a suite
func (s *ExampleTestSuit) TestExample() {
	assert.Equal(s.T(), 5, s.VarShould5)
}

// create normal test, pass our suite to suite.Run
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ExampleTestSuit))
}