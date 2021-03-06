package csl

import "encoding/xml"

// http://docs.citationstyles.org/en/1.0.1/specification.html
// https://github.com/citation-style-language/schema/blob/v1.0.1/csl.rnc

// Style holds CSL info
type Style struct {
	XMLName xml.Name `xml:"style"`
	InheritableNameAttributes
	Xmlns                     string       `xml:"xmlns,attr"`
	Class                     string       `xml:"class,attr"` //"in-text" | "note"
	Version                   string       `xml:"version,attr"`
	DemoteNonDroppingParticle string       `xml:"demote-non-dropping-particle,attr"`
	PageRangeFormat           string       `xml:"page-range-format,attr"`
	DefaultLocale             string       `xml:"default-locale,attr"`
	Info                      Info         `xml:"info"`         //n=1
	Locale                    []Locale     `xml:"locale"`       //n=0+
	Macro                     []Macro      `xml:"macro"`        //n=0+
	Citation                  Citation     `xml:"citation"`     //n=1
	Bibliography              Bibliography `xml:"bibliography"` //n=1?
}

// PersonalDetails holds contact info
type PersonalDetails struct {
	Name  string `xml:"name"`
	Email string `xml:"email"`
	URI   string `xml:"uri"`
}

// Contributor details
type Contributor struct {
	PersonalDetails
}

// Author details
type Author struct {
	PersonalDetails
}

// Translator details
type Translator struct {
	PersonalDetails
}

// Info holds style meta-data for both dependent and independent styles
type Info struct {
	Title       string        `xml:"title"`
	TitleShort  string        `xml:"title-short"`
	ID          string        `xml:"id"`
	Author      []Author      `xml:"author"`
	Contributor []Contributor `xml:"contributor"`
	Category    []struct {
		CitationFormat string `xml:"citation-format,attr"`
		Field          string `xml:"field,attr"`
	} `xml:"category"`
	Summary string `xml:"summary"`
	Updated string `xml:"updated"`
	Rights  struct {
		License string `xml:"license,attr"`
	} `xml:"rights"`
	Link []struct {
		Href string `xml:"href,attr"`
		Rel  string `xml:"rel,attr"`
	} `xml:"link"`
}

// Locale used to (re)define localized terms, dates and options.
type Locale struct {
	XMLName xml.Name `xml:"locale"`
	Text    string   `xml:",chardata"`
	Lang    string   `xml:"lang,attr"`
	Terms   struct {
		Term []Term `xml:"term"`
	} `xml:"terms"`
}

// Term element can either hold a basic string, or "cs:single" and
//  "cs:multiple" child elements to give singular and plural forms of the term.
type Term struct {
	Text     string `xml:",chardata"`
	Name     string `xml:"name,attr"`
	Form     string `xml:"form,attr"`
	Single   string `xml:"single"`
	Multiple string `xml:"multiple"`
}

//Citation describes the formatting of in-text or end/footnote citations
type Citation struct {
	InheritableNameAttributes
	Collapse string `xml:"collapse,attr"`
	Sort     Sort   `xml:"sort"` //n=0+
	Layout   Layout `xml:"layout"`
}

// Bibliography describes the formatting of bibliographic entries
type Bibliography struct {
	InheritableNameAttributes
	// white spacing
	SecondFieldAlign string `xml:"second-field-align,attr"`
	HangingIndent    string `xml:"hanging-indent,attr"`
	EntrySpacing     string `xml:"entry-spacing,attr"`
	LineSpacing      string `xml:"line-spacing,attr"`
	Layout           Layout `xml:"layout"`
}

// InheritableNameAttributes Attributes for the cs:names and cs:name elements may also be set on cs:style, cs:citation and cs:bibliography. The available inheritable attributes for cs:name are and, delimiter-precedes-et-al, delimiter-precedes-last, et-al-min, et-al-use-first, et-al-use-last, et-al-subsequent-min, et-al-subsequent-use-first, initialize, initialize-with, name-as-sort-order and sort-separator
type InheritableNameAttributes struct {
	And                   string `xml:"and,attr"`
	EtAlMin               string `xml:"et-al-min,attr"`
	DelimiterPrecedesLast string `xml:"delimiter-precedes-last,attr"`
	EtAlUseFirst          string `xml:"et-al-use-first,attr"`
	EtAlUseLast           string `xml:"et-al-use-last,attr"`
	SortSeparator         string `xml:"sort-separator,attr"`
	InitializeWith        string `xml:"initialize-with,attr"`
	NameAsSortOrder       string `xml:"name-as-sort-order,attr"`
}

// Affixes provides prefix and suffix attributes
type Affixes struct {
	Prefix string `xml:"prefix,attr"` //eg prefix="(" suffix=")" add () around citation
	Suffix string `xml:"suffix,attr"`
}

// Layout required child of Citation and Bibiliography; contains >0 rendering elements
type Layout struct {
	Affixes
	Delimiter string `xml:"delimiter,attr"` // separates neighboring cites
	// children elements determines the format of each cite
	RenderingElement
}

// Sort specifies how cites and bibliographic entries should be sorted
type Sort struct {
	Key []struct {
		Variable string `xml:"variable,attr"`
		Marco    string `xml:"macro,attr"`
	} `xml:"key"`
}

// RenderingElement holds info on one or more rendering elements
type RenderingElement struct {
	Names  []Names  `xml:"names"`
	Date   []Date   `xml:"date"`
	Label  []Label  `xml:"label"`
	Text   []Text   `xml:"text"`
	Number []Number `xml:"number"`
	Choose []Choose `xml:"choose"`
	Group  []Group  `xml:"group"`
}

// Number outputs the number variable selected with the required variable attribute
type Number struct {
	Affixes
	Variable string `xml:"variable,attr"`
	Form     string `xml:"form,attr"` //"numeric" | "ordinal" | "long-ordinal" | "roman"
}

// Date renders date
type Date struct {
	Affixes
	Form      string     `xml:"form,attr"`       //localized date format "text" or "numeric"
	DateParts string     `xml:"date-parts,attr"` //defaults to "year-month-day” | "year-month” | "year"
	Delimiter string     `xml:"delimiter,attr"`  // delimits non-empty pieces of output
	TextCase  string     `xml:"text-case,attr"`
	DatePart  []DatePart `xml:"date-part"` //used to override locale defaults
}

// DatePart renders a data part
type DatePart struct {
	Affixes
	Name         string `xml:"name,attr"` // day | month | year
	Form         string `xml:"form,attr"`
	StripPeriods string `xml:"strip-periods,attr"` //if “true” (default=“false”), any periods in the rendered text are removed.
	TextCase     string `xml:"text-case,attr"`
}

// Text whose attributes determines what is rendered
type Text struct {
	Affixes
	Quotes       string `xml:"quotes,attr"`        // “true” (“false” is default), the rendered text is wrapped in quotes
	StripPeriods string `xml:"strip-periods,attr"` //if “true” (default=“false”), any periods in the rendered text are removed.
	TextCase     string `xml:"text-case,attr"`
	CData        string `xml:",chardata"`
	Variable     string `xml:"variable,attr"` //renders the text of one of the standard variables
	Macro        string `xml:"macro,attr"`    // renders the text output of a macro
	Term         string `xml:"term,attr"`     //renders a term
	Value        string `xml:"value,attr"`    //renders the value of the attribute
}

// Names outputs the contents of one or more name variables specified in the variable attr
type Names struct {
	Affixes
	InheritableNameAttributes
	Delimiter  string     `xml:"delimiter,attr"` // delimits name lists from different name variables
	Variable   string     `xml:"variable,attr"`
	Name       Name       `xml:"name"`       //n=1?
	EtAl       EtAl       `xml:"et-al"`      //n=1?
	Label      []Label    `xml:"label"`      //n=0+, must be after name and et al
	Substitute Substitute `xml:"substitute"` //n=1?, must be the last element
}

// Name an optional child of Names used to describe the formatting of individual names
type Name struct {
	Affixes
	InheritableNameAttributes
	Delimiter string `xml:"delimiter,attr"` // delimits non-empty pieces of output
}

// NamePart ?
type NamePart struct {
	Affixes
	TextCase string `xml:"text-case,attr"`
}

// EtAl specifies the term used for et-al abbreviation and its formatting.
type EtAl struct {
	Term string `xml:"term,attr"` // "et-al" | "and others"
}

// Label ?
type Label struct {
	Affixes
	StripPeriods string `xml:"strip-periods,attr"` //if “true” (default=“false”), any periods in the rendered text are removed.
	Form         string `xml:"form,attr"`
	TextCase     string `xml:"text-case,attr"`
}

// Substitute adds substitution in case the name variables specified in the parent cs:names element are empty
type Substitute struct { //n=1?, must be the last element
	RenderingElement
}

// Group endering element must contain one or more rendering elements (with the exception of Layout)
type Group struct {
	Affixes
	Delimiter string `xml:"delimiter,attr"` // delimits non-empty pieces of output
	RenderingElement
}

// Macro holds formatting instructions (1 or more rendering elements)
type Macro struct {
	Name string `xml:"name,attr"`
	RenderingElement
}

// Choose used to conditionally render rendering elements.
type Choose struct {
	If     If   `xml:"if"`      //n=1
	Elseif []If `xml:"else-if"` //n=0+
	Else   Else `xml:"else"`    //n=0+
}

//
type Else struct {
	RenderingElement
}

type If struct {
	Condition []Condition `xml:"condition"`
	// Match sets the testing logic.
	Match string `xml:"match,attr"` // "all", "any", "none"
	RenderingElement
}

type Condition struct {
	Disambiguate string `xml:"disambiguate,attr"`
	IsNumeric    string `xml:"is-numeric,attr"`
}
