package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/Roditu/BE_RS_TEST/util"
	"github.com/stretchr/testify/require"
)

func createRandomPerson(t *testing.T) Person {
	arg := CreatePersonParams{
		Name:			util.RandomPerson(),
		Ambition: sql.NullString{String: util.RandomAmbition(), Valid: true},
	}

	person, err := testQueries.CreatePerson(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, person)

	require.Equal(t, arg.Name, person.Name)
	require.Equal(t, arg.Ambition.String, person.Ambition)

	return person
}

func TestCreate(t *testing.T) {
	createRandomPerson(t)
}

func TestGetPerson(t *testing.T) {
	newPerson := createRandomPerson(t)
	person, err := testQueries.GetPerson(context.Background(), newPerson.ID)
	require.NoError(t, err)
	require.NotEmpty(t, person)

	require.Equal(t, newPerson.Name, person.Name)
	require.Equal(t, newPerson.Ambition.String, person.Ambition)
}