import { defineStore } from 'pinia';
import { Ref, ref, watch } from 'vue';
import { Lab } from 'src/models/lab/lab';
import { Lab3b } from 'src/models/lab/lab3b';
import { lab3cService } from 'src/services/lab3c';
import { Lab3c } from 'src/models/lab/lab3c';

export const useLab3cStore = defineStore('lab3c', () => {
  const userId: Ref<number | null> = ref(null);
  const variant: Ref<Lab.Variant | null> = ref(null);
  const info: Ref<Lab.Info | null> = ref(null);
  const normalStep: Ref<number> = ref(
    Number(localStorage.getItem('3c-normal-step')) || 1
  );
  const matrix: Ref<number[][]> = ref([]);

  const getLabVariant = async () => {
    const res = await lab3cService.getVariant();
    if (res.data) {
      userId.value = res.data.user_id;
      variant.value = res.data.variant;
      matrix.value = res.data.matrix;
    }
  };

  const sendAlternativeMatrix = async (data: Lab3c.AlternativeMatrix) => {
    const res = await lab3cService.sendAlternativeMatrix(data);
    return res;
  };

  const sendCriteriaMatrix = async (data: Lab3c.CriteriaMatirx) => {
    const res = await lab3cService.sendCriteriaMatrix(data);
    return res;
  };

  const sendAlternativeMatrices = async (data: Lab3c.AlternativeMatrices) => {
    const res = await lab3cService.sendAlternativeMatrices(data);
    return res;
  };

  const sendCriteriaEstimation = async (data: Lab3c.CriteriaEstimation) => {
    const res = await lab3cService.sendCriteriaEstimation(data);
    return res;
  };

  const sendEstimation = async (data: Lab3c.CriteriaEstimation) => {
    const res = await lab3cService.sendEstimation(data);
    return res;
  };

  const sendArea = async (data: Lab3c.Area) => {
    const res = await lab3cService.sendArea(data);
    return res;
  };

  const sendLine = async (data: Lab3c.LineParameters) => {
    const res = await lab3cService.sendLine(data);
    return res;
  };

  const sendQuadratic = async (data: Lab3c.KvLine) => {
    const res = await lab3cService.sendQuadratic(data);
    return res;
  };

  const sendResult = async (data: Lab3c.Result) => {
    const res = await lab3cService.sendResult(data);
    return res;
  };

  const getInfo = async () => {
    const res = await lab3cService.getInfo();
    if (res.data) {
      info.value = res.data;
    }
  };

  const increment0Step = async () => {
    await lab3cService.increment0Step();
  };

  const increment2Step = async () => {
    await lab3cService.increment2Step();
  };

  watch(normalStep, () => {
    localStorage.setItem('3c-normal-step', normalStep.value.toString());
  });

  return {
    getLabVariant,
    sendAlternativeMatrix,
    getInfo,
    userId,
    variant,
    info,
    sendAlternativeMatrices,
    sendCriteriaMatrix,
    sendCriteriaEstimation,
    sendEstimation,
    sendArea,
    sendLine,
    sendQuadratic,
    sendResult,
    increment2Step,
    increment0Step,
    matrix,
    normalStep,
  };
});
