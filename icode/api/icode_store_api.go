package api

import "github.com/it-chain/it-chain-Engine/icode/domain/model"

//Api to import or store icode from outside
type ItCodeStoreApi interface {

	//get icode from outside
	Clone(repositoryUrl string) (*model.ICodeMeta, error)

	//push code to auth repo
	Push(icodeMeta *model.ICodeMeta) error
}