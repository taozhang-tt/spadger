package xconvey

import "github.com/smartystreets/goconvey/convey"

var (
	// equality assertions
	ShouldEqual          = convey.ShouldEqual
	ShouldNotEqual       = convey.ShouldNotEqual
	ShouldAlmostEqual    = convey.ShouldAlmostEqual
	ShouldNotAlmostEqual = convey.ShouldNotAlmostEqual
	ShouldResemble       = convey.ShouldResemble
	ShouldNotResemble    = convey.ShouldNotResemble
	ShouldPointTo        = convey.ShouldPointTo
	ShouldNotPointTo     = convey.ShouldNotPointTo
	ShouldBeNil          = convey.ShouldBeNil
	ShouldNotBeNil       = convey.ShouldNotBeNil
	ShouldBeTrue         = convey.ShouldBeTrue
	ShouldBeFalse        = convey.ShouldBeFalse
	ShouldBeZeroValue    = convey.ShouldBeZeroValue
	ShouldNotBeZeroValue = convey.ShouldNotBeZeroValue

	// numeric comparison assertions
	ShouldBeGreaterThan          = convey.ShouldBeGreaterThan
	ShouldBeGreaterThanOrEqualTo = convey.ShouldBeGreaterThanOrEqualTo
	ShouldBeLessThan             = convey.ShouldBeLessThan
	ShouldBeLessThanOrEqualTo    = convey.ShouldBeLessThanOrEqualTo
	ShouldBeBetween              = convey.ShouldBeBetween
	ShouldNotBeBetween           = convey.ShouldNotBeBetween
	ShouldBeBetweenOrEqual       = convey.ShouldBeBetweenOrEqual
	ShouldNotBeBetweenOrEqual    = convey.ShouldNotBeBetweenOrEqual

	// container assertions
	ShouldContain       = convey.ShouldContain
	ShouldNotContain    = convey.ShouldNotContain
	ShouldContainKey    = convey.ShouldContainKey
	ShouldNotContainKey = convey.ShouldNotContainKey
	ShouldBeIn          = convey.ShouldBeIn
	ShouldNotBeIn       = convey.ShouldNotBeIn
	ShouldBeEmpty       = convey.ShouldBeEmpty
	ShouldNotBeEmpty    = convey.ShouldNotBeEmpty
	ShouldHaveLength    = convey.ShouldHaveLength

	// string assertions
	ShouldStartWith           = convey.ShouldStartWith
	ShouldNotStartWith        = convey.ShouldNotStartWith
	ShouldEndWith             = convey.ShouldEndWith
	ShouldNotEndWith          = convey.ShouldNotEndWith
	ShouldBeBlank             = convey.ShouldBeBlank
	ShouldNotBeBlank          = convey.ShouldNotBeBlank
	ShouldContainSubstring    = convey.ShouldContainSubstring
	ShouldNotContainSubstring = convey.ShouldNotContainSubstring

	// panic recovery assertions
	ShouldPanic        = convey.ShouldPanic
	ShouldNotPanic     = convey.ShouldNotPanic
	ShouldPanicWith    = convey.ShouldPanicWith
	ShouldNotPanicWith = convey.ShouldNotPanicWith

	// type-checking assertions
	ShouldHaveSameTypeAs    = convey.ShouldHaveSameTypeAs
	ShouldNotHaveSameTypeAs = convey.ShouldNotHaveSameTypeAs
	ShouldImplement         = convey.ShouldImplement
	ShouldNotImplement      = convey.ShouldNotImplement

	// time assertions
	ShouldHappenBefore         = convey.ShouldHappenBefore
	ShouldHappenOnOrBefore     = convey.ShouldHappenOnOrBefore
	ShouldHappenAfter          = convey.ShouldHappenAfter
	ShouldHappenOnOrAfter      = convey.ShouldHappenOnOrAfter
	ShouldHappenBetween        = convey.ShouldHappenBetween
	ShouldHappenOnOrBetween    = convey.ShouldHappenOnOrBetween
	ShouldNotHappenOnOrBetween = convey.ShouldNotHappenOnOrBetween
	ShouldHappenWithin         = convey.ShouldHappenWithin
	ShouldNotHappenWithin      = convey.ShouldNotHappenWithin
	ShouldBeChronological      = convey.ShouldBeChronological

	// error assertions
	ShouldBeError = convey.ShouldBeError
	ShouldWrap    = convey.ShouldWrap
)

const (
	FailureContinues = convey.FailureContinues
	FailureHalts     = convey.FailureHalts
	FailureInherits  = convey.FailureInherits
	StackError       = convey.StackError
	StackFail        = convey.StackFail
	StackInherits    = convey.StackInherits
)

var (
	Convey                = convey.Convey
	SkipConvey            = convey.SkipConvey
	FocusConvey           = convey.FocusConvey
	Reset                 = convey.Reset
	So                    = convey.So
	SkipSo                = convey.SkipSo
	SetDefaultFailureMode = convey.SetDefaultFailureMode
	XPrint                = convey.Print
	XPrintf               = convey.Printf
	XPrintln              = convey.Println
)

type (
	C           = convey.C
	FailureMode = convey.FailureMode
	StackMode   = convey.StackMode
)
