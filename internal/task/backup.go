package task

import (
	"github.com/xlshiz/bonclay/internal/core"
	"github.com/xlshiz/bonclay/internal/file"
)

// Backup spec ('source:target' pairs) defined in the configuration spec to copy
// the sources to the targets.
func backupSpec(config *core.Configuration) []string {
	errors := make([]string, 0, len(config.Spec))
	for src, dst := range config.Spec {
		err := file.CopySpec(src, dst, config.Backup.Overwrite)
		if err != nil {
			errors = append(errors, err.Error())
			core.WriteTaskFailure(src, dst)
			continue
		}

		core.WriteTaskSuccess(src, dst)
	}

	return errors
}

// Backup glob defined in the configuration spec to copy
// the sources to the targets.
func backupGlob(config *core.Configuration) []string {
	errors := make([]string, 0, len(config.Spec))
	for src, dst := range config.Glob {
		err := file.CopyGlob(src, dst, config.Backup.Overwrite)
		if err != nil {
			errors = append(errors, err.Error())
			core.WriteTaskFailure(src, dst.Dst)
			continue
		}

		core.WriteTaskSuccess(src+" <-"+dst.Filter+"->", dst.Dst)
	}

	return errors
}

// Backup spec and glob
func Backup(config *core.Configuration) {
	core.WriteTaskHeader("backup")

	specErrors := backupSpec(config)
	globErrors := backupGlob(config)

	errors := append(specErrors, globErrors...)
	core.WriteTaskFooter("backup", len(errors) == 0)
	core.WriteTaskErrors(errors)
}
