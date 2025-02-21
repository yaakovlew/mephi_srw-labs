import { defineStore } from 'pinia';
import { labService } from '../services/lab';
import { Ref, ref } from 'vue';
import { Lab } from 'src/models/lab/lab';
import { data } from 'autoprefixer';

export const useLabStore = defineStore('lab', () => {
  const userId: Ref<number | null> = ref(null);
  const variant: Ref<Lab.Variant | null> = ref(null);
  const info: Ref<Lab.Info | null> = ref(null);

  const getLabVariant = async () => {
    const res = await labService.getVariant();
    if (res.data) {
      userId.value = res.data.user_id;
      variant.value = res.data.variant;
    }
  };

  const sendAlternativeSet = async (data: Lab.AlternativeSet) => {
    const res = await labService.sendAlternativeSet(data);
    return res;
  };

  const sendAlternativeDiffMatrices = async (
    data: Lab.AlternativeDiffMatrices
  ) => {
    const res = await labService.sendAlternativeDiffMatrices(data);
    return res;
  };

  const getInfo = async () => {
    const res = await labService.getInfo();
    if (res.data) {
      info.value = res.data;
    }
  };

  const sendSecondNonDominated = async (data: Lab.NonDominated) => {
    const res = await labService.sendSecondNonDominated(data);
    return res;
  };

  const sendNonDominated = async (data: Lab.NonDominated) => {
    const res = await labService.sendNonDominated(data);
    return res;
  };

  const sendIntersections = async (data: Lab.Intersection) => {
    const res = await labService.sendIntersection(data);
    return res;
  };

  const sendCoff = async (data: Lab.Coff) => {
    const res = await labService.sendCoff(data);
    return res;
  };

  const sendResult = async (data: Lab.Result) => {
    const res = await labService.sendResult(data);
    return res;
  };

  const incrementSecondStep = async () => {
    await labService.incrementSecondStep();
  };

  return {
    getLabVariant,
    sendAlternativeSet,
    getInfo,
    sendSecondNonDominated,
    sendAlternativeDiffMatrices,
    sendIntersections,
    sendNonDominated,
    sendCoff,
    sendResult,
    incrementSecondStep,
    userId,
    variant,
    info,
  };
});
