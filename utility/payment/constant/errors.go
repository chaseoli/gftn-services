// © Copyright IBM Corporation 2020. All rights reserved.
// SPDX-License-Identifier: Apache2.0
// 
package constant

var (
	STATUS_CODE_DEFAULT                      = 1000
	STATUS_CODE_PARSE_FAIL                   = 1001
	STATUS_CODE_REQ_VALIDATE_FAIL            = 1002
	STATUS_CODE_DECODE_FAIL                  = 1003
	STATUS_CODE_REQUEST_CLOSE                = 1004
	STATUS_CODE_XML_VALIDATE_FAIL            = 1005
	STATUS_CODE_DUP_ID                       = 1006
	STATUS_CODE_ALREADY_SETTLED              = 1007
	STATUS_CODE_RFI_OR_OFI_NOT_IN_WL         = 1008
	STATUS_CODE_OFI_NOT_IN_WL                = 1009
	STATUS_CODE_SIGN_PAYLOAD_FAIL            = 1010
	STATUS_CODE_INTERNAL_ERROR               = 1011
	STATUS_CODE_WRONG_ORIGINAL_INCONSISTENCE = 1012
	STATUS_CODE_OFI_SIGNATURE_FAIL           = 1013
	STATUS_CODE_RFI_SIGNATURE_FAIL           = 1014
	STATUS_CODE_INTERNAL_ERROR_PARSE         = 1015
	STATUS_CODE_CALLBACK_FAIL                = 1018
	STATUS_CODE_WRONG_SETTLEMENT_METHOD      = 1019
	STATUS_CODE_ACCOUNT_NOT_EXIST            = 1020
	STATUS_CODE_WRONG_ACCOUNT_NAME           = 1021
	STATUS_CODE_WRONG_AUTHEN                 = 1022
	STATUS_CODE_WRONG_AUTHOR                 = 1023
	STATUS_CODE_WRONG_ID                     = 1024
	STATUS_CODE_WRONG_FI                     = 1025
	STATUS_CODE_JWT_FAIL                     = 1026
	STATUS_CODE_BLOCKLIST                    = 1027
	STATUS_CODE_WRONG_ASSET_ISSUER           = 1028
	STATUS_CODE_WRONG_ORIGINAL_ID            = 1029
	STATUS_CODE_WRONG_INSTRUCTION_ID         = 1030
	STATUS_CODE_WRONG_MSG_NAME_ID            = 1031
	STATUS_CODE_WRONG_RESPONSE_CODE          = 1032
	STATUS_CODE_WRONG_CCY                    = 1033
	STATUS_CODE_MSG_TYPE_VALIDATE_FAIL       = 1034
	STATUS_CODE_INSUFFICIENT_AMOUNT          = 1035
	STATUS_CODE_OFI_SEND_TO_KAFKA            = 1101
	STATUS_CODE_RFI_SEND_TO_KAFKA            = 1102
	STATUS_CODE_TX_SEND_TO_STELLAR           = 1103
	STATUS_CODE_CXL_RES_SEND_TO_KAFKA        = 1104
	STATUS_CODE_RDO_REQ_SEND_TO_KAFKA        = 1105
	STATUS_CODE_FED_COMP_RJCT                = 1201
	STATUS_CODE_PAYMENT_CXL_RJCT             = 1202
	STATUS_CODE_SUBMIT_TO_KAFKA_FAIL         = 1203
	STATUS_CODE_UPDATE_FIREBASE_FAIL         = 1204
	STATUS_CODE_UNIDENTIFIED_REASON_CODE     = 1205
	STATUS_CODE_DO_INTERNAL_TRANSFER_ERROR   = 1206
	STATUS_CODE_EMPTY_ADDRESS                = 1207
	STATUS_CODE_WRONG_MSG_ID                 = 1208
	STATUS_CODE_ORIGINAL_REQUEST_NOT_DONE    = 1209
	STATUS_CODE_ORIGINAL_REQUEST_NOT_INIT    = 1210
	STATUS_CODE_ASSET_REDEMPTION_RJCT        = 1211
	STATUS_CODE_REDEMPTION_AMOUNT_MISMATCH   = 1212
	STATUS_CODE_ALREADY_REPLIED              = 1213
	STATUS_CODE_UNABLE_TO_APPLY_NOT_INIT     = 1214
)