package piles

func positionMinimum(pile []int) int {
	position := 0
	for i, nombre := range pile {
		if nombre < pile[position] {
			position = i
		}
	}
	return position
}

func Trier(nombres []int) []string {
	p := Piles{A: nombres}
	var instructions []string
	jouer := func(instruction string) {
		p.Appliquer(instruction)
		instructions = append(instructions, instruction)
	}

	if EstTriee(p.A) {
		return instructions
	}

	if len(p.A) <= 5 {
		trierPetitePile(&p, jouer)
	} else {
		trierParBits(&p, jouer)
	}
	return instructions
}

func trierPetitePile(p *Piles, jouer func(string)) {
	for len(p.A) > 3 && !EstTriee(p.A) {
		position := positionMinimum(p.A)
		if position <= len(p.A)/2 {
			for i := 0; i < position; i++ {
				jouer("ra")
			}
		} else {
			for i := len(p.A) - position; i > 0; i-- {
				jouer("rra")
			}
		}
		if !EstTriee(p.A) {
			jouer("pb")
		}
	}

	if len(p.A) == 3 {
		premier, deuxieme, troisieme := p.A[0], p.A[1], p.A[2]
		if premier > deuxieme && premier > troisieme {
			jouer("ra")
		} else if deuxieme > premier && deuxieme > troisieme {
			jouer("rra")
		}
	}
	if !EstTriee(p.A) {
		jouer("sa")
	}
	for len(p.B) > 0 {
		jouer("pa")
	}
}

func trierParBits(p *Piles, jouer func(string)) {
	minimum := p.A[positionMinimum(p.A)]
	var bitsUtiles uint
	for _, nombre := range p.A {
		bitsUtiles |= uint(nombre) - uint(minimum)
	}

	for masque := uint(1); masque != 0 && masque <= bitsUtiles; masque <<= 1 {
		if bitsUtiles&masque == 0 {
			continue
		}
		taille := len(p.A)
		for i := 0; i < taille; i++ {
			ecart := uint(p.A[0]) - uint(minimum)
			if ecart&masque == 0 {
				jouer("pb")
			} else {
				jouer("ra")
			}
		}
		for len(p.B) > 0 {
			jouer("pa")
		}
		if EstTriee(p.A) {
			return
		}
	}
}
