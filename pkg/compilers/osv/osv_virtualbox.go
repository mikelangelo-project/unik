package osv

import (
	"io/ioutil"
	"os"

	"github.com/emc-advanced-dev/pkg/errors"
	"github.com/emc-advanced-dev/unik/pkg/types"
)

const OSV_VIRTUALBOX_MEMORY = 512

type VirtualboxCompilerHelper struct {
	CompilerHelperBase
}

func (b *VirtualboxCompilerHelper) Convert(params ConvertParams) (*types.RawImage, error) {
	// Convert to WMDK format.
	resultFile, err := ioutil.TempFile("", "osv-boot.vmdk.")
	if err != nil {
		return nil, errors.New("failed to create tmpfile for result", err)
	}
	defer func() {
		if err != nil && !params.CompileParams.NoCleanup {
			os.Remove(resultFile.Name())
		}
	}()
	if err := os.Rename(params.CapstanImagePath, resultFile.Name()); err != nil {
		return nil, errors.New("failed to rename result file", err)
	}

	return &types.RawImage{
		LocalImagePath: resultFile.Name(),
		StageSpec: types.StageSpec{
			ImageFormat: types.ImageFormat_QCOW2,
		},
		RunSpec: types.RunSpec{
			DeviceMappings: []types.DeviceMapping{
				types.DeviceMapping{MountPoint: "/", DeviceName: "/dev/sda1"},
			},
			StorageDriver:         types.StorageDriver_SATA,
			DefaultInstanceMemory: OSV_VIRTUALBOX_MEMORY,
		},
	}, nil
}
