package impl

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/mniak/bench/domain"
	"github.com/mniak/bench/internal/mocks"
	"github.com/stretchr/testify/suite"
)

type ToolchainProducerSuite struct {
	suite.Suite
}

func TestToolchainProducer(t *testing.T) {
	suite.Run(t, new(ToolchainProducerSuite))
}

func (suite *ToolchainProducerSuite) Test_WhenFileHasRunnableExtension_AndSourceExists_ShouldFind() {
	ctrl := gomock.NewController(suite.T())
	defer ctrl.Finish()

	binaryExt := "." + gofakeit.Word()
	sourceExt := "." + gofakeit.Word()
	filenameWithoutExt := gofakeit.Word()

	tempdir, err := ioutil.TempDir(os.TempDir(), "test_*")
	suite.Require().NoError(err)
	defer os.RemoveAll(tempdir)

	binaryFile, err := os.Create(filepath.Join(tempdir, filenameWithoutExt+binaryExt))
	suite.Require().NoError(err)
	defer binaryFile.Close()

	sourceFile, err := os.Create(filepath.Join(tempdir, filenameWithoutExt+sourceExt))
	suite.Require().NoError(err)
	defer sourceFile.Close()

	tchain := mocks.NewMockToolchain(ctrl)
	tchain.EXPECT().
		InputExtensions().
		Return([]string{sourceExt})
	tchain.EXPECT().
		OutputExtension().
		Return(binaryExt)

	sut := NewToolchainProducerFromExtensionMap([]domain.Toolchain{
		tchain,
	})

	result, err := sut.Produce(binaryFile.Name())
	suite.Require().NoError(err)
	suite.Same(result, tchain)
}
