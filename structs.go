package boticordgo

import (
	"net/http"
)

const BotiCordURL string = "https://api.boticord.top/v3"

var Headers []map[string]interface{}

func ResponseErrorCodes() ResponseErrorCodesStruct {
	return ResponseErrorCodesStruct{
		ServerError:     500,
		TooManyRequests: 429,
		NotFound:        404,
		Forbidden:       403,
		UnAuthorized:    401,
		BadRequest:      400,
	}
}

func HTTPErrorCodes() HTTPErrorCodesStruct {
	return HTTPErrorCodesStruct{
		UnknownError:                0,
		InternalServerError:         1,
		RateLimited:                 2,
		NotFound:                    3,
		Forbidden:                   4,
		BadRequest:                  5,
		UnAuthorized:                6,
		RPCError:                    7,
		WSError:                     8,
		ThirdPartyFail:              9,
		UnknownUser:                 10,
		ShortDomainTaken:            11,
		UnknownShortDomain:          12,
		UnknownLibrary:              13,
		TokenInvalid:                14,
		UnknownResource:             15,
		UnknownTag:                  16,
		PermissionDenied:            17,
		UnknownComment:              18,
		UnknownBot:                  19,
		UnknownServer:               20,
		UnknownBadge:                21,
		UserAlreadyHasBadge:         22,
		InvalidInviteCode:           23,
		ServerAlreadyExists:         24,
		BotNotPresentOnQueueServer:  25,
		UnknownUP:                   26,
		TooManyUPS:                  27,
		InvalidStatus:               28,
		UnknownReport:               29,
		UnSupportedMediaType:        30,
		UnknownApplication:          31,
		AutomatedRequestsNotAllowed: 32,
	}
}

func ResourceStatus() ResourceStatusStruct {
	return ResourceStatusStruct{
		Hidden:  0,
		Public:  1,
		Banned:  2,
		Pending: 3,
	}
}

type HTTPErrorCodesStruct struct {
	UnknownError                int
	InternalServerError         int
	RateLimited                 int
	NotFound                    int
	Forbidden                   int
	BadRequest                  int
	UnAuthorized                int
	RPCError                    int
	WSError                     int
	ThirdPartyFail              int
	UnknownUser                 int
	ShortDomainTaken            int
	UnknownShortDomain          int
	UnknownLibrary              int
	TokenInvalid                int
	UnknownResource             int
	UnknownTag                  int
	PermissionDenied            int
	UnknownComment              int
	UnknownBot                  int
	UnknownServer               int
	UnknownBadge                int
	UserAlreadyHasBadge         int
	InvalidInviteCode           int
	ServerAlreadyExists         int
	BotNotPresentOnQueueServer  int
	UnknownUP                   int
	TooManyUPS                  int
	InvalidStatus               int
	UnknownReport               int
	UnSupportedMediaType        int
	UnknownApplication          int
	AutomatedRequestsNotAllowed int
}

type ResponseErrorCodesStruct struct {
	ServerError     int
	TooManyRequests int
	NotFound        int
	Forbidden       int
	UnAuthorized    int
	BadRequest      int
}

type ResourceStatusStruct struct {
	Hidden  int
	Public  int
	Banned  int
	Pending int
}

type Client struct {
	BaseURL    string
	Token      string
	HTTPClient *http.Client
}

type StatsPayload struct {
	Servers int `json:"servers"`
	Shards  int `json:"shards"`
	Users   int `json:"users"`
}

type BotInfo struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Servers int    `json:"servers"`
}

type BotWritenLibrary struct {
	Discord4J  int
	Discordcr  int
	DiscordGo  int
	Discordoo  int
	DSharpPlus int
	DiscordJS  int
	DiscordNET int
}
