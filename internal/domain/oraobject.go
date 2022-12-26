package domain

import "errors"

const (
	OracleTablespaceType = iota
	OracleDirectoryType
	OracleDbLinkType
	OracleUserType
	OracleSynonymType
	OracleContextType
	OracleSequencesType
	OracleTypeType
	OracleTableType
	OracleMLogType
	OracleMViewType
	OraclePackageType
	OracleViewType
	OracleTriggerType
	OracleVTaskType
	OracleRowType
	OracleRoleType
	OracleFunctionType
	OracleVClogType
	OracleScriptsBeforeType
	OracleScriptsAfterType
	OracleScriptsMigrationType
)

var UnknownObjectType = errors.New("can't extract object type from folder")
var FileNotExists = errors.New("file changed in git commit not exists")

type OracleObject struct {
	EpicModuleName string
	ModuleName     string
	ObjectName     string
	ObjectType     int
	Schema         string
	Server         string
	File           File
	Errors         []string
	InstallOrder   int
}