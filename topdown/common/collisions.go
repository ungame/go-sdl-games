package common

import "math"

func Collides(c1, c2 Circle) bool {
	sum := math.Pow(c2.Center.X-c1.Center.X, 2) + math.Pow(c2.Center.Y-c1.Center.Y, 2)
	distance := math.Sqrt(sum)
	return distance <= c1.Radius+c2.Radius
}

func CheckCollisions(elements []*Element) error {
	for i := 0; i < len(elements) - 1; i++ {
		for j := i + 1; j < len(elements) ; j++ {

			before := elements[i]
			next := elements[j]

			if !before.Active || !next.Active {
				continue
			}

			for _, c1 := range before.Collisions {
				for _, c2 := range next.Collisions {
					if Collides(c1, c2) {
						if err := before.Collision(next); err != nil {
							return err
						}
						if err := next.Collision(before); err != nil {
							return err
						}
					}
				}
			}

		}
	}
	return nil
}