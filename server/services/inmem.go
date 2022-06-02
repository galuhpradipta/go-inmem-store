package services

import (
	"errors"
	"strings"

	"github.com/galuhpradipta/go-inmem-db/server/store"
)

type InMemHandler interface {
	Run(cmd string) (string, error)
}

type inMemHandler struct {
	store store.StoreHandler
}

const (
	cmdSet    = "SET"
	cmdDump   = "DUMP"
	cmdDelete = "DELETE"
	cmdRename = "RENAME"
	cmdHelp   = "HELP"

	helpRes = `
Usage:
	<command> [arguments]

The commands are:
	SET 	[key] [value] 	- set the value of the key
	DUMP 	[key] 		- get the value of the key
	DELETE 	[key] 		- delete the key
;`
)

var (
	errCmdNotFound     = errors.New("command not found")
	errInvalidCommand  = errors.New("invalid command")
	errKeyAlreadyExist = errors.New("key already exist")
	errNotFound        = errors.New("key value not found")
)

func NewInMemHandler(store store.StoreHandler) InMemHandler {
	return &inMemHandler{
		store: store,
	}
}

func (i *inMemHandler) Run(cmd string) (string, error) {
	cmds := strings.Split(cmd, " ")
	headerCommand := cmds[0]
	switch headerCommand {
	case cmdHelp:
		return helpRes, nil
	case cmdSet:
		return i.set(cmds)
	case cmdDump:
		return i.dump(cmds)
	case cmdRename:
		return i.rename(cmds)
	case cmdDelete:
		return i.delete(cmds)
	}

	return "", errCmdNotFound
}

func (i *inMemHandler) set(cmds []string) (string, error) {
	if len(cmds) != 3 {
		return "", errInvalidCommand
	}
	k, v := cmds[1], cmds[2]

	if i.store.Get(k) != "" {
		return "", errKeyAlreadyExist
	}

	i.store.Set(k, v)
	return "set command success", nil
}

func (i *inMemHandler) dump(cmds []string) (string, error) {
	if len(cmds) != 2 {
		return "", errInvalidCommand
	}

	value := i.store.Get(cmds[1])
	if value == "" {
		return "", errNotFound
	}

	return value, nil
}

func (i *inMemHandler) rename(cmds []string) (string, error) {
	if len(cmds) != 3 {
		return "", errInvalidCommand
	}

	oldKey, newKey := cmds[1], cmds[2]
	value := i.store.Get(oldKey)
	if value == "" {
		return "", errNotFound
	}

	i.store.Delete(oldKey)
	i.store.Set(newKey, value)

	return "rename command success", nil
}

func (i *inMemHandler) delete(cmds []string) (string, error) {
	if len(cmds) != 2 {
		return "", errInvalidCommand
	}

	i.store.Delete(cmds[1])
	return "delete command success", nil
}
