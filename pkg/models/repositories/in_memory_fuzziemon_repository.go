package repositories

import "github.com/imsoka/fuzziedeck-model/pkg/models"


type InMemoryFuzziemonRepository struct {

}

func (r InMemoryFuzziemonRepository) GetAll() ([]*models.Fuzziemon, error) {
    fzz := make([]*models.Fuzziemon, 0)
    stats := [6]uint8{1, 2, 3, 4, 5, 6}

    fzz = append(fzz, models.NewFuzziemon(
        "Nuchitas",
        "El primer fuzziemon",
        1,
        stats, 
    ))

    fzz = append(fzz, models.NewFuzziemon(
        "Solete",
        "El segundo fuzziemon",
        2,
        stats, 
    ))

    fzz = append(fzz, models.NewFuzziemon(
        "Manguito",
        "El tercer fuzziemon",
        3,
        stats, 
    ))

    return fzz, nil
}

