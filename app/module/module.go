package module

import (
	"go_plate/app/module/account"

	"go.uber.org/fx"
)

// Registry of all modules
var NewModule = fx.Options(
	account.NewAccountModule,
)
