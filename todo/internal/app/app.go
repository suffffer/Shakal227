package app

import db "github.com/suffffer/Shakal227/internal/database"

func RunApp() error {
	if err = db.Run(); err != nil {
		return err
	}
	if err = StartServer(); err != nil {
		return err
	}
	return nil
}
