package impl

import (
	"io/ioutil"
	"os"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/mniak/bench/internal/mocks"
	"github.com/stretchr/testify/suite"
)

type ToolchainProducerSuite struct {
	suite.Suite
}

func (suite *ToolchainProducerSuite) WhenFileHasRunnableExtension_AndSourceExists_ShouldFind() {
	ctrl := gomock.NewController(suite.T())
	defer ctrl.Finish()

	binaryExt := ".a" + gofakeit.Word()
	sourceExt := ".b" + gofakeit.Word()
	filenameWithoutExt := gofakeit.Word()

	tempdir, err := ioutil.TempDir(os.TempDir(), "test_*")
	suite.Require().NoError(err)
	defer os.RemoveAll(tempdir)

	binaryFile, err := os.CreateTemp(tempdir, filenameWithoutExt+binaryExt)
	suite.Require().NoError(err)
	defer binaryFile.Close()

	sourceFile, err := os.CreateTemp(tempdir, filenameWithoutExt+sourceExt)
	suite.Require().NoError(err)
	defer sourceFile.Close()

	toolchain := mocks.NewMockToolchain(ctrl)
	toolchain.EXPECT().
		Build(sourceFile.Name()).
		Return(binaryFile.Name(), nil)

	sut := NewToolchainProducer()

	sut.Produce(binaryFile.Name())
}
