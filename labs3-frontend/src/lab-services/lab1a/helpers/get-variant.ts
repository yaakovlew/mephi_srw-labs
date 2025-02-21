import { getRandomInt } from 'src/lab-services/base/get-random-number';
import { lab1aVariants } from 'src/mock/lab1a';
import { Lab1a } from 'src/models/lab/lab1a';

export const getVariant = (
  choosenVariantNumber?: number
): Lab1a.VariantInfo => {
  const totalNumberCount = Object.keys(lab1aVariants).length;
  const variantNumber =
    choosenVariantNumber ?? getRandomInt(1, totalNumberCount);
  const variant = lab1aVariants[variantNumber];

  return {
    variantNumber,
    variant,
  };
};
