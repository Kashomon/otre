package latex

var gooePackages = `\usepackage{gooemacs}`

var gooeFontDefsBase = `
\font\tenpoint=%s10
\font\tenpointeleven=%s10 at 11pt
\font\eightpoint=%s8
\font\eightpointnine=%s8 at 9pt`

var gooeSizeDefs = `% Basic defs
\newdimen\bigRaise
\bigRaise=4.3pt
\newdimen\smallRaise
\smallRaise=3.5pt
\newdimen\inlineRaise
\inlineRaise=3.5pt`

var gooeBigBoardDefs = `% Big Board Defs
% Empty intersection label (BigBoard). Usage: \eLblBig{A}
\def\eLblBig#1{\leavevmode\hbox to \goIntWd{
  \hss\raise\bigRaise\hbox{\rm \tenpointeleven{#1}}\hss}}
% White Stone + Black Label (BigBoard). Usage: \goWsLblBig{A}
\def\goWsLblBig#1{\setbox0=\hbox{\0??!}
  \rlap{\0??!}\raise\bigRaise\hbox to \wd0{\hss\tenpointeleven{#1}\hss}}
% Black Stone + White Label (BigBoard). Usage: \goBsLblBig{A}
\def\goBsLblBig#1{\setbox0=\hbox{\0??@}
  \rlap{\0??@}\raise\bigRaise\hbox to \wd0{\hss\color{white}
  \tenpointeleven{#1}\color{white}\hss}}
`

var gooeNormalBoardDefs = `% Normal Board Defs
% Empty Label. Usage: \eLbl{A}
\def\eLbl#1{\leavevmode\hbox to \goIntWd{
  \hss\raise\smallRaise\hbox{\rm \tenpoint{#1}}\hss}}
% White Stone + Black Label (SmallBoard). Usage: \goWsLbl{A}
\def\goWsLbl#1{\setbox0=\hbox{\0??!}
  \rlap{\0??!}\raise\smallRaise\hbox to \wd0{\hss\eightpointnine{#1}\hss}}
% Black Stone + White Label (SmallBoard). Usage: \goBsLbl{A}
\def\goBsLbl#1{\setbox0=\hbox{\0??@}
  \rlap{\0??@}\raise\smallRaise\hbox to \wd0{\hss\color{white}
  \eightpointnine{#1}\color{white}\hss}}
`

const gooeInlineDefs = `% Inline Defs
% Inline White stone + Black label. Usage: \goinWsLbl{A}
\def\goinWsLbl#1{\textstone{\goo {\setbox0=\hbox{\0??!}
  \rlap{\0??!}\raise\inlineRaise\hbox to \wd0{\hss\eightpoint{#1}\hss}}}}
% Inline Black stone + White label. Usage: \goinBsLbl{A}
\def\goinBsLbl#1{\textstone{\goo {\setbox0=\hbox{\0??@}
  \rlap{\0??@}\raise\inlineRaise\hbox to \wd0{\hss\color{white}
  \eightpoint{#1}\color{white}\hss}}}}
% Inline stonechar. Usage: \goinChar{\0??S} -- Square Stone.
\def\goinChar#1{\textstone{\goo {#1}}}
`

var gooeBigBoard = `{\bgoo\n%s\n}`
var gooeBoardDef = `{\goo\n%s\n}`

var gooemacsMap = map[GoChar]string{
	// Lines
	TL_CORNER:  "\\0??<",
	TR_CORNER:  "\\0??>",
	BL_CORNER:  "\\0??,",
	BR_CORNER:  "\\0??.",
	TOP_EDGE:   "\\0??(",
	BOT_EDGE:   "\\0??)",
	LEFT_EDGE:  "\\0??[",
	RIGHT_EDGE: "\\0??]",
	CENTER:     "\\0??+",

	// Starpoint
	STARPOINT: "\\0??*",

	// Stones
	BSTONE: "\\0??@",
	WSTONE: "\\0??!",

	// Marks and StoneMarks
	BSTONE_TRIANGLE: "\\0??S",
	WSTONE_TRIANGLE: "\\0??s",
	TRIANGLE:        "\\0??3",
	BSTONE_SQUARE:   "\\0??S",
	WSTONE_SQUARE:   "\\0??s",
	SQUARE:          "\\0??2",
	BSTONE_CIRCLE:   "\\0??C",
	WSTONE_CIRCLE:   "\\0??c",
	CIRCLE:          "\\0??1",
	BSTONE_XMARK:    "\\0??X",
	WSTONE_XMARK:    "\\0??x",
	XMARK:           "\\0??4",

	// BigBoard Numbering
	// Example: \003+.  Note, the board decides which 'color' is selected
	BSTONE_LABEL_BIG: "\\goBsLblBig{%s}",
	WSTONE_LABEL_BIG: "\\goWsLblBig{%s}",
	EMPTY_LABEL_BIG:  "\\eLblBig{%s}",

	// SmalBoard Numbering
	BSTONE_LABEL: "\\goBsLbl{%s}",
	WSTONE_LABEL: "\\goWsLbl{%s}",
	EMPTY_LABEL:  "\\eLbl{%s}",

	// Formatting
	BSTONE_INLINE:     "\\goinBsLbl{%s}",
	WSTONE_INLINE:     "\\goinWsLbl{%s}",
	MISC_STONE_INLINE: "\\goinChar{%s}",
}
