// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package parser2v2

import (
	"fmt"
	gordfParser "github.com/spdx/gordf/rdfloader/parser"
	"github.com/spdx/gordf/rdfwriter"
	"github.com/spdx/tools-golang/spdx"
)

func (parser *rdfParser2_2) getExtractedLicensingInfoFromNode(node *gordfParser.Node) (lic ExtractedLicensingInfo, err error) {
	associatedTriples := rdfwriter.FilterTriples(parser.gordfParserObj.Triples, &node.ID, nil, nil)
	var restTriples []*gordfParser.Triple
	for _, triple := range associatedTriples {
		switch triple.Predicate.ID {
		case SPDX_EXTRACTED_TEXT:
			lic.extractedText = triple.Object.ID
		default:
			restTriples = append(restTriples, triple)
		}
	}
	lic.SimpleLicensingInfo, err = parser.getSimpleLicensingInfoFromTriples(restTriples)
	if err != nil {
		return lic, fmt.Errorf("error setting simple licensing information of extracted licensing info: %s", err)
	}
	return lic, nil
}

func (parser *rdfParser2_2) extractedLicenseToOtherLicense(extLicense ExtractedLicensingInfo) (othLic spdx.OtherLicense2_2) {
	othLic.Identifier = extLicense.licenseID
	othLic.ExtractedText = extLicense.extractedText
	othLic.Comment = extLicense.comment
	othLic.CrossReferences = extLicense.seeAlso
	othLic.Name = extLicense.name
	return othLic
}
