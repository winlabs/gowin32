/*
 * Copyright (c) 2014 MongoDB, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the license is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package wrappers

import (
	"syscall"
	"unsafe"
)

const (
	VER_SUITE_SMALLBUSINESS            = 0x00000001
	VER_SUITE_ENTERPRISE               = 0x00000002
	VER_SUITE_BACKOFFICE               = 0x00000004
	VER_SUITE_COMMUNICATIONS           = 0x00000008
	VER_SUITE_TERMINAL                 = 0x00000010
	VER_SUITE_SMALLBUSINESS_RESTRICTED = 0x00000020
	VER_SUITE_EMBEDDEDNT               = 0x00000040
	VER_SUITE_DATACENTER               = 0x00000080
	VER_SUITE_SINGLEUSERTS             = 0x00000100
	VER_SUITE_PERSONAL                 = 0x00000200
	VER_SUITE_BLADE                    = 0x00000400
	VER_SUITE_EMBEDDED_RESTRICTED      = 0x00000800
	VER_SUITE_SECURITY_APPLIANCE       = 0x00001000
	VER_SUITE_STORAGE_SERVER           = 0x00002000
	VER_SUITE_COMPUTE_SERVER           = 0x00004000
	VER_SUITE_WH_SERVER                = 0x00008000
)

const (
	LANG_NEUTRAL   = 0x00
	LANG_INVARIANT = 0x7F

	LANG_AFRIKAANS     = 0x36
	LANG_ALBANIAN      = 0x1C
	LANG_ALSATIAN      = 0x84
	LANG_AMHARIC       = 0x5E
	LANG_ARABIC        = 0x01
	LANG_ARMENIAN      = 0x2B
	LANG_ASSAMESE      = 0x4D
	LANG_AZERI         = 0x2C
	LANG_BASHKIR       = 0x6D
	LANG_BASQUE        = 0x2D
	LANG_BELARUSIAN    = 0x23
	LANG_BENGALI       = 0x45
	LANG_BRETON        = 0x7E
	LANG_BOSNIAN       = 0x1A
	LANG_BULGARIAN     = 0x02
	LANG_CATALAN       = 0x03
	LANG_CHINESE       = 0x04
	LANG_CORSICAN      = 0x83
	LANG_CROATIAN      = 0x1A
	LANG_CZECH         = 0x05
	LANG_DANISH        = 0x06
	LANG_DARI          = 0x8C
	LANG_DIVEHI        = 0x65
	LANG_DUTCH         = 0x13
	LANG_ENGLISH       = 0x09
	LANG_ESTONIAN      = 0x25
	LANG_FAEROESE      = 0x38
	LANG_FARSI         = 0x29
	LANG_FILIPINO      = 0x64
	LANG_FINNISH       = 0x0B
	LANG_FRENCH        = 0x0C
	LANG_FRISIAN       = 0x62
	LANG_GALICIAN      = 0x56
	LANG_GEORGIAN      = 0x37
	LANG_GERMAN        = 0x07
	LANG_GREEK         = 0x08
	LANG_GREENLANDIC   = 0x6F
	LANG_GUJARATI      = 0x47
	LANG_HAUSA         = 0x68
	LANG_HEBREW        = 0x0D
	LANG_HINDI         = 0x39
	LANG_HUNGARIAN     = 0x0E
	LANG_ICELANDIC     = 0x0F
	LANG_IGBO          = 0x70
	LANG_INDONESIAN    = 0x21
	LANG_INUKTITUT     = 0x5D
	LANG_IRISH         = 0x3C
	LANG_ITALIAN       = 0x10
	LANG_JAPANESE      = 0x11
	LANG_KANNADA       = 0x4B
	LANG_KASHMIRI      = 0x60
	LANG_KAZAK         = 0x3F
	LANG_KHMER         = 0x53
	LANG_KICHE         = 0x86
	LANG_KINYARWANDA   = 0x87
	LANG_KONKANI       = 0x57
	LANG_KOREAN        = 0x12
	LANG_KYRGYZ        = 0x40
	LANG_LAO           = 0x54
	LANG_LATVIAN       = 0x26
	LANG_LITHUANIAN    = 0x27
	LANG_LOWER_SORBIAN = 0x2E
	LANG_LUXEMBOURGISH = 0x6E
	LANG_MACEDONIAN    = 0x2F
	LANG_MALAY         = 0x3E
	LANG_MALAYALAM     = 0x4C
	LANG_MALTESE       = 0x3A
	LANG_MANIPURI      = 0x58
	LANG_MAORI         = 0x81
	LANG_MAPUDUNGUN    = 0x7A
	LANG_MARATHI       = 0x4E
	LANG_MOHAWK        = 0x7C
	LANG_MONGOLIAN     = 0x50
	LANG_NEPALI        = 0x61
	LANG_NORWEGIAN     = 0x14
	LANG_OCCITAN       = 0x82
	LANG_ORIYA         = 0x48
	LANG_PASHTO        = 0x63
	LANG_PERSIAN       = 0x29
	LANG_POLISH        = 0x15
	LANG_PORTUGUESE    = 0x16
	LANG_PUNJABI       = 0x46
	LANG_QUECHUA       = 0x6B
	LANG_ROMANIAN      = 0x18
	LANG_ROMANSH       = 0x17
	LANG_RUSSIAN       = 0x19
	LANG_SAMI          = 0x3B
	LANG_SANSKRIT      = 0x4F
	LANG_SERBIAN       = 0x1F
	LANG_SINDHI        = 0x59
	LANG_SINHALESE     = 0x5B
	LANG_SLOVAK        = 0x1B
	LANG_SLOVENIAN     = 0x24
	LANG_SOTHO         = 0x6C
	LANG_SPANISH       = 0x0A
	LANG_SWAHILI       = 0x41
	LANG_SWEDISH       = 0x1D
	LANG_SYRIAC        = 0x5A
	LANG_TAJIK         = 0x28
	LANG_TAMAZIGHT     = 0x5F
	LANG_TAMIL         = 0x49
	LANG_TATAR         = 0x44
	LANG_TELUGU        = 0x4A
	LANG_THAI          = 0x1E
	LANG_TIBETAN       = 0x51
	LANG_TIGRIGNA      = 0x73
	LANG_TSWANA        = 0x32
	LANG_TURKISH       = 0x1F
	LANG_TURKMEN       = 0x42
	LANG_UIGHUR        = 0x80
	LANG_UKRAINIAN     = 0x22
	LANG_UPPER_SORBIAN = 0x2E
	LANG_URDU          = 0x20
	LANG_UZBEK         = 0x43
	LANG_VIETNAMESE    = 0x2A
	LANG_WELSH         = 0x52
	LANG_WOLOF         = 0x88
	LANG_XHOSA         = 0x34
	LANG_YAKUT         = 0x85
	LANG_YI            = 0x78
	LANG_YORUBA        = 0x6A
	LANG_ZULU          = 0x35
)

const (
	SUBLANG_NEUTRAL            = 0x00
	SUBLANG_DEFAULT            = 0x01
	SUBLANG_SYS_DEFAULT        = 0x02
	SUBLANG_CUSTOM_DEFAULT     = 0x03
	SUBLANG_CUSTOM_UNSPECIFIED = 0x04
	SUBLANG_UI_CUSTOM_DEFAULT  = 0x05

	SUBLANG_AFRIKAANS_SOUTH_AFRICA              = 0x01
	SUBLANG_ALBANIAN_ALBANIA                    = 0x01
	SUBLANG_ALSATIAN_FRANCE                     = 0x01
	SUBLANG_AMHARIC_ETHIOPIA                    = 0x01
	SUBLANG_ARABIC_SAUDI_ARABIA                 = 0x01
	SUBLANG_ARABIC_IRAQ                         = 0x02
	SUBLANG_ARABIC_EGYPT                        = 0x03
	SUBLANG_ARABIC_LIBYA                        = 0x04
	SUBLANG_ARABIC_ALGERIA                      = 0x05
	SUBLANG_ARABIC_MOROCCO                      = 0x06
	SUBLANG_ARABIC_TUNISIA                      = 0x07
	SUBLANG_ARABIC_OMAN                         = 0x08
	SUBLANG_ARABIC_YEMEN                        = 0x09
	SUBLANG_ARABIC_SYRIA                        = 0x0A
	SUBLANG_ARABIC_JORDAN                       = 0x0B
	SUBLANG_ARABIC_LEBANON                      = 0x0C
	SUBLANG_ARABIC_KUWAIT                       = 0x0D
	SUBLANG_ARABIC_UAE                          = 0x0E
	SUBLANG_ARABIC_BAHRAIN                      = 0x0F
	SUBLANG_ARABIC_QATAR                        = 0x10
	SUBLANG_ARMENIAN_ARMENIA                    = 0x01
	SUBLANG_ASSAMESE_INDIA                      = 0x01
	SUBLANG_AZERI_LATIN                         = 0x01
	SUBLANG_AZERI_CYRILLIC                      = 0x01
	SUBLANG_BASHKIR_RUSSIA                      = 0x01
	SUBLANG_BASQUE_BASQUE                       = 0x01
	SUBLANG_BELARUSIAN_BELARUS                  = 0x01
	SUBLANG_BENGALI_INDIA                       = 0x01
	SUBLANG_BENGALI_BANGLADESH                  = 0x02
	SUBLANG_BOSNIAN_BOSNIA_HERZEGOVINA_LATIN    = 0x05
	SUBLANG_BOSNIAN_BOSNIA_HERZEGOVINA_CYRILLIC = 0x08
	SUBLANG_BRETON_FRANCE                       = 0x01
	SUBLANG_BULGARIAN_BULGARIA                  = 0x01
	SUBLANG_CATALAN_CATALAN                     = 0x01
	SUBLANG_CHINESE_TRADITIONAL                 = 0x01
	SUBLANG_CHINESE_SIMPLIFIED                  = 0x02
	SUBLANG_CHINESE_HONGKONG                    = 0x03
	SUBLANG_CHINESE_SINGAPORE                   = 0x04
	SUBLANG_CHINESE_MACAU                       = 0x05
	SUBLANG_CORSICAN_FRANCE                     = 0x01
	SUBLANG_CZECH_CZECH_REPUBLIC                = 0x01
	SUBLANG_CROATIAN_CROATIA                    = 0x01
	SUBLANG_CROATIAN_BOSNIA_HERVEGOVINA_LATIN   = 0x04
	SUBLANG_DANISH_DENMARK                      = 0x01
	SUBLANG_DARI_AFGHANISTAN                    = 0x01
	SUBLANG_DIVEHI_MALDIVES                     = 0x01
	SUBLANG_DUTCH                               = 0x01
	SUBLANG_DUTCH_BELGIAN                       = 0x02
	SUBLANG_ENGLISH_US                          = 0x01
	SUBLANG_ENGLISH_UK                          = 0x02
	SUBLANG_ENGLISH_AUS                         = 0x03
	SUBLANG_ENGLISH_CAN                         = 0x04
	SUBLANG_ENGLISH_NZ                          = 0x05
	SUBLANG_ENGLISH_EIRE                        = 0x06
	SUBLANG_ENGLISH_SOUTH_AFRICA                = 0x07
	SUBLANG_ENGLISH_JAMAICA                     = 0x08
	SUBLANG_ENGLISH_CARIBBEAN                   = 0x09
	SUBLANG_ENGLISH_BELIZE                      = 0x0A
	SUBLANG_ENGLISH_TRINIDAD                    = 0x0B
	SUBLANG_ENGLISH_ZIMBABWE                    = 0x0C
	SUBLANG_ENGLISH_PHILIPPINES                 = 0x0D
	SUBLANG_ENGLISH_INDIA                       = 0x10
	SUBLANG_ENGLISH_MALAYSIA                    = 0x11
	SUBLANG_ENGLISH_SINGAPORE                   = 0x12
	SUBLANG_ESTONIAN_ESTONIA                    = 0x01
	SUBLANG_FAEROESE_FAERO_ISLANDS              = 0x01
	SUBLANG_FILIPINO_PHILIPPINES                = 0x01
	SUBLANG_FINNISH_FINLAND                     = 0x01
	SUBLANG_FRENCH                              = 0x01
	SUBLANG_FRENCH_BELGIAN                      = 0x02
	SUBLANG_FRENCH_CANADIAN                     = 0x03
	SUBLANG_FRENCH_SWISS                        = 0x04
	SUBLANG_FRENCH_LUXEMBOURG                   = 0x05
	SUBLANG_FRENCH_MONACO                       = 0x06
	SUBLANG_FRISIAN_NETHERLANDS                 = 0x01
	SUBLANG_GALICIAN_GALICIAN                   = 0x01
	SUBLANG_GEORGIAN_GEORGIA                    = 0x01
	SUBLANG_GERMAN                              = 0x01
	SUBLANG_GERMAN_SWISS                        = 0x02
	SUBLANG_GERMAN_AUSTRIAN                     = 0x03
	SUBLANG_GERMAN_LUXEMBOURG                   = 0x04
	SUBLANG_GERMAN_LIECHTENSTEIN                = 0x05
	SUBLANG_GREEK_GREECE                        = 0x01
	SUBLANG_GREENLANDIC_GREENLAND               = 0x02
	SUBLANG_GUJARATI_INDIA                      = 0x01
	SUBLANG_HAUSA_NIGERIA_LATIN                 = 0x01
	SUBLANG_HEBREW_ISRAEL                       = 0x01
	SUBLANG_HINDI_INDIA                         = 0x01
	SUBLANG_HUNGARIAN_HUNGARY                   = 0x01
	SUBLANG_ICELANDIC_ICELAND                   = 0x01
	SUBLANG_IGBO_NIGERIA                        = 0x01
	SUBLANG_INDONESIAN_INDONESIA                = 0x01
	SUBLANG_INUKTITUT_CANADA                    = 0x01
	SUBLANG_INUKTITUT_CANADA_LATIN              = 0x02
	SUBLANG_IRISH_IRELAND                       = 0x02
	SUBLANG_ITALIAN                             = 0x01
	SUBLANG_ITALIAN_SWISS                       = 0x02
	SUBLANG_JAPANESE_JAPAN                      = 0x01
	SUBLANG_KANNADA_INDIA                       = 0x01
	SUBLANG_KASHMIRI_SASIA                      = 0x02
	SUBLANG_KASHMIRI_INDIA                      = 0x02
	SUBLANG_KAZAK_KAZAKHSTAN                    = 0x01
	SUBLANG_KHMER_CAMBODIA                      = 0x01
	SUBLANG_KICHE_GUATEMALA                     = 0x01
	SUBLANG_KINYARWANDA_RWANDA                  = 0x01
	SUBLANG_KONKANI_INDIA                       = 0x01
	SUBLANG_KOREAN                              = 0x01
	SUBLANG_KYRGYZ_KYRGYZSTAN                   = 0x01
	SUBLANG_LAO_LAO                             = 0x01
	SUBLANG_LATVIAN_LATVIA                      = 0x01
	SUBLANG_LITHUANIAN                          = 0x01
	SUBLANG_LOWER_SORBIAN_GERMANY               = 0x02
	SUBLANG_LUXEMBOURGISH_LUXEMBOURG            = 0x01
	SUBLANG_MACEDONIAN_MACEDONIA                = 0x01
	SUBLANG_MALAY_MALAYSIA                      = 0x01
	SUBLANG_MALAY_BRUNEI_DARUSSALAM             = 0x02
	SUBLANG_MALAYALAM_INDIA                     = 0x01
	SUBLANG_MALTESE_MALTA                       = 0x01
	SUBLANG_MAORI_NEW_ZEALAND                   = 0x01
	SUBLANG_MAPUDUNGUN_CHILE                    = 0x01
	SUBLANG_MARATHI_INDIA                       = 0x01
	SUBLANG_MOHAWK_MOHAWK                       = 0x01
	SUBLANG_MONGOLIAN_CYRILLIC_MONGOLIA         = 0x01
	SUBLANG_MONGOLIAN_PRC                       = 0x02
	SUBLANG_NEPALI_INDIA                        = 0x02
	SUBLANG_NEPALI_NEPAL                        = 0x01
	SUBLANG_NORWEGIAN_BOKMAL                    = 0x01
	SUBLANG_NORWEGIAN_NYNORSK                   = 0x02
	SUBLANG_OCCITAN_FRANCE                      = 0x01
	SUBLANG_ORIYA_INDIA                         = 0x01
	SUBLANG_PASHTO_AFGHANISTAN                  = 0x01
	SUBLANG_PERSIAN_IRAN                        = 0x01
	SUBLANG_POLISH_POLAND                       = 0x01
	SUBLANG_PORTUGUESE                          = 0x02
	SUBLANG_PORTUGUESE_BRAZILIAN                = 0x01
	SUBLANG_PUNJABI_INDIA                       = 0x01
	SUBLANG_QUECHUA_BOLIVIA                     = 0x01
	SUBLANG_QUECHUA_ECUADOR                     = 0x02
	SUBLANG_QUECHUA_PERU                        = 0x03
	SUBLANG_ROMANIAN_ROMANIA                    = 0x01
	SUBLANG_ROMANSH_SWITZERLAND                 = 0x01
	SUBLANG_RUSSIAN_RUSSIA                      = 0x01
	SUBLANG_SAMI_NORTHERN_NORWAY                = 0x01
	SUBLANG_SAMI_NORTHERN_SWEDEN                = 0x02
	SUBLANG_SAMI_NORTHERN_FINLAND               = 0x03
	SUBLANG_SAMI_LULE_NORWAY                    = 0x04
	SUBLANG_SAMI_LULE_SWEDEN                    = 0x05
	SUBLANG_SAMI_SOUTHERN_NORWAY                = 0x06
	SUBLANG_SAMI_SOUTHERN_SWEDEN                = 0x07
	SUBLANG_SAMI_SKOLT_FINLAND                  = 0x08
	SUBLANG_SAMI_INARI_FINLAND                  = 0x09
	SUBLANG_SANSKRIT_INDIA                      = 0x01
	SUBLANG_SERBIAN_BOSNIA_HERZEGOVINA_LATIN    = 0x06
	SUBLANG_SERBIAN_BOSNIA_HERZEGOVINA_CYRILLIC = 0x07
	SUBLANG_SERBIAN_CROATIA                     = 0x01
	SUBLANG_SERBIAN_LATIN                       = 0x02
	SUBLANG_SERBIAN_CYRILLIC                    = 0x03
	SUBLANG_SINDHI_INDIA                        = 0x01
	SUBLANG_SINDHI_PAKISTAN                     = 0x02
	SUBLANG_SINDHI_AFGHANISTAN                  = 0x02
	SUBLANG_SINHALESE_SRI_LANKA                 = 0x01
	SUBLANG_SOTHO_NORTHERN_SOUTH_AFRICA         = 0x01
	SUBLANG_SLOVAK_SLOVAKIA                     = 0x01
	SUBLANG_SLOVENIAN_SLOVENIA                  = 0x01
	SUBLANG_SPANISH                             = 0x01
	SUBLANG_SPANISH_MEXICAN                     = 0x02
	SUBLANG_SPANISH_MODERN                      = 0x03
	SUBLANG_SPANISH_GUATEMALA                   = 0x04
	SUBLANG_SPANISH_COSTA_RICA                  = 0x05
	SUBLANG_SPANISH_PANAMA                      = 0x06
	SUBLANG_SPANISH_DOMINICAN_REPUBLIC          = 0x07
	SUBLANG_SPANISH_VENEZUELA                   = 0x08
	SUBLANG_SPANISH_COLOMBIA                    = 0x09
	SUBLANG_SPANISH_PERU                        = 0x0A
	SUBLANG_SPANISH_ARGENTINA                   = 0x0B
	SUBLANG_SPANISH_ECUADOR                     = 0x0C
	SUBLANG_SPANISH_CHILE                       = 0x0D
	SUBLANG_SPANISH_URUGUAY                     = 0x0E
	SUBLANG_SPANISH_PARAGUAY                    = 0x0F
	SUBLANG_SPANISH_BOLIVIA                     = 0x10
	SUBLANG_SPANISH_EL_SALVADOR                 = 0x11
	SUBLANG_SPANISH_HONDURAS                    = 0x12
	SUBLANG_SPANISH_NICARAGUA                   = 0x13
	SUBLANG_SPANISH_PEURTO_RICO                 = 0x14
	SUBLANG_SPANISH_US                          = 0x15
	SUBLANG_SWEDISH                             = 0x01
	SUBLANG_SWEDISH_FINLAND                     = 0x02
	SUBLANG_SYRIAC_SYRIA                        = 0x01
	SUBLANG_TAJIK_TAJIKISTAN                    = 0x01
	SUBLANG_TAMAZIGHT_ALGERIA_LATIN             = 0x02
	SUBLANG_TAMIL_INDIA                         = 0x01
	SUBLANG_TATAR_RUSSIA                        = 0x01
	SUBLANG_TELUGU_INDIA                        = 0x01
	SUBLANG_THAI_THAILAND                       = 0x01
	SUBLANG_TIBETAN_PRC                         = 0x01
	SUBLANG_TIGRIGNA_ERITREA                    = 0x02
	SUBLANG_TSWANA_SOUTH_AFRICA                 = 0x01
	SUBLANG_TURKISH_TURKEY                      = 0x01
	SUBLANG_TURKMEN_TURKMENISTAN                = 0x01
	SUBLANG_UIGHUR_PRC                          = 0x01
	SUBLANG_UKRAINIAN_UKRAINE                   = 0x01
	SUBLANG_UPPER_SORBIAN_GERMANY               = 0x01
	SUBLANG_URDU_PAKISTAN                       = 0x01
	SUBLANG_URDU_INDIA                          = 0x02
	SUBLANG_UZBEK_LATIN                         = 0x01
	SUBLANG_UZBEK_CYRILLIC                      = 0x02
	SUBLANG_VIETNAMESE_VIETNAM                  = 0x01
	SUBLANG_WELSH_UNITED_KINGDOM                = 0x01
	SUBLANG_WOLOF_SENEGAL                       = 0x01
	SUBLANG_XHOSA_SOUTH_AFRICA                  = 0x01
	SUBLANG_YAKUT_RUSSIA                        = 0x01
	SUBLANG_YI_PRC                              = 0x01
	SUBLANG_YORUBA_NIGERIA                      = 0x01
	SUBLANG_ZULU_SOUTH_AFRICA                   = 0x01
)

const (
	SORT_DEFAULT                = 0x0
	SORT_JAPANESE_XJIS          = 0x0
	SORT_JAPANESE_UNICODE       = 0x1
	SORT_JAPANESE_RADICALSTROKE = 0x4
	SORT_CHINESE_BIG5           = 0x0
	SORT_CHINESE_PRCP           = 0x0
	SORT_CHINESE_UNICODE        = 0x1
	SORT_CHINESE_PRC            = 0x2
	SORT_CHINESE_BOPOMOFO       = 0x3
	SORT_CHINESE_RADICALSTROKE  = 0x4
	SORT_KOREAN_KSC             = 0x0
	SORT_KOREAN_UNICODE         = 0x1
	SORT_GERMAN_PHONE_BOOK      = 0x1
	SORT_HUNGARIAN_DEFAULT      = 0x0
	SORT_HUNGARIAN_TECHNICAL    = 0x1
	SORT_GEORGIAN_TRADITIONAL   = 0x0
	SORT_GEORGIAN_MODERN        = 0x1
)

func MakeLCID(languageID uint16, sortID uint16) uint32 {
	return (uint32)(sortID) << 16 | uint32(languageID)
}

func MakeSortLCID(languageID uint16, sortID uint16, sortVersion uint16) uint32 {
	return MakeLCID(languageID, sortID) | (uint32)(sortVersion) << 20
}

func LangIDFromLCID(lcid uint32) uint16 {
	return uint16(lcid)
}

func SortIDFromLCID(lcid uint32) uint16 {
	return uint16((lcid >> 16) & 0xF)
}

func SortVersionFromLCID(lcid uint32) uint16 {
	return uint16((lcid >> 20) & 0xF)
}

func MakeLangID(primaryLanguage uint16, subLanguage uint16) uint16 {
	return (subLanguage << 10) | primaryLanguage
}

func PrimaryLangID(lgid uint16) uint16 {
	return lgid & 0x03FF
}

func SubLangID(lgid uint16) uint16 {
	return lgid >> 10
}

var (
	LANG_SYSTEM_DEFAULT = MakeLangID(LANG_NEUTRAL, SUBLANG_SYS_DEFAULT)
	LANG_USER_DEFAULT   = MakeLangID(LANG_NEUTRAL, SUBLANG_DEFAULT)
)

var (
	LOCALE_SYSTEM_DEFAULT     = MakeLCID(LANG_SYSTEM_DEFAULT, SORT_DEFAULT)
	LOCALE_USER_DEFAULT       = MakeLCID(LANG_USER_DEFAULT, SORT_DEFAULT)
	LOCALE_CUSTOM_DEFAULT     = MakeLCID(MakeLangID(LANG_NEUTRAL, SUBLANG_CUSTOM_DEFAULT), SORT_DEFAULT)
	LOCALE_CUSTOM_UNSPECIFIED = MakeLCID(MakeLangID(LANG_NEUTRAL, SUBLANG_CUSTOM_UNSPECIFIED), SORT_DEFAULT)
	LOCALE_CUSTOM_UI_DEFAULT  = MakeLCID(MakeLangID(LANG_NEUTRAL, SUBLANG_UI_CUSTOM_DEFAULT), SORT_DEFAULT)
	LOCALE_NEUTRAL            = MakeLCID(MakeLangID(LANG_NEUTRAL, SUBLANG_NEUTRAL), SORT_DEFAULT)
	LOCALE_INVARIANT          = MakeLCID(MakeLangID(LANG_INVARIANT, SUBLANG_NEUTRAL), SORT_DEFAULT)
)

type SIDIdentifierAuthority struct {
	Value [6]byte
}

var (
	SECURITY_NULL_SID_AUTHORITY        = SIDIdentifierAuthority{[6]byte{0, 0, 0, 0, 0, 0}}
	SECURITY_WORLD_SID_AUTHORITY       = SIDIdentifierAuthority{[6]byte{0, 0, 0, 0, 0, 1}}
	SECURITY_LOCAL_SID_AUTHORITY       = SIDIdentifierAuthority{[6]byte{0, 0, 0, 0, 0, 2}}
	SECURITY_CREATOR_SID_AUTHORITY     = SIDIdentifierAuthority{[6]byte{0, 0, 0, 0, 0, 3}}
	SECURITY_NT_AUTHORITY              = SIDIdentifierAuthority{[6]byte{0, 0, 0, 0, 0, 5}}
	SECURITY_MANDATORY_LABEL_AUTHORITY = SIDIdentifierAuthority{[6]byte{0, 0, 0, 0, 0, 16}}
)

const (
	SECURITY_NULL_RID          = 0x00000000
	SECURITY_WORLD_RID         = 0x00000000
	SECURITY_LOCAL_RID         = 0x00000000
	SECURITY_LOCAL_LOGON_RID   = 0x00000001
	SECURITY_CREATOR_OWNER_RID = 0x00000000
	SECURITY_CREATOR_GROUP_RID = 0x00000001
)

const (
	SECURITY_DIALUP_RID                 = 0x00000001
	SECURITY_NETWORK_RID                = 0x00000002
	SECURITY_BATCH_RID                  = 0x00000003
	SECURITY_INTERACTIVE_RID            = 0x00000004
	SECURITY_LOGON_IDS_RID              = 0x00000005
	SECURITY_SERVICE_RID                = 0x00000006
	SECURITY_ANONYMOUS_LOGON_RID        = 0x00000007
	SECURITY_PROXY_RID                  = 0x00000008
	SECURITY_ENTERPRISE_CONTROLLERS_RID = 0x00000009
	SECURITY_PRINCIPAL_SELF_RID         = 0x0000000A
	SECURITY_AUTHENTICATED_USER_RID     = 0x0000000B
	SECURITY_RESTRICTED_CODE_RID        = 0x0000000C
	SECURITY_TERMINAL_SERVER_RID        = 0x0000000D
	SECURITY_LOCAL_SYSTEM_RID           = 0x00000012
	SECURITY_LOCAL_SERVICE_RID          = 0x00000013
	SECURITY_NETWORK_SERVICE_RID        = 0x00000014
	SECURITY_NT_NON_UNIQUE              = 0x00000015
	SECURITY_BUILTIN_DOMAIN_RID         = 0x00000020
)

const (
	DOMAIN_USER_RID_ADMIN = 0x000001F4
	DOMAIN_USER_RID_GUEST = 0x000001F5
)

const (
	DOMAIN_GROUP_RID_ADMINS               = 0x00000200
	DOMAIN_GROUP_RID_USERS                = 0x00000201
	DOMAIN_GROUP_RID_GUESTS               = 0x00000202
	DOMAIN_GROUP_RID_COMPUTERS            = 0x00000203
	DOMAIN_GROUP_RID_CONTROLLERS          = 0x00000204
	DOMAIN_GROUP_RID_CERT_ADMINS          = 0x00000205
	DOMAIN_GROUP_RID_SCHEMA_ADMINS        = 0x00000206
	DOMAIN_GROUP_RID_ENTERPRISE_ADMINS    = 0x00000207
	DOMAIN_GROUP_RID_POLICY_ADMINS        = 0x00000208
	DOMAIN_GROUP_RID_READONLY_CONTROLLERS = 0x00000209
)

const (
	DOMAIN_ALIAS_RID_ADMINS                         = 0x00000220
	DOMAIN_ALIAS_RID_USERS                          = 0x00000221
	DOMAIN_ALIAS_RID_GUESTS                         = 0x00000222
	DOMAIN_ALIAS_RID_POWER_USERS                    = 0x00000223
	DOMAIN_ALIAS_RID_ACCOUNT_OPS                    = 0x00000224
	DOMAIN_ALIAS_RID_SYSTEM_OPS                     = 0x00000225
	DOMAIN_ALIAS_RID_PRINT_OPS                      = 0x00000226
	DOMAIN_ALIAS_RID_BACKUP_OPS                     = 0x00000227
	DOMAIN_ALIAS_RID_REPLICATOR                     = 0x00000228
	DOMAIN_ALIAS_RID_RAS_SERVERS                    = 0x00000229
	DOMAIN_ALIAS_RID_PREW2KCOMPACCESS               = 0x0000022A
	DOMAIN_ALIAS_RID_REMOTE_DESKTOP_USERS           = 0x0000022B
	DOMAIN_ALIAS_RID_NETWORK_CONFIGURATION_OPS      = 0x0000022C
	DOMAIN_ALIAS_RID_INCOMING_FOREST_TRUST_BUILDERS = 0x0000022D
	DOMAIN_ALIAS_RID_MONITORING_USERS               = 0x0000022E
	DOMAIN_ALIAS_RID_LOGGING_USERS                  = 0x0000022F
	DOMAIN_ALIAS_RID_AUTHORIZATIONACCESS            = 0x00000230
	DOMAIN_ALIAS_RID_TS_LICENSE_SERVERS             = 0x00000231
	DOMAIN_ALIAS_RID_DCOM_USERS                     = 0x00000232
	DOMAIN_ALIAS_RID_IUSERS                         = 0x00000238
	DOMAIN_ALIAS_RID_CRYPTO_OPERATORS               = 0x00000239
	DOMAIN_ALIAS_RID_CACHEABLE_PRINCIPALS_GROUP     = 0x0000023B
	DOMAIN_ALIAS_RID_NON_CACHEABLE_PRINCIPALS_GROUP = 0x0000023C
	DOMAIN_ALIAS_RID_EVENT_LOG_READERS_GROUP        = 0x0000023D
	DOMAIN_ALIAS_RID_CERTSVC_DCOM_ACCESS_GROUP      = 0x0000023E
)

const (
	SECURITY_MANDATORY_UNTRUSTED_RID         = 0x00000000
	SECURITY_MANDATORY_LOW_RID               = 0x00001000
	SECURITY_MANDATORY_MEDIUM_RID            = 0x00002000
	SECURITY_MANDATORY_MEDIUM_PLUS_RID       = SECURITY_MANDATORY_MEDIUM_RID + 0x00000100
	SECURITY_MANDATORY_HIGH_RID              = 0x00003000
	SECURITY_MANDATORY_SYSTEM_RID            = 0x00004000
	SECURITY_MANDATORY_PROTECTED_PROCESS_RID = 0x00005000
)

type TokenOwnerData struct {
	Owner *syscall.SID
}

const (
	OWNER_SECURITY_INFORMATION            = 0x00000001
	GROUP_SECURITY_INFORMATION            = 0x00000002
	DACL_SECURITY_INFORMATION             = 0x00000004
	SACL_SECURITY_INFORMATION             = 0x00000008
	LABEL_SECURITY_INFORMATION            = 0x00000010
	PROTECTED_DACL_SECURITY_INFORMATION   = 0x80000000
	PROTECTED_SACL_SECURITY_INFORMATION   = 0x40000000
	UNPROTECTED_DACL_SECURITY_INFORMATION = 0x20000000
	UNPROTECTED_SACL_SECURITY_INFORMATION = 0x10000000
)

const (
	PROCESS_TERMINATE                 = 0x0001
	PROCESS_CREATE_THREAD             = 0x0002
	PROCESS_SET_SESSIONID             = 0x0004
	PROCESS_VM_OPERATION              = 0x0008
	PROCESS_VM_READ                   = 0x0010
	PROCESS_VM_WRITE                  = 0x0020
	PROCESS_DUP_HANDLE                = 0x0040
	PROCESS_CREATE_PROCESS            = 0x0080
	PROCESS_SET_QUOTA                 = 0x0100
	PROCESS_SET_INFORMATION           = 0x0200
	PROCESS_QUERY_INFORMATION         = 0x0400
	PROCESS_SUSPEND_RESUME            = 0x0800
	PROCESS_QUERY_LIMITED_INFORMATION = 0x1000
	PROCESS_ALL_ACCESS                = syscall.STANDARD_RIGHTS_REQUIRED | syscall.SYNCHRONIZE | 0xFFFF
)

type OSVERSIONINFO struct {
	OSVersionInfoSize uint32
	MajorVersion      uint32
	MinorVersion      uint32
	BuildNumber       uint32
	PlatformId        uint32
	CSDVersion        [128]uint16
}

type OSVERSIONINFOEX struct {
	OSVERSIONINFO
	ServicePackMajor uint16
	ServicePackMinor uint16
	SuiteMask        uint16
	ProductType      uint8
	Reserved         uint8
}

const (
	VER_EQUAL         = 1
	VER_GREATER       = 2
	VER_GREATER_EQUAL = 3
	VER_LESS          = 4
	VER_LESS_EQUAL    = 5
	VER_AND           = 6
	VER_OR            = 7
)

const (
	VER_MINORVERSION     = 0x00000001
	VER_MAJORVERSION     = 0x00000002
	VER_BUILDNUMBER      = 0x00000004
	VER_PLATFORMID       = 0x00000008
	VER_SERVICEPACKMINOR = 0x00000010
	VER_SERVICEPACKMAJOR = 0x00000020
	VER_SUITENAME        = 0x00000040
	VER_PRODUCT_TYPE     = 0x00000080
)

const (
	VER_NT_WORKSTATION       = 0x00000001
	VER_NT_DOMAIN_CONTROLLER = 0x00000002
	VER_NT_SERVER            = 0x00000003
)

const (
	VER_PLATFORM_WIN32s        = 0
	VER_PLATFORM_WIN32_WINDOWS = 1
	VER_PLATFORM_WIN32_NT      = 2
)

const (
	REG_OPTION_NON_VOLATILE   = 0x00000000
	REG_OPTION_VOLATILE       = 0x00000001
	REG_OPTION_CREATE_LINK    = 0x00000002
	REG_OPTION_BACKUP_RESTORE = 0x00000004
)

const (
	REG_CREATED_NEW_KEY     = 0x00000001
	REG_OPENED_EXISTING_KEY = 0x00000002
)

const (
	SERVICE_KERNEL_DRIVER       = 0x00000001
	SERVICE_FILE_SYSTEM_DRIVER  = 0x00000002
	SERVICE_ADAPTER             = 0x00000004
	SERVICE_RECOGNIZER_DRIVER   = 0x00000008
	SERVICE_DRIVER              = SERVICE_KERNEL_DRIVER | SERVICE_FILE_SYSTEM_DRIVER | SERVICE_RECOGNIZER_DRIVER
	SERVICE_WIN32_OWN_PROCESS   = 0x00000010
	SERVICE_WIN32_SHARE_PROCESS = 0x00000020
	SERVICE_WIN32               = SERVICE_WIN32_OWN_PROCESS | SERVICE_WIN32_SHARE_PROCESS
	SERVICE_INTERACTIVE_PROCESS = 0x00000100
)

const (
	SERVICE_BOOT_START   = 0x00000000
	SERVICE_SYSTEM_START = 0x00000001
	SERVICE_AUTO_START   = 0x00000002
	SERVICE_DEMAND_START = 0x00000003
	SERVICE_DISABLED     = 0x00000004
)

const (
	SERVICE_ERROR_IGNORE   = 0x00000000
	SERVICE_ERROR_NORMAL   = 0x00000001
	SERVICE_ERROR_SEVERE   = 0x00000002
	SERVICE_ERROR_CRITICAL = 0x00000003
)

var (
	procRtlMoveMemory       = modkernel32.NewProc("RtlMoveMemory")
	procVerSetConditionMask = modkernel32.NewProc("VerSetConditionMask")
)

func RtlMoveMemory(destination *byte, source *byte, length uintptr) {
	procRtlMoveMemory.Call(
		uintptr(unsafe.Pointer(destination)),
		uintptr(unsafe.Pointer(source)),
		length)
}

func VerSetConditionMask(conditionMask uint64, typeBitMask uint32, condition uint8) uint64 {
	r1, _, _ := procVerSetConditionMask.Call(
		uintptr(conditionMask),
		uintptr(typeBitMask),
		uintptr(condition))
	return uint64(r1)
}
