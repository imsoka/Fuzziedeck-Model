package repositories

import (
	"strconv"

	"github.com/imsoka/fuzziedeck-model/pkg/models"
)

type InMemoryFuzziemonRepositoryInterface struct {

}

func (r InMemoryFuzziemonRepositoryInterface) GetAll() []*models.Fuzziemon {
    var fuzziemons []*models.Fuzziemon;
    var i int;

    stats := [6]uint8{5, 20, 54, 92, 54, 82};

    i = 0;
    for i < 5 {
        fuzziemon, _ := models.NewFuzziemon(
            "Fuzziemon" + strconv.Itoa(i), 
            "Este es un fuzziemon de prueba",
            uint(i),
            stats,
        );

        fuzziemons = append(fuzziemons, fuzziemon);
        i++;
    }

    return fuzziemons;
}

func(r InMemoryFuzziemonRepositoryInterface) GetById(id string) *models.Fuzziemon {

    return nil;
}
