package oracle

//#cgo CFLAGS: -I /Users/xiebo/oracle/source/sdk/include
//#cgo LDFLAGS: -L /Users/xiebo/oracle/lib -lclntsh
//#include "stdlib.h"
//#include "oci.h"
//#include "stdio.h"
import "C"

const (
	/*-----------------------------  Various Modes ------------------------------*/
	OCI_DEFAULT = C.OCI_DEFAULT
	/*-------------OCIInitialize Modes / OCICreateEnvironment Modes -------------*/
	OCI_THREADED                  = C.OCI_THREADED
	OCI_OBJECT                    = C.OCI_OBJECT
	OCI_EVENTS                    = C.OCI_EVENTS
	OCI_RESERVED1                 = C.OCI_RESERVED1
	OCI_SHARED                    = C.OCI_SHARED
	OCI_RESERVED2                 = C.OCI_RESERVED2
	OCI_NO_UCB                    = C.OCI_NO_UCB
	OCI_NO_MUTEX                  = C.OCI_NO_MUTEX
	OCI_SHARED_EXT                = C.OCI_SHARED_EXT
	OCI_ALWAYS_BLOCKING           = C.OCI_ALWAYS_BLOCKING
	OCI_USE_LDAP                  = C.OCI_USE_LDAP
	OCI_REG_LDAPONLY              = C.OCI_REG_LDAPONLY
	OCI_UTF16                     = C.OCI_UTF16
	OCI_AFC_PAD_ON                = C.OCI_AFC_PAD_ON
	OCI_ENVCR_RESERVED3           = C.OCI_ENVCR_RESERVED3
	OCI_NEW_LENGTH_SEMANTICS      = C.OCI_NEW_LENGTH_SEMANTICS
	OCI_NO_MUTEX_STMT             = C.OCI_NO_MUTEX_STMT
	OCI_MUTEX_ENV_ONLY            = C.OCI_MUTEX_ENV_ONLY
	OCI_SUPPRESS_NLS_VALIDATION   = C.OCI_SUPPRESS_NLS_VALIDATION
	OCI_MUTEX_TRY                 = C.OCI_MUTEX_TRY
	OCI_NCHAR_LITERAL_REPLACE_ON  = C.OCI_NCHAR_LITERAL_REPLACE_ON
	OCI_NCHAR_LITERAL_REPLACE_OFF = C.OCI_NCHAR_LITERAL_REPLACE_OFF
	OCI_ENABLE_NLS_VALIDATION     = C.OCI_ENABLE_NLS_VALIDATION
	OCI_ENVCR_RESERVED4           = C.OCI_ENVCR_RESERVED4
	OCI_ENVCR_RESERVED5           = C.OCI_ENVCR_RESERVED5
	OCI_ENVCR_RESERVED6           = C.OCI_ENVCR_RESERVED6
	OCI_ENVCR_RESERVED7           = C.OCI_ENVCR_RESERVED7
	OCI_SECURE_NOTIFICATION       = C.OCI_SECURE_NOTIFICATION
	OCI_DISABLE_DIAG              = C.OCI_DISABLE_DIAG

	/*-----------------------------Handle Types----------------------------------*/
	OCI_HTYPE_FIRST                = C.OCI_HTYPE_FIRST
	OCI_HTYPE_ENV                  = C.OCI_HTYPE_ENV
	OCI_HTYPE_ERROR                = C.OCI_HTYPE_ERROR
	OCI_HTYPE_SVCCTX               = C.OCI_HTYPE_SVCCTX
	OCI_HTYPE_STMT                 = C.OCI_HTYPE_STMT
	OCI_HTYPE_BIND                 = C.OCI_HTYPE_BIND
	OCI_HTYPE_DEFINE               = C.OCI_HTYPE_DEFINE
	OCI_HTYPE_DESCRIBE             = C.OCI_HTYPE_DESCRIBE
	OCI_HTYPE_SERVER               = C.OCI_HTYPE_SERVER
	OCI_HTYPE_SESSION              = C.OCI_HTYPE_SESSION
	OCI_HTYPE_AUTHINFO             = C.OCI_HTYPE_AUTHINFO
	OCI_HTYPE_TRANS                = C.OCI_HTYPE_TRANS
	OCI_HTYPE_COMPLEXOBJECT        = C.OCI_HTYPE_COMPLEXOBJECT
	OCI_HTYPE_SECURITY             = C.OCI_HTYPE_SECURITY
	OCI_HTYPE_SUBSCRIPTION         = C.OCI_HTYPE_SUBSCRIPTION
	OCI_HTYPE_DIRPATH_CTX          = C.OCI_HTYPE_DIRPATH_CTX
	OCI_HTYPE_DIRPATH_COLUMN_ARRAY = C.OCI_HTYPE_DIRPATH_COLUMN_ARRAY
	OCI_HTYPE_DIRPATH_STREAM       = C.OCI_HTYPE_DIRPATH_STREAM
	OCI_HTYPE_PROC                 = C.OCI_HTYPE_PROC
	OCI_HTYPE_DIRPATH_FN_CTX       = C.OCI_HTYPE_DIRPATH_FN_CTX
	OCI_HTYPE_DIRPATH_FN_COL_ARRAY = C.OCI_HTYPE_DIRPATH_FN_COL_ARRAY
	OCI_HTYPE_XADSESSION           = C.OCI_HTYPE_XADSESSION
	OCI_HTYPE_XADTABLE             = C.OCI_HTYPE_XADTABLE
	OCI_HTYPE_XADFIELD             = C.OCI_HTYPE_XADFIELD
	OCI_HTYPE_XADGRANULE           = C.OCI_HTYPE_XADGRANULE
	OCI_HTYPE_XADRECORD            = C.OCI_HTYPE_XADRECORD
	OCI_HTYPE_XADIO                = C.OCI_HTYPE_XADIO
	OCI_HTYPE_CPOOL                = C.OCI_HTYPE_CPOOL
	OCI_HTYPE_SPOOL                = C.OCI_HTYPE_SPOOL
	OCI_HTYPE_ADMIN                = C.OCI_HTYPE_ADMIN
	OCI_HTYPE_EVENT                = C.OCI_HTYPE_EVENT
	/*------------------------Error Return Values--------------------------------*/
	OCI_SUCCESS              = C.OCI_SUCCESS
	OCI_SUCCESS_WITH_INFO    = C.OCI_SUCCESS_WITH_INFO
	OCI_RESERVED_FOR_INT_USE = C.OCI_RESERVED_FOR_INT_USE
	OCI_NO_DATA              = C.OCI_NO_DATA
	OCI_ERROR                = C.OCI_ERROR
	OCI_INVALID_HANDLE       = C.OCI_INVALID_HANDLE
	OCI_NEED_DATA            = C.OCI_NEED_DATA
	OCI_STILL_EXECUTING      = C.OCI_STILL_EXECUTING

	OCI_ATTR_SERVER   = C.OCI_ATTR_SERVER
	OCI_ATTR_SESSION  = C.OCI_ATTR_SESSION
	OCI_ATTR_USERNAME = C.OCI_ATTR_USERNAME
	OCI_ATTR_PASSWORD = C.OCI_ATTR_PASSWORD

	OCI_CRED_RDBMS = C.OCI_CRED_RDBMS

	OCI_NTV_SYNTAX = C.OCI_NTV_SYNTAX

	OCI_TYPECODE_CHAR     = C.OCI_TYPECODE_CHAR
	OCI_TYPECODE_VARCHAR2 = C.OCI_TYPECODE_VARCHAR2
	OCI_TYPECODE_VARCHAR  = C.OCI_TYPECODE_VARCHAR
	OCI_TYPECODE_INTEGER  = C.OCI_TYPECODE_INTEGER
	OCI_TYPECODE_SMALLINT = C.OCI_TYPECODE_SMALLINT
	OCI_TYPECODE_OCTET    = C.OCI_TYPECODE_OCTET

	OCI_DTYPE_PARAM = C.OCI_DTYPE_PARAM

	SQLT_STR = C.SQLT_STR
	SQLT_FLT = C.SQLT_FLT
	SQLT_INT = C.SQLT_INT
	SQLT_DAT = C.SQLT_DAT
)
