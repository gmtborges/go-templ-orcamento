// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("AutoCategories", testAutoCategories)
	t.Run("AutoOffers", testAutoOffers)
	t.Run("AutoStores", testAutoStores)
	t.Run("Biddings", testBiddings)
	t.Run("Companies", testCompanies)
	t.Run("Employers", testEmployers)
	t.Run("Users", testUsers)
}

func TestDelete(t *testing.T) {
	t.Run("AutoCategories", testAutoCategoriesDelete)
	t.Run("AutoOffers", testAutoOffersDelete)
	t.Run("AutoStores", testAutoStoresDelete)
	t.Run("Biddings", testBiddingsDelete)
	t.Run("Companies", testCompaniesDelete)
	t.Run("Employers", testEmployersDelete)
	t.Run("Users", testUsersDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("AutoCategories", testAutoCategoriesQueryDeleteAll)
	t.Run("AutoOffers", testAutoOffersQueryDeleteAll)
	t.Run("AutoStores", testAutoStoresQueryDeleteAll)
	t.Run("Biddings", testBiddingsQueryDeleteAll)
	t.Run("Companies", testCompaniesQueryDeleteAll)
	t.Run("Employers", testEmployersQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("AutoCategories", testAutoCategoriesSliceDeleteAll)
	t.Run("AutoOffers", testAutoOffersSliceDeleteAll)
	t.Run("AutoStores", testAutoStoresSliceDeleteAll)
	t.Run("Biddings", testBiddingsSliceDeleteAll)
	t.Run("Companies", testCompaniesSliceDeleteAll)
	t.Run("Employers", testEmployersSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("AutoCategories", testAutoCategoriesExists)
	t.Run("AutoOffers", testAutoOffersExists)
	t.Run("AutoStores", testAutoStoresExists)
	t.Run("Biddings", testBiddingsExists)
	t.Run("Companies", testCompaniesExists)
	t.Run("Employers", testEmployersExists)
	t.Run("Users", testUsersExists)
}

func TestFind(t *testing.T) {
	t.Run("AutoCategories", testAutoCategoriesFind)
	t.Run("AutoOffers", testAutoOffersFind)
	t.Run("AutoStores", testAutoStoresFind)
	t.Run("Biddings", testBiddingsFind)
	t.Run("Companies", testCompaniesFind)
	t.Run("Employers", testEmployersFind)
	t.Run("Users", testUsersFind)
}

func TestBind(t *testing.T) {
	t.Run("AutoCategories", testAutoCategoriesBind)
	t.Run("AutoOffers", testAutoOffersBind)
	t.Run("AutoStores", testAutoStoresBind)
	t.Run("Biddings", testBiddingsBind)
	t.Run("Companies", testCompaniesBind)
	t.Run("Employers", testEmployersBind)
	t.Run("Users", testUsersBind)
}

func TestOne(t *testing.T) {
	t.Run("AutoCategories", testAutoCategoriesOne)
	t.Run("AutoOffers", testAutoOffersOne)
	t.Run("AutoStores", testAutoStoresOne)
	t.Run("Biddings", testBiddingsOne)
	t.Run("Companies", testCompaniesOne)
	t.Run("Employers", testEmployersOne)
	t.Run("Users", testUsersOne)
}

func TestAll(t *testing.T) {
	t.Run("AutoCategories", testAutoCategoriesAll)
	t.Run("AutoOffers", testAutoOffersAll)
	t.Run("AutoStores", testAutoStoresAll)
	t.Run("Biddings", testBiddingsAll)
	t.Run("Companies", testCompaniesAll)
	t.Run("Employers", testEmployersAll)
	t.Run("Users", testUsersAll)
}

func TestCount(t *testing.T) {
	t.Run("AutoCategories", testAutoCategoriesCount)
	t.Run("AutoOffers", testAutoOffersCount)
	t.Run("AutoStores", testAutoStoresCount)
	t.Run("Biddings", testBiddingsCount)
	t.Run("Companies", testCompaniesCount)
	t.Run("Employers", testEmployersCount)
	t.Run("Users", testUsersCount)
}

func TestHooks(t *testing.T) {
	t.Run("AutoCategories", testAutoCategoriesHooks)
	t.Run("AutoOffers", testAutoOffersHooks)
	t.Run("AutoStores", testAutoStoresHooks)
	t.Run("Biddings", testBiddingsHooks)
	t.Run("Companies", testCompaniesHooks)
	t.Run("Employers", testEmployersHooks)
	t.Run("Users", testUsersHooks)
}

func TestInsert(t *testing.T) {
	t.Run("AutoCategories", testAutoCategoriesInsert)
	t.Run("AutoCategories", testAutoCategoriesInsertWhitelist)
	t.Run("AutoOffers", testAutoOffersInsert)
	t.Run("AutoOffers", testAutoOffersInsertWhitelist)
	t.Run("AutoStores", testAutoStoresInsert)
	t.Run("AutoStores", testAutoStoresInsertWhitelist)
	t.Run("Biddings", testBiddingsInsert)
	t.Run("Biddings", testBiddingsInsertWhitelist)
	t.Run("Companies", testCompaniesInsert)
	t.Run("Companies", testCompaniesInsertWhitelist)
	t.Run("Employers", testEmployersInsert)
	t.Run("Employers", testEmployersInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
}

func TestReload(t *testing.T) {
	t.Run("AutoCategories", testAutoCategoriesReload)
	t.Run("AutoOffers", testAutoOffersReload)
	t.Run("AutoStores", testAutoStoresReload)
	t.Run("Biddings", testBiddingsReload)
	t.Run("Companies", testCompaniesReload)
	t.Run("Employers", testEmployersReload)
	t.Run("Users", testUsersReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("AutoCategories", testAutoCategoriesReloadAll)
	t.Run("AutoOffers", testAutoOffersReloadAll)
	t.Run("AutoStores", testAutoStoresReloadAll)
	t.Run("Biddings", testBiddingsReloadAll)
	t.Run("Companies", testCompaniesReloadAll)
	t.Run("Employers", testEmployersReloadAll)
	t.Run("Users", testUsersReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("AutoCategories", testAutoCategoriesSelect)
	t.Run("AutoOffers", testAutoOffersSelect)
	t.Run("AutoStores", testAutoStoresSelect)
	t.Run("Biddings", testBiddingsSelect)
	t.Run("Companies", testCompaniesSelect)
	t.Run("Employers", testEmployersSelect)
	t.Run("Users", testUsersSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("AutoCategories", testAutoCategoriesUpdate)
	t.Run("AutoOffers", testAutoOffersUpdate)
	t.Run("AutoStores", testAutoStoresUpdate)
	t.Run("Biddings", testBiddingsUpdate)
	t.Run("Companies", testCompaniesUpdate)
	t.Run("Employers", testEmployersUpdate)
	t.Run("Users", testUsersUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("AutoCategories", testAutoCategoriesSliceUpdateAll)
	t.Run("AutoOffers", testAutoOffersSliceUpdateAll)
	t.Run("AutoStores", testAutoStoresSliceUpdateAll)
	t.Run("Biddings", testBiddingsSliceUpdateAll)
	t.Run("Companies", testCompaniesSliceUpdateAll)
	t.Run("Employers", testEmployersSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
}
