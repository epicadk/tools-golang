// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package builder2v1

import (
	"fmt"

	"github.com/spdx/tools-golang/spdx"
	"github.com/spdx/tools-golang/utils"
)

// BuildPackageSection2_1 creates an SPDX Package (version 2.1), returning
// that package or error if any is encountered. Arguments:
//   - packageName: name of package / directory
//   - dirRoot: path to directory to be analyzed
//   - pathsIgnore: slice of strings for filepaths to ignore
func BuildPackageSection2_1(packageName string, dirRoot string, pathsIgnore []string) (*spdx.Package2_1, error) {
	// build the file section first, so we'll have it available
	// for calculating the package verification code
	filepaths, err := utils.GetAllFilePaths(dirRoot, pathsIgnore)
	if err != nil {
		return nil, err
	}

	files := map[spdx.ElementID]*spdx.File2_1{}
	fileNumber := 0
	for _, fp := range filepaths {
		newFile, err := BuildFileSection2_1(fp, dirRoot, fileNumber)
		if err != nil {
			return nil, err
		}
		files[newFile.SPDXIdentifier] = newFile
		fileNumber++
	}

	// get the verification code
	code, err := utils.GetVerificationCode2_1(files, "")
	if err != nil {
		return nil, err
	}

	// now build the package section
	pkg := &spdx.Package2_1{
		Name:                      packageName,
		SPDXIdentifier:            spdx.ElementID(fmt.Sprintf("Package-%s", packageName)),
		DownloadLocation:          "NOASSERTION",
		FilesAnalyzed:             true,
		IsFilesAnalyzedTagPresent: true,
		VerificationCode:          code,
		LicenseConcluded:          "NOASSERTION",
		LicenseInfoFromFiles:      []string{},
		LicenseDeclared:           "NOASSERTION",
		CopyrightText:             "NOASSERTION",
		Files:                     files,
	}

	return pkg, nil
}
