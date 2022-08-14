package p3_suite

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"math/rand"
	"testing"
	"time"
)

type _Suite struct {
	suite.Suite
}

// SetupSuite() 和 TearDownSuite() 全局只会执行一次
// SetupTest() TearDownTest() BeforeTest() AfterTest() 对套件中的每个测试执行一次
func (s *_Suite) AfterTest(suiteName, testName string) {
	fmt.Printf("AferTest: suiteName=%s,testName=%s\n", suiteName, testName)
}

func (s *_Suite) BeforeTest(suiteName, testName string) {
	fmt.Printf("BeforeTest: suiteName=%s,testName=%s\n", suiteName, testName)
}

// SetupSuite() 仅执行一次
func (s *_Suite) SetupSuite() {
	fmt.Printf("SetupSuite() ...\n")
}

// TearDownSuite() 仅执行一次
func (s *_Suite) TearDownSuite() {
	fmt.Printf("TearDowmnSuite()...\n")
}

func (s *_Suite) SetupTest() {
	fmt.Printf("SetupTest()... \n")
}

func (s *_Suite) TearDownTest() {
	fmt.Printf("TearDownTest()... \n")
}

func (s *_Suite) TestSimpleRand() {
	fmt.Printf("TestSimpleRand()... \n")
	ret := SimpleRand(1, 10) // 4.
	assert.Equal(s.T(), ret, int64(10))
}

func (s *_Suite) TestCalculate() {
	fmt.Printf("TestCalculate()... \n")
	ret := Calculate(1) //9.
	assert.Equal(s.T(), ret, 0)
}

// 让 go test 执行测试
func TestAll(t *testing.T) {
	suite.Run(t, new(_Suite))
}

func Calculate(i int) int {
	return 11 - i
}

func SimpleRand(a, b int64) int64 {
	rand.Seed(time.Now().UnixNano())
	n := rand.Int63n(b - a)
	return n + a
}
