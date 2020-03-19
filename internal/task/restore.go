package task

import (
	"github.com/xlshiz/bonclay/internal/core"
	"github.com/xlshiz/bonclay/internal/file"
)

// Restore uses 'source:target' pairs defined in the configuration spec to copy
// the targets to the sources. It is the reverse of Backup.
func restoreSpec(config *core.Configuration) []string {
	errors := make([]string, 0, len(config.Spec))
	for dst, src := range config.Spec {
		err := file.CopySpec(src, dst, config.Restore.Overwrite)
		if err != nil {
			errors = append(errors, err.Error())
			core.WriteTaskFailure(src, dst)
			continue
		}

		core.WriteTaskSuccess(src, dst)
	}

	return errors
}

// Restore uses 'source:target' pairs defined in the configuration spec to copy
// the targets to the sources. It is the reverse of Backup.
func restoreGlob(config *core.Configuration) []string {
	errors := make([]string, 0, len(config.Spec))
	for dst, src := range config.Glob {
		err := file.RestoreGlob(src, dst, config.Backup.Overwrite)
		if err != nil {
			errors = append(errors, err.Error())
			core.WriteTaskFailure(src.Dst, dst)
			continue
		}

		core.WriteTaskSuccess(src.Dst, dst)
	}

	return errors
}

// Restore uses 'source:target' pairs defined in the configuration spec to copy
// the targets to the sources. It is the reverse of Backup.
func Restore(config *core.Configuration) {
	core.WriteTaskHeader("restore")

	specErrors := restoreSpec(config)
	globErrors := restoreGlob(config)
	errors := append(specErrors, globErrors...)

	core.WriteTaskFooter("restore", len(errors) == 0)
	core.WriteTaskErrors(errors)
}
