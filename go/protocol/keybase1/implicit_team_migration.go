// Auto-generated by avdl-compiler v1.3.24 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/implicit_team_migration.avdl

package keybase1

import (
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type StartMigrationArg struct {
	Folder Folder `codec:"folder" json:"folder"`
}

type FinalizeMigrationArg struct {
	Folder Folder `codec:"folder" json:"folder"`
}

type ImplicitTeamMigrationInterface interface {
	StartMigration(context.Context, Folder) error
	FinalizeMigration(context.Context, Folder) error
}

func ImplicitTeamMigrationProtocol(i ImplicitTeamMigrationInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.implicitTeamMigration",
		Methods: map[string]rpc.ServeHandlerDescription{
			"startMigration": {
				MakeArg: func() interface{} {
					ret := make([]StartMigrationArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]StartMigrationArg)
					if !ok {
						err = rpc.NewTypeError((*[]StartMigrationArg)(nil), args)
						return
					}
					err = i.StartMigration(ctx, (*typedArgs)[0].Folder)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"finalizeMigration": {
				MakeArg: func() interface{} {
					ret := make([]FinalizeMigrationArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]FinalizeMigrationArg)
					if !ok {
						err = rpc.NewTypeError((*[]FinalizeMigrationArg)(nil), args)
						return
					}
					err = i.FinalizeMigration(ctx, (*typedArgs)[0].Folder)
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type ImplicitTeamMigrationClient struct {
	Cli rpc.GenericClient
}

func (c ImplicitTeamMigrationClient) StartMigration(ctx context.Context, folder Folder) (err error) {
	__arg := StartMigrationArg{Folder: folder}
	err = c.Cli.Call(ctx, "keybase.1.implicitTeamMigration.startMigration", []interface{}{__arg}, nil)
	return
}

func (c ImplicitTeamMigrationClient) FinalizeMigration(ctx context.Context, folder Folder) (err error) {
	__arg := FinalizeMigrationArg{Folder: folder}
	err = c.Cli.Call(ctx, "keybase.1.implicitTeamMigration.finalizeMigration", []interface{}{__arg}, nil)
	return
}
