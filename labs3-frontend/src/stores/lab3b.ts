import { defineStore } from 'pinia';
import { lab3bService } from '../services/lab3b';
import { Ref, ref } from 'vue';
import { Lab } from 'src/models/lab/lab';
import { Lab3b } from 'src/models/lab/lab3b';

export const useLab3bStore = defineStore('lab3b', () => {
  const userId: Ref<number | null> = ref(null);
  const variant: Ref<Lab.Variant | null> = ref(null);
  const info: Ref<Lab.Info | null> = ref(null);
  const matrix: Ref<number[][]> = ref([]);
  const nextMatricesHeaders: Ref<string[][]> = ref([]);

  const getLabVariant = async () => {
    const res = await lab3bService.getVariant();
    if (res.data) {
      userId.value = res.data.user_id;
      variant.value = res.data.variant;
      matrix.value = res.data.matrix;
      nextMatricesHeaders.value = res.data.next_matrices_headers;
    }
  };

  const sendRuleNumber = async (data: Lab3b.RuleNumber) => {
    const res = await lab3bService.sendRuleNumber(data);
    return res;
  };

  const sendAllMatrices = async (data: Lab3b.sendAllMatrices) => {
    const res = await lab3bService.sendAllMatrices(data);
    return res;
  };

  const getInfo = async () => {
    const res = await lab3bService.getInfo();
    if (res.data) {
      info.value = res.data;
    }
  };

  const sendAllMatricesIntersection = async (
    data: Lab3b.sendAllMatricesIntersection
  ) => {
    const res = await lab3bService.sendAllMatricesIntersection(data);
    return res;
  };

  const sendLevelSet = async (data: Lab3b.SendLevelSet) => {
    const res = await lab3bService.sendLevelSet(data);
    return res;
  };

  const sendRuleValue = async (data: unknown) => {
    const res = await lab3bService.sendRuleValue(data);
    return res;
  };

  const sendCoff = async (data: Lab.Coff) => {
    const res = await lab3bService.sendCoff(data);
    return res;
  };

  const sendResult = async (data: Lab.Result) => {
    const res = await lab3bService.sendResult(data);
    return res;
  };

  const increment0Step = async () => {
    await lab3bService.increment0Step();
  };

  const increment2Step = async () => {
    await lab3bService.increment2Step();
  };

  return {
    getLabVariant,
    sendRuleNumber,
    getInfo,
    sendAllMatricesIntersection,
    sendAllMatrices,
    sendRuleValue,
    sendLevelSet,
    sendCoff,
    sendResult,
    userId,
    variant,
    info,
    matrix,
    nextMatricesHeaders,
    increment0Step,
    increment2Step,
  };
});
