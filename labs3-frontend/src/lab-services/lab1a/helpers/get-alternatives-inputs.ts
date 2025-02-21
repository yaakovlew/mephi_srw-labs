import { Lab1a } from 'src/models/lab/lab1a';

export const getAlternativesInputs = (
  alternatives: string[]
): Lab1a.PrioritiesAlternatives => ({
  alternatives: alternatives.map(() => 0),
  choosenAlternativeIndex: 0,
});
