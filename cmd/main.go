package main

import (
	"github.com/MachadoMichael/notifications/infra"
	"github.com/MachadoMichael/notifications/infra/database"
	"github.com/MachadoMichael/notifications/pkg/route"
	"github.com/MachadoMichael/notifications/pkg/whatsapp"
)

func main() {
	infra.Init()
	whatsapp.Init()
	database.Init()
	route.Init()
}
