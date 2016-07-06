// This file was generated by counterfeiter
package fakes

import "sync"

type IPTables struct {
	ExistsStub        func(table, chain string, rulespec ...string) (bool, error)
	existsMutex       sync.RWMutex
	existsArgsForCall []struct {
		table    string
		chain    string
		rulespec []string
	}
	existsReturns struct {
		result1 bool
		result2 error
	}
	InsertStub        func(table, chain string, pos int, rulespec ...string) error
	insertMutex       sync.RWMutex
	insertArgsForCall []struct {
		table    string
		chain    string
		pos      int
		rulespec []string
	}
	insertReturns struct {
		result1 error
	}
	AppendUniqueStub        func(table, chain string, rulespec ...string) error
	appendUniqueMutex       sync.RWMutex
	appendUniqueArgsForCall []struct {
		table    string
		chain    string
		rulespec []string
	}
	appendUniqueReturns struct {
		result1 error
	}
	DeleteStub        func(table, chain string, rulespec ...string) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		table    string
		chain    string
		rulespec []string
	}
	deleteReturns struct {
		result1 error
	}
	ListStub        func(table, chain string) ([]string, error)
	listMutex       sync.RWMutex
	listArgsForCall []struct {
		table string
		chain string
	}
	listReturns struct {
		result1 []string
		result2 error
	}
	NewChainStub        func(table, chain string) error
	newChainMutex       sync.RWMutex
	newChainArgsForCall []struct {
		table string
		chain string
	}
	newChainReturns struct {
		result1 error
	}
	ClearChainStub        func(table, chain string) error
	clearChainMutex       sync.RWMutex
	clearChainArgsForCall []struct {
		table string
		chain string
	}
	clearChainReturns struct {
		result1 error
	}
	DeleteChainStub        func(table, chain string) error
	deleteChainMutex       sync.RWMutex
	deleteChainArgsForCall []struct {
		table string
		chain string
	}
	deleteChainReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *IPTables) Exists(table string, chain string, rulespec ...string) (bool, error) {
	fake.existsMutex.Lock()
	fake.existsArgsForCall = append(fake.existsArgsForCall, struct {
		table    string
		chain    string
		rulespec []string
	}{table, chain, rulespec})
	fake.recordInvocation("Exists", []interface{}{table, chain, rulespec})
	fake.existsMutex.Unlock()
	if fake.ExistsStub != nil {
		return fake.ExistsStub(table, chain, rulespec...)
	} else {
		return fake.existsReturns.result1, fake.existsReturns.result2
	}
}

func (fake *IPTables) ExistsCallCount() int {
	fake.existsMutex.RLock()
	defer fake.existsMutex.RUnlock()
	return len(fake.existsArgsForCall)
}

func (fake *IPTables) ExistsArgsForCall(i int) (string, string, []string) {
	fake.existsMutex.RLock()
	defer fake.existsMutex.RUnlock()
	return fake.existsArgsForCall[i].table, fake.existsArgsForCall[i].chain, fake.existsArgsForCall[i].rulespec
}

func (fake *IPTables) ExistsReturns(result1 bool, result2 error) {
	fake.ExistsStub = nil
	fake.existsReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *IPTables) Insert(table string, chain string, pos int, rulespec ...string) error {
	fake.insertMutex.Lock()
	fake.insertArgsForCall = append(fake.insertArgsForCall, struct {
		table    string
		chain    string
		pos      int
		rulespec []string
	}{table, chain, pos, rulespec})
	fake.recordInvocation("Insert", []interface{}{table, chain, pos, rulespec})
	fake.insertMutex.Unlock()
	if fake.InsertStub != nil {
		return fake.InsertStub(table, chain, pos, rulespec...)
	} else {
		return fake.insertReturns.result1
	}
}

func (fake *IPTables) InsertCallCount() int {
	fake.insertMutex.RLock()
	defer fake.insertMutex.RUnlock()
	return len(fake.insertArgsForCall)
}

func (fake *IPTables) InsertArgsForCall(i int) (string, string, int, []string) {
	fake.insertMutex.RLock()
	defer fake.insertMutex.RUnlock()
	return fake.insertArgsForCall[i].table, fake.insertArgsForCall[i].chain, fake.insertArgsForCall[i].pos, fake.insertArgsForCall[i].rulespec
}

func (fake *IPTables) InsertReturns(result1 error) {
	fake.InsertStub = nil
	fake.insertReturns = struct {
		result1 error
	}{result1}
}

func (fake *IPTables) AppendUnique(table string, chain string, rulespec ...string) error {
	fake.appendUniqueMutex.Lock()
	fake.appendUniqueArgsForCall = append(fake.appendUniqueArgsForCall, struct {
		table    string
		chain    string
		rulespec []string
	}{table, chain, rulespec})
	fake.recordInvocation("AppendUnique", []interface{}{table, chain, rulespec})
	fake.appendUniqueMutex.Unlock()
	if fake.AppendUniqueStub != nil {
		return fake.AppendUniqueStub(table, chain, rulespec...)
	} else {
		return fake.appendUniqueReturns.result1
	}
}

func (fake *IPTables) AppendUniqueCallCount() int {
	fake.appendUniqueMutex.RLock()
	defer fake.appendUniqueMutex.RUnlock()
	return len(fake.appendUniqueArgsForCall)
}

func (fake *IPTables) AppendUniqueArgsForCall(i int) (string, string, []string) {
	fake.appendUniqueMutex.RLock()
	defer fake.appendUniqueMutex.RUnlock()
	return fake.appendUniqueArgsForCall[i].table, fake.appendUniqueArgsForCall[i].chain, fake.appendUniqueArgsForCall[i].rulespec
}

func (fake *IPTables) AppendUniqueReturns(result1 error) {
	fake.AppendUniqueStub = nil
	fake.appendUniqueReturns = struct {
		result1 error
	}{result1}
}

func (fake *IPTables) Delete(table string, chain string, rulespec ...string) error {
	fake.deleteMutex.Lock()
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		table    string
		chain    string
		rulespec []string
	}{table, chain, rulespec})
	fake.recordInvocation("Delete", []interface{}{table, chain, rulespec})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(table, chain, rulespec...)
	} else {
		return fake.deleteReturns.result1
	}
}

func (fake *IPTables) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *IPTables) DeleteArgsForCall(i int) (string, string, []string) {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].table, fake.deleteArgsForCall[i].chain, fake.deleteArgsForCall[i].rulespec
}

func (fake *IPTables) DeleteReturns(result1 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *IPTables) List(table string, chain string) ([]string, error) {
	fake.listMutex.Lock()
	fake.listArgsForCall = append(fake.listArgsForCall, struct {
		table string
		chain string
	}{table, chain})
	fake.recordInvocation("List", []interface{}{table, chain})
	fake.listMutex.Unlock()
	if fake.ListStub != nil {
		return fake.ListStub(table, chain)
	} else {
		return fake.listReturns.result1, fake.listReturns.result2
	}
}

func (fake *IPTables) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *IPTables) ListArgsForCall(i int) (string, string) {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return fake.listArgsForCall[i].table, fake.listArgsForCall[i].chain
}

func (fake *IPTables) ListReturns(result1 []string, result2 error) {
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *IPTables) NewChain(table string, chain string) error {
	fake.newChainMutex.Lock()
	fake.newChainArgsForCall = append(fake.newChainArgsForCall, struct {
		table string
		chain string
	}{table, chain})
	fake.recordInvocation("NewChain", []interface{}{table, chain})
	fake.newChainMutex.Unlock()
	if fake.NewChainStub != nil {
		return fake.NewChainStub(table, chain)
	} else {
		return fake.newChainReturns.result1
	}
}

func (fake *IPTables) NewChainCallCount() int {
	fake.newChainMutex.RLock()
	defer fake.newChainMutex.RUnlock()
	return len(fake.newChainArgsForCall)
}

func (fake *IPTables) NewChainArgsForCall(i int) (string, string) {
	fake.newChainMutex.RLock()
	defer fake.newChainMutex.RUnlock()
	return fake.newChainArgsForCall[i].table, fake.newChainArgsForCall[i].chain
}

func (fake *IPTables) NewChainReturns(result1 error) {
	fake.NewChainStub = nil
	fake.newChainReturns = struct {
		result1 error
	}{result1}
}

func (fake *IPTables) ClearChain(table string, chain string) error {
	fake.clearChainMutex.Lock()
	fake.clearChainArgsForCall = append(fake.clearChainArgsForCall, struct {
		table string
		chain string
	}{table, chain})
	fake.recordInvocation("ClearChain", []interface{}{table, chain})
	fake.clearChainMutex.Unlock()
	if fake.ClearChainStub != nil {
		return fake.ClearChainStub(table, chain)
	} else {
		return fake.clearChainReturns.result1
	}
}

func (fake *IPTables) ClearChainCallCount() int {
	fake.clearChainMutex.RLock()
	defer fake.clearChainMutex.RUnlock()
	return len(fake.clearChainArgsForCall)
}

func (fake *IPTables) ClearChainArgsForCall(i int) (string, string) {
	fake.clearChainMutex.RLock()
	defer fake.clearChainMutex.RUnlock()
	return fake.clearChainArgsForCall[i].table, fake.clearChainArgsForCall[i].chain
}

func (fake *IPTables) ClearChainReturns(result1 error) {
	fake.ClearChainStub = nil
	fake.clearChainReturns = struct {
		result1 error
	}{result1}
}

func (fake *IPTables) DeleteChain(table string, chain string) error {
	fake.deleteChainMutex.Lock()
	fake.deleteChainArgsForCall = append(fake.deleteChainArgsForCall, struct {
		table string
		chain string
	}{table, chain})
	fake.recordInvocation("DeleteChain", []interface{}{table, chain})
	fake.deleteChainMutex.Unlock()
	if fake.DeleteChainStub != nil {
		return fake.DeleteChainStub(table, chain)
	} else {
		return fake.deleteChainReturns.result1
	}
}

func (fake *IPTables) DeleteChainCallCount() int {
	fake.deleteChainMutex.RLock()
	defer fake.deleteChainMutex.RUnlock()
	return len(fake.deleteChainArgsForCall)
}

func (fake *IPTables) DeleteChainArgsForCall(i int) (string, string) {
	fake.deleteChainMutex.RLock()
	defer fake.deleteChainMutex.RUnlock()
	return fake.deleteChainArgsForCall[i].table, fake.deleteChainArgsForCall[i].chain
}

func (fake *IPTables) DeleteChainReturns(result1 error) {
	fake.DeleteChainStub = nil
	fake.deleteChainReturns = struct {
		result1 error
	}{result1}
}

func (fake *IPTables) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.existsMutex.RLock()
	defer fake.existsMutex.RUnlock()
	fake.insertMutex.RLock()
	defer fake.insertMutex.RUnlock()
	fake.appendUniqueMutex.RLock()
	defer fake.appendUniqueMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	fake.newChainMutex.RLock()
	defer fake.newChainMutex.RUnlock()
	fake.clearChainMutex.RLock()
	defer fake.clearChainMutex.RUnlock()
	fake.deleteChainMutex.RLock()
	defer fake.deleteChainMutex.RUnlock()
	return fake.invocations
}

func (fake *IPTables) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}