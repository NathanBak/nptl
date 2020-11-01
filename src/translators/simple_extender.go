package translators

import "context"

type SimpleExtender struct{}

func (e *SimpleExtender) Extend(ctx context.Context, source, target, prefix, suffix []rune, extenderRune rune) ([]rune, error) {
	response := []rune{}
	response = append(response, prefix...)
	response = append(response, target...)
	response = append(response, extenderRune)

	targetLength := int((float32(len(source)) * 1.2))

	for len(response) < targetLength-1 {
		response = append(response, extenderRune)
	}

	response = append(response, suffix...)

	return response, nil
}
