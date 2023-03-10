package domain

import (
	"fmt"
)

const (
	AddAction    int = 1
	DeleteAction int = 2
	ModifyAction int = 3
	RenameAction int = 4
)

const (
	Data int = iota
	Install
	ErrorLog
	WarningLog
)

const (
	ErrorLogFileName              string = "error_log.txt"
	WarningLogFileName            string = "warning_log.txt"
	VtbsCoreInstallName           string = "10_VTBS_CORE.sql"
	VtbsHpffmInstallName          string = "20_VTBS_HPFFM.sql"
	VtbsAdeskHpffmInstallName     string = "30_VTBS_ADESK_HPFFM.sql"
	VtbsXAlarisHpffmInstallName   string = "40_VTBS_X_ALARIS_HPFFM.sql"
	VtbsBiHpffmInstallName        string = "50_VTBS_BI_HPFFM.sql"
	VtbsCoreMigrationName         string = "90_VTBS_CORE_MIGRATION.sql"
	VtbsHpffmMigrationName        string = "91_VTBS_HPFFM_MIGRATION.sql"
	VtbsAdeskHpffmMigrationName   string = "92_VTBS_ADESK_HPFFM_MIGRATION.sql"
	VtbsXAlarisHpffmMigrationName string = "93_VTBS_X_ALARIS_HPFFM_MIGRATION.sql"
	VtbsBiHpffmMigrationName      string = "94_VTBS_BI_HPFFM_MIGRATION.sql"
)

// GitFileDetails add and handle initialAction to handle add -> modify -> delete sequence
type GitFileDetails struct {
	InitialName string
	InitialPath string
	Comment     string
	Action      int
	New         bool
}

type File struct {
	Name       string
	Path       string
	FileLines  []string
	GitDetails GitFileDetails
}

func (f File) FormShortInfoStr() string {
	return fmt.Sprintf("Name %s Path %s  InitialName %s InitialPath %s Comment %s Action %s New %t",
		f.Name, f.Path, f.GitDetails.InitialName, f.GitDetails.InitialPath, ActionNameDict[f.GitDetails.Action], f.GitDetails.Comment, f.GitDetails.New)
}

type OracleFile struct {
	OracleDataType int
	FileDetails    File
}

var ActionNameDict = map[int]string{
	AddAction:    "Add",
	DeleteAction: "Delete",
	ModifyAction: "Modify",
	RenameAction: "Rename",
}
