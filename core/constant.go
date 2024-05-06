package core

import "net/http"

const (
	GAME_SECRET_KEY = "secret_key"
	ENCRYPTION_KEY  = "encryption_key"
)

// TABLES
const (
	// MASTER
	TBL_USER_WALLET = "user_wallets"

	// USER
	TBL_USER             = "users"
	TBL_USER_INVITE      = "user_invites"
	TBL_USER_RECENT_GAME = "user_recent_games"
	TBL_USER_PREFERENCE  = "user_preferences"
	TBL_USER_SESSION     = "user_sessions"

	// GAME PORTAL
	TBL_GAME_PORTAL_FEEDBACK = "game_portal_feedbacks"
	TBL_GAME_PORTAL_REWARD   = "game_portal_rewards"
	TBL_GAME_PORTAL_AVATAR   = "game_portal_avatars"
	TBL_GAME_PORTAL_STICKER  = "game_portal_stickers"
	TBL_GAME_PORTAL_BG_IMG   = "game_portal_bg_imgs"

	// DASHBOARD
	TBL_DEV         = "developers"
	TBL_DEV_INVITE  = "dev_invites"
	TBL_DEV_SESSION = "dev_sessions"

	// GAME PLAYER
	TBL_GAME_PLAYER          = "game_players"
	TBL_GAME_PLAYER_DATA     = "game_player_data"
	TBL_GAME_PLAYER_PURCHASE = "game_player_purchases"
	TBL_GAME_PLAYER_SESSION  = "game_player_session"

	// GAME
	TBL_GAME          = "games"
	TBL_GAME_COMMENT  = "game_comments"
	TBL_GAME_RATING   = "game_ratings"
	TBL_GAME_FEEDBACK = "game_feedbacks"

	// IN-GAME
	TBL_GAME_CATALOG           = "game_catalogs"
	TBL_GAME_DELETED_CATALOG   = "game_deleted_catalog_items"
	TBL_GAME_ITEM              = "game_items"
	TBL_GAME_CURRENCY          = "game_currencies"
	TBL_GAME_STATISTIC         = "game_statistics"
	TBL_GAME_STAT_HISTORY      = "game_statistic_histories"
	TBL_GAME_STAT_PRIZE        = "game_statistic_prizes"
	TBL_GAME_STAT_PRIZE_ACTION = "game_statistics_prize_actions"
)

// Response Codes
const (
	// Success Codes
	Code_Success    = "Success"
	Code_Registered = "RegisterSuccess"
	Code_LoggedIn   = "LoginSuccess"
	Code_Updated    = "UpdateSuccess"
	Code_Deleted    = "DeleteSuccess"
	Code_Created    = "CreateSuccess"
	Code_Uploaded   = "UploadSuccess"

	// Error Codes
	Code_WrongCredential = "WrongCredential"

	Code_NotFound         = "NotFound"
	Code_UserNotFound     = "UserNotFound"
	Code_DevNotFound      = "DevNotFound"
	Code_EmailNotFound    = "EmailNotFound"
	Code_GameNotFound     = "GameNotFound"
	Code_CommentNotFound  = "CommentNotFound"
	Code_StickerNotFound  = "StickerNotFound"
	Code_AvatarNotFound   = "AvatarNotFound"
	Code_CatalogNotFound  = "CatalogNotFound"
	Code_ItemNotFound     = "ItemNotFound"
	Code_StatNotFound     = "StatNotFound"
	Code_CurrencyNotFound = "CurrencyNotFounc"
	Code_BgNotFound       = "BgNotFound"

	Code_InvalidId    = "InvalidId"
	Code_InvalidBody  = "InvalidBody"
	Code_InvalidURL   = "InvalidURL"
	Code_InvalidInput = "InvalidInput"
	Code_InvalidDate  = "InvalidDate"

	Code_AlreadyExists               = "AlreadyExists"
	Code_EmailAlreadyExists          = "EmailAlreadyExists"
	Code_NicknameAlreadyExists       = "NicknameAlreadyExists"
	Code_ItemIdAlreadyExists         = "ItemIdAlreadyExists"
	Code_CatalogVersionAlreadyExists = "CatalogVersionAlreadyExists"
	Code_StatCodeAlreadyExists       = "StatCodeAlreadyExists"

	Code_SessionExpired = "SessionExpired"
	Code_LinkExpired    = "LinkExpired"
	Code_CodeExpired    = "CodeExpired"

	Code_Unauthorized   = "Unauthorized"
	Code_Forbidden      = "Forbidden"
	Code_RequestTimeout = "RequestTimeout"
	Code_TooManyRequest = "TooManyRequest"
	Code_LimitExceed    = "LimitExceed"

	Code_InternalServerError = "InternalServerError"
)

// Status Codes
const (
	// Success Status Codes
	Status_Success = http.StatusOK
	Status_Created = http.StatusCreated

	// Error Status Codes
	Status_WrongCredential     = http.StatusNotFound
	Status_NotFound            = http.StatusNotFound
	Status_InvalidInput        = http.StatusBadRequest
	Status_BadRequest          = http.StatusBadRequest
	Status_Forbidden           = http.StatusForbidden
	Status_AlreadyExists       = http.StatusForbidden
	Status_LimitExceed         = http.StatusForbidden
	Status_Unauthorized        = http.StatusUnauthorized
	Status_RequestTimeout      = http.StatusRequestTimeout
	Status_TooManyRequest      = http.StatusTooManyRequests
	Status_InternalServerError = http.StatusInternalServerError
)
