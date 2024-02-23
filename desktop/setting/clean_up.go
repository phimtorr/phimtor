package setting

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func CleanUpStorage(setting Settings) error {
	if setting.GetDeleteAfterClosed() {
		dif, err := os.ReadDir(setting.GetCurrentDataDir())
		if err != nil {
			return fmt.Errorf("read data directory: %w", err)
		}
		var errs []error
		for _, f := range dif {
			if err := os.RemoveAll(filepath.Join(setting.GetCurrentDataDir(), f.Name())); err != nil {
				errs = append(errs, fmt.Errorf("remove %s: %w", f.Name(), err))
			}
		}
		return errors.Join(errs...)
	}
	return nil
}
