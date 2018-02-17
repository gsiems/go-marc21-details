// Copyright 2017-2018 Gregory Siems. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package details

import "github.com/gsiems/go-marc21/pkg/marc21"

/*
http://www.loc.gov/marc/bibliographic/bdintro.html

    Leader - Data elements that primarily provide information for the
    processing of the record. The data elements contain numbers or
    coded values and are identified by relative character position. The
    Leader is fixed in length at 24 character positions and is the
    first field of a MARC record.

Also:
    http://www.loc.gov/marc/holdings/hdleader.html
    http://www.loc.gov/marc/authority/adleader.html
    http://www.loc.gov/marc/classification/cdleader.html
    http://www.loc.gov/marc/community/cileader.html

While the general leader layout is the same for the different MARC formats
there are differences.

MARC 21 Bibliography
    00-04 - Record length
    05 - Record status
    06 - Type of record
    07 - Bibliographic level
    08 - Type of control
    09 - Character coding scheme
    10 - Indicator count
    11 - Subfield code count
    12-16 - Base address of data
    17 - Encoding level
    18 - Descriptive cataloging form
    19 - Multipart resource record level
    20 - Length of the length-of-field portion
    21 - Length of the starting-character-position portion
    22 - Length of the implementation-defined portion
    23 - Undefined

MARC 21 Holdings
    00-04 - Record length
    05 - Record status
    06 - Type of record
    07-08 - Undefined character positions
    09 - Character coding scheme
    10 - Indicator count
    11 - Subfield code length
    12-16 - Base address of data
    17 - Encoding level
    18 - Item information in record
    19 - Undefined character position
    20 - Length of the length-of-field portion
    21 - Length of the starting-character-position portion
    22 - Length of the implementation-defined portion
    23 - Undefined

MARC 21 Authority
    00-04 - Record length
    05 - Record status
    06 - Type of record
    07-08 - Undefined character positions
    09 - Character coding scheme
    10 - Indicator count
    11 - Subfield code length
    12-16 - Base address of data
    17 - Encoding level
    18 - Punctuation policy
    19 - Undefined
    20 - Length of the length-of-field portion
    21 - Length of the starting-character-position portion
    22 - Length of the implementation-defined portion
    23 - Undefined

MARC 21 Classification
    00-04 - Record length
    05 - Record status
    06 - Type of record
    07-08 - Undefined character positions
    09 - Character coding scheme
    10 - Indicator count
    11 - Subfield code length
    12-16 - Base address of data
    17 - Encoding level
    18-19 - Undefined character positions
    20 - Length of the length-of-field portion
    21 - Length of the starting-character-position portion
    22 - Length of the implementation-defined portion
    23 - Undefined

MARC 21 Community Information
    00-04 - Record length
    05 - Record status
    06 - Type of record
    07 - Kind of data
    08 - Undefined character position
    09 - Character coding scheme
    10 - Indicator count
    11 - Subfield code length
    12-16 - Base address of data
    17-19 - Undefined character positions
    20 - Length of the length-of-field portion
    21 - Length of the starting-character-position portion
    22 - Length of the implementation-defined portion
    23 - Undefined

*/

type LdrDesc map[string]CodeValue

var characterCodingScheme = map[string]string{
	" ": "MARC-8",
	"a": "UCS/Unicode",
}

////////////////////////////////////////////////////////////////////////
// Authority
////////////////////////////////////////////////////////////////////////
var authorityRecordStatus = map[string]string{
	"a": "Increase in encoding level",
	"c": "Corrected or revised",
	"d": "Deleted",
	"n": "New",
	"o": "Obsolete",
	"s": "Deleted; heading split into two or more headings",
	"x": "Deleted; heading replaced by another heading",
}
var authorityTypeOfRecord = map[string]string{
	"z": "Authority data",
}
var authorityEncodingLevel = map[string]string{
	"n": "Complete authority record",
	"o": "Incomplete authority record",
}
var authorityPunctuationPolicy = map[string]string{
	" ": "No information provided",
	"c": "Punctuation omitted",
	"i": "Punctuation included",
	"u": "Unknown",
}

func parseAuthorityLdr(s string) (ldr LdrDesc) {

	ldr = make(LdrDesc)

	var c string
	var l string

	ldr["(00/05) Record length"] = CodeValue{Code: pluckBytes(s, 0, 5), Label: "", Offset: 0, Width: 5}

	c, l = codeLookup(authorityRecordStatus, s, 5, 1)
	ldr["(05/01) Record status"] = CodeValue{Code: c, Label: l, Offset: 5, Width: 1}

	c, l = codeLookup(authorityTypeOfRecord, s, 6, 1)
	ldr["(06/01) Type of record"] = CodeValue{Code: c, Label: l, Offset: 6, Width: 1}

	ldr["(07/02) Undefined character positions"] = CodeValue{Code: pluckBytes(s, 7, 2), Label: "", Offset: 7, Width: 2}

	c, l = codeLookup(characterCodingScheme, s, 9, 1)
	ldr["(09/01) Character coding scheme"] = CodeValue{Code: c, Label: l, Offset: 9, Width: 1}

	ldr["(10/01) Indicator count"] = CodeValue{Code: pluckBytes(s, 10, 1), Label: "", Offset: 10, Width: 1}
	ldr["(11/01) Subfield code length"] = CodeValue{Code: pluckBytes(s, 11, 1), Label: "", Offset: 11, Width: 1}
	ldr["(12/05) Base address of data"] = CodeValue{Code: pluckBytes(s, 12, 5), Label: "", Offset: 12, Width: 5}

	c, l = codeLookup(authorityEncodingLevel, s, 17, 1)
	ldr["(17/01) Encoding level"] = CodeValue{Code: c, Label: l, Offset: 17, Width: 1}

	c, l = codeLookup(authorityPunctuationPolicy, s, 18, 1)
	ldr["(18/01) Punctuation policy"] = CodeValue{Code: c, Label: l, Offset: 18, Width: 1}

	ldr["(19/01) Undefined"] = CodeValue{Code: pluckBytes(s, 19, 1), Label: "", Offset: 19, Width: 1}
	ldr["(20/01) Length of the length-of-field portion"] = CodeValue{Code: pluckBytes(s, 20, 1), Label: "", Offset: 20, Width: 1}
	ldr["(21/01) Length of the starting-character-position portion"] = CodeValue{Code: pluckBytes(s, 21, 1), Label: "", Offset: 21, Width: 1}
	ldr["(22/01) Length of the implementation-defined portion"] = CodeValue{Code: pluckBytes(s, 22, 1), Label: "", Offset: 22, Width: 1}
	ldr["(23/01) Undefined"] = CodeValue{Code: pluckBytes(s, 23, 1), Label: "", Offset: 23, Width: 1}

	return ldr
}

////////////////////////////////////////////////////////////////////////
// Bibliography
////////////////////////////////////////////////////////////////////////
var bibliographyRecordStatus = map[string]string{
	"a": "Increase in encoding level",
	"c": "Corrected or revised",
	"d": "Deleted",
	"n": "New",
	"p": "Increase in encoding level from prepublication",
}
var bibliographyTypeOfRecord = map[string]string{
	"a": "Language material",
	"c": "Notated music",
	"d": "Manuscript notated music",
	"e": "Cartographic material",
	"f": "Manuscript cartographic material",
	"g": "Projected medium",
	"i": "Nonmusical sound recording",
	"j": "Musical sound recording",
	"k": "Two-dimensional nonprojectable graphic",
	"m": "Computer file",
	"o": "Kit",
	"p": "Mixed materials",
	"r": "Three-dimensional artifact or naturally occurring object",
	"t": "Manuscript language material",
}
var bibliographyBibliographicLevel = map[string]string{
	"a": "Monographic component part",
	"b": "Serial component part",
	"c": "Collection",
	"d": "Subunit",
	"i": "Integrating resource",
	"m": "Monograph/Item",
	"s": "Serial",
}
var bibliographyTypeOfControl = map[string]string{
	" ": "No specified type",
	"a": "Archival",
}
var bibliographyEncodingLevel = map[string]string{
	" ": "Full level",
	"1": "Full level, material not examined",
	"2": "Less-than-full level, material not examined",
	"3": "Abbreviated level",
	"4": "Core level",
	"5": "Partial (preliminary) level",
	"7": "Minimal level",
	"8": "Prepublication level",
	"u": "Unknown",
	"z": "Not applicable",
}
var bibliographyDescriptiveCatalogingForm = map[string]string{
	" ": "Non-ISBD",
	"a": "AACR 2",
	"c": "ISBD punctuation omitted",
	"i": "ISBD punctuation included",
	"n": "Non-ISBD punctuation omitted",
	"u": "Unknown",
}
var bibliographyMultipartResourceRecordLevel = map[string]string{
	" ": "Not specified or not applicable",
	"a": "Set",
	"b": "Part with independent title",
	"c": "Part with dependent title",
}

func parseBibliographyLdr(s string) (ldr LdrDesc) {

	ldr = make(LdrDesc)

	var c string
	var l string

	ldr["(00/05) Logical record length"] = CodeValue{Code: pluckBytes(s, 0, 5), Label: "", Offset: 0, Width: 5}

	c, l = codeLookup(bibliographyRecordStatus, s, 5, 1)
	ldr["(05/01) Record status"] = CodeValue{Code: c, Label: l, Offset: 5, Width: 1}

	c, l = codeLookup(bibliographyTypeOfRecord, s, 6, 1)
	ldr["(06/01) Type of record"] = CodeValue{Code: c, Label: l, Offset: 6, Width: 1}

	c, l = codeLookup(bibliographyBibliographicLevel, s, 7, 1)
	ldr["(07/01) Bibliographic level"] = CodeValue{Code: c, Label: l, Offset: 7, Width: 1}

	c, l = codeLookup(bibliographyTypeOfControl, s, 8, 1)
	ldr["(08/01) Type of control"] = CodeValue{Code: c, Label: l, Offset: 8, Width: 1}

	c, l = codeLookup(characterCodingScheme, s, 9, 1)
	ldr["(09/01) Character coding scheme"] = CodeValue{Code: c, Label: l, Offset: 9, Width: 1}

	ldr["(10/01) Indicator count"] = CodeValue{Code: pluckBytes(s, 10, 1), Label: "", Offset: 10, Width: 1}
	ldr["(11/01) Subfield code count"] = CodeValue{Code: pluckBytes(s, 11, 1), Label: "", Offset: 11, Width: 1}
	ldr["(12/05) Base address of data"] = CodeValue{Code: pluckBytes(s, 12, 5), Label: "", Offset: 12, Width: 5}

	c, l = codeLookup(bibliographyEncodingLevel, s, 17, 1)
	ldr["(17/01) Encoding level"] = CodeValue{Code: c, Label: l, Offset: 17, Width: 1}

	c, l = codeLookup(bibliographyDescriptiveCatalogingForm, s, 18, 1)
	ldr["(18/01) Descriptive cataloging form"] = CodeValue{Code: c, Label: l, Offset: 18, Width: 1}

	c, l = codeLookup(bibliographyMultipartResourceRecordLevel, s, 19, 1)
	ldr["(19/01) Multipart resource record level"] = CodeValue{Code: c, Label: l, Offset: 19, Width: 1}

	ldr["(20/01) Length of the length-of-field portion"] = CodeValue{Code: pluckBytes(s, 20, 1), Label: "", Offset: 20, Width: 1}
	ldr["(21/01) Length of the starting-character-position portion"] = CodeValue{Code: pluckBytes(s, 21, 1), Label: "", Offset: 21, Width: 1}
	ldr["(22/01) Length of the implementation-defined portion"] = CodeValue{Code: pluckBytes(s, 22, 1), Label: "", Offset: 22, Width: 1}
	ldr["(23/01) Undefined Entry map character position"] = CodeValue{Code: pluckBytes(s, 23, 1), Label: "", Offset: 23, Width: 1}

	return ldr
}

////////////////////////////////////////////////////////////////////////
// Classification
////////////////////////////////////////////////////////////////////////
var classificationRecordStatus = map[string]string{
	"a": "Increase in encoding level",
	"c": "Corrected or revised",
	"d": "Deleted",
	"n": "New",
}
var classificationTypeOfRecord = map[string]string{
	"w": "Classification data",
}
var classificationEncodingLevel = map[string]string{
	"n": "Complete classification record",
	"o": "Incomplete classification record",
}

func parseClassificationLdr(s string) (ldr LdrDesc) {

	ldr = make(LdrDesc)

	var c string
	var l string

	ldr["(00/05) Record length"] = CodeValue{Code: pluckBytes(s, 0, 5), Label: "", Offset: 0, Width: 5}

	c, l = codeLookup(classificationRecordStatus, s, 5, 1)
	ldr["(05/01) Record status"] = CodeValue{Code: c, Label: l, Offset: 5, Width: 1}

	c, l = codeLookup(classificationTypeOfRecord, s, 6, 1)
	ldr["(06/01) Type of record"] = CodeValue{Code: c, Label: l, Offset: 6, Width: 1}

	ldr["(07/02) Undefined character positions"] = CodeValue{Code: pluckBytes(s, 7, 2), Label: "", Offset: 7, Width: 2}

	c, l = codeLookup(characterCodingScheme, s, 9, 1)
	ldr["(09/01) Character coding scheme"] = CodeValue{Code: c, Label: l, Offset: 9, Width: 1}

	ldr["(10/01) Indicator count"] = CodeValue{Code: pluckBytes(s, 10, 1), Label: "", Offset: 10, Width: 1}
	ldr["(11/01) Subfield code length"] = CodeValue{Code: pluckBytes(s, 11, 1), Label: "", Offset: 11, Width: 1}
	ldr["(12/05) Base address of data"] = CodeValue{Code: pluckBytes(s, 12, 5), Label: "", Offset: 12, Width: 5}

	c, l = codeLookup(classificationEncodingLevel, s, 17, 1)
	ldr["(17/01) Encoding level"] = CodeValue{Code: c, Label: l, Offset: 17, Width: 1}

	ldr["(18/02) Undefined character positions"] = CodeValue{Code: pluckBytes(s, 18, 2), Label: "", Offset: 18, Width: 2}
	ldr["(20/01) Length of the length-of-field portion"] = CodeValue{Code: pluckBytes(s, 20, 1), Label: "", Offset: 20, Width: 1}
	ldr["(21/01) Length of the starting-character-position portion"] = CodeValue{Code: pluckBytes(s, 21, 1), Label: "", Offset: 21, Width: 1}
	ldr["(22/01) Length of the implementation-defined portion"] = CodeValue{Code: pluckBytes(s, 22, 1), Label: "", Offset: 22, Width: 1}
	ldr["(23/01) Undefined"] = CodeValue{Code: pluckBytes(s, 23, 1), Label: "", Offset: 23, Width: 1}

	return ldr
}

////////////////////////////////////////////////////////////////////////
// Community
////////////////////////////////////////////////////////////////////////
var communityRecordStatus = map[string]string{
	"c": "Corrected or revised",
	"d": "Deleted",
	"n": "New",
}
var communityTypeOfRecord = map[string]string{
	"q": "Community information",
}
var communityKindOfData = map[string]string{
	"n": "Individual",
	"o": "Organization",
	"p": "Program or service",
	"q": "Event",
	"z": "Other",
}

func parseCommunityLdr(s string) (ldr LdrDesc) {

	ldr = make(LdrDesc)

	var c string
	var l string

	ldr["(00/05) Record length"] = CodeValue{Code: pluckBytes(s, 0, 5), Label: "", Offset: 0, Width: 5}

	c, l = codeLookup(communityRecordStatus, s, 5, 1)
	ldr["(05/01) Record status"] = CodeValue{Code: c, Label: l, Offset: 5, Width: 1}

	c, l = codeLookup(communityTypeOfRecord, s, 6, 1)
	ldr["(06/01) Type of record"] = CodeValue{Code: c, Label: l, Offset: 6, Width: 1}

	c, l = codeLookup(communityKindOfData, s, 7, 1)
	ldr["(07/01) Kind of data"] = CodeValue{Code: c, Label: l, Offset: 7, Width: 1}

	ldr["(08/01) Undefined character positions"] = CodeValue{Code: pluckBytes(s, 8, 1), Label: "", Offset: 8, Width: 1}

	c, l = codeLookup(characterCodingScheme, s, 9, 1)
	ldr["(09/01) Character coding scheme"] = CodeValue{Code: c, Label: l, Offset: 9, Width: 1}

	ldr["(10/01) Indicator count"] = CodeValue{Code: pluckBytes(s, 10, 1), Label: "", Offset: 10, Width: 1}
	ldr["(11/01) Subfield code length"] = CodeValue{Code: pluckBytes(s, 11, 1), Label: "", Offset: 11, Width: 1}
	ldr["(12/05) Base address of data"] = CodeValue{Code: pluckBytes(s, 12, 5), Label: "", Offset: 12, Width: 5}
	ldr["(17/03) Undefined character positions"] = CodeValue{Code: pluckBytes(s, 17, 3), Label: "", Offset: 17, Width: 3}
	ldr["(20/01) Length of the length-of-field portion"] = CodeValue{Code: pluckBytes(s, 20, 1), Label: "", Offset: 20, Width: 1}
	ldr["(21/01) Length of the starting-character-position portion"] = CodeValue{Code: pluckBytes(s, 21, 1), Label: "", Offset: 21, Width: 1}
	ldr["(22/01) Length of the implementation-defined portion"] = CodeValue{Code: pluckBytes(s, 22, 1), Label: "", Offset: 22, Width: 1}
	ldr["(23/01) Undefined"] = CodeValue{Code: pluckBytes(s, 23, 1), Label: "", Offset: 23, Width: 1}

	return ldr
}

////////////////////////////////////////////////////////////////////////
// Holdings
////////////////////////////////////////////////////////////////////////
var holdingsRecordStatus = map[string]string{
	"c": "Corrected or revised",
	"d": "Deleted",
	"n": "New",
}
var holdingsTypeOfRecord = map[string]string{
	"u": "Unknown",
	"v": "Multipart item holdings",
	"x": "Single-part item holdings",
	"y": "Serial item holdings",
}
var holdingsEncodingLevel = map[string]string{
	"1": "Holdings level 1",
	"2": "Holdings level 2",
	"3": "Holdings level 3",
	"4": "Holdings level 4",
	"5": "Holdings level 4 with piece designation",
	"m": "Mixed level",
	"u": "Unknown",
	"z": "Other level",
}
var holdingsItemInformationInRecord = map[string]string{
	"i": "Item information",
	"n": "No item information",
}

func parseHoldingsLdr(s string) (ldr LdrDesc) {

	ldr = make(LdrDesc)

	var c string
	var l string

	ldr["(00/05) Record length"] = CodeValue{Code: pluckBytes(s, 0, 5), Label: "", Offset: 0, Width: 5}

	c, l = codeLookup(holdingsRecordStatus, s, 5, 1)
	ldr["(05/01) Record status"] = CodeValue{Code: c, Label: l, Offset: 5, Width: 1}

	c, l = codeLookup(holdingsTypeOfRecord, s, 6, 1)
	ldr["(06/01) Type of record"] = CodeValue{Code: c, Label: l, Offset: 6, Width: 1}

	ldr["(07/02) Undefined character positions"] = CodeValue{Code: pluckBytes(s, 7, 2), Label: "", Offset: 7, Width: 2}

	c, l = codeLookup(characterCodingScheme, s, 9, 1)
	ldr["(09/01) Character coding scheme"] = CodeValue{Code: c, Label: l, Offset: 9, Width: 1}

	ldr["(10/01) Indicator count"] = CodeValue{Code: pluckBytes(s, 10, 1), Label: "", Offset: 10, Width: 1}
	ldr["(11/01) Subfield code length"] = CodeValue{Code: pluckBytes(s, 11, 1), Label: "", Offset: 11, Width: 1}
	ldr["(12/05) Base address of data"] = CodeValue{Code: pluckBytes(s, 12, 5), Label: "", Offset: 12, Width: 5}

	c, l = codeLookup(holdingsEncodingLevel, s, 17, 1)
	ldr["(17/01) Encoding level"] = CodeValue{Code: c, Label: l, Offset: 17, Width: 1}

	c, l = codeLookup(holdingsItemInformationInRecord, s, 18, 1)
	ldr["(18/01) Item information in record"] = CodeValue{Code: c, Label: l, Offset: 18, Width: 1}

	ldr["(19/01) Undefined character position"] = CodeValue{Code: pluckBytes(s, 19, 1), Label: "", Offset: 19, Width: 1}
	ldr["(20/01) Length of the length-of-field portion"] = CodeValue{Code: pluckBytes(s, 20, 1), Label: "", Offset: 20, Width: 1}
	ldr["(21/01) Length of the starting-character-position portion"] = CodeValue{Code: pluckBytes(s, 21, 1), Label: "", Offset: 21, Width: 1}
	ldr["(22/01) Length of the implementation-defined portion"] = CodeValue{Code: pluckBytes(s, 22, 1), Label: "", Offset: 22, Width: 1}
	ldr["(23/01) Undefined"] = CodeValue{Code: pluckBytes(s, 23, 1), Label: "", Offset: 23, Width: 1}

	return ldr
}

func ParseLeader(rec marc21.Record) (ldr LdrDesc) {

	rf := rec.RecordFormat()

	switch rf {
	case marc21.Bibliography:
		return parseBibliographyLdr(rec.Leader.Text)
	case marc21.Holdings:
		return parseHoldingsLdr(rec.Leader.Text)
	case marc21.Community:
		return parseCommunityLdr(rec.Leader.Text)
	case marc21.Authority:
		return parseAuthorityLdr(rec.Leader.Text)
	case marc21.Classification:
		return parseClassificationLdr(rec.Leader.Text)
	}

	ldr = make(LdrDesc)

	return ldr
}
