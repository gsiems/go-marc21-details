package details

// Auto-generated code. Do not edit.

////////////////////////////////////////////////////////////////////////
// Authority
var authorityLdrRecordStatus = map[string]string{
	"a": "Increase in encoding level",
	"c": "Corrected or revised",
	"d": "Deleted",
	"n": "New",
	"o": "Obsolete",
	"s": "Deleted; heading split into two or more headings",
	"x": "Deleted; heading replaced by another heading",
}
var authorityLdrTypeOfRecord = map[string]string{
	"z": "Authority data",
}
var authorityLdrCharacterCodingScheme = map[string]string{
	" ": "MARC-8",
	"a": "UCS/Unicode",
}
var authorityLdrEncodingLevel = map[string]string{
	"n": "Complete authority record",
	"o": "Incomplete authority record",
}
var authorityLdrPunctuationPolicy = map[string]string{
	" ": "No information provided",
	"c": "Punctuation omitted",
	"i": "Punctuation included",
	"u": "Unknown",
}

// parseAuthorityLdr parses leader data for Authority records data
func parseAuthorityLdr(s string) (ldr LdrDesc) {

	ldr = make(LdrDesc)

	var c string
	var l string
	ldr["(00/05) Record length"] = CodeValue{Code: pluckBytes(s, 0, 5), Label: "", Offset: 0, Width: 5}
	c, l = codeLookup(authorityLdrRecordStatus, s, 5, 1)
	ldr["(05/01) Record status"] = CodeValue{Code: c, Label: l, Offset: 5, Width: 1}
	c, l = codeLookup(authorityLdrTypeOfRecord, s, 6, 1)
	ldr["(06/01) Type of record"] = CodeValue{Code: c, Label: l, Offset: 6, Width: 1}
	ldr["(07/02) Undefined character positions"] = CodeValue{Code: pluckBytes(s, 7, 2), Label: "", Offset: 7, Width: 2}
	c, l = codeLookup(authorityLdrCharacterCodingScheme, s, 9, 1)
	ldr["(09/01) Character coding scheme"] = CodeValue{Code: c, Label: l, Offset: 9, Width: 1}
	ldr["(10/01) Indicator count"] = CodeValue{Code: pluckBytes(s, 10, 1), Label: "", Offset: 10, Width: 1}
	ldr["(11/01) Subfield code length"] = CodeValue{Code: pluckBytes(s, 11, 1), Label: "", Offset: 11, Width: 1}
	ldr["(12/05) Base address of data"] = CodeValue{Code: pluckBytes(s, 12, 5), Label: "", Offset: 12, Width: 5}
	c, l = codeLookup(authorityLdrEncodingLevel, s, 17, 1)
	ldr["(17/01) Encoding level"] = CodeValue{Code: c, Label: l, Offset: 17, Width: 1}
	c, l = codeLookup(authorityLdrPunctuationPolicy, s, 18, 1)
	ldr["(18/01) Punctuation policy"] = CodeValue{Code: c, Label: l, Offset: 18, Width: 1}
	ldr["(19/01) Undefined"] = CodeValue{Code: pluckBytes(s, 19, 1), Label: "", Offset: 19, Width: 1}
	// (20/04) Entry map
	ldr["(20/01) Length of the length-of-field portion"] = CodeValue{Code: pluckBytes(s, 20, 1), Label: "", Offset: 20, Width: 1}
	ldr["(21/01) Length of the starting-character-position portion"] = CodeValue{Code: pluckBytes(s, 21, 1), Label: "", Offset: 21, Width: 1}
	ldr["(22/01) Length of the implementation-defined portion"] = CodeValue{Code: pluckBytes(s, 22, 1), Label: "", Offset: 22, Width: 1}
	ldr["(23/01) Undefined"] = CodeValue{Code: pluckBytes(s, 23, 1), Label: "", Offset: 23, Width: 1}

	return ldr
}

////////////////////////////////////////////////////////////////////////
// Bibliography
var bibliographyLdrRecordStatus = map[string]string{
	"a": "Increase in encoding level",
	"c": "Corrected or revised",
	"d": "Deleted",
	"n": "New",
	"p": "Increase in encoding level from prepublication",
}
var bibliographyLdrTypeOfRecord = map[string]string{
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
	"p": "Mixed material",
	"r": "Three-dimensional artifact or naturally occurring object",
	"t": "Manuscript language material",
}
var bibliographyLdrBibliographicLevel = map[string]string{
	"a": "Monographic component part",
	"b": "Serial component part",
	"c": "Collection",
	"d": "Subunit",
	"i": "Integrating resource",
	"m": "Monograph/item",
	"s": "Serial",
}
var bibliographyLdrTypeOfControl = map[string]string{
	" ": "No specific type",
	"a": "Archival",
}
var bibliographyLdrCharacterCodingScheme = map[string]string{
	" ": "MARC-8",
	"a": "UCS/Unicode",
}
var bibliographyLdrEncodingLevel = map[string]string{
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
var bibliographyLdrDescriptiveCatalogingForm = map[string]string{
	" ": "Non-ISBD",
	"a": "AACR 2",
	"c": "ISBD punctuation omitted",
	"i": "ISBD punctuation included",
	"n": "Non-ISBD punctuation omitted",
	"u": "Unknown",
}
var bibliographyLdrMultipartResourceRecordLevel = map[string]string{
	" ": "Not specified or not applicable",
	"a": "Set",
	"b": "Part with independent title",
	"c": "Part with dependent title",
}

// parseBibliographyLdr parses leader data for Bibliography records data
func parseBibliographyLdr(s string) (ldr LdrDesc) {

	ldr = make(LdrDesc)

	var c string
	var l string
	ldr["(00/05) Logical record length"] = CodeValue{Code: pluckBytes(s, 0, 5), Label: "", Offset: 0, Width: 5}
	c, l = codeLookup(bibliographyLdrRecordStatus, s, 5, 1)
	ldr["(05/01) Record status"] = CodeValue{Code: c, Label: l, Offset: 5, Width: 1}
	c, l = codeLookup(bibliographyLdrTypeOfRecord, s, 6, 1)
	ldr["(06/01) Type of record"] = CodeValue{Code: c, Label: l, Offset: 6, Width: 1}
	c, l = codeLookup(bibliographyLdrBibliographicLevel, s, 7, 1)
	ldr["(07/01) Bibliographic level"] = CodeValue{Code: c, Label: l, Offset: 7, Width: 1}
	c, l = codeLookup(bibliographyLdrTypeOfControl, s, 8, 1)
	ldr["(08/01) Type of control"] = CodeValue{Code: c, Label: l, Offset: 8, Width: 1}
	c, l = codeLookup(bibliographyLdrCharacterCodingScheme, s, 9, 1)
	ldr["(09/01) Character coding scheme"] = CodeValue{Code: c, Label: l, Offset: 9, Width: 1}
	ldr["(10/01) Indicator count"] = CodeValue{Code: pluckBytes(s, 10, 1), Label: "", Offset: 10, Width: 1}
	ldr["(11/01) Subfield code count"] = CodeValue{Code: pluckBytes(s, 11, 1), Label: "", Offset: 11, Width: 1}
	ldr["(12/05) Base address of data"] = CodeValue{Code: pluckBytes(s, 12, 5), Label: "", Offset: 12, Width: 5}
	c, l = codeLookup(bibliographyLdrEncodingLevel, s, 17, 1)
	ldr["(17/01) Encoding level"] = CodeValue{Code: c, Label: l, Offset: 17, Width: 1}
	c, l = codeLookup(bibliographyLdrDescriptiveCatalogingForm, s, 18, 1)
	ldr["(18/01) Descriptive cataloging form"] = CodeValue{Code: c, Label: l, Offset: 18, Width: 1}
	c, l = codeLookup(bibliographyLdrMultipartResourceRecordLevel, s, 19, 1)
	ldr["(19/01) Multipart resource record level"] = CodeValue{Code: c, Label: l, Offset: 19, Width: 1}
	// (20/04) Entry map
	ldr["(20/01) Length of the length-of-field portion"] = CodeValue{Code: pluckBytes(s, 20, 1), Label: "", Offset: 20, Width: 1}
	ldr["(21/01) Length of the starting-character-position portion"] = CodeValue{Code: pluckBytes(s, 21, 1), Label: "", Offset: 21, Width: 1}
	ldr["(22/01) Length of the implementation-defined portion"] = CodeValue{Code: pluckBytes(s, 22, 1), Label: "", Offset: 22, Width: 1}
	ldr["(23/01) Undefined Entry map character position"] = CodeValue{Code: pluckBytes(s, 23, 1), Label: "", Offset: 23, Width: 1}

	return ldr
}

////////////////////////////////////////////////////////////////////////
// Classification
var classificationLdrRecordStatus = map[string]string{
	"a": "Increase in encoding level",
	"c": "Corrected or revised",
	"d": "Deleted",
	"n": "New",
}
var classificationLdrTypeOfRecord = map[string]string{
	"w": "Classification data",
}
var classificationLdrCharacterCodingScheme = map[string]string{
	" ": "MARC 8",
	"a": "UCS/Unicode",
}
var classificationLdrEncodingLevel = map[string]string{
	"n": "Complete classification record",
	"o": "Incomplete classification record",
}

// parseClassificationLdr parses leader data for Classification records data
func parseClassificationLdr(s string) (ldr LdrDesc) {

	ldr = make(LdrDesc)

	var c string
	var l string
	ldr["(00/05) Record length"] = CodeValue{Code: pluckBytes(s, 0, 5), Label: "", Offset: 0, Width: 5}
	c, l = codeLookup(classificationLdrRecordStatus, s, 5, 1)
	ldr["(05/01) Record status"] = CodeValue{Code: c, Label: l, Offset: 5, Width: 1}
	c, l = codeLookup(classificationLdrTypeOfRecord, s, 6, 1)
	ldr["(06/01) Type of record"] = CodeValue{Code: c, Label: l, Offset: 6, Width: 1}
	ldr["(07/02) Undefined character positions"] = CodeValue{Code: pluckBytes(s, 7, 2), Label: "", Offset: 7, Width: 2}
	c, l = codeLookup(classificationLdrCharacterCodingScheme, s, 9, 1)
	ldr["(09/01) Character coding scheme"] = CodeValue{Code: c, Label: l, Offset: 9, Width: 1}
	ldr["(10/01) Indicator count"] = CodeValue{Code: pluckBytes(s, 10, 1), Label: "", Offset: 10, Width: 1}
	ldr["(11/01) Subfield code length"] = CodeValue{Code: pluckBytes(s, 11, 1), Label: "", Offset: 11, Width: 1}
	ldr["(12/05) Base address of data"] = CodeValue{Code: pluckBytes(s, 12, 5), Label: "", Offset: 12, Width: 5}
	c, l = codeLookup(classificationLdrEncodingLevel, s, 17, 1)
	ldr["(17/01) Encoding level"] = CodeValue{Code: c, Label: l, Offset: 17, Width: 1}
	ldr["(18/02) Undefined character positions"] = CodeValue{Code: pluckBytes(s, 18, 2), Label: "", Offset: 18, Width: 2}
	// (20/04) Entry map
	ldr["(20/01) Length of the length-of-field portion"] = CodeValue{Code: pluckBytes(s, 20, 1), Label: "", Offset: 20, Width: 1}
	ldr["(21/01) Length of the starting-character-position portion"] = CodeValue{Code: pluckBytes(s, 21, 1), Label: "", Offset: 21, Width: 1}
	ldr["(22/01) Length of the implementation-defined portion"] = CodeValue{Code: pluckBytes(s, 22, 1), Label: "", Offset: 22, Width: 1}
	ldr["(23/01) Undefined"] = CodeValue{Code: pluckBytes(s, 23, 1), Label: "", Offset: 23, Width: 1}

	return ldr
}

////////////////////////////////////////////////////////////////////////
// Community
var communityLdrRecordStatus = map[string]string{
	"c": "Corrected or revised",
	"d": "Deleted",
	"n": "New",
}
var communityLdrTypeOfRecord = map[string]string{
	"q": "Community information",
}
var communityLdrKindOfData = map[string]string{
	"n": "Individual",
	"o": "Organization",
	"p": "Program or service",
	"q": "Event",
	"z": "Other",
}
var communityLdrCharacterCodingScheme = map[string]string{
	" ": "MARC-8",
	"a": "UCS/Unicode",
}

// parseCommunityLdr parses leader data for Community records data
func parseCommunityLdr(s string) (ldr LdrDesc) {

	ldr = make(LdrDesc)

	var c string
	var l string
	ldr["(00/05) Record length"] = CodeValue{Code: pluckBytes(s, 0, 5), Label: "", Offset: 0, Width: 5}
	c, l = codeLookup(communityLdrRecordStatus, s, 5, 1)
	ldr["(05/01) Record status"] = CodeValue{Code: c, Label: l, Offset: 5, Width: 1}
	c, l = codeLookup(communityLdrTypeOfRecord, s, 6, 1)
	ldr["(06/01) Type of record"] = CodeValue{Code: c, Label: l, Offset: 6, Width: 1}
	c, l = codeLookup(communityLdrKindOfData, s, 7, 1)
	ldr["(07/01) Kind of data"] = CodeValue{Code: c, Label: l, Offset: 7, Width: 1}
	ldr["(08/01) Undefined character position"] = CodeValue{Code: pluckBytes(s, 8, 1), Label: "", Offset: 8, Width: 1}
	c, l = codeLookup(communityLdrCharacterCodingScheme, s, 9, 1)
	ldr["(09/01) Character coding scheme"] = CodeValue{Code: c, Label: l, Offset: 9, Width: 1}
	ldr["(10/01) Indicator count"] = CodeValue{Code: pluckBytes(s, 10, 1), Label: "", Offset: 10, Width: 1}
	ldr["(11/01) Subfield code length"] = CodeValue{Code: pluckBytes(s, 11, 1), Label: "", Offset: 11, Width: 1}
	ldr["(12/05) Base address of data"] = CodeValue{Code: pluckBytes(s, 12, 5), Label: "", Offset: 12, Width: 5}
	ldr["(17/03) Undefined character positions"] = CodeValue{Code: pluckBytes(s, 17, 3), Label: "", Offset: 17, Width: 3}
	// (20/04) Entry map
	ldr["(20/01) Length of the length-of-field portion"] = CodeValue{Code: pluckBytes(s, 20, 1), Label: "", Offset: 20, Width: 1}
	ldr["(21/01) Length of the starting-character-position portion"] = CodeValue{Code: pluckBytes(s, 21, 1), Label: "", Offset: 21, Width: 1}
	ldr["(22/01) Length of the implementation-defined portion"] = CodeValue{Code: pluckBytes(s, 22, 1), Label: "", Offset: 22, Width: 1}
	ldr["(23/01) Undefined"] = CodeValue{Code: pluckBytes(s, 23, 1), Label: "", Offset: 23, Width: 1}

	return ldr
}

////////////////////////////////////////////////////////////////////////
// Holdings
var holdingsLdrRecordStatus = map[string]string{
	"c": "Corrected or revised",
	"d": "Deleted",
	"n": "New",
}
var holdingsLdrTypeOfRecord = map[string]string{
	"u": "Unknown",
	"v": "Multipart item holdings",
	"x": "Single-part item holdings",
	"y": "Serial item holdings",
}
var holdingsLdrCharacterCodingScheme = map[string]string{
	" ": "MARC-8",
	"a": "UCS/Unicode",
}
var holdingsLdrEncodingLevel = map[string]string{
	"1": "Holdings level 1",
	"2": "Holdings level 2",
	"3": "Holdings level 3",
	"4": "Holdings level 4",
	"5": "Holdings level 4 with piece designation",
	"m": "Mixed level",
	"u": "Unknown",
	"z": "Other level",
}
var holdingsLdrItemInformationInRecord = map[string]string{
	"i": "Item information",
	"n": "No item information",
}

// parseHoldingsLdr parses leader data for Holdings records data
func parseHoldingsLdr(s string) (ldr LdrDesc) {

	ldr = make(LdrDesc)

	var c string
	var l string
	ldr["(00/05) Record length"] = CodeValue{Code: pluckBytes(s, 0, 5), Label: "", Offset: 0, Width: 5}
	c, l = codeLookup(holdingsLdrRecordStatus, s, 5, 1)
	ldr["(05/01) Record status"] = CodeValue{Code: c, Label: l, Offset: 5, Width: 1}
	c, l = codeLookup(holdingsLdrTypeOfRecord, s, 6, 1)
	ldr["(06/01) Type of record"] = CodeValue{Code: c, Label: l, Offset: 6, Width: 1}
	ldr["(07/02) Undefined character positions"] = CodeValue{Code: pluckBytes(s, 7, 2), Label: "", Offset: 7, Width: 2}
	c, l = codeLookup(holdingsLdrCharacterCodingScheme, s, 9, 1)
	ldr["(09/01) Character coding scheme"] = CodeValue{Code: c, Label: l, Offset: 9, Width: 1}
	ldr["(10/01) Indicator count"] = CodeValue{Code: pluckBytes(s, 10, 1), Label: "", Offset: 10, Width: 1}
	ldr["(11/01) Subfield code length"] = CodeValue{Code: pluckBytes(s, 11, 1), Label: "", Offset: 11, Width: 1}
	ldr["(12/05) Base address of data"] = CodeValue{Code: pluckBytes(s, 12, 5), Label: "", Offset: 12, Width: 5}
	c, l = codeLookup(holdingsLdrEncodingLevel, s, 17, 1)
	ldr["(17/01) Encoding level"] = CodeValue{Code: c, Label: l, Offset: 17, Width: 1}
	c, l = codeLookup(holdingsLdrItemInformationInRecord, s, 18, 1)
	ldr["(18/01) Item information in record"] = CodeValue{Code: c, Label: l, Offset: 18, Width: 1}
	ldr["(19/01) Undefined character position"] = CodeValue{Code: pluckBytes(s, 19, 1), Label: "", Offset: 19, Width: 1}
	// (20/04) Entry map
	ldr["(20/01) Length of the length-of-field portion"] = CodeValue{Code: pluckBytes(s, 20, 1), Label: "", Offset: 20, Width: 1}
	ldr["(21/01) Length of the starting-character-position portion"] = CodeValue{Code: pluckBytes(s, 21, 1), Label: "", Offset: 21, Width: 1}
	ldr["(22/01) Length of the implementation-defined portion"] = CodeValue{Code: pluckBytes(s, 22, 1), Label: "", Offset: 22, Width: 1}
	ldr["(23/01) Undefined"] = CodeValue{Code: pluckBytes(s, 23, 1), Label: "", Offset: 23, Width: 1}

	return ldr
}
