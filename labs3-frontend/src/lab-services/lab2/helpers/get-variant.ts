import { getRandomInt } from 'src/lab-services/base/get-random-number';
import { lab1aVariants } from 'src/mock/lab1a';
import { lab2Variants } from 'src/mock/lab2';
import { Lab2 } from 'src/models/lab/lab2';

export const getVariant = (choosenVariantNumber?: number): Lab2.VariantInfo => {
  const totalNumberCount = Object.keys(lab1aVariants).length;
  const variantNumber =
    choosenVariantNumber ?? getRandomInt(1, totalNumberCount);
  const variant = lab2Variants[variantNumber];

  return {
    variantNumber,
    variant,
  };
};
